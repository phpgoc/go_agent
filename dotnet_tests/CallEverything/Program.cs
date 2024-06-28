// See https://aka.ms/new-console-template for more information
using System.Threading.Tasks;
using CallEverything;


class Program
{
    static async Task Main(string[] args)
    {
        await Task.Delay(0); //当把所有await都注释时，会警告这个方法没有await，所以加上这个
        GrpcCaller caller = new GrpcCaller("http://localhost:50051");
// GrpcCaller caller = new GrpcCaller("http://192.168.31.244:50051");
        
        // caller.CallHelloWorld();
        // caller.CallGetApacheInfo();
        // caller.CallGetSysInfo();
        caller.CallGetAllNetworkConnect();
        // caller.CallGetUserList();
        // caller.CallGetNetworkInterface();
        // await caller.CallFileDownload("/etc/hosts", "D:/hosts.txt");
    }
}

