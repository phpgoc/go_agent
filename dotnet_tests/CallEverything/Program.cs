// See https://aka.ms/new-console-template for more information
using System.Threading.Tasks;
using AgentProto;
using CallEverything;


class Program
{
    static async Task Main(string[] args)
    {
        await Task.Delay(0); //当把所有await都注释时，会警告这个方法没有await，所以加上这个
         GrpcCaller caller = new GrpcCaller("http://localhost:50051");
//        GrpcCaller caller = new GrpcCaller("http://172.17.0.1:50051");
// GrpcCaller caller = new GrpcCaller("http://192.168.31.244:50051");

          // caller.CallGetApacheInfo();
         // caller.CallGetSysInfo();
         // caller.CallGetProcessList(true);
//         caller.CallGetUserList();
//         caller.CallGetShellHistory("yangdianqing");
//         caller.CallGetShellHistory("");
         caller.CallGetNetworkInterface();
         caller.CallGetAllNetworkConnect();
         caller.CallGetNetworkBindList(Protocol.All, "");
//         caller.CallGetNetworkBindList(Protocol.Tcp, "");
//         caller.CallGetNetworkBindList(Protocol.All, "lo");
//         caller.CallGetNetworkBindList(Protocol.All, "Loopback Pseudo-Interface 1");
//         await caller.CallFileDownload("/etc/hosts", "D:/hosts.txt");
         // caller.GetNginxInfo();
         // caller.MysqlDump();

    }
}

