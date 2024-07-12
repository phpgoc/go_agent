using AgentProto.Database;

namespace CallEverything;

using AgentProto;
using Grpc.Net.Client;
using GrpcLib;
using Newtonsoft.Json;
using System;

public class GrpcCaller
{
    private GrpcChannel _channel;
    private CallGrpcLib _callGrpcLib;
    private JsonSerializerSettings _jsonBeautifierSetting;

    public GrpcCaller(string address)
    {
        _channel = GrpcChannel.ForAddress(address);
        _callGrpcLib = new CallGrpcLib(_channel);
        _jsonBeautifierSetting = new JsonSerializerSettings
        {
            Formatting = Formatting.Indented,
        };
    }

    private void PrintJson(object obj)
    {
        Console.WriteLine(JsonConvert.SerializeObject(obj, _jsonBeautifierSetting));
    }

    public void CallHelloWorld()
    {
        Console.WriteLine("call helloworld");
        string greetResult = _callGrpcLib.Echo("Alice");
        PrintJson(greetResult);
    }

    public void CallGetApacheInfo()
    {
        Console.WriteLine("call getApacheInfo");
        GetApacheInfoResponse apacheInfo = _callGrpcLib.GetApacheInfo();
        PrintJson(apacheInfo);
    }

    public void CallGetSysInfo()
    {
        Console.WriteLine("call getSysInfo");
        GetSysInfoResponse sysInfo = _callGrpcLib.GetSysInfo();
        PrintJson(sysInfo);
    }

    public void CallGetUserList()
    {
        Console.WriteLine("call getUserList");
        UserListResponse userList = _callGrpcLib.GetUserList();
        PrintJson(userList);
    }
    
    public void CallGetShellHistory(string name)
    {
        Console.WriteLine("call getShellHistory");
        GetShellHistoryResponse shellHistory = _callGrpcLib.GetShellHistory(name);
        PrintJson(shellHistory);
    }
    
    public void CallGetProcessList(bool withThread = false)
    {
        Console.WriteLine("call getProcessList");
        GetProcessListResponse processList = _callGrpcLib.GetProcessList(withThread);
        PrintJson(processList);
    }
    
    
    public void CallGetNetworkInterface()
    {
        Console.WriteLine("call getNetworkInterface");
        GetNetworkInterfaceResponse networkInterfaceList = _callGrpcLib.GetNetworkInterface();
        PrintJson(networkInterfaceList);
    }
    
    public void CallGetAllNetworkConnect()
    {
        Console.WriteLine("call getAllNetworkConnect");
        GetAllNetworkConnectResponse allNetworkConnect = _callGrpcLib.GetAllNetworkConnect();
        PrintJson(allNetworkConnect);
    }
    
    public void CallGetNetworkBindList(Protocol protocol, string interfaceName)
    {
        Console.WriteLine("call getNetworkBindList");
        GetNetworkBindListResponse networkBindList = _callGrpcLib.GetNetworkBindList(protocol, interfaceName);
        PrintJson(networkBindList);
    }
    
    
    public async Task CallFileDownload(string remote, string local)
    {
        Console.WriteLine("call fileDownload");
        await _callGrpcLib.FileDownload(remote, local);
    }
    
    public void GetNginxInfo()
    {
        Console.WriteLine("call getNginxInfo");
        GetNginxInfoResponse nginxInfo = _callGrpcLib.GetNginxInfo();
        PrintJson(nginxInfo);
    }

    public void MysqlDump()
    {
        Console.WriteLine("call mysqlDump");
        ConnectionInfo? connectionInfo = new ConnectionInfo
        {
            Host = "localhost",
            Port = 3306,
            Username = "root",
            Password = "123456",
        };
        // connectionInfo = null;
        MysqlDumpResponse mysqlDump = _callGrpcLib.MysqlDump(connectionInfo, true,false);
        PrintJson(mysqlDump);
    }
}