// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Inputs
{

    public sealed class BdsAutoScalingConfigurationPolicyRuleMetricArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Updatable) Allowed value is CPU_UTILIZATION.
        /// </summary>
        [Input("metricType", required: true)]
        public Input<string> MetricType { get; set; } = null!;

        /// <summary>
        /// (Updatable) An autoscale action is triggered when a performance metric meets or exceeds a threshold.
        /// </summary>
        [Input("threshold", required: true)]
        public Input<Inputs.BdsAutoScalingConfigurationPolicyRuleMetricThresholdArgs> Threshold { get; set; } = null!;

        public BdsAutoScalingConfigurationPolicyRuleMetricArgs()
        {
        }
    }
}