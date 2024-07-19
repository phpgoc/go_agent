package file

import (
	"errors"
	"github.com/shirou/gopsutil/v4/disk"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"os/exec"
)

func platformDiskMirror(req *pb.DiskMirrorRequest, resStream pb.FileService_DiskMirrorServer) error {
	device := req.Device
	disks, err := disk.Partitions(false)
	if err != nil {
		utils.LogError(err.Error())
		return err
	}
	found := false
	for _, d := range disks {
		if d.Device == device {
			found = true
			break
		}
	}
	if !found {
		utils.LogError("device not found")
		return errors.New("device not found")
	}
	cmd := exec.Command("dd", "if="+device, "of=/dev/stdout")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		utils.LogError("Failed to create stdout pipe: " + err.Error())
		return err
	}

	// 启动命令
	if err = cmd.Start(); err != nil {
		utils.LogError("Failed to start command: " + err.Error())
		return err
	}
	var chunk [1024]byte

	//md, ok := metadata.FromIncomingContext(resStream.Context()) // 可以获取metadata，相当于在接收或发送数据之前互相发送一些元数据
	for {
		n, err := stdout.Read(chunk[:])
		if err != nil {
			break
		}
		if n == 0 {
			break
		}
		err = resStream.Send(&pb.DiskMirrorResponse{Chunk: chunk[:n]})
		if err != nil {
			utils.LogError(err.Error())
			return err
		}
	}
	return nil
}
