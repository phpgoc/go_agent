package network

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/v4/net"
	pb "go-agent/agent_proto"
	"go-agent/utils"
	"syscall"
)

func (s *Server) GetAllNetworkConnect(_ context.Context, _ *pb.GetAllNetworkConnectRequest) (*pb.GetAllNetworkConnectResponse, error) {
	utils.LogInfo("called GetAllNetworkConnect")
	res := &pb.GetAllNetworkConnectResponse{}
	connects, err := net.Connections("all")

	if err != nil {
		return utils.SetResponseErrorAndLogMessageGeneric(res, err.Error(), pb.ResponseCode_UNKNOWN_SERVER_ERROR)
	}
	connectNumForLog := 0

	for _, connect := range connects {
		//默认不要unix socket
		if connect.Family == syscall.AF_UNIX {
			continue
		}
		connectNumForLog++
		thisOne := &pb.NetworkConnect{
			Family: family2string(connect.Family),
			Type:   socketType2string(connect.Type),
			Local: &pb.Address{
				Ip:   connect.Laddr.IP,
				Port: connect.Laddr.Port,
			},
			Remote: &pb.Address{
				Ip:   connect.Raddr.IP,
				Port: connect.Raddr.Port,
			},
			Status: connect.Status, // Status应该只在tcp连接中有
			Pid:    connect.Pid,
		}
		res.NetworkConnects = append(res.NetworkConnects, thisOne)

	}

	utils.LogInfo(fmt.Sprintf("connections nums: %d", connectNumForLog))

	return res, nil
}
