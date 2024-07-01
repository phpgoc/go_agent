package network

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/v4/net"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"strings"
)

func (s *Server) GetNetworkInterface(_ context.Context, _ *pb.GetNetworkInterfaceRequest) (*pb.GetNetworkInterfaceResponse, error) {
	utils.LogInfo("called GetNetworkInterface")
	res, err := innerGetNetworkInterface()
	if err != nil {
		res.Message = err.Error()
		return res, err
	}
	utils.LogInfo(fmt.Sprintf("get network interface len: %d", len(res.NetworkInterfaces)))
	return res, nil

}

func innerGetNetworkInterface() (*pb.GetNetworkInterfaceResponse, error) {
	res := &pb.GetNetworkInterfaceResponse{}
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
