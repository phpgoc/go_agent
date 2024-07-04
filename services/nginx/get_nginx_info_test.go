package nginx

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pb "go-agent/agent_proto"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"
)

func TestInsertNginxInfoSimple(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	configFile := filepath.Join(dir, "test_data/simple/nginx.conf")

	var res pb.GetNginxInfoResponse
	insertNginxInfo(configFile, filepath.Join(dir, "test_data", "simple"), &res)
	require.Equal(t, 1, len(res.NginxInstances), "Expected 1 NginxInstance, got %d", len(res.NginxInstances))

	assert.True(t, strings.HasSuffix(res.NginxInstances[0].ErrorLog.FilePath, filepath.Join("simple", "error.log")),
		"Expected error log file path to end with simple/error.log, got %s", res.NginxInstances[0].ErrorLog.FilePath)
	require.Equal(t, 1, len(res.NginxInstances[0].Servers), "Expected 1 Server, got %d", len(res.NginxInstances[0].Servers))
	assert.Equal(t, "_", res.NginxInstances[0].Servers[0].ServerName, "Expected _ ServerName, got %s", res.NginxInstances[0].Servers[0].ServerName)
	assert.True(t, slices.Contains(res.NginxInstances[0].Servers[0].Listens, "8080"))
}

func TestInsertNginxInfoNotAbsolutePathConfig(t *testing.T) {
	var res pb.GetNginxInfoResponse
	insertNginxInfo("test_data/simple/nginx.conf", "", &res)
	assert.Equal(t, 0, len(res.NginxInstances), "Expected 0 NginxInstance, got %d", len(res.NginxInstances))
}

func TestInsertNginxInfoInclude(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	configFile := filepath.Join(dir, "test_data/include/nginx.conf")

	var res pb.GetNginxInfoResponse
	insertNginxInfo(configFile, filepath.Join(dir, "test_data", "include"), &res)
	require.Equal(t, 1, len(res.NginxInstances), "Expected 1 NginxInstance, got %d", len(res.NginxInstances))

	assert.True(t, strings.HasSuffix(res.NginxInstances[0].ErrorLog.FilePath, filepath.Join("include", "error.log")),
		"Expected error log file path to end with include/nested/error.log, got %s", res.NginxInstances[0].ErrorLog.FilePath)
	require.Equal(t, 1, len(res.NginxInstances[0].Servers), "Expected 1 Server, got %d", len(res.NginxInstances[0].Servers))
	assert.Equal(t, "a.com", res.NginxInstances[0].Servers[0].ServerName, "Expected _ ServerName, got %s", res.NginxInstances[0].Servers[0].ServerName)
	assert.True(t, slices.Contains(res.NginxInstances[0].Servers[0].Listens, "8081"))
}
