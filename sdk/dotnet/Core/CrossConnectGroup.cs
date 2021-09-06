// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Core
{
    /// <summary>
    /// This resource provides the Cross Connect Group resource in Oracle Cloud Infrastructure Core service.
    /// 
    /// Creates a new cross-connect group to use with Oracle Cloud Infrastructure
    /// FastConnect. For more information, see
    /// [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).
    /// 
    /// For the purposes of access control, you must provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the
    /// compartment where you want the cross-connect group to reside. If you're
    /// not sure which compartment to use, put the cross-connect group in the
    /// same compartment with your VCN. For more information about
    /// compartments and access control, see
    /// [Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
    /// For information about OCIDs, see
    /// [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
    /// 
    /// You may optionally specify a *display name* for the cross-connect group.
    /// It does not have to be unique, and you can change it. Avoid entering confidential information.
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
    ///         var testCrossConnectGroup = new Oci.Core.CrossConnectGroup("testCrossConnectGroup", new Oci.Core.CrossConnectGroupArgs
    ///         {
    ///             CompartmentId = @var.Compartment_id,
    ///             CustomerReferenceName = @var.Cross_connect_group_customer_reference_name,
    ///             DefinedTags = 
    ///             {
    ///                 { "Operations.CostCenter", "42" },
    ///             },
    ///             DisplayName = @var.Cross_connect_group_display_name,
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
    /// CrossConnectGroups can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import oci:core/crossConnectGroup:CrossConnectGroup test_cross_connect_group "id"
    /// ```
    /// </summary>
    [OciResourceType("oci:core/crossConnectGroup:CrossConnectGroup")]
    public partial class CrossConnectGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// (Updatable) The OCID of the compartment to contain the cross-connect group.
        /// </summary>
        [Output("compartmentId")]
        public Output<string> CompartmentId { get; private set; } = null!;

        /// <summary>
        /// (Updatable) A reference name or identifier for the physical fiber connection that this cross-connect group uses.
        /// </summary>
        [Output("customerReferenceName")]
        public Output<string> CustomerReferenceName { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        [Output("definedTags")]
        public Output<ImmutableDictionary<string, object>> DefinedTags { get; private set; } = null!;

        /// <summary>
        /// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        [Output("freeformTags")]
        public Output<ImmutableDictionary<string, object>> FreeformTags { get; private set; } = null!;

        /// <summary>
        /// The cross-connect group's current state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The date and time the cross-connect group was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;


        /// <summary>
        /// Create a CrossConnectGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CrossConnectGroup(string name, CrossConnectGroupArgs args, CustomResourceOptions? options = null)
            : base("oci:core/crossConnectGroup:CrossConnectGroup", name, args ?? new CrossConnectGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CrossConnectGroup(string name, Input<string> id, CrossConnectGroupState? state = null, CustomResourceOptions? options = null)
            : base("oci:core/crossConnectGroup:CrossConnectGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CrossConnectGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CrossConnectGroup Get(string name, Input<string> id, CrossConnectGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new CrossConnectGroup(name, id, state, options);
        }
    }

    public sealed class CrossConnectGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Updatable) The OCID of the compartment to contain the cross-connect group.
        /// </summary>
        [Input("compartmentId", required: true)]
        public Input<string> CompartmentId { get; set; } = null!;

        /// <summary>
        /// (Updatable) A reference name or identifier for the physical fiber connection that this cross-connect group uses.
        /// </summary>
        [Input("customerReferenceName")]
        public Input<string>? CustomerReferenceName { get; set; }

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
        /// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
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

        public CrossConnectGroupArgs()
        {
        }
    }

    public sealed class CrossConnectGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Updatable) The OCID of the compartment to contain the cross-connect group.
        /// </summary>
        [Input("compartmentId")]
        public Input<string>? CompartmentId { get; set; }

        /// <summary>
        /// (Updatable) A reference name or identifier for the physical fiber connection that this cross-connect group uses.
        /// </summary>
        [Input("customerReferenceName")]
        public Input<string>? CustomerReferenceName { get; set; }

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
        /// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
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
        /// The cross-connect group's current state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The date and time the cross-connect group was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
        /// </summary>
        [Input("timeCreated")]
        public Input<string>? TimeCreated { get; set; }

        public CrossConnectGroupState()
        {
        }
    }
}