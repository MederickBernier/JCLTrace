using JCLTrace.UI.Models;
using System.Collections.Generic;

namespace JCLTrace.UI.Models;

public class JclJobMap
{
    public List<JobNode> Jobs { get; set; } = new();
}