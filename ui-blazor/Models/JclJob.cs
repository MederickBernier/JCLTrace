using System.Collections.Generic;
namespace JCLTrace.UI.Models;

public class JclJob
{
    public string Name { get; set; }
    public string File { get; set; }
    public List<JclJobStep> Steps { get; set; } = new();
}