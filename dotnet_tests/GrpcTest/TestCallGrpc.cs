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
        public void TestEcho()
        {
            // Arrange
            string name = "Alice";
            string expected = "Hello Alice";

            // Act
            string result = grpcLib.Echo(name);

            // Assert
            Assert.AreEqual(expected, result);
        }

        [TestMethod]
        public void TestIsLinux()
        {
            GetSysInfoResponse res = grpcLib.GetSysInfo();

            Assert.AreEqual(res.Caption, "ubuntu");
            Assert.AreEqual(res.Timezone, "+8");
        }
    }
}