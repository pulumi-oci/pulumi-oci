// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetCoreInstancePools
    {
        /// <summary>
        /// This data source provides the list of Instance Pools in Oracle Cloud Infrastructure Core service.
        /// 
        /// Lists the instance pools in the specified compartment.
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
        ///         var testInstancePools = Output.Create(Oci.GetCoreInstancePools.InvokeAsync(new Oci.GetCoreInstancePoolsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             DisplayName = @var.Instance_pool_display_name,
        ///             State = @var.Instance_pool_state,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCoreInstancePoolsResult> InvokeAsync(GetCoreInstancePoolsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCoreInstancePoolsResult>("oci:index/getCoreInstancePools:GetCoreInstancePools", args ?? new GetCoreInstancePoolsArgs(), options.WithVersion());
    }


    public sealed class GetCoreInstancePoolsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// A filter to return only resources that match the given display name exactly.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("filters")]
        private List<Inputs.GetCoreInstancePoolsFilterArgs>? _filters;
        public List<Inputs.GetCoreInstancePoolsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetCoreInstancePoolsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// A filter to only return resources that match the given lifecycle state. The state value is case-insensitive.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetCoreInstancePoolsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCoreInstancePoolsResult
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance pool.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// The display name of the VNIC. This is also use to match against the instance configuration defined secondary VNIC.
        /// </summary>
        public readonly string? DisplayName;
        public readonly ImmutableArray<Outputs.GetCoreInstancePoolsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of instance_pools.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCoreInstancePoolsInstancePoolResult> InstancePools;
        /// <summary>
        /// The current state of the instance pool.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private GetCoreInstancePoolsResult(
            string compartmentId,

            string? displayName,

            ImmutableArray<Outputs.GetCoreInstancePoolsFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetCoreInstancePoolsInstancePoolResult> instancePools,

            string? state)
        {
            CompartmentId = compartmentId;
            DisplayName = displayName;
            Filters = filters;
            Id = id;
            InstancePools = instancePools;
            State = state;
        }
    }
}