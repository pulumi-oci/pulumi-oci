// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Marketplace.Outputs
{

    [OutputType]
    public sealed class GetPublicationSupportedOperatingSystemResult
    {
        /// <summary>
        /// name of the operating system
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetPublicationSupportedOperatingSystemResult(string name)
        {
            Name = name;
        }
    }
}