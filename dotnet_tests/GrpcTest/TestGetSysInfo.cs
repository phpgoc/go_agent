using AgentProto;
using GrpcLib;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace GrpcTest
{

    [TestClass]
    public class TestGetSysInfo
    {


        [TestMethod]
        public void TestIsLinux()
        {
            GetSysInfoResponse res = GetSysInfoLib.GetSysInfoCSharp(Utils.GetGlobalChannel());

            Assert.AreEqual(res.Caption, "ubuntu");
            Assert.AreEqual(res.Timezone, "+8");
        }
    }
}
