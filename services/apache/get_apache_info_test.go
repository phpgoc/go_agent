package apache

import (
	"github.com/stretchr/testify/require"
	pb "go-agent/agent_proto"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestInsertApacheInstanceWithSimpleConfig(t *testing.T) {
	// Assuming simple.config is located in the same directory as this test file
	configFilePath, err := filepath.Abs("./test_data/simple.config")
	if err != nil {
		t.Fatalf("Failed to get absolute path of simple.config: %v", err)
	}

	// Mock httpdRoot - can be the directory of simple.config or any relevant path
	httpdRoot := filepath.Dir(configFilePath)

	// Create a new GetApacheInfoResponse instance
	response := &pb.GetApacheInfoResponse{}

	// Mock environment variables map if needed
	envMap := make(map[string]string)
	// Populate envMap as necessary, e.g., envMap["VAR_NAME"] = "value"

	// Call the function under test
	err = insertApacheInstance(configFilePath, httpdRoot, response, envMap)
	if err != nil {
		t.Errorf("insertApacheInstance returned an error: %v", err)
	}

	require.Equal(t, 3, len(response.VirtualHosts), "Expected 3 virtual hosts in the response")
	require.True(t, reflect.DeepEqual([]string{"8888"}, response.Listens), "Expected listens to be [8888]")
	require.True(t, reflect.DeepEqual([]string{"abc.com"}, response.ServerNames), "Expected server names to be [abc.com]")
	var equalNum = 0

	for _, virtualHost := range response.VirtualHosts {
		if virtualHost.Root == "/www/example1" {

			require.True(t,
				reflect.DeepEqual(virtualHost.ServerNames, []string{"www.example1.com"}), "Expected server names to be [www.example1.com]")
			require.True(t, reflect.DeepEqual(virtualHost.Listens, []string{"*:80", "*:1888"}), "Expected listens to be [*:80, *:1888]")
			equalNum++
		} else if virtualHost.Root == "/www/example2" {

			require.True(t, reflect.DeepEqual(virtualHost.ServerNames, []string{"www.example2.org"}), "Expected server names to be [www.example2.org]")
			require.True(t, reflect.DeepEqual(virtualHost.Listens, []string{"*:80"}), "Expected listens to be [*:80]")
			require.True(t, strings.HasSuffix(virtualHost.CustomLog.FilePath, "example2-access_log"), "Expected custom log file path to contain logs/example2-access_log")
			equalNum++
		} else if virtualHost.Root == "/www/example3" {

			require.True(t, reflect.DeepEqual(virtualHost.ServerNames, []string{"www.example3.org"}), "Expected server names to be [www.example3.org]")
			require.True(t, reflect.DeepEqual(virtualHost.Listens, []string{"*"}), "Expected listens to be [*]")
			require.True(t, strings.HasSuffix(virtualHost.CustomLog.FilePath, "example3-access_log"), "Expected custom log file path to contain logs/example3-access_log")

			equalNum++
		}
	}
	require.Equal(t, 3, equalNum, "Expected 3 virtual hosts to match the expected configuration")

}

func TestInsertApacheInstanceWithInclude(t *testing.T) {
	configFilePath, err := filepath.Abs("./test_data/include/apache.config")
	if err != nil {
		t.Fatalf("Failed to get absolute path of simple.config: %v", err)
	}
	httpdRoot := filepath.Dir(configFilePath)

	response := &pb.GetApacheInfoResponse{}

	// Mock environment variables map if needed
	envMap := map[string]string{"INCLUDE_PATH": httpdRoot}

	err = insertApacheInstance(configFilePath, httpdRoot, response, envMap)
	require.True(t, reflect.DeepEqual([]string{"8081", "8082"}, response.Listens), "Expected listens to be [8081, 8082]")

	require.Equal(t, 4, len(response.ConfigFiles), "Expected 4 config files in the response")

	require.Equal(t, 2, len(response.VirtualHosts), "Expected 2 virtual hosts in the response")
}

func TestInsertApacheInstanceInLogAndOutLog(t *testing.T) {
	configFilePath, err := filepath.Abs("./test_data/inlog_outlog/apache.config")
	if err != nil {
		t.Fatalf("Failed to get absolute path of simple.config: %v", err)
	}
	httpdRoot := filepath.Dir(configFilePath)

	response := &pb.GetApacheInfoResponse{}

	envMap := map[string]string{}
	err = insertApacheInstance(configFilePath, httpdRoot, response, envMap)

	require.True(t, true)
	require.True(t, len(response.ErrorLogs) == 1, "Expected 1 global error log")
	require.True(t, strings.HasSuffix(response.ErrorLogs[0].FilePath, "error_out.log"), "Expected error log to end with error_out.log")
	require.True(t, len(response.CustomLogs) == 1, "Expected 1 global custom log")
	require.True(t, strings.HasSuffix(response.CustomLogs[0].FilePath, "custom_out.log"), "Expected custom log to end with custom_out.log")
	require.True(t, len(response.VirtualHosts) == 1, "Expected 1 virtual host")
	require.True(t, response.VirtualHosts[0].ErrorLog != nil, "Expected virtual host to have an error log")
	require.True(t, strings.HasSuffix(response.VirtualHosts[0].ErrorLog.FilePath, "error_in.log"), "Expected virtual host error log to end with error_in.log")
	require.True(t, response.VirtualHosts[0].CustomLog != nil, "Expected virtual host to have a custom log")
	require.True(t, strings.HasSuffix(response.VirtualHosts[0].CustomLog.FilePath, "custom_in.log"), "Expected virtual host custom log to end with custom_in.log")
}

func TestInsertApacheInstanceWithDeepInclude(t *testing.T) {

	configFilePath, err := filepath.Abs("./test_data/double_include/apache.config")
	if err != nil {
		t.Fatalf("Failed to get absolute path of simple.config: %v", err)
	}
	httpdRoot := filepath.Dir(configFilePath)

	response := &pb.GetApacheInfoResponse{}

	envMap := map[string]string{"TEST_ROOT": httpdRoot}
	err = insertApacheInstance(configFilePath, httpdRoot, response, envMap)
	require.True(t, 3 == len(response.ConfigFiles), "Expected 3 config files in the response")
	require.True(t, reflect.DeepEqual([]string{"8081", "8082", "8083"}, response.Listens), "Expected listens to be [8081, 8082, 8083]")
}
