package nginx

import (
	"github.com/stretchr/testify/assert"
	pb "go-agent/agent_proto"
	"path/filepath"
	"runtime"
	"testing"
)

func TestInsertNginxInfoSimple(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	configFile := filepath.Join(dir, "test_data/simple/nginx.conf")

	var res pb.GetNginxInfoResponse
	insertNginxInfo(configFile, dir, &res)

	if len(res.NginxInstances) == 0 {
		t.Errorf("Expected at least one NginxInstance, got %d", len(res.NginxInstances))
	}
	assert.Equal(t, "error.log", res.NginxInstances[0].ErrorLog.FilePath)
}

func TestInsertNginxInfoNotAbsolutePathConfig(t *testing.T) {
	var res pb.GetNginxInfoResponse
	insertNginxInfo("test_data/simple/nginx.conf", "", &res)
	assert.Equal(t, 0, len(res.NginxInstances), "Expected 0 NginxInstance, got %d", len(res.NginxInstances))
}
