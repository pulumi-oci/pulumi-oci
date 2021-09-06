// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Opsi
{
    public static class GetEnterpriseManagerBridge
    {
        /// <summary>
        /// This data source provides details about a specific Enterprise Manager Bridge resource in Oracle Cloud Infrastructure Opsi service.
        /// 
        /// Gets details of an Operations Insights Enterprise Manager bridge.
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
        ///         var testEnterpriseManagerBridge = Output.Create(Oci.Opsi.GetEnterpriseManagerBridge.InvokeAsync(new Oci.Opsi.GetEnterpriseManagerBridgeArgs
        ///         {
        ///             EnterpriseManagerBridgeId = oci_opsi_enterprise_manager_bridge.Test_enterprise_manager_bridge.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEnterpriseManagerBridgeResult> InvokeAsync(GetEnterpriseManagerBridgeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEnterpriseManagerBridgeResult>("oci:opsi/getEnterpriseManagerBridge:getEnterpriseManagerBridge", args ?? new GetEnterpriseManagerBridgeArgs(), options.WithVersion());
    }


    public sealed class GetEnterpriseManagerBridgeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique Enterprise Manager bridge identifier
        /// </summary>
        [Input("enterpriseManagerBridgeId", required: true)]
        public string EnterpriseManagerBridgeId { get; set; } = null!;

        public GetEnterpriseManagerBridgeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEnterpriseManagerBridgeResult
    {
        /// <summary>
        /// Compartment identifier of the Enterprise Manager bridge
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// Description of Enterprise Manager Bridge
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// User-friedly name of Enterprise Manager Bridge that does not have to be unique.
        /// </summary>
        public readonly string DisplayName;
        public readonly string EnterpriseManagerBridgeId;
        /// <summary>
        /// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// Enterprise Manager bridge identifier
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
        /// </summary>
        public readonly string LifecycleDetails;
        /// <summary>
        /// Object Storage Bucket Name
        /// </summary>
        public readonly string ObjectStorageBucketName;
        /// <summary>
        /// A message describing status of the object storage bucket of this resource. For example, it can be used to provide actionable information about the permission and content validity of the bucket.
        /// </summary>
        public readonly string ObjectStorageBucketStatusDetails;
        /// <summary>
        /// Object Storage Namespace Name
        /// </summary>
        public readonly string ObjectStorageNamespaceName;
        /// <summary>
        /// The current state of the Enterprise Manager bridge.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> SystemTags;
        /// <summary>
        /// The time the the Enterprise Manager bridge was first created. An RFC3339 formatted datetime string
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The time the Enterprise Manager bridge was updated. An RFC3339 formatted datetime string
        /// </summary>
        public readonly string TimeUpdated;

        [OutputConstructor]
        private GetEnterpriseManagerBridgeResult(
            string compartmentId,

            ImmutableDictionary<string, object> definedTags,

            string description,

            string displayName,

            string enterpriseManagerBridgeId,

            ImmutableDictionary<string, object> freeformTags,

            string id,

            string lifecycleDetails,

            string objectStorageBucketName,

            string objectStorageBucketStatusDetails,

            string objectStorageNamespaceName,

            string state,

            ImmutableDictionary<string, object> systemTags,

            string timeCreated,

            string timeUpdated)
        {
            CompartmentId = compartmentId;
            DefinedTags = definedTags;
            Description = description;
            DisplayName = displayName;
            EnterpriseManagerBridgeId = enterpriseManagerBridgeId;
            FreeformTags = freeformTags;
            Id = id;
            LifecycleDetails = lifecycleDetails;
            ObjectStorageBucketName = objectStorageBucketName;
            ObjectStorageBucketStatusDetails = objectStorageBucketStatusDetails;
            ObjectStorageNamespaceName = objectStorageNamespaceName;
            State = state;
            SystemTags = systemTags;
            TimeCreated = timeCreated;
            TimeUpdated = timeUpdated;
        }
    }
}