// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Database
{
    public static class GetAutonomousDatabaseDataguardAssociations
    {
        /// <summary>
        /// This data source provides the list of Autonomous Database Dataguard Associations in Oracle Cloud Infrastructure Database service.
        /// 
        /// Gets a list of the Autonomous Data Guard-enabled databases associated with the specified Autonomous Database.
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
        ///         var testAutonomousDatabaseDataguardAssociations = Output.Create(Oci.Database.GetAutonomousDatabaseDataguardAssociations.InvokeAsync(new Oci.Database.GetAutonomousDatabaseDataguardAssociationsArgs
        ///         {
        ///             AutonomousDatabaseId = oci_database_autonomous_database.Test_autonomous_database.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAutonomousDatabaseDataguardAssociationsResult> InvokeAsync(GetAutonomousDatabaseDataguardAssociationsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAutonomousDatabaseDataguardAssociationsResult>("oci:database/getAutonomousDatabaseDataguardAssociations:getAutonomousDatabaseDataguardAssociations", args ?? new GetAutonomousDatabaseDataguardAssociationsArgs(), options.WithVersion());
    }


    public sealed class GetAutonomousDatabaseDataguardAssociationsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        [Input("autonomousDatabaseId", required: true)]
        public string AutonomousDatabaseId { get; set; } = null!;

        [Input("filters")]
        private List<Inputs.GetAutonomousDatabaseDataguardAssociationsFilterArgs>? _filters;
        public List<Inputs.GetAutonomousDatabaseDataguardAssociationsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetAutonomousDatabaseDataguardAssociationsFilterArgs>());
            set => _filters = value;
        }

        public GetAutonomousDatabaseDataguardAssociationsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAutonomousDatabaseDataguardAssociationsResult
    {
        /// <summary>
        /// The list of autonomous_database_dataguard_associations.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAutonomousDatabaseDataguardAssociationsAutonomousDatabaseDataguardAssociationResult> AutonomousDatabaseDataguardAssociations;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database that has a relationship with the peer Autonomous Database.
        /// </summary>
        public readonly string AutonomousDatabaseId;
        public readonly ImmutableArray<Outputs.GetAutonomousDatabaseDataguardAssociationsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAutonomousDatabaseDataguardAssociationsResult(
            ImmutableArray<Outputs.GetAutonomousDatabaseDataguardAssociationsAutonomousDatabaseDataguardAssociationResult> autonomousDatabaseDataguardAssociations,

            string autonomousDatabaseId,

            ImmutableArray<Outputs.GetAutonomousDatabaseDataguardAssociationsFilterResult> filters,

            string id)
        {
            AutonomousDatabaseDataguardAssociations = autonomousDatabaseDataguardAssociations;
            AutonomousDatabaseId = autonomousDatabaseId;
            Filters = filters;
            Id = id;
        }
    }
}