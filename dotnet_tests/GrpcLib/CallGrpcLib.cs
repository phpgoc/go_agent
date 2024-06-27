using System.Net.NetworkInformation;
using AgentProto;
using Grpc.Net.Client;

namespace GrpcLib
{
    public class CallGrpcLib
    {
        private GrpcChannel channel;

        public CallGrpcLib(GrpcChannel channel)
        {
            this.channel = channel;
        }

        public string Echo(string name)
        {
            var client = new Greeter.GreeterClient(channel);
            var reply = client.SayHello(new HelloRequest { Name = name });
            return reply.Message;
        }
        
        public GetApacheInfoResponse GetApacheInfo()
        {
            var client = new GetApacheInfo.GetApacheInfoClient(channel);
            var reply = client.GetApacheInfo(new GetApacheInfoRequest { });
            return reply;
        }

        public GetSysInfoResponse GetSysInfo()
        {
            var client = new GetSysInfo.GetSysInfoClient(channel);
            var reply = client.GetSysInfo(new GetSysInfoRequest { });
            return reply;
        }
        
        public UserListResponse GetUserList()
        {
            var client = new GetUserList.GetUserListClient(channel);
            var reply = client.GetUserList(new UserListRequest { });
            return reply;
        } 
        
        public NetworkInterfaceResponse GetNetworkInterface()
        {
            var client = new Network.NetworkClient(channel);
            var reply = client.GetNetworkInterface(new NetworkInterfaceRequest { });
            return reply;
        }
    }
}