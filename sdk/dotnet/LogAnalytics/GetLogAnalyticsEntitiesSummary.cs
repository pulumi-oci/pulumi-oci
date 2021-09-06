// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.LogAnalytics
{
    public static class GetLogAnalyticsEntitiesSummary
    {
        /// <summary>
        /// This data source provides details about a specific Log Analytics Entities Summary resource in Oracle Cloud Infrastructure Log Analytics service.
        /// 
        /// Returns log analytics entities count summary report.
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
        ///         var testLogAnalyticsEntitiesSummary = Output.Create(Oci.LogAnalytics.GetLogAnalyticsEntitiesSummary.InvokeAsync(new Oci.LogAnalytics.GetLogAnalyticsEntitiesSummaryArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             Namespace = @var.Log_analytics_entities_summary_namespace,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLogAnalyticsEntitiesSummaryResult> InvokeAsync(GetLogAnalyticsEntitiesSummaryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLogAnalyticsEntitiesSummaryResult>("oci:loganalytics/getLogAnalyticsEntitiesSummary:getLogAnalyticsEntitiesSummary", args ?? new GetLogAnalyticsEntitiesSummaryArgs(), options.WithVersion());
    }


    public sealed class GetLogAnalyticsEntitiesSummaryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the compartment in which to list resources.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// The Logging Analytics namespace used for the request.
        /// </summary>
        [Input("namespace", required: true)]
        public string Namespace { get; set; } = null!;

        public GetLogAnalyticsEntitiesSummaryArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLogAnalyticsEntitiesSummaryResult
    {
        /// <summary>
        /// Total number of ACTIVE entities
        /// </summary>
        public readonly int ActiveEntitiesCount;
        /// <summary>
        /// Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// Entities with log collection enabled
        /// </summary>
        public readonly int EntitiesWithHasLogsCollectedCount;
        /// <summary>
        /// Entities with management agent
        /// </summary>
        public readonly int EntitiesWithManagementAgentCount;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Namespace;

        [OutputConstructor]
        private GetLogAnalyticsEntitiesSummaryResult(
            int activeEntitiesCount,

            string compartmentId,

            int entitiesWithHasLogsCollectedCount,

            int entitiesWithManagementAgentCount,

            string id,

            string @namespace)
        {
            ActiveEntitiesCount = activeEntitiesCount;
            CompartmentId = compartmentId;
            EntitiesWithHasLogsCollectedCount = entitiesWithHasLogsCollectedCount;
            EntitiesWithManagementAgentCount = entitiesWithManagementAgentCount;
            Id = id;
            Namespace = @namespace;
        }
    }
}