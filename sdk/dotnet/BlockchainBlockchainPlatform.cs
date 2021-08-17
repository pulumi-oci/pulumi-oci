// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    /// <summary>
    /// This resource provides the Blockchain Platform resource in Oracle Cloud Infrastructure Blockchain service.
    /// 
    /// Creates a new Blockchain Platform.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Oci = Pulumi.Oci;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testBlockchainPlatform = new Oci.BlockchainBlockchainPlatform("testBlockchainPlatform", new Oci.BlockchainBlockchainPlatformArgs
    ///         {
    ///             CompartmentId = @var.Compartment_id,
    ///             ComputeShape = @var.Blockchain_platform_compute_shape,
    ///             DisplayName = @var.Blockchain_platform_display_name,
    ///             IdcsAccessToken = @var.Blockchain_platform_idcs_access_token,
    ///             PlatformRole = @var.Blockchain_platform_platform_role,
    ///             CaCertArchiveText = @var.Blockchain_platform_ca_cert_archive_text,
    ///             DefinedTags = 
    ///             {
    ///                 { "foo-namespace.bar-key", "value" },
    ///             },
    ///             Description = @var.Blockchain_platform_description,
    ///             FederatedUserId = oci_identity_user.Test_user.Id,
    ///             FreeformTags = 
    ///             {
    ///                 { "bar-key", "value" },
    ///             },
    ///             IsByol = @var.Blockchain_platform_is_byol,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// BlockchainPlatforms can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import oci:index/blockchainBlockchainPlatform:BlockchainBlockchainPlatform test_blockchain_platform "id"
    /// ```
    /// </summary>
    [OciResourceType("oci:index/blockchainBlockchainPlatform:BlockchainBlockchainPlatform")]
    public partial class BlockchainBlockchainPlatform : Pulumi.CustomResource
    {
        /// <summary>
        /// Base64 encoded text in ASCII character set of a Thirdparty CA Certificates archive file. The Archive file is a zip file containing third part CA Certificates, the ca key and certificate files used when issuing enrollment certificates (ECerts) and transaction certificates (TCerts). The chainfile (if it exists) contains the certificate chain which should be trusted for this CA, where the 1st in the chain is always the root CA certificate. File list in zip file [ca-cert.pem,ca-key.pem,ca-chain.pem(optional)].
        /// </summary>
        [Output("caCertArchiveText")]
        public Output<string> CaCertArchiveText { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Compartment Identifier
        /// </summary>
        [Output("compartmentId")]
        public Output<string> CompartmentId { get; private set; } = null!;

        /// <summary>
        /// Blockchain Platform component details.
        /// </summary>
        [Output("componentDetails")]
        public Output<Outputs.BlockchainBlockchainPlatformComponentDetails> ComponentDetails { get; private set; } = null!;

        /// <summary>
        /// Compute shape - STANDARD or ENTERPRISE_SMALL or ENTERPRISE_MEDIUM or ENTERPRISE_LARGE or ENTERPRISE_EXTRA_LARGE
        /// </summary>
        [Output("computeShape")]
        public Output<string> ComputeShape { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
        /// </summary>
        [Output("definedTags")]
        public Output<ImmutableDictionary<string, object>> DefinedTags { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Platform Instance Description
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Platform Instance Display name, can be renamed
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Identifier for a federated user
        /// </summary>
        [Output("federatedUserId")]
        public Output<string> FederatedUserId { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
        /// </summary>
        [Output("freeformTags")]
        public Output<ImmutableDictionary<string, object>> FreeformTags { get; private set; } = null!;

        /// <summary>
        /// List of OcpuUtilization for all hosts
        /// </summary>
        [Output("hostOcpuUtilizationInfos")]
        public Output<ImmutableArray<Outputs.BlockchainBlockchainPlatformHostOcpuUtilizationInfo>> HostOcpuUtilizationInfos { get; private set; } = null!;

        /// <summary>
        /// IDCS access token with Identity Domain Administrator role
        /// </summary>
        [Output("idcsAccessToken")]
        public Output<string> IdcsAccessToken { get; private set; } = null!;

        /// <summary>
        /// Bring your own license
        /// </summary>
        [Output("isByol")]
        public Output<bool> IsByol { get; private set; } = null!;

        /// <summary>
        /// True for multi-AD blockchain plaforms, false for single-AD
        /// </summary>
        [Output("isMultiAd")]
        public Output<bool> IsMultiAd { get; private set; } = null!;

        /// <summary>
        /// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
        /// </summary>
        [Output("lifecycleDetails")]
        public Output<string> LifecycleDetails { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Type of Load Balancer shape - LB_100_MBPS or LB_400_MBPS. Default is LB_100_MBPS.
        /// </summary>
        [Output("loadBalancerShape")]
        public Output<string> LoadBalancerShape { get; private set; } = null!;

        /// <summary>
        /// Role of platform - founder or participant
        /// </summary>
        [Output("platformRole")]
        public Output<string> PlatformRole { get; private set; } = null!;

        /// <summary>
        /// Type of Platform shape - DEFAULT or CUSTOM
        /// </summary>
        [Output("platformShapeType")]
        public Output<string> PlatformShapeType { get; private set; } = null!;

        /// <summary>
        /// Number of replicas of service components like Rest Proxy, CA and Console
        /// </summary>
        [Output("replicas")]
        public Output<Outputs.BlockchainBlockchainPlatformReplicas> Replicas { get; private set; } = null!;

        /// <summary>
        /// Service endpoint URL, valid post-provisioning
        /// </summary>
        [Output("serviceEndpoint")]
        public Output<string> ServiceEndpoint { get; private set; } = null!;

        /// <summary>
        /// The version of the Platform Instance.
        /// </summary>
        [Output("serviceVersion")]
        public Output<string> ServiceVersion { get; private set; } = null!;

        /// <summary>
        /// The current state of the Platform Instance.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Storage size in TBs
        /// </summary>
        [Output("storageSizeInTbs")]
        public Output<double> StorageSizeInTbs { get; private set; } = null!;

        /// <summary>
        /// Storage used in TBs
        /// </summary>
        [Output("storageUsedInTbs")]
        public Output<double> StorageUsedInTbs { get; private set; } = null!;

        /// <summary>
        /// The time the the Platform Instance was created. An RFC3339 formatted datetime string
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// The time the Platform Instance was updated. An RFC3339 formatted datetime string
        /// </summary>
        [Output("timeUpdated")]
        public Output<string> TimeUpdated { get; private set; } = null!;

        /// <summary>
        /// Number of total OCPUs allocated to the platform cluster
        /// </summary>
        [Output("totalOcpuCapacity")]
        public Output<int> TotalOcpuCapacity { get; private set; } = null!;


        /// <summary>
        /// Create a BlockchainBlockchainPlatform resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BlockchainBlockchainPlatform(string name, BlockchainBlockchainPlatformArgs args, CustomResourceOptions? options = null)
            : base("oci:index/blockchainBlockchainPlatform:BlockchainBlockchainPlatform", name, args ?? new BlockchainBlockchainPlatformArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BlockchainBlockchainPlatform(string name, Input<string> id, BlockchainBlockchainPlatformState? state = null, CustomResourceOptions? options = null)
            : base("oci:index/blockchainBlockchainPlatform:BlockchainBlockchainPlatform", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BlockchainBlockchainPlatform resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BlockchainBlockchainPlatform Get(string name, Input<string> id, BlockchainBlockchainPlatformState? state = null, CustomResourceOptions? options = null)
        {
            return new BlockchainBlockchainPlatform(name, id, state, options);
        }
    }

    public sealed class BlockchainBlockchainPlatformArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base64 encoded text in ASCII character set of a Thirdparty CA Certificates archive file. The Archive file is a zip file containing third part CA Certificates, the ca key and certificate files used when issuing enrollment certificates (ECerts) and transaction certificates (TCerts). The chainfile (if it exists) contains the certificate chain which should be trusted for this CA, where the 1st in the chain is always the root CA certificate. File list in zip file [ca-cert.pem,ca-key.pem,ca-chain.pem(optional)].
        /// </summary>
        [Input("caCertArchiveText")]
        public Input<string>? CaCertArchiveText { get; set; }

        /// <summary>
        /// (Updatable) Compartment Identifier
        /// </summary>
        [Input("compartmentId", required: true)]
        public Input<string> CompartmentId { get; set; } = null!;

        /// <summary>
        /// Compute shape - STANDARD or ENTERPRISE_SMALL or ENTERPRISE_MEDIUM or ENTERPRISE_LARGE or ENTERPRISE_EXTRA_LARGE
        /// </summary>
        [Input("computeShape", required: true)]
        public Input<string> ComputeShape { get; set; } = null!;

        [Input("definedTags")]
        private InputMap<object>? _definedTags;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
        /// </summary>
        public InputMap<object> DefinedTags
        {
            get => _definedTags ?? (_definedTags = new InputMap<object>());
            set => _definedTags = value;
        }

        /// <summary>
        /// (Updatable) Platform Instance Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Platform Instance Display name, can be renamed
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Identifier for a federated user
        /// </summary>
        [Input("federatedUserId")]
        public Input<string>? FederatedUserId { get; set; }

        [Input("freeformTags")]
        private InputMap<object>? _freeformTags;

        /// <summary>
        /// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
        /// </summary>
        public InputMap<object> FreeformTags
        {
            get => _freeformTags ?? (_freeformTags = new InputMap<object>());
            set => _freeformTags = value;
        }

        /// <summary>
        /// IDCS access token with Identity Domain Administrator role
        /// </summary>
        [Input("idcsAccessToken", required: true)]
        public Input<string> IdcsAccessToken { get; set; } = null!;

        /// <summary>
        /// Bring your own license
        /// </summary>
        [Input("isByol")]
        public Input<bool>? IsByol { get; set; }

        /// <summary>
        /// (Updatable) Type of Load Balancer shape - LB_100_MBPS or LB_400_MBPS. Default is LB_100_MBPS.
        /// </summary>
        [Input("loadBalancerShape")]
        public Input<string>? LoadBalancerShape { get; set; }

        /// <summary>
        /// Role of platform - founder or participant
        /// </summary>
        [Input("platformRole", required: true)]
        public Input<string> PlatformRole { get; set; } = null!;

        /// <summary>
        /// Number of replicas of service components like Rest Proxy, CA and Console
        /// </summary>
        [Input("replicas")]
        public Input<Inputs.BlockchainBlockchainPlatformReplicasArgs>? Replicas { get; set; }

        /// <summary>
        /// Storage size in TBs
        /// </summary>
        [Input("storageSizeInTbs")]
        public Input<double>? StorageSizeInTbs { get; set; }

        /// <summary>
        /// Number of total OCPUs allocated to the platform cluster
        /// </summary>
        [Input("totalOcpuCapacity")]
        public Input<int>? TotalOcpuCapacity { get; set; }

        public BlockchainBlockchainPlatformArgs()
        {
        }
    }

    public sealed class BlockchainBlockchainPlatformState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base64 encoded text in ASCII character set of a Thirdparty CA Certificates archive file. The Archive file is a zip file containing third part CA Certificates, the ca key and certificate files used when issuing enrollment certificates (ECerts) and transaction certificates (TCerts). The chainfile (if it exists) contains the certificate chain which should be trusted for this CA, where the 1st in the chain is always the root CA certificate. File list in zip file [ca-cert.pem,ca-key.pem,ca-chain.pem(optional)].
        /// </summary>
        [Input("caCertArchiveText")]
        public Input<string>? CaCertArchiveText { get; set; }

        /// <summary>
        /// (Updatable) Compartment Identifier
        /// </summary>
        [Input("compartmentId")]
        public Input<string>? CompartmentId { get; set; }

        /// <summary>
        /// Blockchain Platform component details.
        /// </summary>
        [Input("componentDetails")]
        public Input<Inputs.BlockchainBlockchainPlatformComponentDetailsGetArgs>? ComponentDetails { get; set; }

        /// <summary>
        /// Compute shape - STANDARD or ENTERPRISE_SMALL or ENTERPRISE_MEDIUM or ENTERPRISE_LARGE or ENTERPRISE_EXTRA_LARGE
        /// </summary>
        [Input("computeShape")]
        public Input<string>? ComputeShape { get; set; }

        [Input("definedTags")]
        private InputMap<object>? _definedTags;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
        /// </summary>
        public InputMap<object> DefinedTags
        {
            get => _definedTags ?? (_definedTags = new InputMap<object>());
            set => _definedTags = value;
        }

        /// <summary>
        /// (Updatable) Platform Instance Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Platform Instance Display name, can be renamed
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Identifier for a federated user
        /// </summary>
        [Input("federatedUserId")]
        public Input<string>? FederatedUserId { get; set; }

        [Input("freeformTags")]
        private InputMap<object>? _freeformTags;

        /// <summary>
        /// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
        /// </summary>
        public InputMap<object> FreeformTags
        {
            get => _freeformTags ?? (_freeformTags = new InputMap<object>());
            set => _freeformTags = value;
        }

        [Input("hostOcpuUtilizationInfos")]
        private InputList<Inputs.BlockchainBlockchainPlatformHostOcpuUtilizationInfoGetArgs>? _hostOcpuUtilizationInfos;

        /// <summary>
        /// List of OcpuUtilization for all hosts
        /// </summary>
        public InputList<Inputs.BlockchainBlockchainPlatformHostOcpuUtilizationInfoGetArgs> HostOcpuUtilizationInfos
        {
            get => _hostOcpuUtilizationInfos ?? (_hostOcpuUtilizationInfos = new InputList<Inputs.BlockchainBlockchainPlatformHostOcpuUtilizationInfoGetArgs>());
            set => _hostOcpuUtilizationInfos = value;
        }

        /// <summary>
        /// IDCS access token with Identity Domain Administrator role
        /// </summary>
        [Input("idcsAccessToken")]
        public Input<string>? IdcsAccessToken { get; set; }

        /// <summary>
        /// Bring your own license
        /// </summary>
        [Input("isByol")]
        public Input<bool>? IsByol { get; set; }

        /// <summary>
        /// True for multi-AD blockchain plaforms, false for single-AD
        /// </summary>
        [Input("isMultiAd")]
        public Input<bool>? IsMultiAd { get; set; }

        /// <summary>
        /// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
        /// </summary>
        [Input("lifecycleDetails")]
        public Input<string>? LifecycleDetails { get; set; }

        /// <summary>
        /// (Updatable) Type of Load Balancer shape - LB_100_MBPS or LB_400_MBPS. Default is LB_100_MBPS.
        /// </summary>
        [Input("loadBalancerShape")]
        public Input<string>? LoadBalancerShape { get; set; }

        /// <summary>
        /// Role of platform - founder or participant
        /// </summary>
        [Input("platformRole")]
        public Input<string>? PlatformRole { get; set; }

        /// <summary>
        /// Type of Platform shape - DEFAULT or CUSTOM
        /// </summary>
        [Input("platformShapeType")]
        public Input<string>? PlatformShapeType { get; set; }

        /// <summary>
        /// Number of replicas of service components like Rest Proxy, CA and Console
        /// </summary>
        [Input("replicas")]
        public Input<Inputs.BlockchainBlockchainPlatformReplicasGetArgs>? Replicas { get; set; }

        /// <summary>
        /// Service endpoint URL, valid post-provisioning
        /// </summary>
        [Input("serviceEndpoint")]
        public Input<string>? ServiceEndpoint { get; set; }

        /// <summary>
        /// The version of the Platform Instance.
        /// </summary>
        [Input("serviceVersion")]
        public Input<string>? ServiceVersion { get; set; }

        /// <summary>
        /// The current state of the Platform Instance.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Storage size in TBs
        /// </summary>
        [Input("storageSizeInTbs")]
        public Input<double>? StorageSizeInTbs { get; set; }

        /// <summary>
        /// Storage used in TBs
        /// </summary>
        [Input("storageUsedInTbs")]
        public Input<double>? StorageUsedInTbs { get; set; }

        /// <summary>
        /// The time the the Platform Instance was created. An RFC3339 formatted datetime string
        /// </summary>
        [Input("timeCreated")]
        public Input<string>? TimeCreated { get; set; }

        /// <summary>
        /// The time the Platform Instance was updated. An RFC3339 formatted datetime string
        /// </summary>
        [Input("timeUpdated")]
        public Input<string>? TimeUpdated { get; set; }

        /// <summary>
        /// Number of total OCPUs allocated to the platform cluster
        /// </summary>
        [Input("totalOcpuCapacity")]
        public Input<int>? TotalOcpuCapacity { get; set; }

        public BlockchainBlockchainPlatformState()
        {
        }
    }
}