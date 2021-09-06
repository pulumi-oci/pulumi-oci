// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Bds.Outputs
{

    [OutputType]
    public sealed class GetAutoScalingConfigurationPolicyRuleMetricThresholdResult
    {
        /// <summary>
        /// This value is the minimum period of time the metric value meets or exceeds the threshold value before the action is triggered. The value is in minutes.
        /// </summary>
        public readonly int DurationInMinutes;
        /// <summary>
        /// The comparison operator to use. Options are greater than (GT) or less than (LT).
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// Integer non-negative value. 0 &lt; value &lt; 100
        /// </summary>
        public readonly int Value;

        [OutputConstructor]
        private GetAutoScalingConfigurationPolicyRuleMetricThresholdResult(
            int durationInMinutes,

            string @operator,

            int value)
        {
            DurationInMinutes = durationInMinutes;
            Operator = @operator;
            Value = value;
        }
    }
}