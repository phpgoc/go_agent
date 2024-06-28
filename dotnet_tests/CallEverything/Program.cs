// See https://aka.ms/new-console-template for more information
using System.Threading.Tasks;
using CallEverything;


class Program
{
    static async Task Main(string[] args)
    {
        GrpcCaller caller = new GrpcCaller("http://localhost:50051");
// GrpcCaller caller = new GrpcCaller("http://192.168.31.244:50051");

        caller.CallHelloWorld();
        caller.CallGetApacheInfo();
        caller.CallGetSysInfo();
        caller.CallGetUserList();
        caller.CallGetNetworkInterface();
        await caller.CallFileDownload("/etc/hosts", "D:/hosts.txt");
    }
}

