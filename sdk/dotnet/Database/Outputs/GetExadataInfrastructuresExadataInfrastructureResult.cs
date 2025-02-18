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
    public sealed class GetExadataInfrastructuresExadataInfrastructureResult
    {
        /// <summary>
        /// The requested number of additional storage servers activated for the Exadata infrastructure.
        /// </summary>
        public readonly int ActivatedStorageCount;
        public readonly string ActivationFile;
        /// <summary>
        /// The requested number of additional storage servers for the Exadata infrastructure.
        /// </summary>
        public readonly int AdditionalStorageCount;
        /// <summary>
        /// The CIDR block for the Exadata administration network.
        /// </summary>
        public readonly string AdminNetworkCidr;
        /// <summary>
        /// The IP address for the first control plane server.
        /// </summary>
        public readonly string CloudControlPlaneServer1;
        /// <summary>
        /// The IP address for the second control plane server.
        /// </summary>
        public readonly string CloudControlPlaneServer2;
        /// <summary>
        /// The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// The number of compute servers for the Exadata infrastructure.
        /// </summary>
        public readonly int ComputeCount;
        /// <summary>
        /// The list of contacts for the Exadata infrastructure.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetExadataInfrastructuresExadataInfrastructureContactResult> Contacts;
        /// <summary>
        /// The corporate network proxy for access to the control plane network.
        /// </summary>
        public readonly string CorporateProxy;
        /// <summary>
        /// The number of enabled CPU cores.
        /// </summary>
        public readonly int CpusEnabled;
        public readonly bool CreateAsync;
        /// <summary>
        /// The CSI Number of the Exadata infrastructure.
        /// </summary>
        public readonly string CsiNumber;
        /// <summary>
        /// Size, in terabytes, of the DATA disk group.
        /// </summary>
        public readonly double DataStorageSizeInTbs;
        /// <summary>
        /// The local node storage allocated in GBs.
        /// </summary>
        public readonly int DbNodeStorageSizeInGbs;
        /// <summary>
        /// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// A filter to return only resources that match the entire display name given. The match is not case sensitive.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The list of DNS server IP addresses. Maximum of 3 allowed.
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// The gateway for the control plane network.
        /// </summary>
        public readonly string Gateway;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The CIDR block for the Exadata InfiniBand interconnect.
        /// </summary>
        public readonly string InfiniBandNetworkCidr;
        /// <summary>
        /// Additional information about the current lifecycle state.
        /// </summary>
        public readonly string LifecycleDetails;
        /// <summary>
        /// A field to capture ‘Maintenance SLO Status’ for the Exadata infrastructure with values ‘OK’, ‘DEGRADED’. Default is ‘OK’ when the infrastructure is provisioned.
        /// </summary>
        public readonly string MaintenanceSloStatus;
        /// <summary>
        /// The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window.
        /// </summary>
        public readonly Outputs.GetExadataInfrastructuresExadataInfrastructureMaintenanceWindowResult MaintenanceWindow;
        /// <summary>
        /// The total number of CPU cores available.
        /// </summary>
        public readonly int MaxCpuCount;
        /// <summary>
        /// The total available DATA disk group size.
        /// </summary>
        public readonly double MaxDataStorageInTbs;
        /// <summary>
        /// The total local node storage available in GBs.
        /// </summary>
        public readonly int MaxDbNodeStorageInGbs;
        /// <summary>
        /// The total memory available in GBs.
        /// </summary>
        public readonly int MaxMemoryInGbs;
        /// <summary>
        /// The memory allocated in GBs.
        /// </summary>
        public readonly int MemorySizeInGbs;
        /// <summary>
        /// The netmask for the control plane network.
        /// </summary>
        public readonly string Netmask;
        /// <summary>
        /// The list of NTP server IP addresses. Maximum of 3 allowed.
        /// </summary>
        public readonly ImmutableArray<string> NtpServers;
        /// <summary>
        /// The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance.
        /// </summary>
        public readonly string Shape;
        /// <summary>
        /// A filter to return only resources that match the given lifecycle state exactly.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The number of Exadata storage servers for the Exadata infrastructure.
        /// </summary>
        public readonly int StorageCount;
        /// <summary>
        /// The date and time the Exadata infrastructure was created.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The time zone of the Exadata infrastructure. For details, see [Exadata Infrastructure Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
        /// </summary>
        public readonly string TimeZone;

        [OutputConstructor]
        private GetExadataInfrastructuresExadataInfrastructureResult(
            int activatedStorageCount,

            string activationFile,

            int additionalStorageCount,

            string adminNetworkCidr,

            string cloudControlPlaneServer1,

            string cloudControlPlaneServer2,

            string compartmentId,

            int computeCount,

            ImmutableArray<Outputs.GetExadataInfrastructuresExadataInfrastructureContactResult> contacts,

            string corporateProxy,

            int cpusEnabled,

            bool createAsync,

            string csiNumber,

            double dataStorageSizeInTbs,

            int dbNodeStorageSizeInGbs,

            ImmutableDictionary<string, object> definedTags,

            string displayName,

            ImmutableArray<string> dnsServers,

            ImmutableDictionary<string, object> freeformTags,

            string gateway,

            string id,

            string infiniBandNetworkCidr,

            string lifecycleDetails,

            string maintenanceSloStatus,

            Outputs.GetExadataInfrastructuresExadataInfrastructureMaintenanceWindowResult maintenanceWindow,

            int maxCpuCount,

            double maxDataStorageInTbs,

            int maxDbNodeStorageInGbs,

            int maxMemoryInGbs,

            int memorySizeInGbs,

            string netmask,

            ImmutableArray<string> ntpServers,

            string shape,

            string state,

            int storageCount,

            string timeCreated,

            string timeZone)
        {
            ActivatedStorageCount = activatedStorageCount;
            ActivationFile = activationFile;
            AdditionalStorageCount = additionalStorageCount;
            AdminNetworkCidr = adminNetworkCidr;
            CloudControlPlaneServer1 = cloudControlPlaneServer1;
            CloudControlPlaneServer2 = cloudControlPlaneServer2;
            CompartmentId = compartmentId;
            ComputeCount = computeCount;
            Contacts = contacts;
            CorporateProxy = corporateProxy;
            CpusEnabled = cpusEnabled;
            CreateAsync = createAsync;
            CsiNumber = csiNumber;
            DataStorageSizeInTbs = dataStorageSizeInTbs;
            DbNodeStorageSizeInGbs = dbNodeStorageSizeInGbs;
            DefinedTags = definedTags;
            DisplayName = displayName;
            DnsServers = dnsServers;
            FreeformTags = freeformTags;
            Gateway = gateway;
            Id = id;
            InfiniBandNetworkCidr = infiniBandNetworkCidr;
            LifecycleDetails = lifecycleDetails;
            MaintenanceSloStatus = maintenanceSloStatus;
            MaintenanceWindow = maintenanceWindow;
            MaxCpuCount = maxCpuCount;
            MaxDataStorageInTbs = maxDataStorageInTbs;
            MaxDbNodeStorageInGbs = maxDbNodeStorageInGbs;
            MaxMemoryInGbs = maxMemoryInGbs;
            MemorySizeInGbs = memorySizeInGbs;
            Netmask = netmask;
            NtpServers = ntpServers;
            Shape = shape;
            State = state;
            StorageCount = storageCount;
            TimeCreated = timeCreated;
            TimeZone = timeZone;
        }
    }
}
