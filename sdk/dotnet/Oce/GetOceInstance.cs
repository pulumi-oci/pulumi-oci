// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Oce
{
    public static class GetOceInstance
    {
        /// <summary>
        /// This data source provides details about a specific Oce Instance resource in Oracle Cloud Infrastructure Content and Experience service.
        /// 
        /// Gets a OceInstance by identifier
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
        ///         var testOceInstance = Output.Create(Oci.Oce.GetOceInstance.InvokeAsync(new Oci.Oce.GetOceInstanceArgs
        ///         {
        ///             OceInstanceId = oci_oce_oce_instance.Test_oce_instance.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOceInstanceResult> InvokeAsync(GetOceInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOceInstanceResult>("oci:oce/getOceInstance:getOceInstance", args ?? new GetOceInstanceArgs(), options.WithVersion());
    }


    public sealed class GetOceInstanceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// unique OceInstance identifier
        /// </summary>
        [Input("oceInstanceId", required: true)]
        public string OceInstanceId { get; set; } = null!;

        public GetOceInstanceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOceInstanceResult
    {
        /// <summary>
        /// Admin Email for Notification
        /// </summary>
        public readonly string AdminEmail;
        /// <summary>
        /// Compartment Identifier
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// OceInstance description, can be updated
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// Unique GUID identifier that is immutable on creation
        /// </summary>
        public readonly string Guid;
        /// <summary>
        /// Unique identifier that is immutable on creation
        /// </summary>
        public readonly string Id;
        public readonly string IdcsAccessToken;
        /// <summary>
        /// IDCS Tenancy Identifier
        /// </summary>
        public readonly string IdcsTenancy;
        /// <summary>
        /// Flag indicating whether the instance access is private or public
        /// </summary>
        public readonly string InstanceAccessType;
        /// <summary>
        /// Flag indicating whether the instance license is new cloud or bring your own license
        /// </summary>
        public readonly string InstanceLicenseType;
        /// <summary>
        /// Instance type based on its usage
        /// </summary>
        public readonly string InstanceUsageType;
        /// <summary>
        /// OceInstance Name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Object Storage Namespace of tenancy
        /// </summary>
        public readonly string ObjectStorageNamespace;
        public readonly string OceInstanceId;
        /// <summary>
        /// SERVICE data. Example: `{"service": {"IDCS": "value"}}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> Service;
        /// <summary>
        /// The current state of the file system.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
        /// </summary>
        public readonly string StateMessage;
        /// <summary>
        /// Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> SystemTags;
        /// <summary>
        /// Tenancy Identifier
        /// </summary>
        public readonly string TenancyId;
        /// <summary>
        /// Tenancy Name
        /// </summary>
        public readonly string TenancyName;
        /// <summary>
        /// The time the the OceInstance was created. An RFC3339 formatted datetime string
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The time the OceInstance was updated. An RFC3339 formatted datetime string
        /// </summary>
        public readonly string TimeUpdated;
        /// <summary>
        /// Upgrade schedule type representing service to be upgraded immediately whenever latest version is released or delay upgrade of the service to previous released version
        /// </summary>
        public readonly string UpgradeSchedule;
        /// <summary>
        /// Web Application Firewall(WAF) primary domain
        /// </summary>
        public readonly string WafPrimaryDomain;

        [OutputConstructor]
        private GetOceInstanceResult(
            string adminEmail,

            string compartmentId,

            ImmutableDictionary<string, object> definedTags,

            string description,

            ImmutableDictionary<string, object> freeformTags,

            string guid,

            string id,

            string idcsAccessToken,

            string idcsTenancy,

            string instanceAccessType,

            string instanceLicenseType,

            string instanceUsageType,

            string name,

            string objectStorageNamespace,

            string oceInstanceId,

            ImmutableDictionary<string, object> service,

            string state,

            string stateMessage,

            ImmutableDictionary<string, object> systemTags,

            string tenancyId,

            string tenancyName,

            string timeCreated,

            string timeUpdated,

            string upgradeSchedule,

            string wafPrimaryDomain)
        {
            AdminEmail = adminEmail;
            CompartmentId = compartmentId;
            DefinedTags = definedTags;
            Description = description;
            FreeformTags = freeformTags;
            Guid = guid;
            Id = id;
            IdcsAccessToken = idcsAccessToken;
            IdcsTenancy = idcsTenancy;
            InstanceAccessType = instanceAccessType;
            InstanceLicenseType = instanceLicenseType;
            InstanceUsageType = instanceUsageType;
            Name = name;
            ObjectStorageNamespace = objectStorageNamespace;
            OceInstanceId = oceInstanceId;
            Service = service;
            State = state;
            StateMessage = stateMessage;
            SystemTags = systemTags;
            TenancyId = tenancyId;
            TenancyName = tenancyName;
            TimeCreated = timeCreated;
            TimeUpdated = timeUpdated;
            UpgradeSchedule = upgradeSchedule;
            WafPrimaryDomain = wafPrimaryDomain;
        }
    }
}