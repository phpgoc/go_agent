// See https://aka.ms/new-console-template for more information

using AgentProto;
using Grpc.Net.Client;
using GrpcLib;

Console.WriteLine("Hello, World!");

GrpcChannel channel = GrpcChannel.ForAddress ("http://192.168.31.244:50051");
CallGrpcLib callGrpcLib = new CallGrpcLib (channel);

// call helloworld
string greetResult = callGrpcLib.Echo( "Alice");
Console.WriteLine(greetResult);

//call getSysInfo
GetSysInfoResponse sysInfo = callGrpcLib.GetSysInfo();
Console.WriteLine(sysInfo);