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
    public sealed class GetDatascienceModelDeploymentShapesModelDeploymentShapeResult
    {
        /// <summary>
        /// The number of cores associated with this model deployment shape.
        /// </summary>
        public readonly int CoreCount;
        /// <summary>
        /// The amount of memory in GBs associated with this model deployment shape.
        /// </summary>
        public readonly int MemoryInGbs;
        /// <summary>
        /// The name of the model deployment shape.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetDatascienceModelDeploymentShapesModelDeploymentShapeResult(
            int coreCount,

            int memoryInGbs,

            string name)
        {
            CoreCount = coreCount;
            MemoryInGbs = memoryInGbs;
            Name = name;
        }
    }
}