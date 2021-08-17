// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Inputs
{

    public sealed class OptimizerProfileTargetTagsItemGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Updatable) The name of the tag definition.
        /// </summary>
        [Input("tagDefinitionName", required: true)]
        public Input<string> TagDefinitionName { get; set; } = null!;

        /// <summary>
        /// (Updatable) The name of the tag namespace.
        /// </summary>
        [Input("tagNamespaceName", required: true)]
        public Input<string> TagNamespaceName { get; set; } = null!;

        /// <summary>
        /// (Updatable) The tag value type.
        /// </summary>
        [Input("tagValueType", required: true)]
        public Input<string> TagValueType { get; set; } = null!;

        [Input("tagValues")]
        private InputList<string>? _tagValues;

        /// <summary>
        /// (Updatable) The list of tag values.
        /// </summary>
        public InputList<string> TagValues
        {
            get => _tagValues ?? (_tagValues = new InputList<string>());
            set => _tagValues = value;
        }

        public OptimizerProfileTargetTagsItemGetArgs()
        {
        }
    }
}