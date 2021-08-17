// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetKmsVault
    {
        /// <summary>
        /// This data source provides details about a specific Vault resource in Oracle Cloud Infrastructure Kms service.
        /// 
        /// Gets the specified vault's configuration information.
        /// 
        /// As a provisioning operation, this call is subject to a Key Management limit that applies to
        /// the total number of requests across all provisioning read operations. Key Management might
        /// throttle this call to reject an otherwise valid request when the total rate of provisioning
        /// read operations exceeds 10 requests per second for a given tenancy.
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
        ///         var testVault = Output.Create(Oci.GetKmsVault.InvokeAsync(new Oci.GetKmsVaultArgs
        ///         {
        ///             VaultId = oci_kms_vault.Test_vault.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetKmsVaultResult> InvokeAsync(GetKmsVaultArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKmsVaultResult>("oci:index/getKmsVault:GetKmsVault", args ?? new GetKmsVaultArgs(), options.WithVersion());
    }


    public sealed class GetKmsVaultArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The OCID of the vault.
        /// </summary>
        [Input("vaultId", required: true)]
        public string VaultId { get; set; } = null!;

        public GetKmsVaultArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKmsVaultResult
    {
        /// <summary>
        /// The OCID of the compartment that contains a particular vault.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// The service endpoint to perform cryptographic operations against. Cryptographic operations include [Encrypt](https://docs.cloud.oracle.com/iaas/api/#/en/key/latest/EncryptedData/Encrypt), [Decrypt](https://docs.cloud.oracle.com/iaas/api/#/en/key/latest/DecryptedData/Decrypt), and [GenerateDataEncryptionKey](https://docs.cloud.oracle.com/iaas/api/#/en/key/latest/GeneratedKey/GenerateDataEncryptionKey) operations.
        /// </summary>
        public readonly string CryptoEndpoint;
        /// <summary>
        /// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// A user-friendly name for the vault. It does not have to be unique, and it is changeable. Avoid entering confidential information.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// The OCID of the vault.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A boolean that will be true when vault is primary, and will be false when vault is a replica from a primary vault.
        /// </summary>
        public readonly bool IsPrimary;
        /// <summary>
        /// The service endpoint to perform management operations against. Management operations include "Create," "Update," "List," "Get," and "Delete" operations.
        /// </summary>
        public readonly string ManagementEndpoint;
        /// <summary>
        /// Vault replica details
        /// </summary>
        public readonly Outputs.GetKmsVaultReplicaDetailsResult ReplicaDetails;
        /// <summary>
        /// Details where vault was backed up.
        /// </summary>
        public readonly Outputs.GetKmsVaultRestoreFromFileResult RestoreFromFile;
        /// <summary>
        /// Details where vault was backed up
        /// </summary>
        public readonly Outputs.GetKmsVaultRestoreFromObjectStoreResult RestoreFromObjectStore;
        /// <summary>
        /// When flipped, triggers restore if restore options are provided. Values of 0 or 1 are supported
        /// </summary>
        public readonly bool RestoreTrigger;
        /// <summary>
        /// The OCID of the vault from which this vault was restored, if it was restored from a backup file.  If you restore a vault to the same region, the vault retains the same OCID that it had when you  backed up the vault.
        /// </summary>
        public readonly string RestoredFromVaultId;
        /// <summary>
        /// The vault's current lifecycle state.  Example: `DELETED`
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The date and time this vault was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z`
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// An optional property to indicate when to delete the vault, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z`
        /// </summary>
        public readonly string TimeOfDeletion;
        public readonly string VaultId;
        /// <summary>
        /// The type of vault. Each type of vault stores the key with different degrees of isolation and has different options and pricing.
        /// </summary>
        public readonly string VaultType;

        [OutputConstructor]
        private GetKmsVaultResult(
            string compartmentId,

            string cryptoEndpoint,

            ImmutableDictionary<string, object> definedTags,

            string displayName,

            ImmutableDictionary<string, object> freeformTags,

            string id,

            bool isPrimary,

            string managementEndpoint,

            Outputs.GetKmsVaultReplicaDetailsResult replicaDetails,

            Outputs.GetKmsVaultRestoreFromFileResult restoreFromFile,

            Outputs.GetKmsVaultRestoreFromObjectStoreResult restoreFromObjectStore,

            bool restoreTrigger,

            string restoredFromVaultId,

            string state,

            string timeCreated,

            string timeOfDeletion,

            string vaultId,

            string vaultType)
        {
            CompartmentId = compartmentId;
            CryptoEndpoint = cryptoEndpoint;
            DefinedTags = definedTags;
            DisplayName = displayName;
            FreeformTags = freeformTags;
            Id = id;
            IsPrimary = isPrimary;
            ManagementEndpoint = managementEndpoint;
            ReplicaDetails = replicaDetails;
            RestoreFromFile = restoreFromFile;
            RestoreFromObjectStore = restoreFromObjectStore;
            RestoreTrigger = restoreTrigger;
            RestoredFromVaultId = restoredFromVaultId;
            State = state;
            TimeCreated = timeCreated;
            TimeOfDeletion = timeOfDeletion;
            VaultId = vaultId;
            VaultType = vaultType;
        }
    }
}