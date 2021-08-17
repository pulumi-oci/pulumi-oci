// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetCoreBootVolumeBackups
    {
        /// <summary>
        /// This data source provides the list of Boot Volume Backups in Oracle Cloud Infrastructure Core service.
        /// 
        /// Lists the boot volume backups in the specified compartment. You can filter the results by boot volume.
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
        ///         var testBootVolumeBackups = Output.Create(Oci.GetCoreBootVolumeBackups.InvokeAsync(new Oci.GetCoreBootVolumeBackupsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             BootVolumeId = oci_core_boot_volume.Test_boot_volume.Id,
        ///             DisplayName = @var.Boot_volume_backup_display_name,
        ///             SourceBootVolumeBackupId = oci_core_boot_volume_backup.Test_boot_volume_backup.Id,
        ///             State = @var.Boot_volume_backup_state,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCoreBootVolumeBackupsResult> InvokeAsync(GetCoreBootVolumeBackupsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCoreBootVolumeBackupsResult>("oci:index/getCoreBootVolumeBackups:GetCoreBootVolumeBackups", args ?? new GetCoreBootVolumeBackupsArgs(), options.WithVersion());
    }


    public sealed class GetCoreBootVolumeBackupsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The OCID of the boot volume.
        /// </summary>
        [Input("bootVolumeId")]
        public string? BootVolumeId { get; set; }

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
        private List<Inputs.GetCoreBootVolumeBackupsFilterArgs>? _filters;
        public List<Inputs.GetCoreBootVolumeBackupsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetCoreBootVolumeBackupsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// A filter to return only resources that originated from the given source boot volume backup.
        /// </summary>
        [Input("sourceBootVolumeBackupId")]
        public string? SourceBootVolumeBackupId { get; set; }

        /// <summary>
        /// A filter to only return resources that match the given lifecycle state. The state value is case-insensitive.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetCoreBootVolumeBackupsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCoreBootVolumeBackupsResult
    {
        /// <summary>
        /// The list of boot_volume_backups.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCoreBootVolumeBackupsBootVolumeBackupResult> BootVolumeBackups;
        /// <summary>
        /// The OCID of the boot volume.
        /// </summary>
        public readonly string? BootVolumeId;
        /// <summary>
        /// The OCID of the compartment that contains the boot volume backup.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// A user-friendly name for the boot volume backup. Does not have to be unique and it's changeable. Avoid entering confidential information.
        /// </summary>
        public readonly string? DisplayName;
        public readonly ImmutableArray<Outputs.GetCoreBootVolumeBackupsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The OCID of the source boot volume backup.
        /// </summary>
        public readonly string? SourceBootVolumeBackupId;
        /// <summary>
        /// The current state of a boot volume backup.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private GetCoreBootVolumeBackupsResult(
            ImmutableArray<Outputs.GetCoreBootVolumeBackupsBootVolumeBackupResult> bootVolumeBackups,

            string? bootVolumeId,

            string compartmentId,

            string? displayName,

            ImmutableArray<Outputs.GetCoreBootVolumeBackupsFilterResult> filters,

            string id,

            string? sourceBootVolumeBackupId,

            string? state)
        {
            BootVolumeBackups = bootVolumeBackups;
            BootVolumeId = bootVolumeId;
            CompartmentId = compartmentId;
            DisplayName = displayName;
            Filters = filters;
            Id = id;
            SourceBootVolumeBackupId = sourceBootVolumeBackupId;
            State = state;
        }
    }
}