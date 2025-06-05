using System.Collections.Generic;
namespace JCLTrace.UI.Models;

public class JclJobStep
{
    public string Name { get; set; }
    public string Type { get; set; }
    public string Proc { get; set; }
    public List<string> Calls { get; set; } = new();
    public List<string> Outputs { get; set; } = new();
}