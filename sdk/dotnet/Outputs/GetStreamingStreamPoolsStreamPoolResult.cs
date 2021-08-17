// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Outputs
{

    [OutputType]
    public sealed class GetStreamingStreamPoolsStreamPoolResult
    {
        /// <summary>
        /// The OCID of the compartment.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// Custom Encryption Key which will be used for encryption by all the streams in the pool.
        /// </summary>
        public readonly Outputs.GetStreamingStreamPoolsStreamPoolCustomEncryptionKeyResult CustomEncryptionKey;
        /// <summary>
        /// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations": {"CostCenter": "42"}}'
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// The FQDN used to access the streams inside the stream pool (same FQDN as the messagesEndpoint attribute of a [Stream](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/Stream) object). If the stream pool is private, the FQDN is customized and can only be accessed from inside the associated subnetId, otherwise the FQDN is publicly resolvable. Depending on which protocol you attempt to use, you need to either prepend https or append the Kafka port.
        /// </summary>
        public readonly string EndpointFqdn;
        /// <summary>
        /// Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// A filter to return only resources that match the given ID exactly.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// True if the stream pool is private, false otherwise. The associated endpoint and subnetId of a private stream pool can be retrieved through the [GetStreamPool](https://docs.cloud.oracle.com/iaas/api/#/en/streaming/20180418/StreamPool/GetStreamPool) API.
        /// </summary>
        public readonly bool IsPrivate;
        /// <summary>
        /// Settings for the Kafka compatibility layer.
        /// </summary>
        public readonly Outputs.GetStreamingStreamPoolsStreamPoolKafkaSettingsResult KafkaSettings;
        /// <summary>
        /// Any additional details about the current state of the stream.
        /// </summary>
        public readonly string LifecycleStateDetails;
        /// <summary>
        /// A filter to return only resources that match the given name exactly.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional settings if the stream pool is private.
        /// </summary>
        public readonly Outputs.GetStreamingStreamPoolsStreamPoolPrivateEndpointSettingsResult PrivateEndpointSettings;
        /// <summary>
        /// A filter to only return resources that match the given lifecycle state. The state value is case-insensitive.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The date and time the stream pool was created, expressed in in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2018-04-20T00:00:07.405Z`
        /// </summary>
        public readonly string TimeCreated;

        [OutputConstructor]
        private GetStreamingStreamPoolsStreamPoolResult(
            string compartmentId,

            Outputs.GetStreamingStreamPoolsStreamPoolCustomEncryptionKeyResult customEncryptionKey,

            ImmutableDictionary<string, object> definedTags,

            string endpointFqdn,

            ImmutableDictionary<string, object> freeformTags,

            string id,

            bool isPrivate,

            Outputs.GetStreamingStreamPoolsStreamPoolKafkaSettingsResult kafkaSettings,

            string lifecycleStateDetails,

            string name,

            Outputs.GetStreamingStreamPoolsStreamPoolPrivateEndpointSettingsResult privateEndpointSettings,

            string state,

            string timeCreated)
        {
            CompartmentId = compartmentId;
            CustomEncryptionKey = customEncryptionKey;
            DefinedTags = definedTags;
            EndpointFqdn = endpointFqdn;
            FreeformTags = freeformTags;
            Id = id;
            IsPrivate = isPrivate;
            KafkaSettings = kafkaSettings;
            LifecycleStateDetails = lifecycleStateDetails;
            Name = name;
            PrivateEndpointSettings = privateEndpointSettings;
            State = state;
            TimeCreated = timeCreated;
        }
    }
}