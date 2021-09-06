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
    public sealed class GetAutonomousExadataInfrastructureOcpuByWorkloadTypeResult
    {
        /// <summary>
        /// The total number of OCPU cores in use for Autonomous Data Warehouse databases in the infrastructure instance.
        /// </summary>
        public readonly double Adw;
        /// <summary>
        /// The total number of OCPU cores in use for Autonomous Transaction Processing databases in the infrastructure instance.
        /// </summary>
        public readonly double Atp;

        [OutputConstructor]
        private GetAutonomousExadataInfrastructureOcpuByWorkloadTypeResult(
            double adw,

            double atp)
        {
            Adw = adw;
            Atp = atp;
        }
    }
}