package nginx

import (
	"github.com/stretchr/testify/assert"
	pb "go-agent/agent_proto"
	"path/filepath"
	"runtime"
	"testing"
)

func TestInsertNginxInfo(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	configFile := filepath.Join(dir, "test_data/simple/nginx.conf")

	var res pb.GetNginxInfoResponse
	insertNginxInfo(configFile, &res)

	// Example assertion: check if the response contains expected values
	// This is a simplistic check; your actual checks should be based on the expected outcome from parsing nginx.conf
	if len(res.NginxInstances) == 0 {
		t.Errorf("Expected at least one NginxInstance, got %d", len(res.NginxInstances))
	}
	assert.Equal(t, "error.log", res.NginxInstances[0].ErrorLog.FilePath)

	// Add more assertions as needed based on the expected content of nginx.conf
}
