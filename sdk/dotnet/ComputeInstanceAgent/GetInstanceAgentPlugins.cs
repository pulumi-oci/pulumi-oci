// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.ComputeInstanceAgent
{
    public static class GetInstanceAgentPlugins
    {
        /// <summary>
        /// This data source provides the list of Instance Agent Plugins in Oracle Cloud Infrastructure Compute Instance Agent service.
        /// 
        /// The API to get one or more plugin information.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Oci = Pulumi.Oci;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var testInstanceAgentPlugins = Output.Create(Oci.ComputeInstanceAgent.GetInstanceAgentPlugins.InvokeAsync(new Oci.ComputeInstanceAgent.GetInstanceAgentPluginsArgs
        ///         {
        ///             InstanceagentId = oci_computeinstanceagent_instanceagent.Test_instanceagent.Id,
        ///             Name = @var.Instance_agent_plugin_name,
        ///             Status = @var.Instance_agent_plugin_status,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceAgentPluginsResult> InvokeAsync(GetInstanceAgentPluginsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceAgentPluginsResult>("oci:computeinstanceagent/getInstanceAgentPlugins:getInstanceAgentPlugins", args ?? new GetInstanceAgentPluginsArgs(), options.WithVersion());
    }


    public sealed class GetInstanceAgentPluginsArgs : Pulumi.InvokeArgs
    {
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetInstanceAgentPluginsFilterArgs>? _filters;
        public List<Inputs.GetInstanceAgentPluginsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetInstanceAgentPluginsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The OCID of the instance.
        /// </summary>
        [Input("instanceagentId", required: true)]
        public string InstanceagentId { get; set; } = null!;

        /// <summary>
        /// The plugin name
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The plugin status
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetInstanceAgentPluginsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceAgentPluginsResult
    {
        public readonly string CompartmentId;
        public readonly ImmutableArray<Outputs.GetInstanceAgentPluginsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of instance_agent_plugins.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceAgentPluginsInstanceAgentPluginResult> InstanceAgentPlugins;
        public readonly string InstanceagentId;
        /// <summary>
        /// The plugin name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The plugin status Specified the plugin state on the instance * `RUNNING` - The plugin is in running state * `STOPPED` - The plugin is in stopped state * `NOT_SUPPORTED` - The plugin is not supported on this platform * `INVALID` - The plugin state is not recognizable by the service
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetInstanceAgentPluginsResult(
            string compartmentId,

            ImmutableArray<Outputs.GetInstanceAgentPluginsFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetInstanceAgentPluginsInstanceAgentPluginResult> instanceAgentPlugins,

            string instanceagentId,

            string? name,

            string? status)
        {
            CompartmentId = compartmentId;
            Filters = filters;
            Id = id;
            InstanceAgentPlugins = instanceAgentPlugins;
            InstanceagentId = instanceagentId;
            Name = name;
            Status = status;
        }
    }
}