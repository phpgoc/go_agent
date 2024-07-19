using AgentProto;
using GrpcLib;


namespace GrpcTest
{
    [TestClass]
    public class TestCallGrpc
    {
        private CallGrpcLib grpcLib;

        TestCallGrpc()
        {
            this.grpcLib = new CallGrpcLib(Utils.GetGlobalChannel());
        }



        [TestMethod]
        public void TestIsLinux()
        {
            GetSysInfoResponse res = grpcLib.GetSysInfo();

            Assert.AreEqual(res.Timezone, "+8");
        }
    }
}