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
    public sealed class GetDevopsDeploymentsDeploymentCollectionItemDeploymentArgumentsItemResult
    {
        /// <summary>
        /// Name of the step.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// value of the argument.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetDevopsDeploymentsDeploymentCollectionItemDeploymentArgumentsItemResult(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}