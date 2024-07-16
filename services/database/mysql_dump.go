package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	pb "go-agent/agent_proto"
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
var linkTargetPath string

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
		return utils.SetResponseErrorAndLogMessageGeneric(response, "mysqld not found", pb.ResponseCode_UNSUPPORTED)
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
			response.Code = pb.ResponseCode_UNKNOWN_SERVER_ERROR
			return response, nil
			//不用下边的是因为已经在platformRestartMysqlSkipGrantTables里面打印了
			//return utils.SetResponseErrorAndLogMessageGeneric(response, err.Error(), pb.ResponseCode_UNKNOWN_SERVER_ERROR)
		}
		var code pb.ResponseCode
		outFile, code = platformUseMysqldump(mysqldCmd)
		if code != pb.ResponseCode_OK {

			response.Message = err.Error()
			response.Code = code
			return response, nil
		}
		err = platformRestartMysqlDefault()
		if err != nil {
			return utils.SetResponseErrorAndLogMessageGeneric(response, err.Error(), pb.ResponseCode_UNKNOWN_SERVER_ERROR)
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

			return utils.SetResponseErrorAndLogMessageGeneric(response, fmt.Sprintf("Cannot connect to %s:%d with username %s and password %s", key.Host, key.Port, key.Username, key.Password),
				pb.ResponseCode_UNKNOWN_SERVER_ERROR)
		}
		outFile, err = doDump(db, &key, config)
	}

	// 失败file是空,这个赋值也没问题
	connectionInfoKey2path[key] = outFile

	response.Filepath = outFile
	if err != nil {
		response.Message = err.Error()
		response.Code = pb.ResponseCode_UNKNOWN_SERVER_ERROR
	}
	return response, nil
}

func doDump(db *sql.DB, key *ConnectionInfoForKey, config *mysql.Config) (string, error) {

	err := os.MkdirAll(agent_runtime.OutDir, 0700)
	if err != nil {
		return "", utils.LogErrorThrough(&err)
	}
	fileName := sqlName(key.Host)
	databases, err := getDatabases(db)
	if err != nil {
		return "", utils.LogErrorThrough(&err)
	}
	f, _ := os.Create(fileName)
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			utils.LogError(err.Error())
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
			utils.LogWarn(err.Error())
			continue
		}
		_, err = f.WriteString(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`;\n", database))
		_, err = f.WriteString(fmt.Sprintf("USE `%s`;\n", database))
		tables, err := getTables(eachDb)
		if err != nil {
			utils.LogError(err.Error())
			err = eachDb.Close()
			if err != nil {
				return "", utils.LogErrorThrough(&err)
			} // 在执行continue之前关闭数据库连接
			continue
		}
		//这里怎么insert 区分不同的数据库
		// Write database name

		// Dump each table
		for _, table := range tables {
			err = dumpCreateTableSQL(eachDb, table, f)
			if err != nil {
				return "", utils.LogErrorThrough(&err)
			}
			err = dumpTable(eachDb, table, f)
			if err != nil {
				utils.LogError(fmt.Sprintf("Error dumping table %s: %v\n", table, err))
			}
		}

		err = eachDb.Close()
		if err != nil {
			return "", utils.LogErrorThrough(&err)
		}
	}

	return fileName, nil
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
			utils.LogWarn(err.Error())
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
		utils.LogError(err.Error())
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
			utils.LogError(err.Error())
		}
	}(rows)

	columns, err := rows.Columns()
	if err != nil {
		return utils.LogErrorThrough(&err)
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
			return utils.LogErrorThrough(&err)
		}

		buffer = append(buffer, values)
		if len(buffer) >= 1000 {
			if err := writeInserts(tableName, columns, buffer, file); err != nil {
				return utils.LogErrorThrough(&err)
			}
			buffer = buffer[:0] // Clear the buffer
		}
	}

	// Insert any remaining rows
	if len(buffer) > 0 {
		if err := writeInserts(tableName, columns, buffer, file); err != nil {
			return utils.LogErrorThrough(&err)
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

func copyBakAndReplaceWithSkipGrantTables(original, bak string) error {

	sourceFileContent, err := utils.ReadFile(original)

	//判断是否是软链接
	linkTargetPath, err = os.Readlink(original)
	if err != nil {
		//copy
		err = utils.CopyFile(original, bak)
	} else {
		//unlink
		err = os.Remove(original)
		if err != nil {
			utils.LogError(fmt.Sprintf("Failed to remove link %s: %s", original, err))
		}
	}
	if err != nil {
		return err
	}

	newFile, _ := os.OpenFile(original, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	defer func(newFile *os.File) {
		err = newFile.Close()
		if err != nil {
			utils.LogError(err.Error())
		}
	}(newFile)

	foundMysqld := false
	for _, line := range strings.Split(sourceFileContent, "\n") {

		// Check if we're in the [mysqld] section
		if strings.TrimSpace(line) == "[mysqld]" {
			foundMysqld = true
			line += "\n skip-grant-tables"
		}
		_, _ = newFile.WriteString(line + "\n")

	}
	if !foundMysqld {
		//linux不会出现这种情况,因为这个文件是通过find [mysqld]找到的
		return fmt.Errorf("no [mysqld] section found in %s", original)
	}

	return nil
}
