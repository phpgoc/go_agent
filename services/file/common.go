package file

import (
	"errors"
	"fmt"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"os"
)

type Server struct {
	pb.UnimplementedFileServer
}

func (s *Server) DownloadFile(req *pb.DownloadFileRequest, resStream pb.File_DownloadFileServer) error {
	utils.LogInfo(fmt.Sprintf("called DownloadFile, filename: %s", req.Filename))
	var chunk [1024]byte
	fileName := req.Filename
	if fileName == "" {
		utils.LogError("filename is empty")
		return errors.New("filename is empty")
	}
	// assert file exists in os
	fileReader, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		//log error 可以记录得更详细
		utils.LogError(err.Error())
		return errors.New("file not found")
	}
	//md, ok := metadata.FromIncomingContext(resStream.Context()) // 可以获取metadata，相当于在接收或发送数据之前互相发送一些元数据
	for {
		n, err := fileReader.Read(chunk[:])
		if err != nil {
			break
		}
		if n == 0 {
			break
		}
		err = resStream.Send(&pb.DownloadFileResponse{Chunk: chunk[:n]})
		if err != nil {
			utils.LogError(err.Error())
			return err
		}
	}
	return nil
}
