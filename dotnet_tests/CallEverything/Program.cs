// See https://aka.ms/new-console-template for more information

using CallEverything;

GrpcCaller caller = new GrpcCaller("http://localhost:50051");
// GrpcCaller caller = new GrpcCaller("http://192.168.31.244:50051");

// caller.CallHelloWorld();
// caller.CallGetApacheInfo();
caller.CallGetSysInfo();
// caller.CallGetUserList();