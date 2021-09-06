// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Sch.Outputs
{

    [OutputType]
    public sealed class GetServiceConnectorsServiceConnectorCollectionItemSourceResult
    {
        /// <summary>
        /// The type of [cursor](https://docs.cloud.oracle.com/iaas/Content/Streaming/Tasks/using_a_single_consumer.htm#usingcursors), which determines the starting point from which the stream will be consumed.
        /// </summary>
        public readonly Outputs.GetServiceConnectorsServiceConnectorCollectionItemSourceCursorResult Cursor;
        /// <summary>
        /// The type descriminator.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The resources affected by this work request.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceConnectorsServiceConnectorCollectionItemSourceLogSourceResult> LogSources;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream.
        /// </summary>
        public readonly string StreamId;

        [OutputConstructor]
        private GetServiceConnectorsServiceConnectorCollectionItemSourceResult(
            Outputs.GetServiceConnectorsServiceConnectorCollectionItemSourceCursorResult cursor,

            string kind,

            ImmutableArray<Outputs.GetServiceConnectorsServiceConnectorCollectionItemSourceLogSourceResult> logSources,

            string streamId)
        {
            Cursor = cursor;
            Kind = kind;
            LogSources = logSources;
            StreamId = streamId;
        }
    }
}