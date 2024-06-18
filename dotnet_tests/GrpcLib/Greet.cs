using AgentProto;
using Grpc.Net.Client;

namespace GrpcLib;

public class Greet
{
    public static string Echo(GrpcChannel channel,string name)
    {
        var client = new Greeter.GreeterClient(channel);
        var reply = client.SayHello(new HelloRequest { Name = name });
        return reply.Message;
    }
}