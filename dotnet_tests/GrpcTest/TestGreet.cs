using GrpcLib;
namespace GrpcTest;

[TestClass]
public class TestGreet
{
    [TestMethod]
    public void TestEcho()
    {
        // Arrange
        string name = "Alice";
        string expected = "Hello Alice";

        // Act
        string result = Greet.Echo(Utils.GetGlobalChannel(), name);

        // Assert
        Assert.AreEqual(expected, result);
    }
    
    [TestMethod]
    public void TestNoEqual()
    {
        // Arrange
        string name = "1";
        string expected = "Hello 2";

        // Act
        string result = Greet.Echo(Utils.GetGlobalChannel(), name);

        // Assert
        Assert.AreNotEqual(expected, result);
    }
}