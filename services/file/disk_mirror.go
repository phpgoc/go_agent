package file

import (
	"fmt"
	pb "go-agent/agent_proto"
	"go-agent/utils"
)

// DownloadFile 性能还不错 ，肯定和ssd有关，本地拷贝600M用不到 1秒
func (s *Server) DiskMirror(req *pb.DiskMirrorRequest, resStream pb.FileService_DiskMirrorServer) error {
	utils.LogInfo(fmt.Sprintf("called DownloadFile, disk: %s", req.Device))
	return platformDiskMirror(req, resStream)
}
