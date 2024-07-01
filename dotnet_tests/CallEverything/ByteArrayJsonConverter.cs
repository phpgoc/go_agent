using Newtonsoft.Json;

namespace CallEverything;

public class ByteArrayJsonConverter : JsonConverter<byte[]>
{
    public override byte[] ReadJson(JsonReader reader, Type objectType, byte[] existingValue, bool hasExistingValue, JsonSerializer serializer)
    {
        string s = (string)reader.Value;
        return Enumerable.Range(0, s.Length)
            .Where(x => x % 2 == 0)
            .Select(x => Convert.ToByte(s.Substring(x, 2), 16))
            .ToArray();
    }

    public override void WriteJson(JsonWriter writer, byte[] value, JsonSerializer serializer)
    {
        writer.WriteValue(BitConverter.ToString(value).Replace("-", ""));
    }
}