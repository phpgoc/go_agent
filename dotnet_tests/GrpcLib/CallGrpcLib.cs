
using AgentProto;
using Grpc.Core;
using Grpc.Net.Client;
using File = System.IO.File;

namespace GrpcLib
{
    public class CallGrpcLib
    {
        private GrpcChannel channel;

        public CallGrpcLib(GrpcChannel channel)
        {
            this.channel = channel;
        }


        public GetApacheInfoResponse GetApacheInfo()
        {
            var client = new ApacheService.ApacheServiceClient(channel);
            var reply = client.GetApacheInfo(new GetApacheInfoRequest { });
            return reply;
        }

        public GetSysInfoResponse GetSysInfo()
        {

            var client = new SystemService.SystemServiceClient(channel);
            var reply = client.GetSysInfo(new GetSysInfoRequest { });
            return reply;
        }
        
        public UserListResponse GetUserList()
        {
            var client = new SystemService.SystemServiceClient(channel);
            var reply = client.GetUserList(new UserListRequest { });
            return reply;
        } 
        
        public GetShellHistoryResponse GetShellHistory(string name)
        {
            var client = new SystemService.SystemServiceClient(channel);
            var reply = client.GetShellHistory(new GetShellHistoryRequest
            {
                UserName = name
            });
            return reply;
        }
        
        public GetProcessListResponse GetProcessList(bool withThread = false)
        {
            var client = new SystemService.SystemServiceClient(channel);
            var reply = client.GetProcessList(new GetProcessListRequest
            {
                WithThreadTimes = withThread
            });
            return reply;
        }
        
        public GetNetworkInterfaceResponse GetNetworkInterface()
        {
            var client = new NetworkService.NetworkServiceClient(channel);
            var reply = client.GetNetworkInterface(new GetNetworkInterfaceRequest { });
            return reply;
        }

        public GetAllNetworkConnectResponse GetAllNetworkConnect()
        {
            var client = new NetworkService.NetworkServiceClient(channel);
            var reply = client.GetAllNetworkConnect(new GetAllNetworkConnectRequest { });
            return reply;
        }

        public GetNetworkBindListResponse GetNetworkBindList(Protocol protocol, string interfaceName)
        {
            var client = new NetworkService.NetworkServiceClient(channel);
            var reply = client.GetNetworkBindList(new GetNetworkBindListRequest { Protocol = protocol, InterfaceName = interfaceName });
           
            return reply;
        }
        
         public async Task  FileDownload(string remote , string local)
         {

             try
             {
                 //异常视乎只发生在write时，先创建stream，再写一个一个字符，如果正常再重新创建
                 // Get the directory part of the local path
                 string? directory = Path.GetDirectoryName(local);
                 

                 // Check if the directory exists
                 if (directory == null ||  !Directory.Exists(directory))
                 {
                     Console.WriteLine($"The directory {directory} does not exist.");
                     return;
                 }
                 
                 await using FileStream writeStream = File.Create(local);
                 var client = new FileService.FileServiceClient(channel);
                 using var call = client.DownloadFile(new DownloadFileRequest { Filename = remote });
                 await foreach (var res in call.ResponseStream.ReadAllAsync())
                 {
                     if (res.Chunk != null)
                     {
                         await writeStream.WriteAsync(res.Chunk.Memory);
                     }
                 }

                 writeStream.Close();
             }
             catch (Exception e)
             {
                 Console.WriteLine($"An error occurred: {e.Message}");
             }
        }
         
        public GetNginxInfoResponse GetNginxInfo()
        {
            var client = new NginxService.NginxServiceClient(channel);
            var reply = client.GetNginxInfo(new GetNginxInfoRequest { });
            return reply;
        }
        
        public MysqlDumpResponse MysqlDump(ConnectionInfo? connectionInfo, bool skip,bool force )
        {
            var client = new DatabaseService.DatabaseServiceClient(channel);
            var reply = client.MysqlDump(new MysqlDumpRequest
            {
                Force = force,
                SkipGrantTables = skip,
                ConnectionInfo = connectionInfo
            });
            return reply;
        }
    }
}