// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetCoreIpsecConnectionTunnels
    {
        /// <summary>
        /// This data source provides the list of Ip Sec Connection Tunnels in Oracle Cloud Infrastructure Core service.
        /// 
        /// Lists the tunnel information for the specified IPSec connection.
        /// 
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
        ///         var testIpSecConnectionTunnels = Output.Create(Oci.GetCoreIpsecConnectionTunnels.InvokeAsync(new Oci.GetCoreIpsecConnectionTunnelsArgs
        ///         {
        ///             IpsecId = oci_core_ipsec.Test_ipsec.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCoreIpsecConnectionTunnelsResult> InvokeAsync(GetCoreIpsecConnectionTunnelsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCoreIpsecConnectionTunnelsResult>("oci:index/getCoreIpsecConnectionTunnels:GetCoreIpsecConnectionTunnels", args ?? new GetCoreIpsecConnectionTunnelsArgs(), options.WithVersion());
    }


    public sealed class GetCoreIpsecConnectionTunnelsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetCoreIpsecConnectionTunnelsFilterArgs>? _filters;
        public List<Inputs.GetCoreIpsecConnectionTunnelsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetCoreIpsecConnectionTunnelsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IPSec connection.
        /// </summary>
        [Input("ipsecId", required: true)]
        public string IpsecId { get; set; } = null!;

        public GetCoreIpsecConnectionTunnelsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCoreIpsecConnectionTunnelsResult
    {
        public readonly ImmutableArray<Outputs.GetCoreIpsecConnectionTunnelsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of ip_sec_connection_tunnels.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCoreIpsecConnectionTunnelsIpSecConnectionTunnelResult> IpSecConnectionTunnels;
        public readonly string IpsecId;

        [OutputConstructor]
        private GetCoreIpsecConnectionTunnelsResult(
            ImmutableArray<Outputs.GetCoreIpsecConnectionTunnelsFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetCoreIpsecConnectionTunnelsIpSecConnectionTunnelResult> ipSecConnectionTunnels,

            string ipsecId)
        {
            Filters = filters;
            Id = id;
            IpSecConnectionTunnels = ipSecConnectionTunnels;
            IpsecId = ipsecId;
        }
    }
}