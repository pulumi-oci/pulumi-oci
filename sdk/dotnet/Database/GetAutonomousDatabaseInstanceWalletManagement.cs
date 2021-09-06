// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Database
{
    public static class GetAutonomousDatabaseInstanceWalletManagement
    {
        /// <summary>
        /// This data source provides details about a specific Autonomous Database Instance Wallet Management resource in Oracle Cloud Infrastructure Database service.
        /// 
        /// Gets the wallet details for the specified Autonomous Database.
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
        ///         var testAutonomousDatabaseInstanceWalletManagement = Output.Create(Oci.Database.GetAutonomousDatabaseInstanceWalletManagement.InvokeAsync(new Oci.Database.GetAutonomousDatabaseInstanceWalletManagementArgs
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
        public static Task<GetAutonomousDatabaseInstanceWalletManagementResult> InvokeAsync(GetAutonomousDatabaseInstanceWalletManagementArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAutonomousDatabaseInstanceWalletManagementResult>("oci:database/getAutonomousDatabaseInstanceWalletManagement:getAutonomousDatabaseInstanceWalletManagement", args ?? new GetAutonomousDatabaseInstanceWalletManagementArgs(), options.WithVersion());
    }


    public sealed class GetAutonomousDatabaseInstanceWalletManagementArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        [Input("autonomousDatabaseId", required: true)]
        public string AutonomousDatabaseId { get; set; } = null!;

        public GetAutonomousDatabaseInstanceWalletManagementArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAutonomousDatabaseInstanceWalletManagementResult
    {
        /// <summary>
        /// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        public readonly string AutonomousDatabaseId;
        public readonly string Id;
        /// <summary>
        /// Indicates whether to rotate the wallet or not. If `false`, the wallet will not be rotated. The default is `false`.
        /// </summary>
        public readonly bool ShouldRotate;
        /// <summary>
        /// The current lifecycle state of the Autonomous Database wallet.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The date and time the wallet was last rotated.
        /// </summary>
        public readonly string TimeRotated;

        [OutputConstructor]
        private GetAutonomousDatabaseInstanceWalletManagementResult(
            string autonomousDatabaseId,

            string id,

            bool shouldRotate,

            string state,

            string timeRotated)
        {
            AutonomousDatabaseId = autonomousDatabaseId;
            Id = id;
            ShouldRotate = shouldRotate;
            State = state;
            TimeRotated = timeRotated;
        }
    }
}