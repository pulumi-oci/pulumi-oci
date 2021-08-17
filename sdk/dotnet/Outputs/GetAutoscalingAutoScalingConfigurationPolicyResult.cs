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
    public sealed class GetAutoscalingAutoScalingConfigurationPolicyResult
    {
        /// <summary>
        /// The capacity requirements of the autoscaling policy.
        /// </summary>
        public readonly Outputs.GetAutoscalingAutoScalingConfigurationPolicyCapacityResult Capacity;
        /// <summary>
        /// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The schedule for executing the autoscaling policy.
        /// </summary>
        public readonly Outputs.GetAutoscalingAutoScalingConfigurationPolicyExecutionScheduleResult ExecutionSchedule;
        /// <summary>
        /// ID of the condition that is assigned after creation.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether the autoscaling policy is enabled.
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// The type of autoscaling policy.
        /// </summary>
        public readonly string PolicyType;
        /// <summary>
        /// An action that can be executed against a resource.
        /// </summary>
        public readonly Outputs.GetAutoscalingAutoScalingConfigurationPolicyResourceActionResult ResourceAction;
        public readonly ImmutableArray<Outputs.GetAutoscalingAutoScalingConfigurationPolicyRuleResult> Rules;
        /// <summary>
        /// The date and time the autoscaling configuration was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
        /// </summary>
        public readonly string TimeCreated;

        [OutputConstructor]
        private GetAutoscalingAutoScalingConfigurationPolicyResult(
            Outputs.GetAutoscalingAutoScalingConfigurationPolicyCapacityResult capacity,

            string displayName,

            Outputs.GetAutoscalingAutoScalingConfigurationPolicyExecutionScheduleResult executionSchedule,

            string id,

            bool isEnabled,

            string policyType,

            Outputs.GetAutoscalingAutoScalingConfigurationPolicyResourceActionResult resourceAction,

            ImmutableArray<Outputs.GetAutoscalingAutoScalingConfigurationPolicyRuleResult> rules,

            string timeCreated)
        {
            Capacity = capacity;
            DisplayName = displayName;
            ExecutionSchedule = executionSchedule;
            Id = id;
            IsEnabled = isEnabled;
            PolicyType = policyType;
            ResourceAction = resourceAction;
            Rules = rules;
            TimeCreated = timeCreated;
        }
    }
}