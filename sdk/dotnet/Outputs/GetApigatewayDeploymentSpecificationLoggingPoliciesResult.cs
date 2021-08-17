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
    public sealed class GetApigatewayDeploymentSpecificationLoggingPoliciesResult
    {
        /// <summary>
        /// Configures the logging policies for the access logs of an API Deployment.
        /// </summary>
        public readonly Outputs.GetApigatewayDeploymentSpecificationLoggingPoliciesAccessLogResult AccessLog;
        /// <summary>
        /// Configures the logging policies for the execution logs of an API Deployment.
        /// </summary>
        public readonly Outputs.GetApigatewayDeploymentSpecificationLoggingPoliciesExecutionLogResult ExecutionLog;

        [OutputConstructor]
        private GetApigatewayDeploymentSpecificationLoggingPoliciesResult(
            Outputs.GetApigatewayDeploymentSpecificationLoggingPoliciesAccessLogResult accessLog,

            Outputs.GetApigatewayDeploymentSpecificationLoggingPoliciesExecutionLogResult executionLog)
        {
            AccessLog = accessLog;
            ExecutionLog = executionLog;
        }
    }
}