using AgentProto;
using Grpc.Net.Client;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace GrpcLib
{
    public class GetSysInfoLib
    {
        public static GetSysInfoResponse GetSysInfoCSharp(GrpcChannel channel)
        {
            var client = new GetSysInfo.GetSysInfoClient(channel);
            var reply = client.GetSysInfo(new GetSysInfoRequest { });
            return reply;
        }
    }
}
