package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	pb "go-agent/agent_proto/database"
	"go-agent/agent_runtime"
	"go-agent/utils"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

var connectionInfoKey2path = make(map[ConnectionInfoForKey]string)

var defaultConnectionInfo = ConnectionInfoForKey{
	Username:        "root",
	Password:        "",
	Host:            "127.0.0.1",
	Port:            3306,
	SkipGrantTables: true,
}

func (s *Server) MysqlDump(_ context.Context, request *pb.MysqlDumpRequest) (*pb.MysqlDumpResponse, error) {
	response := &pb.MysqlDumpResponse{}
	//var err error = nil
	utils.LogInfo(fmt.Sprintf("Received request: %s", mysqlDumpRequestWrapper{request}))
	//default 如果request传递了,覆盖
	var key = defaultConnectionInfo
	key.SkipGrantTables = request.SkipGrantTables

	//不跳过验证,还不传ConnectionInfo 就全用默认值,连不上就返回
	if !request.SkipGrantTables && request.ConnectionInfo != nil {
		key.Host = request.ConnectionInfo.Host
		key.Port = request.ConnectionInfo.Port
		key.Username = request.ConnectionInfo.Username
		key.Password = request.ConnectionInfo.Password
	}

	mysqldCmd, _ := utils.FindCommandFromPathAndProcessByMatchStringArray([]string{"mysqld"})
	if mysqldCmd == "" {
		response.Message = "mysqld not found"
		utils.LogError(response.Message)
		return response, nil
	}

	if connectionInfoKey2path[key] != "" && !request.Force {
		response.Filepath = connectionInfoKey2path[key]
		return response, nil
	}
	var outFile string
	var err error
	if request.SkipGrantTables {
		err = platformRestartMysqlSkipGrantTables(mysqldCmd)
		if err != nil {
			response.Message = err.Error()
			return response, nil
		}

		outFile, err = platformUseMysqldump(mysqldCmd)
		if err != nil {
			response.Message = err.Error()
		}
		err = platformRestartMysqlDefault()
		if err != nil {
			utils.LogError(err.Error())
		}

	} else {
		config := mysql.NewConfig()

		config.User = key.Username
		config.Passwd = key.Password
		config.Addr = fmt.Sprintf("%s:%d", key.Host, key.Port)
		config.Net = "tcp"
		config.DBName = "mysql"

		var db, _ = sql.Open("mysql", config.FormatDSN())
		utils.LogWarn(config.FormatDSN())

		if !canConnect(db) {
			response.Message = fmt.Sprintf("Cannot connect to %s:%d with username %s and password %s", key.Host, key.Port, key.Username, key.Password)
			// log err 暴露的信息有点多
			utils.LogError(response.Message)
			return response, nil
		}
		outFile = doDump(db, &key, config)
	}

	// 失败file是空,这个赋值也没问题
	connectionInfoKey2path[key] = outFile

	response.Filepath = outFile
	return response, nil
}

func doDump(db *sql.DB, key *ConnectionInfoForKey, config *mysql.Config) string {

	err := os.MkdirAll(agent_runtime.OutDir, 0700)
	if err != nil {
		return ""
	}
	fileName := sqlName(key.Host)
	databases, err := getDatabases(db)
	if err != nil {
		utils.LogError(err.Error())
		return ""
	}
	f, _ := os.Create(fileName)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	//跳过基础4库
	//mysql,information_schema,performance_schema,sys
	basisDatabases := []string{"mysql", "information_schema", "performance_schema", "sys"}
	for _, database := range databases {
		if slices.Contains(basisDatabases, database) {
			continue
		}
		// _, err := db.Exec("USE " + databaseName)
		config.DBName = database
		dsn := config.FormatDSN()
		eachDb, err := sql.Open("mysql", dsn)
		if err != nil {
			utils.LogError(err.Error())
			continue
		}
		_, err = f.WriteString(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`;\n", database))
		_, err = f.WriteString(fmt.Sprintf("USE `%s`;\n", database))
		tables, err := getTables(eachDb)
		if err != nil {
			utils.LogError(err.Error())
			err := eachDb.Close()
			if err != nil {
				return ""
			} // 在执行continue之前关闭数据库连接
			continue
		}
		//这里怎么insert 区分不同的数据库
		// Write database name

		// Dump each table
		for _, table := range tables {
			err := dumpCreateTableSQL(eachDb, table, f)
			if err != nil {
				return ""
			}
			err = dumpTable(eachDb, table, f)
			if err != nil {
				utils.LogError(fmt.Sprintf("Error dumping table %s: %v\n", table, err))
			}
		}

		err = eachDb.Close()
		if err != nil {
			return ""
		}
	}

	return fileName
}

func sqlName(host string) string {
	return filepath.Join(agent_runtime.OutDir, host+"_"+utils.FormatTimeForFileName(time.Now())+".sql")
}

func canConnect(db *sql.DB) bool {

	if err := db.Ping(); err != nil {
		utils.LogError(err.Error())
		return false
	}
	return true
}

func getDatabases(db *sql.DB) ([]string, error) {
	var databases []string
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var databaseName string
	for rows.Next() {
		err := rows.Scan(&databaseName)
		if err != nil {
			return nil, err
		}
		databases = append(databases, databaseName)
	}
	return databases, nil

}

func getTables(db *sql.DB) ([]string, error) {
	var tables []string
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var tableName string
	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}
	return tables, nil
}

func dumpCreateTableSQL(db *sql.DB, tableName string, f *os.File) error {
	query := fmt.Sprintf("SHOW CREATE TABLE %s", tableName)
	row := db.QueryRow(query)
	var tableNameReturned, sqlStatement string
	if err := row.Scan(&tableNameReturned, &sqlStatement); err != nil {
		return err
	}
	_, err := f.WriteString(fmt.Sprintf("%s;\n", sqlStatement))
	return err
}

func dumpTable(db *sql.DB, tableName string, file *os.File) error {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	var buffer [][]sql.RawBytes // Buffer for rows
	scanArgs := make([]interface{}, len(columns))
	for i := range scanArgs {
		scanArgs[i] = new(sql.RawBytes)
	}

	for rows.Next() {
		values := make([]sql.RawBytes, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return err
		}

		buffer = append(buffer, values)
		if len(buffer) >= 1000 {
			if err := writeInserts(tableName, columns, buffer, file); err != nil {
				return err
			}
			buffer = buffer[:0] // Clear the buffer
		}
	}

	// Insert any remaining rows
	if len(buffer) > 0 {
		if err := writeInserts(tableName, columns, buffer, file); err != nil {
			return err
		}
	}

	return nil
}

func writeInserts(tableName string, columns []string, rows [][]sql.RawBytes, file *os.File) error {
	var inserts []string
	for _, row := range rows {
		var valueStrings []string
		for _, value := range row {
			valueStrings = append(valueStrings, fmt.Sprintf("'%s'", escapeString(string(value))))
		}
		inserts = append(inserts, fmt.Sprintf("(%s)", strings.Join(valueStrings, ", ")))
	}
	_, err := file.WriteString(fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;\n", tableName, strings.Join(columns, ", "), strings.Join(inserts, ",\n")))
	return err
}

func escapeString(value string) string {
	// Simple escape, replace ' with '' for SQL values
	return strings.ReplaceAll(value, "'", "''")
}

func copyBakAndReplaceWithSkipGrantTables(origin, bak string) {

	//copy
}
