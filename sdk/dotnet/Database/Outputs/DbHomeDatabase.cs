// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Database.Outputs
{

    [OutputType]
    public sealed class DbHomeDatabase
    {
        /// <summary>
        /// A strong password for SYS, SYSTEM, PDB Admin and TDE Wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numbers, and two special characters. The special characters must be _, \#, or -.
        /// </summary>
        public readonly string AdminPassword;
        /// <summary>
        /// The backup [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        public readonly string? BackupId;
        /// <summary>
        /// The password to open the TDE wallet.
        /// </summary>
        public readonly string? BackupTdePassword;
        /// <summary>
        /// The character set for the database.  The default is AL32UTF8. Allowed values are:
        /// </summary>
        public readonly string? CharacterSet;
        public readonly ImmutableArray<Outputs.DbHomeDatabaseConnectionString> ConnectionStrings;
        /// <summary>
        /// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        public readonly string? DatabaseId;
        /// <summary>
        /// The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
        /// </summary>
        public readonly string? DatabaseSoftwareImageId;
        /// <summary>
        /// (Updatable) Backup Options To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
        /// </summary>
        public readonly Outputs.DbHomeDatabaseDbBackupConfig? DbBackupConfig;
        /// <summary>
        /// The display name of the database to be created from the backup. It must begin with an alphabetic character and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
        /// </summary>
        public readonly string? DbName;
        public readonly string? DbUniqueName;
        /// <summary>
        /// The database workload type.
        /// </summary>
        public readonly string? DbWorkload;
        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? DefinedTags;
        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object>? FreeformTags;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup destination.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Additional information about the current lifecycle state.
        /// </summary>
        public readonly string? LifecycleDetails;
        /// <summary>
        /// The national character set for the database.  The default is AL16UTF16. Allowed values are: AL16UTF16 or UTF8.
        /// </summary>
        public readonly string? NcharacterSet;
        /// <summary>
        /// List of one-off patches for Database Homes.
        /// </summary>
        public readonly ImmutableArray<string> OneOffPatches;
        /// <summary>
        /// The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
        /// </summary>
        public readonly string? PdbName;
        /// <summary>
        /// The current state of the Database Home.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// The optional password to open the TDE wallet. The password must be at least nine characters and contain at least two uppercase, two lowercase, two numeric, and two special characters. The special characters must be _, \#, or -.
        /// </summary>
        public readonly string? TdeWalletPassword;
        /// <summary>
        /// The date and time the Database Home was created.
        /// </summary>
        public readonly string? TimeCreated;
        /// <summary>
        /// The point in time of the original database from which the new database is created. If not specifed, the latest backup is used to create the database.
        /// </summary>
        public readonly string? TimeStampForPointInTimeRecovery;

        [OutputConstructor]
        private DbHomeDatabase(
            string adminPassword,

            string? backupId,

            string? backupTdePassword,

            string? characterSet,

            ImmutableArray<Outputs.DbHomeDatabaseConnectionString> connectionStrings,

            string? databaseId,

            string? databaseSoftwareImageId,

            Outputs.DbHomeDatabaseDbBackupConfig? dbBackupConfig,

            string? dbName,

            string? dbUniqueName,

            string? dbWorkload,

            ImmutableDictionary<string, object>? definedTags,

            ImmutableDictionary<string, object>? freeformTags,

            string? id,

            string? lifecycleDetails,

            string? ncharacterSet,

            ImmutableArray<string> oneOffPatches,

            string? pdbName,

            string? state,

            string? tdeWalletPassword,

            string? timeCreated,

            string? timeStampForPointInTimeRecovery)
        {
            AdminPassword = adminPassword;
            BackupId = backupId;
            BackupTdePassword = backupTdePassword;
            CharacterSet = characterSet;
            ConnectionStrings = connectionStrings;
            DatabaseId = databaseId;
            DatabaseSoftwareImageId = databaseSoftwareImageId;
            DbBackupConfig = dbBackupConfig;
            DbName = dbName;
            DbUniqueName = dbUniqueName;
            DbWorkload = dbWorkload;
            DefinedTags = definedTags;
            FreeformTags = freeformTags;
            Id = id;
            LifecycleDetails = lifecycleDetails;
            NcharacterSet = ncharacterSet;
            OneOffPatches = oneOffPatches;
            PdbName = pdbName;
            State = state;
            TdeWalletPassword = tdeWalletPassword;
            TimeCreated = timeCreated;
            TimeStampForPointInTimeRecovery = timeStampForPointInTimeRecovery;
        }
    }
}