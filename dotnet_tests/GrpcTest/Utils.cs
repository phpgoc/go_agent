using Grpc.Net.Client;

namespace GrpcTest;

// AppContext.SetSwitch("System.Net.Http.SocketsHttpHandler.Http2UnencryptedSupport", true);

public class Utils
{

    private static readonly Lazy<GrpcChannel> _channel = new Lazy<GrpcChannel>(GrpcChannel.ForAddress("http://172.17.0.1:50051"));
    public static GrpcChannel GetGlobalChannel()
    {
        return _channel.Value;
    }
}