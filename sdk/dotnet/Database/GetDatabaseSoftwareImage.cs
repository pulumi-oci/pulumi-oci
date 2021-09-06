// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Database
{
    public static class GetDatabaseSoftwareImage
    {
        /// <summary>
        /// This data source provides details about a specific Database Software Image resource in Oracle Cloud Infrastructure Database service.
        /// 
        /// Gets information about the specified database software image.
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
        ///         var testDatabaseSoftwareImage = Output.Create(Oci.Database.GetDatabaseSoftwareImage.InvokeAsync(new Oci.Database.GetDatabaseSoftwareImageArgs
        ///         {
        ///             DatabaseSoftwareImageId = oci_database_database_software_image.Test_database_software_image.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabaseSoftwareImageResult> InvokeAsync(GetDatabaseSoftwareImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseSoftwareImageResult>("oci:database/getDatabaseSoftwareImage:getDatabaseSoftwareImage", args ?? new GetDatabaseSoftwareImageArgs(), options.WithVersion());
    }


    public sealed class GetDatabaseSoftwareImageArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        [Input("databaseSoftwareImageId", required: true)]
        public string DatabaseSoftwareImageId { get; set; } = null!;

        public GetDatabaseSoftwareImageArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseSoftwareImageResult
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        /// </summary>
        public readonly string CompartmentId;
        public readonly string DatabaseSoftwareImageId;
        /// <summary>
        /// List of one-off patches for Database Homes.
        /// </summary>
        public readonly ImmutableArray<string> DatabaseSoftwareImageIncludedPatches;
        /// <summary>
        /// List of one-off patches for Database Homes.
        /// </summary>
        public readonly ImmutableArray<string> DatabaseSoftwareImageOneOffPatches;
        /// <summary>
        /// The database version with which the database software image is to be built.
        /// </summary>
        public readonly string DatabaseVersion;
        /// <summary>
        /// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// The user-friendly name for the database software image. The name does not have to be unique.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database software image.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// To what shape the image is meant for.
        /// </summary>
        public readonly string ImageShapeFamily;
        /// <summary>
        /// The type of software image. Can be grid or database.
        /// </summary>
        public readonly string ImageType;
        /// <summary>
        /// The patches included in the image and the version of the image
        /// </summary>
        public readonly string IncludedPatchesSummary;
        /// <summary>
        /// True if this Database software image is supported for Upgrade.
        /// </summary>
        public readonly bool IsUpgradeSupported;
        /// <summary>
        /// Detailed message for the lifecycle state.
        /// </summary>
        public readonly string LifecycleDetails;
        /// <summary>
        /// output from lsinventory which will get passed as a string
        /// </summary>
        public readonly string LsInventory;
        /// <summary>
        /// The PSU or PBP or Release Updates. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbVersionSummary/ListDbVersions) operation.
        /// </summary>
        public readonly string PatchSet;
        public readonly string SourceDbHomeId;
        /// <summary>
        /// The current state of the database software image.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The date and time the database software image was created.
        /// </summary>
        public readonly string TimeCreated;

        [OutputConstructor]
        private GetDatabaseSoftwareImageResult(
            string compartmentId,

            string databaseSoftwareImageId,

            ImmutableArray<string> databaseSoftwareImageIncludedPatches,

            ImmutableArray<string> databaseSoftwareImageOneOffPatches,

            string databaseVersion,

            ImmutableDictionary<string, object> definedTags,

            string displayName,

            ImmutableDictionary<string, object> freeformTags,

            string id,

            string imageShapeFamily,

            string imageType,

            string includedPatchesSummary,

            bool isUpgradeSupported,

            string lifecycleDetails,

            string lsInventory,

            string patchSet,

            string sourceDbHomeId,

            string state,

            string timeCreated)
        {
            CompartmentId = compartmentId;
            DatabaseSoftwareImageId = databaseSoftwareImageId;
            DatabaseSoftwareImageIncludedPatches = databaseSoftwareImageIncludedPatches;
            DatabaseSoftwareImageOneOffPatches = databaseSoftwareImageOneOffPatches;
            DatabaseVersion = databaseVersion;
            DefinedTags = definedTags;
            DisplayName = displayName;
            FreeformTags = freeformTags;
            Id = id;
            ImageShapeFamily = imageShapeFamily;
            ImageType = imageType;
            IncludedPatchesSummary = includedPatchesSummary;
            IsUpgradeSupported = isUpgradeSupported;
            LifecycleDetails = lifecycleDetails;
            LsInventory = lsInventory;
            PatchSet = patchSet;
            SourceDbHomeId = sourceDbHomeId;
            State = state;
            TimeCreated = timeCreated;
        }
    }
}