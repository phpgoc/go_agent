// See https://aka.ms/new-console-template for more information

using AgentProto;
using Grpc.Net.Client;
using GrpcLib;
using Newtonsoft.Json;

Console.WriteLine("Hello, World!");

// GrpcChannel channel = GrpcChannel.ForAddress ("http://192.168.31.244:50051");
GrpcChannel channel = GrpcChannel.ForAddress ("http://localhost:50051");
CallGrpcLib callGrpcLib = new CallGrpcLib (channel);
var jsonBeautifierSetting = new JsonSerializerSettings
{
    Formatting = Formatting.Indented,
    
};
void PrintJson(object obj)
{
    Console.WriteLine(JsonConvert.SerializeObject(obj, jsonBeautifierSetting));
}

// call helloworld
Console.WriteLine("call helloworld");
string greetResult = callGrpcLib.Echo( "Alice");
PrintJson(greetResult);

// call getApacheInfo
Console.WriteLine("call getApacheInfo");
GetApacheInfoResponse apacheInfo = callGrpcLib.GetApacheInfo();
PrintJson(apacheInfo);

//call getSysInfo
Console.WriteLine("call getSysInfo");
GetSysInfoResponse sysInfo = callGrpcLib.GetSysInfo();
PrintJson(sysInfo);

//call getUserList
Console.WriteLine("call getUserList");
UserListResponse userList = callGrpcLib.GetUserList();
PrintJson(userList);