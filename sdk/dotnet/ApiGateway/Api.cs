// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.ApiGateway
{
    /// <summary>
    /// This resource provides the Api resource in Oracle Cloud Infrastructure API Gateway service.
    /// 
    /// Creates a new API.
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
    ///         var testApi = new Oci.ApiGateway.Api("testApi", new Oci.ApiGateway.ApiArgs
    ///         {
    ///             CompartmentId = @var.Compartment_id,
    ///             Content = @var.Api_content,
    ///             DefinedTags = 
    ///             {
    ///                 { "Operations.CostCenter", "42" },
    ///             },
    ///             DisplayName = @var.Api_display_name,
    ///             FreeformTags = 
    ///             {
    ///                 { "Department", "Finance" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Apis can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import oci:apigateway/api:Api test_api "id"
    /// ```
    /// </summary>
    [OciResourceType("oci:apigateway/api:Api")]
    public partial class Api : Pulumi.CustomResource
    {
        /// <summary>
        /// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created.
        /// </summary>
        [Output("compartmentId")]
        public Output<string> CompartmentId { get; private set; } = null!;

        /// <summary>
        /// (Updatable) API Specification content in json or yaml format
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        [Output("definedTags")]
        public Output<ImmutableDictionary<string, object>> DefinedTags { get; private set; } = null!;

        /// <summary>
        /// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource`
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        [Output("freeformTags")]
        public Output<ImmutableDictionary<string, object>> FreeformTags { get; private set; } = null!;

        /// <summary>
        /// A message describing the current lifecycleState in more detail. For ACTIVE state it describes if the document has been validated and the possible values are:
        /// * 'New' for just updated API Specifications
        /// * 'Validating' for a document which is being validated.
        /// * 'Valid' the document has been validated without any errors or warnings
        /// * 'Warning' the document has been validated and contains warnings
        /// * 'Error' the document has been validated and contains errors
        /// * 'Failed' the document validation failed
        /// * 'Canceled' the document validation was canceled
        /// </summary>
        [Output("lifecycleDetails")]
        public Output<string> LifecycleDetails { get; private set; } = null!;

        /// <summary>
        /// Type of API Specification file.
        /// </summary>
        [Output("specificationType")]
        public Output<string> SpecificationType { get; private set; } = null!;

        /// <summary>
        /// The current state of the API.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The time this resource was created. An RFC3339 formatted datetime string.
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// The time this resource was last updated. An RFC3339 formatted datetime string.
        /// </summary>
        [Output("timeUpdated")]
        public Output<string> TimeUpdated { get; private set; } = null!;

        /// <summary>
        /// Status of each feature available from the API.
        /// </summary>
        [Output("validationResults")]
        public Output<ImmutableArray<Outputs.ApiValidationResult>> ValidationResults { get; private set; } = null!;


        /// <summary>
        /// Create a Api resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Api(string name, ApiArgs args, CustomResourceOptions? options = null)
            : base("oci:apigateway/api:Api", name, args ?? new ApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Api(string name, Input<string> id, ApiState? state = null, CustomResourceOptions? options = null)
            : base("oci:apigateway/api:Api", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Api resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Api Get(string name, Input<string> id, ApiState? state = null, CustomResourceOptions? options = null)
        {
            return new Api(name, id, state, options);
        }
    }

    public sealed class ApiArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created.
        /// </summary>
        [Input("compartmentId", required: true)]
        public Input<string> CompartmentId { get; set; } = null!;

        /// <summary>
        /// (Updatable) API Specification content in json or yaml format
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        [Input("definedTags")]
        private InputMap<object>? _definedTags;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        public InputMap<object> DefinedTags
        {
            get => _definedTags ?? (_definedTags = new InputMap<object>());
            set => _definedTags = value;
        }

        /// <summary>
        /// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource`
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("freeformTags")]
        private InputMap<object>? _freeformTags;

        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        public InputMap<object> FreeformTags
        {
            get => _freeformTags ?? (_freeformTags = new InputMap<object>());
            set => _freeformTags = value;
        }

        public ApiArgs()
        {
        }
    }

    public sealed class ApiState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created.
        /// </summary>
        [Input("compartmentId")]
        public Input<string>? CompartmentId { get; set; }

        /// <summary>
        /// (Updatable) API Specification content in json or yaml format
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        [Input("definedTags")]
        private InputMap<object>? _definedTags;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        public InputMap<object> DefinedTags
        {
            get => _definedTags ?? (_definedTags = new InputMap<object>());
            set => _definedTags = value;
        }

        /// <summary>
        /// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource`
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("freeformTags")]
        private InputMap<object>? _freeformTags;

        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        public InputMap<object> FreeformTags
        {
            get => _freeformTags ?? (_freeformTags = new InputMap<object>());
            set => _freeformTags = value;
        }

        /// <summary>
        /// A message describing the current lifecycleState in more detail. For ACTIVE state it describes if the document has been validated and the possible values are:
        /// * 'New' for just updated API Specifications
        /// * 'Validating' for a document which is being validated.
        /// * 'Valid' the document has been validated without any errors or warnings
        /// * 'Warning' the document has been validated and contains warnings
        /// * 'Error' the document has been validated and contains errors
        /// * 'Failed' the document validation failed
        /// * 'Canceled' the document validation was canceled
        /// </summary>
        [Input("lifecycleDetails")]
        public Input<string>? LifecycleDetails { get; set; }

        /// <summary>
        /// Type of API Specification file.
        /// </summary>
        [Input("specificationType")]
        public Input<string>? SpecificationType { get; set; }

        /// <summary>
        /// The current state of the API.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The time this resource was created. An RFC3339 formatted datetime string.
        /// </summary>
        [Input("timeCreated")]
        public Input<string>? TimeCreated { get; set; }

        /// <summary>
        /// The time this resource was last updated. An RFC3339 formatted datetime string.
        /// </summary>
        [Input("timeUpdated")]
        public Input<string>? TimeUpdated { get; set; }

        [Input("validationResults")]
        private InputList<Inputs.ApiValidationResultGetArgs>? _validationResults;

        /// <summary>
        /// Status of each feature available from the API.
        /// </summary>
        public InputList<Inputs.ApiValidationResultGetArgs> ValidationResults
        {
            get => _validationResults ?? (_validationResults = new InputList<Inputs.ApiValidationResultGetArgs>());
            set => _validationResults = value;
        }

        public ApiState()
        {
        }
    }
}