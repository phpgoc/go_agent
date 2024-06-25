using AgentProto;
using Grpc.Net.Client;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

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

        public GetSysInfoResponse GetSysInfo()
        {
            var client = new GetSysInfo.GetSysInfoClient(channel);
            var reply = client.GetSysInfo(new GetSysInfoRequest { });
            return reply;
        }
    }
}