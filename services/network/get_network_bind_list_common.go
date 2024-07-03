package network

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/v4/net"
	pb "go-agent/agent_proto"
	"go-agent/agent_runtime"
	"go-agent/utils"
	"slices"
	"strings"
)

func (s *Server) GetNetworkBindList(_ context.Context, req *pb.GetNetworkBindListRequest) (*pb.GetNetworkBindListResponse, error) {
	utils.LogInfo(fmt.Sprintf("called GetNetworkBindList: %v", req))
	res := &pb.GetNetworkBindListResponse{}
	var bindList []net.ConnectionStat = nil
	var err error
	switch req.Protocol {
	case pb.Protocol_ALL:
		bindList, err = net.Connections("all")
	case pb.Protocol_TCP:
		bindList, err = net.Connections("tcp")
	case pb.Protocol_UDP:
		bindList, err = net.Connections("udp")
	}

	if err != nil {
		utils.LogError(err.Error())
		res.Message = err.Error()
		return res, err
	}
	var ethIps []string

	if req.InterfaceName != "" {
		interfaces, err := innerGetNetworkInterface()
		if err != nil {
			res.Message = err.Error()
			return res, err
		}
		for _, i := range interfaces.NetworkInterfaces {
			if i.Name != req.InterfaceName {
				continue
			}
			//没有不包含/的，一定能Split出两个
			for _, v4 := range i.Ipv4 {
				ethIps = append(ethIps, strings.Split(v4, "/")[0])
			}
			for _, v6 := range i.Ipv6 {
				ethIps = append(ethIps, strings.Split(v6, "/")[0])
			}
		}
		if len(ethIps) == 0 {
			res.Message = "no such interface"
			utils.LogWarn(res.Message)
			return res, nil

		}
	}

	uniqueLAddr := make(map[addrAndPort]*networkBindAndPid)

	for _, b := range bindList {
		if ethIps != nil {
			if !slices.Contains(ethIps, b.Laddr.IP) {
				continue
			}
		}
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
		for _, p := range agent_runtime.GetProcesses() {
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

	utils.LogInfo(fmt.Sprintf("GetNetworkBindList: %d", len(res.NetworkBinds)))
	if len(res.NetworkBinds) == 0 {
		res.Message = "This interface has no service bound"
	}
	return res, nil
}
