// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    /// <summary>
    /// This resource provides the Oda Instance resource in Oracle Cloud Infrastructure Digital Assistant service.
    /// 
    /// Starts an asynchronous job to create a Digital Assistant instance.
    /// 
    /// To monitor the status of the job, take the `opc-work-request-id` response
    /// header value and use it to call `GET /workRequests/{workRequestID}`.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Oci = Pulumi.Oci;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testOdaInstance = new Oci.OdaOdaInstance("testOdaInstance", new Oci.OdaOdaInstanceArgs
    ///         {
    ///             CompartmentId = @var.Compartment_id,
    ///             ShapeName = "DEVELOPMENT",
    ///             DefinedTags = 
    ///             {
    ///                 { "foo-namespace.bar-key", "value" },
    ///             },
    ///             Description = @var.Oda_instance_description,
    ///             DisplayName = @var.Oda_instance_display_name,
    ///             FreeformTags = 
    ///             {
    ///                 { "bar-key", "value" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// OdaInstances can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import oci:index/odaOdaInstance:OdaOdaInstance test_oda_instance "id"
    /// ```
    /// </summary>
    [OciResourceType("oci:index/odaOdaInstance:OdaOdaInstance")]
    public partial class OdaOdaInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// (Updatable) Identifier of the compartment.
        /// </summary>
        [Output("compartmentId")]
        public Output<string> CompartmentId { get; private set; } = null!;

        /// <summary>
        /// URL for the connector's endpoint.
        /// </summary>
        [Output("connectorUrl")]
        public Output<string> ConnectorUrl { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}`
        /// </summary>
        [Output("definedTags")]
        public Output<ImmutableDictionary<string, object>> DefinedTags { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Description of the Digital Assistant instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// (Updatable) User-friendly name for the instance. Avoid entering confidential information. You can change this value anytime.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
        /// </summary>
        [Output("freeformTags")]
        public Output<ImmutableDictionary<string, object>> FreeformTags { get; private set; } = null!;

        /// <summary>
        /// The current sub-state of the Digital Assistant instance.
        /// </summary>
        [Output("lifecycleSubState")]
        public Output<string> LifecycleSubState { get; private set; } = null!;

        /// <summary>
        /// Shape or size of the instance.
        /// </summary>
        [Output("shapeName")]
        public Output<string> ShapeName { get; private set; } = null!;

        /// <summary>
        /// The current state of the Digital Assistant instance.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// A message that describes the current state in more detail. For example, actionable information about an instance that's in the `FAILED` state.
        /// </summary>
        [Output("stateMessage")]
        public Output<string> StateMessage { get; private set; } = null!;

        /// <summary>
        /// When the Digital Assistant instance was created. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// When the Digital Assistance instance was last updated. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
        /// </summary>
        [Output("timeUpdated")]
        public Output<string> TimeUpdated { get; private set; } = null!;

        /// <summary>
        /// URL for the Digital Assistant web application that's associated with the instance.
        /// </summary>
        [Output("webAppUrl")]
        public Output<string> WebAppUrl { get; private set; } = null!;


        /// <summary>
        /// Create a OdaOdaInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OdaOdaInstance(string name, OdaOdaInstanceArgs args, CustomResourceOptions? options = null)
            : base("oci:index/odaOdaInstance:OdaOdaInstance", name, args ?? new OdaOdaInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OdaOdaInstance(string name, Input<string> id, OdaOdaInstanceState? state = null, CustomResourceOptions? options = null)
            : base("oci:index/odaOdaInstance:OdaOdaInstance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OdaOdaInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OdaOdaInstance Get(string name, Input<string> id, OdaOdaInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new OdaOdaInstance(name, id, state, options);
        }
    }

    public sealed class OdaOdaInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Updatable) Identifier of the compartment.
        /// </summary>
        [Input("compartmentId", required: true)]
        public Input<string> CompartmentId { get; set; } = null!;

        [Input("definedTags")]
        private InputMap<object>? _definedTags;

        /// <summary>
        /// (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}`
        /// </summary>
        public InputMap<object> DefinedTags
        {
            get => _definedTags ?? (_definedTags = new InputMap<object>());
            set => _definedTags = value;
        }

        /// <summary>
        /// (Updatable) Description of the Digital Assistant instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (Updatable) User-friendly name for the instance. Avoid entering confidential information. You can change this value anytime.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("freeformTags")]
        private InputMap<object>? _freeformTags;

        /// <summary>
        /// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
        /// </summary>
        public InputMap<object> FreeformTags
        {
            get => _freeformTags ?? (_freeformTags = new InputMap<object>());
            set => _freeformTags = value;
        }

        /// <summary>
        /// Shape or size of the instance.
        /// </summary>
        [Input("shapeName", required: true)]
        public Input<string> ShapeName { get; set; } = null!;

        /// <summary>
        /// The current state of the Digital Assistant instance.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public OdaOdaInstanceArgs()
        {
        }
    }

    public sealed class OdaOdaInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Updatable) Identifier of the compartment.
        /// </summary>
        [Input("compartmentId")]
        public Input<string>? CompartmentId { get; set; }

        /// <summary>
        /// URL for the connector's endpoint.
        /// </summary>
        [Input("connectorUrl")]
        public Input<string>? ConnectorUrl { get; set; }

        [Input("definedTags")]
        private InputMap<object>? _definedTags;

        /// <summary>
        /// (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}`
        /// </summary>
        public InputMap<object> DefinedTags
        {
            get => _definedTags ?? (_definedTags = new InputMap<object>());
            set => _definedTags = value;
        }

        /// <summary>
        /// (Updatable) Description of the Digital Assistant instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (Updatable) User-friendly name for the instance. Avoid entering confidential information. You can change this value anytime.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("freeformTags")]
        private InputMap<object>? _freeformTags;

        /// <summary>
        /// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
        /// </summary>
        public InputMap<object> FreeformTags
        {
            get => _freeformTags ?? (_freeformTags = new InputMap<object>());
            set => _freeformTags = value;
        }

        /// <summary>
        /// The current sub-state of the Digital Assistant instance.
        /// </summary>
        [Input("lifecycleSubState")]
        public Input<string>? LifecycleSubState { get; set; }

        /// <summary>
        /// Shape or size of the instance.
        /// </summary>
        [Input("shapeName")]
        public Input<string>? ShapeName { get; set; }

        /// <summary>
        /// The current state of the Digital Assistant instance.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// A message that describes the current state in more detail. For example, actionable information about an instance that's in the `FAILED` state.
        /// </summary>
        [Input("stateMessage")]
        public Input<string>? StateMessage { get; set; }

        /// <summary>
        /// When the Digital Assistant instance was created. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
        /// </summary>
        [Input("timeCreated")]
        public Input<string>? TimeCreated { get; set; }

        /// <summary>
        /// When the Digital Assistance instance was last updated. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
        /// </summary>
        [Input("timeUpdated")]
        public Input<string>? TimeUpdated { get; set; }

        /// <summary>
        /// URL for the Digital Assistant web application that's associated with the instance.
        /// </summary>
        [Input("webAppUrl")]
        public Input<string>? WebAppUrl { get; set; }

        public OdaOdaInstanceState()
        {
        }
    }
}