// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.ApmSynthetics.Inputs
{

    public sealed class ScriptParameterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If parameter value is default or overwritten.
        /// </summary>
        [Input("isOverwritten")]
        public Input<bool>? IsOverwritten { get; set; }

        /// <summary>
        /// (Updatable) If the parameter value is secret and should be kept confidential, then set isSecret to true.
        /// </summary>
        [Input("isSecret")]
        public Input<bool>? IsSecret { get; set; }

        /// <summary>
        /// (Updatable) Name of the parameter.
        /// </summary>
        [Input("paramName", required: true)]
        public Input<string> ParamName { get; set; } = null!;

        /// <summary>
        /// (Updatable) Value of the parameter.
        /// </summary>
        [Input("paramValue")]
        public Input<string>? ParamValue { get; set; }

        /// <summary>
        /// Details of the script parameters, paramName must be from the script content and these details can be used to overwrite the default parameter present in the script content.
        /// </summary>
        [Input("scriptParameter")]
        public Input<Inputs.ScriptParameterScriptParameterArgs>? ScriptParameter { get; set; }

        public ScriptParameterArgs()
        {
        }
    }
}