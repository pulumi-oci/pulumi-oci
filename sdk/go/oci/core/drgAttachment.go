// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Drg Attachment resource in Oracle Cloud Infrastructure Core service.
//
// Attaches the specified DRG to the specified network resource. A VCN can be attached to only one DRG
// at a time, but a DRG can be attached to more than one VCN. The response includes a `DrgAttachment`
// object with its own [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). For more information about DRGs, see
// [Dynamic Routing Gateways (DRGs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingDRGs.htm).
//
// You may optionally specify a *display name* for the attachment, otherwise a default is provided.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
//
// For the purposes of access control, the DRG attachment is automatically placed into the currently selected compartment.
// For more information about compartments and access control, see
// [Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.NewDrgAttachment(ctx, "testDrgAttachment", &core.DrgAttachmentArgs{
// 			DrgId: pulumi.Any(oci_core_drg.Test_drg.Id),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			DisplayName:     pulumi.Any(_var.Drg_attachment_display_name),
// 			DrgRouteTableId: pulumi.Any(oci_core_drg_route_table.Test_drg_route_table.Id),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			NetworkDetails: &core.DrgAttachmentNetworkDetailsArgs{
// 				Id:           pulumi.Any(_var.Drg_attachment_network_details_id),
// 				Type:         pulumi.Any(_var.Drg_attachment_network_details_type),
// 				RouteTableId: pulumi.Any(oci_core_route_table.Test_route_table.Id),
// 			},
// 			RouteTableId: pulumi.Any(oci_core_route_table.Test_route_table.Id),
// 			VcnId:        pulumi.Any(oci_core_vcn.Test_vcn.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// DrgAttachments can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:core/drgAttachment:DrgAttachment test_drg_attachment "id"
// ```
type DrgAttachment struct {
	pulumi.CustomResourceState

	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the DRG attachment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique. Avoid entering confidential information.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	DrgId pulumi.StringOutput `pulumi:"drgId"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table that is assigned to this attachment.
	DrgRouteTableId pulumi.StringOutput `pulumi:"drgRouteTableId"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export route distribution used to specify how routes in the assigned DRG route table are advertised to the attachment. If this value is null, no routes are advertised through this attachment.
	// This field cannot be set by the user while creating the resource and gets a default value on creation. This can be only be updated to its default value. If this fields needs to be set to null, removeExportDrgRouteDistributionTrigger needs to be used.
	ExportDrgRouteDistributionId pulumi.StringOutput `pulumi:"exportDrgRouteDistributionId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// Indicates whether the DRG attachment and attached network live in a different tenancy than the DRG.  Example: `false`
	IsCrossTenancy pulumi.BoolOutput `pulumi:"isCrossTenancy"`
	// (Updatable)
	NetworkDetails DrgAttachmentNetworkDetailsOutput `pulumi:"networkDetails"`
	// (Updatable) An optional property when set to true during update disables the export of route Distribution by setting exportDrgRouteDistributionId to null.
	RemoveExportDrgRouteDistributionTrigger pulumi.BoolPtrOutput `pulumi:"removeExportDrgRouteDistributionTrigger"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table used by the DRG attachment.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// The DRG attachment's current state.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the DRG attachment was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN. This field is deprecated. Instead, use the `networkDetails` field to specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached resource.
	VcnId pulumi.StringOutput `pulumi:"vcnId"`
}

// NewDrgAttachment registers a new resource with the given unique name, arguments, and options.
func NewDrgAttachment(ctx *pulumi.Context,
	name string, args *DrgAttachmentArgs, opts ...pulumi.ResourceOption) (*DrgAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DrgId == nil {
		return nil, errors.New("invalid value for required argument 'DrgId'")
	}
	var resource DrgAttachment
	err := ctx.RegisterResource("oci:core/drgAttachment:DrgAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDrgAttachment gets an existing DrgAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDrgAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DrgAttachmentState, opts ...pulumi.ResourceOption) (*DrgAttachment, error) {
	var resource DrgAttachment
	err := ctx.ReadResource("oci:core/drgAttachment:DrgAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DrgAttachment resources.
type drgAttachmentState struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the DRG attachment.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	DrgId *string `pulumi:"drgId"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table that is assigned to this attachment.
	DrgRouteTableId *string `pulumi:"drgRouteTableId"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export route distribution used to specify how routes in the assigned DRG route table are advertised to the attachment. If this value is null, no routes are advertised through this attachment.
	// This field cannot be set by the user while creating the resource and gets a default value on creation. This can be only be updated to its default value. If this fields needs to be set to null, removeExportDrgRouteDistributionTrigger needs to be used.
	ExportDrgRouteDistributionId *string `pulumi:"exportDrgRouteDistributionId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Indicates whether the DRG attachment and attached network live in a different tenancy than the DRG.  Example: `false`
	IsCrossTenancy *bool `pulumi:"isCrossTenancy"`
	// (Updatable)
	NetworkDetails *DrgAttachmentNetworkDetails `pulumi:"networkDetails"`
	// (Updatable) An optional property when set to true during update disables the export of route Distribution by setting exportDrgRouteDistributionId to null.
	RemoveExportDrgRouteDistributionTrigger *bool `pulumi:"removeExportDrgRouteDistributionTrigger"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table used by the DRG attachment.
	RouteTableId *string `pulumi:"routeTableId"`
	// The DRG attachment's current state.
	State *string `pulumi:"state"`
	// The date and time the DRG attachment was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *string `pulumi:"timeCreated"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN. This field is deprecated. Instead, use the `networkDetails` field to specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached resource.
	VcnId *string `pulumi:"vcnId"`
}

type DrgAttachmentState struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the DRG attachment.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. Does not have to be unique. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	DrgId pulumi.StringPtrInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table that is assigned to this attachment.
	DrgRouteTableId pulumi.StringPtrInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export route distribution used to specify how routes in the assigned DRG route table are advertised to the attachment. If this value is null, no routes are advertised through this attachment.
	// This field cannot be set by the user while creating the resource and gets a default value on creation. This can be only be updated to its default value. If this fields needs to be set to null, removeExportDrgRouteDistributionTrigger needs to be used.
	ExportDrgRouteDistributionId pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// Indicates whether the DRG attachment and attached network live in a different tenancy than the DRG.  Example: `false`
	IsCrossTenancy pulumi.BoolPtrInput
	// (Updatable)
	NetworkDetails DrgAttachmentNetworkDetailsPtrInput
	// (Updatable) An optional property when set to true during update disables the export of route Distribution by setting exportDrgRouteDistributionId to null.
	RemoveExportDrgRouteDistributionTrigger pulumi.BoolPtrInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table used by the DRG attachment.
	RouteTableId pulumi.StringPtrInput
	// The DRG attachment's current state.
	State pulumi.StringPtrInput
	// The date and time the DRG attachment was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN. This field is deprecated. Instead, use the `networkDetails` field to specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached resource.
	VcnId pulumi.StringPtrInput
}

func (DrgAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*drgAttachmentState)(nil)).Elem()
}

type drgAttachmentArgs struct {
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	DrgId string `pulumi:"drgId"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table that is assigned to this attachment.
	DrgRouteTableId *string `pulumi:"drgRouteTableId"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export route distribution used to specify how routes in the assigned DRG route table are advertised to the attachment. If this value is null, no routes are advertised through this attachment.
	// This field cannot be set by the user while creating the resource and gets a default value on creation. This can be only be updated to its default value. If this fields needs to be set to null, removeExportDrgRouteDistributionTrigger needs to be used.
	ExportDrgRouteDistributionId *string `pulumi:"exportDrgRouteDistributionId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable)
	NetworkDetails *DrgAttachmentNetworkDetails `pulumi:"networkDetails"`
	// (Updatable) An optional property when set to true during update disables the export of route Distribution by setting exportDrgRouteDistributionId to null.
	RemoveExportDrgRouteDistributionTrigger *bool `pulumi:"removeExportDrgRouteDistributionTrigger"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table used by the DRG attachment.
	RouteTableId *string `pulumi:"routeTableId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN. This field is deprecated. Instead, use the `networkDetails` field to specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached resource.
	VcnId *string `pulumi:"vcnId"`
}

// The set of arguments for constructing a DrgAttachment resource.
type DrgAttachmentArgs struct {
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. Does not have to be unique. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	DrgId pulumi.StringInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table that is assigned to this attachment.
	DrgRouteTableId pulumi.StringPtrInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export route distribution used to specify how routes in the assigned DRG route table are advertised to the attachment. If this value is null, no routes are advertised through this attachment.
	// This field cannot be set by the user while creating the resource and gets a default value on creation. This can be only be updated to its default value. If this fields needs to be set to null, removeExportDrgRouteDistributionTrigger needs to be used.
	ExportDrgRouteDistributionId pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable)
	NetworkDetails DrgAttachmentNetworkDetailsPtrInput
	// (Updatable) An optional property when set to true during update disables the export of route Distribution by setting exportDrgRouteDistributionId to null.
	RemoveExportDrgRouteDistributionTrigger pulumi.BoolPtrInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table used by the DRG attachment.
	RouteTableId pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN. This field is deprecated. Instead, use the `networkDetails` field to specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached resource.
	VcnId pulumi.StringPtrInput
}

func (DrgAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*drgAttachmentArgs)(nil)).Elem()
}

type DrgAttachmentInput interface {
	pulumi.Input

	ToDrgAttachmentOutput() DrgAttachmentOutput
	ToDrgAttachmentOutputWithContext(ctx context.Context) DrgAttachmentOutput
}

func (*DrgAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*DrgAttachment)(nil))
}

func (i *DrgAttachment) ToDrgAttachmentOutput() DrgAttachmentOutput {
	return i.ToDrgAttachmentOutputWithContext(context.Background())
}

func (i *DrgAttachment) ToDrgAttachmentOutputWithContext(ctx context.Context) DrgAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrgAttachmentOutput)
}

func (i *DrgAttachment) ToDrgAttachmentPtrOutput() DrgAttachmentPtrOutput {
	return i.ToDrgAttachmentPtrOutputWithContext(context.Background())
}

func (i *DrgAttachment) ToDrgAttachmentPtrOutputWithContext(ctx context.Context) DrgAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrgAttachmentPtrOutput)
}

type DrgAttachmentPtrInput interface {
	pulumi.Input

	ToDrgAttachmentPtrOutput() DrgAttachmentPtrOutput
	ToDrgAttachmentPtrOutputWithContext(ctx context.Context) DrgAttachmentPtrOutput
}

type drgAttachmentPtrType DrgAttachmentArgs

func (*drgAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DrgAttachment)(nil))
}

func (i *drgAttachmentPtrType) ToDrgAttachmentPtrOutput() DrgAttachmentPtrOutput {
	return i.ToDrgAttachmentPtrOutputWithContext(context.Background())
}

func (i *drgAttachmentPtrType) ToDrgAttachmentPtrOutputWithContext(ctx context.Context) DrgAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrgAttachmentPtrOutput)
}

// DrgAttachmentArrayInput is an input type that accepts DrgAttachmentArray and DrgAttachmentArrayOutput values.
// You can construct a concrete instance of `DrgAttachmentArrayInput` via:
//
//          DrgAttachmentArray{ DrgAttachmentArgs{...} }
type DrgAttachmentArrayInput interface {
	pulumi.Input

	ToDrgAttachmentArrayOutput() DrgAttachmentArrayOutput
	ToDrgAttachmentArrayOutputWithContext(context.Context) DrgAttachmentArrayOutput
}

type DrgAttachmentArray []DrgAttachmentInput

func (DrgAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DrgAttachment)(nil)).Elem()
}

func (i DrgAttachmentArray) ToDrgAttachmentArrayOutput() DrgAttachmentArrayOutput {
	return i.ToDrgAttachmentArrayOutputWithContext(context.Background())
}

func (i DrgAttachmentArray) ToDrgAttachmentArrayOutputWithContext(ctx context.Context) DrgAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrgAttachmentArrayOutput)
}

// DrgAttachmentMapInput is an input type that accepts DrgAttachmentMap and DrgAttachmentMapOutput values.
// You can construct a concrete instance of `DrgAttachmentMapInput` via:
//
//          DrgAttachmentMap{ "key": DrgAttachmentArgs{...} }
type DrgAttachmentMapInput interface {
	pulumi.Input

	ToDrgAttachmentMapOutput() DrgAttachmentMapOutput
	ToDrgAttachmentMapOutputWithContext(context.Context) DrgAttachmentMapOutput
}

type DrgAttachmentMap map[string]DrgAttachmentInput

func (DrgAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DrgAttachment)(nil)).Elem()
}

func (i DrgAttachmentMap) ToDrgAttachmentMapOutput() DrgAttachmentMapOutput {
	return i.ToDrgAttachmentMapOutputWithContext(context.Background())
}

func (i DrgAttachmentMap) ToDrgAttachmentMapOutputWithContext(ctx context.Context) DrgAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrgAttachmentMapOutput)
}

type DrgAttachmentOutput struct {
	*pulumi.OutputState
}

func (DrgAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DrgAttachment)(nil))
}

func (o DrgAttachmentOutput) ToDrgAttachmentOutput() DrgAttachmentOutput {
	return o
}

func (o DrgAttachmentOutput) ToDrgAttachmentOutputWithContext(ctx context.Context) DrgAttachmentOutput {
	return o
}

func (o DrgAttachmentOutput) ToDrgAttachmentPtrOutput() DrgAttachmentPtrOutput {
	return o.ToDrgAttachmentPtrOutputWithContext(context.Background())
}

func (o DrgAttachmentOutput) ToDrgAttachmentPtrOutputWithContext(ctx context.Context) DrgAttachmentPtrOutput {
	return o.ApplyT(func(v DrgAttachment) *DrgAttachment {
		return &v
	}).(DrgAttachmentPtrOutput)
}

type DrgAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (DrgAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DrgAttachment)(nil))
}

func (o DrgAttachmentPtrOutput) ToDrgAttachmentPtrOutput() DrgAttachmentPtrOutput {
	return o
}

func (o DrgAttachmentPtrOutput) ToDrgAttachmentPtrOutputWithContext(ctx context.Context) DrgAttachmentPtrOutput {
	return o
}

type DrgAttachmentArrayOutput struct{ *pulumi.OutputState }

func (DrgAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DrgAttachment)(nil))
}

func (o DrgAttachmentArrayOutput) ToDrgAttachmentArrayOutput() DrgAttachmentArrayOutput {
	return o
}

func (o DrgAttachmentArrayOutput) ToDrgAttachmentArrayOutputWithContext(ctx context.Context) DrgAttachmentArrayOutput {
	return o
}

func (o DrgAttachmentArrayOutput) Index(i pulumi.IntInput) DrgAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DrgAttachment {
		return vs[0].([]DrgAttachment)[vs[1].(int)]
	}).(DrgAttachmentOutput)
}

type DrgAttachmentMapOutput struct{ *pulumi.OutputState }

func (DrgAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DrgAttachment)(nil))
}

func (o DrgAttachmentMapOutput) ToDrgAttachmentMapOutput() DrgAttachmentMapOutput {
	return o
}

func (o DrgAttachmentMapOutput) ToDrgAttachmentMapOutputWithContext(ctx context.Context) DrgAttachmentMapOutput {
	return o
}

func (o DrgAttachmentMapOutput) MapIndex(k pulumi.StringInput) DrgAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DrgAttachment {
		return vs[0].(map[string]DrgAttachment)[vs[1].(string)]
	}).(DrgAttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(DrgAttachmentOutput{})
	pulumi.RegisterOutputType(DrgAttachmentPtrOutput{})
	pulumi.RegisterOutputType(DrgAttachmentArrayOutput{})
	pulumi.RegisterOutputType(DrgAttachmentMapOutput{})
}
