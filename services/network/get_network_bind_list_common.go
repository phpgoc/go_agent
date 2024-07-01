package network

import (
	"context"
	"github.com/shirou/gopsutil/v4/net"
	pb "go-agent/agent_proto"
	"go-agent/runtime"
	"go-agent/utils"
)

func (s *Server) GetNetworkBindList(_ context.Context, req *pb.GetNetworkBindListRequest) (*pb.GetNetworkBindListResponse, error) {
	utils.LogInfo("called GetNetworkBindList")
	res := &pb.GetNetworkBindListResponse{}
	var bindList []net.ConnectionStat = nil
	var err error
	if req.Protocol != nil {
		if *req.Protocol == "tcp" {
			bindList, err = net.Connections("tcp")
		} else if *req.Protocol == "udp" {
			bindList, err = net.Connections("udp")
		} else {
			res.Message = "Invalid protocol"
			utils.LogError(res.Message)
			return res, nil
		}
	} else {
		bindList, err = net.Connections("all")

	}
	//bindList, err := net.Connections("all")
	if err != nil {
		utils.LogError(err.Error())
		return nil, err
	}

	uniqueLAddr := make(map[addrAndPort]*networkBindAndPid)

	for _, b := range bindList {
		if uniqueLAddr[addrAndPort{ip: b.Laddr.IP, port: b.Laddr.Port}] == nil {
			uniqueLAddr[addrAndPort{ip: b.Laddr.IP, port: b.Laddr.Port}] =
				&networkBindAndPid{
					bind: &pb.NetworkBind{
						Family: family2string(b.Family),
						Type:   socketType2string(b.Type),
						Bind:   &pb.Address{Ip: b.Laddr.IP, Port: b.Laddr.Port},
						Fd:     b.Fd,
						Status: b.Status,
					},
					pid: b.Pid,
				}
		}
	}
	//从全局processes中找到对应的pid，注意processes不是时时更新的，只在agent进程初始化时录入
	for _, i := range uniqueLAddr {
		for _, p := range runtime.Processes {
			if p.Pid == i.pid {
				exe, err := p.Exe()
				if err != nil {
					continue
				}
				i.bind.Cmd = exe
			}
		}
		res.NetworkBinds = append(res.NetworkBinds, i.bind)
	}

	return res, nil
}
