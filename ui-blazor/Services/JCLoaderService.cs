using System.Text.Json;
using JCLTrace.UI.Models;

namespace JCLTrace.UI.Services;

public class JCLoaderService
{
    private readonly IWebHostEnvironment _env;

    public JCLoaderService(IWebHostEnvironment env)
    {
        _env = env;
    }

    public async Task<JclJobMap?> LoadMockJobMap()
    {
        var fileProvider = _env.WebRootFileProvider;
        var fileInfo = fileProvider.GetFileInfo("mock/job_map.json");

        if (!fileInfo.Exists)
        {
            Console.WriteLine("⚠️ job_map.json not found at runtime.");
            return null;
        }

        try
        {
            using var stream = fileInfo.CreateReadStream();
            using var reader = new StreamReader(stream);
            var json = await reader.ReadToEndAsync();

            var result = JsonSerializer.Deserialize<JclJobMap>(json, new JsonSerializerOptions
            {
                PropertyNameCaseInsensitive = true
            });

            Console.WriteLine($"✅ Loaded job_map.json with {result?.Jobs.Count ?? 0} jobs.");
            return result;
        }
        catch (Exception ex)
        {
            Console.WriteLine($"❌ Failed to load JSON: {ex.Message}");
            return null;
        }
    }
}