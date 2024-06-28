package network

import (
	pb "go-agent/agent_proto"
	"syscall"
)

type Server struct {
	pb.UnimplementedNetworkServiceServer
}

type connectType uint32

func socketType2string(num uint32) string {
	switch connectType(num) {
	case syscall.SOCK_STREAM:
		return "tcp"
	case syscall.SOCK_DGRAM:
		return "udp"

	default:
		return "Unknown"
	}
}

func family2string(family uint32) string {
	switch family {
	case syscall.AF_UNIX:
		return "unix"
	case syscall.AF_INET6:
		return "v6"
	case syscall.AF_INET:
		return "v4"
	default:
		return "Unknown"

	}
}
