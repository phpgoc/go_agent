package network

import (
	"context"
	"github.com/shirou/gopsutil/v4/net"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"strings"
)

type Server struct {
	pb.UnimplementedNetworkServer
}

func (s *Server) GetNetworkInterface(_ context.Context, _ *pb.NetworkInterfaceRequest) (*pb.NetworkInterfaceResponse, error) {
	res := &pb.NetworkInterfaceResponse{}
	interfaces, err := net.Interfaces()
	if err != nil {
		utils.LogError(err.Error())
		return nil, err
	}

	for _, i := range interfaces {
		//ipv4 or ipv6
		thisOne := &pb.NetworkInterface{
			Name:  i.Name,
			Flags: strings.Join(i.Flags, "|"),
		}

		for _, addr := range i.Addrs {
			//通过判断是否有:来判断是ipv4还是ipv6
			if strings.Contains(addr.Addr, ":") {
				thisOne.Ipv6 = append(thisOne.Ipv6, addr.Addr)
			} else {
				thisOne.Ipv4 = append(thisOne.Ipv4, addr.Addr)
			}
		}
		res.NetworkInterfaces = append(res.NetworkInterfaces, thisOne)
	}

	return res, nil
}
