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
	//<VirtualHost *:80 *:1888>
	//    DocumentRoot "/www/example1"
	//    ServerName www.example1.com
	//    # Other directives here
	//</VirtualHost>
	//
	//<VirtualHost *:80>
	//    DocumentRoot "/www/example2"
	//    ServerName www.example2.org
	//     CustomLog logs/example2-access_log  common
	//    # Other directives here
	//</VirtualHost>
	//
	//
	//<VirtualHost *>
	//    DocumentRoot "/www/example3"
	//    ServerName www.example3.org
	//     CustomLog "logs/example3-access_log" common
	//    # Other directives here
	//</VirtualHost>
	for _, virtualHost := range response.VirtualHosts {
		if virtualHost.Root == "/www/example1" {

			require.True(t,
				reflect.DeepEqual(virtualHost.ServerNames, []string{"www.example1.com"}), "Expected server names to be [www.example1.com]")
			require.True(t, reflect.DeepEqual(virtualHost.Listens, []string{"*:80", "*:1888"}), "Expected listens to be [*:80, *:1888]")
			equalNum++
		} else if virtualHost.Root == "/www/example2" {

			require.True(t, reflect.DeepEqual(virtualHost.ServerNames, []string{"www.example2.org"}), "Expected server names to be [www.example2.org]")
			require.True(t, reflect.DeepEqual(virtualHost.Listens, []string{"*:80"}), "Expected listens to be [*:80]")
			require.True(t, strings.Contains(virtualHost.CustomLog.FilePath, "logs/example2-access_log"), "Expected custom log file path to contain logs/example2-access_log")
			equalNum++
		} else if virtualHost.Root == "/www/example3" {

			require.True(t, reflect.DeepEqual(virtualHost.ServerNames, []string{"www.example3.org"}), "Expected server names to be [www.example3.org]")
			require.True(t, reflect.DeepEqual(virtualHost.Listens, []string{"*"}), "Expected listens to be [*]")
			require.True(t, strings.Contains(virtualHost.CustomLog.FilePath, "logs/example3-access_log"), "Expected custom log file path to contain logs/example3-access_log")

			equalNum++
		}
	}
	require.Equal(t, 3, equalNum, "Expected 3 virtual hosts to match the expected configuration")

}
