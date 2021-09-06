// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Database
{
    public static class GetDatabaseSoftwareImages
    {
        /// <summary>
        /// This data source provides the list of Database Software Images in Oracle Cloud Infrastructure Database service.
        /// 
        /// Gets a list of the database software images in the specified compartment.
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
        ///         var testDatabaseSoftwareImages = Output.Create(Oci.Database.GetDatabaseSoftwareImages.InvokeAsync(new Oci.Database.GetDatabaseSoftwareImagesArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             DisplayName = @var.Database_software_image_display_name,
        ///             ImageShapeFamily = @var.Database_software_image_image_shape_family,
        ///             ImageType = @var.Database_software_image_image_type,
        ///             IsUpgradeSupported = @var.Database_software_image_is_upgrade_supported,
        ///             State = @var.Database_software_image_state,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabaseSoftwareImagesResult> InvokeAsync(GetDatabaseSoftwareImagesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseSoftwareImagesResult>("oci:database/getDatabaseSoftwareImages:getDatabaseSoftwareImages", args ?? new GetDatabaseSoftwareImagesArgs(), options.WithVersion());
    }


    public sealed class GetDatabaseSoftwareImagesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// A filter to return only resources that match the entire display name given. The match is not case sensitive.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("filters")]
        private List<Inputs.GetDatabaseSoftwareImagesFilterArgs>? _filters;
        public List<Inputs.GetDatabaseSoftwareImagesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetDatabaseSoftwareImagesFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// A filter to return only resources that match the given image shape family exactly.
        /// </summary>
        [Input("imageShapeFamily")]
        public string? ImageShapeFamily { get; set; }

        /// <summary>
        /// A filter to return only resources that match the given image type exactly.
        /// </summary>
        [Input("imageType")]
        public string? ImageType { get; set; }

        /// <summary>
        /// If provided, filters the results to the set of database versions which are supported for Upgrade.
        /// </summary>
        [Input("isUpgradeSupported")]
        public bool? IsUpgradeSupported { get; set; }

        /// <summary>
        /// A filter to return only resources that match the given lifecycle state exactly.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetDatabaseSoftwareImagesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseSoftwareImagesResult
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// The list of database_software_images.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseSoftwareImagesDatabaseSoftwareImageResult> DatabaseSoftwareImages;
        /// <summary>
        /// The user-friendly name for the database software image. The name does not have to be unique.
        /// </summary>
        public readonly string? DisplayName;
        public readonly ImmutableArray<Outputs.GetDatabaseSoftwareImagesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// To what shape the image is meant for.
        /// </summary>
        public readonly string? ImageShapeFamily;
        /// <summary>
        /// The type of software image. Can be grid or database.
        /// </summary>
        public readonly string? ImageType;
        /// <summary>
        /// True if this Database software image is supported for Upgrade.
        /// </summary>
        public readonly bool? IsUpgradeSupported;
        /// <summary>
        /// The current state of the database software image.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private GetDatabaseSoftwareImagesResult(
            string compartmentId,

            ImmutableArray<Outputs.GetDatabaseSoftwareImagesDatabaseSoftwareImageResult> databaseSoftwareImages,

            string? displayName,

            ImmutableArray<Outputs.GetDatabaseSoftwareImagesFilterResult> filters,

            string id,

            string? imageShapeFamily,

            string? imageType,

            bool? isUpgradeSupported,

            string? state)
        {
            CompartmentId = compartmentId;
            DatabaseSoftwareImages = databaseSoftwareImages;
            DisplayName = displayName;
            Filters = filters;
            Id = id;
            ImageShapeFamily = imageShapeFamily;
            ImageType = imageType;
            IsUpgradeSupported = isUpgradeSupported;
            State = state;
        }
    }
}