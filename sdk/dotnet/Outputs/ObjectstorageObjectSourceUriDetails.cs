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
    public sealed class ObjectstorageObjectSourceUriDetails
    {
        /// <summary>
        /// The name of the bucket for the source object.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// The entity tag to match the target object.
        /// </summary>
        public readonly string? DestinationObjectIfMatchEtag;
        /// <summary>
        /// The entity tag to not match the target object.
        /// </summary>
        public readonly string? DestinationObjectIfNoneMatchEtag;
        /// <summary>
        /// The top-level namespace of the source object.
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// The name of the source object.
        /// </summary>
        public readonly string Object;
        /// <summary>
        /// The region of the source object.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The entity tag to match the source object.
        /// </summary>
        public readonly string? SourceObjectIfMatchEtag;
        /// <summary>
        /// The version id of the object to be restored.
        /// </summary>
        public readonly string? SourceVersionId;

        [OutputConstructor]
        private ObjectstorageObjectSourceUriDetails(
            string bucket,

            string? destinationObjectIfMatchEtag,

            string? destinationObjectIfNoneMatchEtag,

            string @namespace,

            string @object,

            string region,

            string? sourceObjectIfMatchEtag,

            string? sourceVersionId)
        {
            Bucket = bucket;
            DestinationObjectIfMatchEtag = destinationObjectIfMatchEtag;
            DestinationObjectIfNoneMatchEtag = destinationObjectIfNoneMatchEtag;
            Namespace = @namespace;
            Object = @object;
            Region = region;
            SourceObjectIfMatchEtag = sourceObjectIfMatchEtag;
            SourceVersionId = sourceVersionId;
        }
    }
}