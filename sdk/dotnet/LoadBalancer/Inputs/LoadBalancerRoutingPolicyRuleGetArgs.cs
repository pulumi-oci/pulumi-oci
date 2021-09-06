// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.LoadBalancer.Inputs
{

    public sealed class LoadBalancerRoutingPolicyRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<Inputs.LoadBalancerRoutingPolicyRuleActionGetArgs>? _actions;

        /// <summary>
        /// (Updatable) A list of actions to be applied when conditions of the routing rule are met.
        /// </summary>
        public InputList<Inputs.LoadBalancerRoutingPolicyRuleActionGetArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.LoadBalancerRoutingPolicyRuleActionGetArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// (Updatable) A routing rule to evaluate defined conditions against the incoming HTTP request and perform an action.
        /// </summary>
        [Input("condition", required: true)]
        public Input<string> Condition { get; set; } = null!;

        /// <summary>
        /// (Updatable) A unique name for the routing policy rule. Avoid entering confidential information.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public LoadBalancerRoutingPolicyRuleGetArgs()
        {
        }
    }
}