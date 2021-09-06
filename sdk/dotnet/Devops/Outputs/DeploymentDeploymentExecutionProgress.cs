// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Devops.Outputs
{

    [OutputType]
    public sealed class DeploymentDeploymentExecutionProgress
    {
        /// <summary>
        /// Map of stage OCIDs to deploy stage execution progress model.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? DeployStageExecutionProgress;
        /// <summary>
        /// Time the deployment is finished. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
        /// </summary>
        public readonly string? TimeFinished;
        /// <summary>
        /// Time the deployment is started. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
        /// </summary>
        public readonly string? TimeStarted;

        [OutputConstructor]
        private DeploymentDeploymentExecutionProgress(
            ImmutableDictionary<string, object>? deployStageExecutionProgress,

            string? timeFinished,

            string? timeStarted)
        {
            DeployStageExecutionProgress = deployStageExecutionProgress;
            TimeFinished = timeFinished;
            TimeStarted = timeStarted;
        }
    }
}