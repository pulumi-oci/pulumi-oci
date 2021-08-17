// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetMeteringComputationConfiguration
    {
        /// <summary>
        /// This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure Metering Computation service.
        /// 
        /// Returns the configurations list for the UI drop-down list.
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
        ///         var testConfiguration = Output.Create(Oci.GetMeteringComputationConfiguration.InvokeAsync(new Oci.GetMeteringComputationConfigurationArgs
        ///         {
        ///             TenantId = oci_metering_computation_tenant.Test_tenant.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMeteringComputationConfigurationResult> InvokeAsync(GetMeteringComputationConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMeteringComputationConfigurationResult>("oci:index/getMeteringComputationConfiguration:GetMeteringComputationConfiguration", args ?? new GetMeteringComputationConfigurationArgs(), options.WithVersion());
    }


    public sealed class GetMeteringComputationConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// tenant id
        /// </summary>
        [Input("tenantId", required: true)]
        public string TenantId { get; set; } = null!;

        public GetMeteringComputationConfigurationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMeteringComputationConfigurationResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of available configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMeteringComputationConfigurationItemResult> Items;
        public readonly string TenantId;

        [OutputConstructor]
        private GetMeteringComputationConfigurationResult(
            string id,

            ImmutableArray<Outputs.GetMeteringComputationConfigurationItemResult> items,

            string tenantId)
        {
            Id = id;
            Items = items;
            TenantId = tenantId;
        }
    }
}