using System.Collections.Generic;
namespace JCLTrace.UI.Models;

public class JobNode
{
    public string JobName { get; set; }
    public List<JobStep> Steps { get; set; } = new();
    public List<string> Calls { get; set; } = new();
}