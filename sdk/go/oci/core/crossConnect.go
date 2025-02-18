// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Cross Connect resource in Oracle Cloud Infrastructure Core service.
//
// Creates a new cross-connect. Oracle recommends you create each cross-connect in a
// [CrossConnectGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CrossConnectGroup) so you can use link aggregation
// with the connection.
//
// After creating the `CrossConnect` object, you need to go the FastConnect location
// and request to have the physical cable installed. For more information, see
// [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).
//
// For the purposes of access control, you must provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the
// compartment where you want the cross-connect to reside. If you're
// not sure which compartment to use, put the cross-connect in the
// same compartment with your VCN. For more information about
// compartments and access control, see
// [Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
// For information about OCIDs, see
// [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// You may optionally specify a *display name* for the cross-connect.
// It does not have to be unique, and you can change it. Avoid entering confidential information.
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
// 		_, err := core.NewCrossConnect(ctx, "testCrossConnect", &core.CrossConnectArgs{
// 			CompartmentId:         pulumi.Any(_var.Compartment_id),
// 			LocationName:          pulumi.Any(_var.Cross_connect_location_name),
// 			PortSpeedShapeName:    pulumi.Any(_var.Cross_connect_port_speed_shape_name),
// 			CrossConnectGroupId:   pulumi.Any(oci_core_cross_connect_group.Test_cross_connect_group.Id),
// 			CustomerReferenceName: pulumi.Any(_var.Cross_connect_customer_reference_name),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			DisplayName:                          pulumi.Any(_var.Cross_connect_display_name),
// 			FarCrossConnectOrCrossConnectGroupId: pulumi.Any(oci_core_cross_connect_group.Test_cross_connect_group.Id),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			NearCrossConnectOrCrossConnectGroupId: pulumi.Any(oci_core_cross_connect_group.Test_cross_connect_group.Id),
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
// CrossConnects can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:core/crossConnect:CrossConnect test_cross_connect "id"
// ```
type CrossConnect struct {
	pulumi.CustomResourceState

	// (Updatable) The OCID of the compartment to contain the cross-connect.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// The OCID of the cross-connect group to put this cross-connect in.
	CrossConnectGroupId pulumi.StringOutput `pulumi:"crossConnectGroupId"`
	// (Updatable) A reference name or identifier for the physical fiber connection that this cross-connect uses.
	CustomerReferenceName pulumi.StringOutput `pulumi:"customerReferenceName"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on a different router (for the purposes of redundancy), provide the OCID of that existing cross-connect or cross-connect group.
	FarCrossConnectOrCrossConnectGroupId pulumi.StringOutput `pulumi:"farCrossConnectOrCrossConnectGroupId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// (Updatable) Set to true to activate the cross-connect. You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up. Activation indicates to Oracle that the physical connection is ready.
	IsActive pulumi.BoolPtrOutput `pulumi:"isActive"`
	// The name of the FastConnect location where this cross-connect will be installed. To get a list of the available locations, see [ListCrossConnectLocations](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CrossConnectLocation/ListCrossConnectLocations).  Example: `CyrusOne, Chandler, AZ`
	LocationName pulumi.StringOutput `pulumi:"locationName"`
	// If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on the same router, provide the OCID of that existing cross-connect or cross-connect group.
	NearCrossConnectOrCrossConnectGroupId pulumi.StringOutput `pulumi:"nearCrossConnectOrCrossConnectGroupId"`
	// A string identifying the meet-me room port for this cross-connect.
	PortName pulumi.StringOutput `pulumi:"portName"`
	// The port speed for this cross-connect. To get a list of the available port speeds, see [ListCrossConnectPortSpeedShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CrossConnectPortSpeedShape/ListCrossconnectPortSpeedShapes).  Example: `10 Gbps`
	PortSpeedShapeName pulumi.StringOutput `pulumi:"portSpeedShapeName"`
	// The cross-connect's current state.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the cross-connect was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
}

// NewCrossConnect registers a new resource with the given unique name, arguments, and options.
func NewCrossConnect(ctx *pulumi.Context,
	name string, args *CrossConnectArgs, opts ...pulumi.ResourceOption) (*CrossConnect, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.LocationName == nil {
		return nil, errors.New("invalid value for required argument 'LocationName'")
	}
	if args.PortSpeedShapeName == nil {
		return nil, errors.New("invalid value for required argument 'PortSpeedShapeName'")
	}
	var resource CrossConnect
	err := ctx.RegisterResource("oci:core/crossConnect:CrossConnect", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCrossConnect gets an existing CrossConnect resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCrossConnect(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CrossConnectState, opts ...pulumi.ResourceOption) (*CrossConnect, error) {
	var resource CrossConnect
	err := ctx.ReadResource("oci:core/crossConnect:CrossConnect", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CrossConnect resources.
type crossConnectState struct {
	// (Updatable) The OCID of the compartment to contain the cross-connect.
	CompartmentId *string `pulumi:"compartmentId"`
	// The OCID of the cross-connect group to put this cross-connect in.
	CrossConnectGroupId *string `pulumi:"crossConnectGroupId"`
	// (Updatable) A reference name or identifier for the physical fiber connection that this cross-connect uses.
	CustomerReferenceName *string `pulumi:"customerReferenceName"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on a different router (for the purposes of redundancy), provide the OCID of that existing cross-connect or cross-connect group.
	FarCrossConnectOrCrossConnectGroupId *string `pulumi:"farCrossConnectOrCrossConnectGroupId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) Set to true to activate the cross-connect. You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up. Activation indicates to Oracle that the physical connection is ready.
	IsActive *bool `pulumi:"isActive"`
	// The name of the FastConnect location where this cross-connect will be installed. To get a list of the available locations, see [ListCrossConnectLocations](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CrossConnectLocation/ListCrossConnectLocations).  Example: `CyrusOne, Chandler, AZ`
	LocationName *string `pulumi:"locationName"`
	// If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on the same router, provide the OCID of that existing cross-connect or cross-connect group.
	NearCrossConnectOrCrossConnectGroupId *string `pulumi:"nearCrossConnectOrCrossConnectGroupId"`
	// A string identifying the meet-me room port for this cross-connect.
	PortName *string `pulumi:"portName"`
	// The port speed for this cross-connect. To get a list of the available port speeds, see [ListCrossConnectPortSpeedShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CrossConnectPortSpeedShape/ListCrossconnectPortSpeedShapes).  Example: `10 Gbps`
	PortSpeedShapeName *string `pulumi:"portSpeedShapeName"`
	// The cross-connect's current state.
	State *string `pulumi:"state"`
	// The date and time the cross-connect was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *string `pulumi:"timeCreated"`
}

type CrossConnectState struct {
	// (Updatable) The OCID of the compartment to contain the cross-connect.
	CompartmentId pulumi.StringPtrInput
	// The OCID of the cross-connect group to put this cross-connect in.
	CrossConnectGroupId pulumi.StringPtrInput
	// (Updatable) A reference name or identifier for the physical fiber connection that this cross-connect uses.
	CustomerReferenceName pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on a different router (for the purposes of redundancy), provide the OCID of that existing cross-connect or cross-connect group.
	FarCrossConnectOrCrossConnectGroupId pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) Set to true to activate the cross-connect. You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up. Activation indicates to Oracle that the physical connection is ready.
	IsActive pulumi.BoolPtrInput
	// The name of the FastConnect location where this cross-connect will be installed. To get a list of the available locations, see [ListCrossConnectLocations](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CrossConnectLocation/ListCrossConnectLocations).  Example: `CyrusOne, Chandler, AZ`
	LocationName pulumi.StringPtrInput
	// If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on the same router, provide the OCID of that existing cross-connect or cross-connect group.
	NearCrossConnectOrCrossConnectGroupId pulumi.StringPtrInput
	// A string identifying the meet-me room port for this cross-connect.
	PortName pulumi.StringPtrInput
	// The port speed for this cross-connect. To get a list of the available port speeds, see [ListCrossConnectPortSpeedShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CrossConnectPortSpeedShape/ListCrossconnectPortSpeedShapes).  Example: `10 Gbps`
	PortSpeedShapeName pulumi.StringPtrInput
	// The cross-connect's current state.
	State pulumi.StringPtrInput
	// The date and time the cross-connect was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringPtrInput
}

func (CrossConnectState) ElementType() reflect.Type {
	return reflect.TypeOf((*crossConnectState)(nil)).Elem()
}

type crossConnectArgs struct {
	// (Updatable) The OCID of the compartment to contain the cross-connect.
	CompartmentId string `pulumi:"compartmentId"`
	// The OCID of the cross-connect group to put this cross-connect in.
	CrossConnectGroupId *string `pulumi:"crossConnectGroupId"`
	// (Updatable) A reference name or identifier for the physical fiber connection that this cross-connect uses.
	CustomerReferenceName *string `pulumi:"customerReferenceName"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on a different router (for the purposes of redundancy), provide the OCID of that existing cross-connect or cross-connect group.
	FarCrossConnectOrCrossConnectGroupId *string `pulumi:"farCrossConnectOrCrossConnectGroupId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) Set to true to activate the cross-connect. You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up. Activation indicates to Oracle that the physical connection is ready.
	IsActive *bool `pulumi:"isActive"`
	// The name of the FastConnect location where this cross-connect will be installed. To get a list of the available locations, see [ListCrossConnectLocations](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CrossConnectLocation/ListCrossConnectLocations).  Example: `CyrusOne, Chandler, AZ`
	LocationName string `pulumi:"locationName"`
	// If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on the same router, provide the OCID of that existing cross-connect or cross-connect group.
	NearCrossConnectOrCrossConnectGroupId *string `pulumi:"nearCrossConnectOrCrossConnectGroupId"`
	// The port speed for this cross-connect. To get a list of the available port speeds, see [ListCrossConnectPortSpeedShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CrossConnectPortSpeedShape/ListCrossconnectPortSpeedShapes).  Example: `10 Gbps`
	PortSpeedShapeName string `pulumi:"portSpeedShapeName"`
}

// The set of arguments for constructing a CrossConnect resource.
type CrossConnectArgs struct {
	// (Updatable) The OCID of the compartment to contain the cross-connect.
	CompartmentId pulumi.StringInput
	// The OCID of the cross-connect group to put this cross-connect in.
	CrossConnectGroupId pulumi.StringPtrInput
	// (Updatable) A reference name or identifier for the physical fiber connection that this cross-connect uses.
	CustomerReferenceName pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on a different router (for the purposes of redundancy), provide the OCID of that existing cross-connect or cross-connect group.
	FarCrossConnectOrCrossConnectGroupId pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) Set to true to activate the cross-connect. You activate it after the physical cabling is complete, and you've confirmed the cross-connect's light levels are good and your side of the interface is up. Activation indicates to Oracle that the physical connection is ready.
	IsActive pulumi.BoolPtrInput
	// The name of the FastConnect location where this cross-connect will be installed. To get a list of the available locations, see [ListCrossConnectLocations](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CrossConnectLocation/ListCrossConnectLocations).  Example: `CyrusOne, Chandler, AZ`
	LocationName pulumi.StringInput
	// If you already have an existing cross-connect or cross-connect group at this FastConnect location, and you want this new cross-connect to be on the same router, provide the OCID of that existing cross-connect or cross-connect group.
	NearCrossConnectOrCrossConnectGroupId pulumi.StringPtrInput
	// The port speed for this cross-connect. To get a list of the available port speeds, see [ListCrossConnectPortSpeedShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CrossConnectPortSpeedShape/ListCrossconnectPortSpeedShapes).  Example: `10 Gbps`
	PortSpeedShapeName pulumi.StringInput
}

func (CrossConnectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*crossConnectArgs)(nil)).Elem()
}

type CrossConnectInput interface {
	pulumi.Input

	ToCrossConnectOutput() CrossConnectOutput
	ToCrossConnectOutputWithContext(ctx context.Context) CrossConnectOutput
}

func (*CrossConnect) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossConnect)(nil))
}

func (i *CrossConnect) ToCrossConnectOutput() CrossConnectOutput {
	return i.ToCrossConnectOutputWithContext(context.Background())
}

func (i *CrossConnect) ToCrossConnectOutputWithContext(ctx context.Context) CrossConnectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossConnectOutput)
}

func (i *CrossConnect) ToCrossConnectPtrOutput() CrossConnectPtrOutput {
	return i.ToCrossConnectPtrOutputWithContext(context.Background())
}

func (i *CrossConnect) ToCrossConnectPtrOutputWithContext(ctx context.Context) CrossConnectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossConnectPtrOutput)
}

type CrossConnectPtrInput interface {
	pulumi.Input

	ToCrossConnectPtrOutput() CrossConnectPtrOutput
	ToCrossConnectPtrOutputWithContext(ctx context.Context) CrossConnectPtrOutput
}

type crossConnectPtrType CrossConnectArgs

func (*crossConnectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossConnect)(nil))
}

func (i *crossConnectPtrType) ToCrossConnectPtrOutput() CrossConnectPtrOutput {
	return i.ToCrossConnectPtrOutputWithContext(context.Background())
}

func (i *crossConnectPtrType) ToCrossConnectPtrOutputWithContext(ctx context.Context) CrossConnectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossConnectPtrOutput)
}

// CrossConnectArrayInput is an input type that accepts CrossConnectArray and CrossConnectArrayOutput values.
// You can construct a concrete instance of `CrossConnectArrayInput` via:
//
//          CrossConnectArray{ CrossConnectArgs{...} }
type CrossConnectArrayInput interface {
	pulumi.Input

	ToCrossConnectArrayOutput() CrossConnectArrayOutput
	ToCrossConnectArrayOutputWithContext(context.Context) CrossConnectArrayOutput
}

type CrossConnectArray []CrossConnectInput

func (CrossConnectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CrossConnect)(nil)).Elem()
}

func (i CrossConnectArray) ToCrossConnectArrayOutput() CrossConnectArrayOutput {
	return i.ToCrossConnectArrayOutputWithContext(context.Background())
}

func (i CrossConnectArray) ToCrossConnectArrayOutputWithContext(ctx context.Context) CrossConnectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossConnectArrayOutput)
}

// CrossConnectMapInput is an input type that accepts CrossConnectMap and CrossConnectMapOutput values.
// You can construct a concrete instance of `CrossConnectMapInput` via:
//
//          CrossConnectMap{ "key": CrossConnectArgs{...} }
type CrossConnectMapInput interface {
	pulumi.Input

	ToCrossConnectMapOutput() CrossConnectMapOutput
	ToCrossConnectMapOutputWithContext(context.Context) CrossConnectMapOutput
}

type CrossConnectMap map[string]CrossConnectInput

func (CrossConnectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CrossConnect)(nil)).Elem()
}

func (i CrossConnectMap) ToCrossConnectMapOutput() CrossConnectMapOutput {
	return i.ToCrossConnectMapOutputWithContext(context.Background())
}

func (i CrossConnectMap) ToCrossConnectMapOutputWithContext(ctx context.Context) CrossConnectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossConnectMapOutput)
}

type CrossConnectOutput struct {
	*pulumi.OutputState
}

func (CrossConnectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossConnect)(nil))
}

func (o CrossConnectOutput) ToCrossConnectOutput() CrossConnectOutput {
	return o
}

func (o CrossConnectOutput) ToCrossConnectOutputWithContext(ctx context.Context) CrossConnectOutput {
	return o
}

func (o CrossConnectOutput) ToCrossConnectPtrOutput() CrossConnectPtrOutput {
	return o.ToCrossConnectPtrOutputWithContext(context.Background())
}

func (o CrossConnectOutput) ToCrossConnectPtrOutputWithContext(ctx context.Context) CrossConnectPtrOutput {
	return o.ApplyT(func(v CrossConnect) *CrossConnect {
		return &v
	}).(CrossConnectPtrOutput)
}

type CrossConnectPtrOutput struct {
	*pulumi.OutputState
}

func (CrossConnectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossConnect)(nil))
}

func (o CrossConnectPtrOutput) ToCrossConnectPtrOutput() CrossConnectPtrOutput {
	return o
}

func (o CrossConnectPtrOutput) ToCrossConnectPtrOutputWithContext(ctx context.Context) CrossConnectPtrOutput {
	return o
}

type CrossConnectArrayOutput struct{ *pulumi.OutputState }

func (CrossConnectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CrossConnect)(nil))
}

func (o CrossConnectArrayOutput) ToCrossConnectArrayOutput() CrossConnectArrayOutput {
	return o
}

func (o CrossConnectArrayOutput) ToCrossConnectArrayOutputWithContext(ctx context.Context) CrossConnectArrayOutput {
	return o
}

func (o CrossConnectArrayOutput) Index(i pulumi.IntInput) CrossConnectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CrossConnect {
		return vs[0].([]CrossConnect)[vs[1].(int)]
	}).(CrossConnectOutput)
}

type CrossConnectMapOutput struct{ *pulumi.OutputState }

func (CrossConnectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CrossConnect)(nil))
}

func (o CrossConnectMapOutput) ToCrossConnectMapOutput() CrossConnectMapOutput {
	return o
}

func (o CrossConnectMapOutput) ToCrossConnectMapOutputWithContext(ctx context.Context) CrossConnectMapOutput {
	return o
}

func (o CrossConnectMapOutput) MapIndex(k pulumi.StringInput) CrossConnectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CrossConnect {
		return vs[0].(map[string]CrossConnect)[vs[1].(string)]
	}).(CrossConnectOutput)
}

func init() {
	pulumi.RegisterOutputType(CrossConnectOutput{})
	pulumi.RegisterOutputType(CrossConnectPtrOutput{})
	pulumi.RegisterOutputType(CrossConnectArrayOutput{})
	pulumi.RegisterOutputType(CrossConnectMapOutput{})
}
