// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Database
{
    public static class GetDbVersions
    {
        /// <summary>
        /// This data source provides the list of Db Versions in Oracle Cloud Infrastructure Database service.
        /// 
        /// Gets a list of supported Oracle Database versions.
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
        ///         var testDbVersions = Output.Create(Oci.Database.GetDbVersions.InvokeAsync(new Oci.Database.GetDbVersionsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             DbSystemId = oci_database_db_system.Test_db_system.Id,
        ///             DbSystemShape = @var.Db_version_db_system_shape,
        ///             IsDatabaseSoftwareImageSupported = @var.Db_version_is_database_software_image_supported,
        ///             IsUpgradeSupported = @var.Db_version_is_upgrade_supported,
        ///             StorageManagement = @var.Db_version_storage_management,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDbVersionsResult> InvokeAsync(GetDbVersionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDbVersionsResult>("oci:database/getDbVersions:getDbVersions", args ?? new GetDbVersionsArgs(), options.WithVersion());
    }


    public sealed class GetDbVersionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). If provided, filters the results to the set of database versions which are supported for the DB system.
        /// </summary>
        [Input("dbSystemId")]
        public string? DbSystemId { get; set; }

        /// <summary>
        /// If provided, filters the results to the set of database versions which are supported for the given shape.
        /// </summary>
        [Input("dbSystemShape")]
        public string? DbSystemShape { get; set; }

        [Input("filters")]
        private List<Inputs.GetDbVersionsFilterArgs>? _filters;
        public List<Inputs.GetDbVersionsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDbVersionsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// If true, filters the results to the set of Oracle Database versions that are supported for Oracle Cloud Infrastructure database software images.
        /// </summary>
        [Input("isDatabaseSoftwareImageSupported")]
        public bool? IsDatabaseSoftwareImageSupported { get; set; }

        /// <summary>
        /// If provided, filters the results to the set of database versions which are supported for Upgrade.
        /// </summary>
        [Input("isUpgradeSupported")]
        public bool? IsUpgradeSupported { get; set; }

        /// <summary>
        /// The DB system storage management option. Used to list database versions available for that storage manager. Valid values are:
        /// * ASM - Automatic storage management
        /// * LVM - Logical volume management
        /// </summary>
        [Input("storageManagement")]
        public string? StorageManagement { get; set; }

        public GetDbVersionsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDbVersionsResult
    {
        public readonly string CompartmentId;
        public readonly string? DbSystemId;
        public readonly string? DbSystemShape;
        /// <summary>
        /// The list of db_versions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDbVersionsDbVersionResult> DbVersions;
        public readonly ImmutableArray<Outputs.GetDbVersionsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IsDatabaseSoftwareImageSupported;
        /// <summary>
        /// True if this version of the Oracle Database software is supported for Upgrade.
        /// </summary>
        public readonly bool? IsUpgradeSupported;
        public readonly string? StorageManagement;

        [OutputConstructor]
        private GetDbVersionsResult(
            string compartmentId,

            string? dbSystemId,

            string? dbSystemShape,

            ImmutableArray<Outputs.GetDbVersionsDbVersionResult> dbVersions,

            ImmutableArray<Outputs.GetDbVersionsFilterResult> filters,

            string id,

            bool? isDatabaseSoftwareImageSupported,

            bool? isUpgradeSupported,

            string? storageManagement)
        {
            CompartmentId = compartmentId;
            DbSystemId = dbSystemId;
            DbSystemShape = dbSystemShape;
            DbVersions = dbVersions;
            Filters = filters;
            Id = id;
            IsDatabaseSoftwareImageSupported = isDatabaseSoftwareImageSupported;
            IsUpgradeSupported = isUpgradeSupported;
            StorageManagement = storageManagement;
        }
    }
}