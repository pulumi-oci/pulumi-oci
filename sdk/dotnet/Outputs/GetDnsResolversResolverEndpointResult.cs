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
    public sealed class GetDnsResolversResolverEndpointResult
    {
        /// <summary>
        /// The OCID of the compartment the resource belongs to.
        /// </summary>
        public readonly string CompartmentId;
        public readonly string EndpointType;
        public readonly string ForwardingAddress;
        public readonly bool IsForwarding;
        public readonly bool IsListening;
        public readonly string ListeningAddress;
        public readonly string Name;
        /// <summary>
        /// The canonical absolute URL of the resource.
        /// </summary>
        public readonly string Self;
        /// <summary>
        /// The state of a resource.
        /// </summary>
        public readonly string State;
        public readonly string SubnetId;
        /// <summary>
        /// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.
        /// </summary>
        public readonly string TimeUpdated;

        [OutputConstructor]
        private GetDnsResolversResolverEndpointResult(
            string compartmentId,

            string endpointType,

            string forwardingAddress,

            bool isForwarding,

            bool isListening,

            string listeningAddress,

            string name,

            string self,

            string state,

            string subnetId,

            string timeCreated,

            string timeUpdated)
        {
            CompartmentId = compartmentId;
            EndpointType = endpointType;
            ForwardingAddress = forwardingAddress;
            IsForwarding = isForwarding;
            IsListening = isListening;
            ListeningAddress = listeningAddress;
            Name = name;
            Self = self;
            State = state;
            SubnetId = subnetId;
            TimeCreated = timeCreated;
            TimeUpdated = timeUpdated;
        }
    }
}