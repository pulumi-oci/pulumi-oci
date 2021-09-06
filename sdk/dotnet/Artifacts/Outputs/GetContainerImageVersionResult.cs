// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Artifacts.Outputs
{

    [OutputType]
    public sealed class GetContainerImageVersionResult
    {
        /// <summary>
        /// The OCID of the user or principal that pushed the version.
        /// </summary>
        public readonly string CreatedBy;
        /// <summary>
        /// The creation time of the version.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The version name.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetContainerImageVersionResult(
            string createdBy,

            string timeCreated,

            string version)
        {
            CreatedBy = createdBy;
            TimeCreated = timeCreated;
            Version = version;
        }
    }
}