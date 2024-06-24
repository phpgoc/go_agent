package get_sys_info

import (
	"context"
	"github.com/elastic/go-sysinfo"
	pb "go-agent/agent_proto"
	"go-agent/utils"
)

type GetSysInfoServer struct {
	pb.UnimplementedGetSysInfoServer
}

func (s *GetSysInfoServer) GetApacheInfo(_ context.Context, _in *pb.GetSysInfoRequest) (*pb.GetSysInfoResponse, error) {
	host, err := sysinfo.Host()
	var res pb.GetSysInfoResponse
	if err != nil {
		res.Message = err.Error()
		return &res, err
	}
	var info = host.Info()
	// why info is nil
	//

	//print(info)
	res.Timezone = info.Timezone
	res.Hostname = info.Hostname
	res.BootTime = utils.FormatTime(info.BootTime)
	res.SysArch = info.Architecture

	print(info.Hostname)
	//print(res)
	println("GetSysInfoServer GetApacheInfo")
	return &res, nil

}
