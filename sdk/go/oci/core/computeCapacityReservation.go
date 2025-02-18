// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Compute Capacity Reservation resource in Oracle Cloud Infrastructure Core service.
//
// Creates a new compute capacity reservation in the specified compartment and availability domain.
// Compute capacity reservations let you reserve instances in a compartment.
// When you launch an instance using this reservation, you are assured that you have enough space for your instance,
// and you won't get out of capacity errors.
// For more information, see [Reserved Capacity](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm).
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
// 		_, err := core.NewComputeCapacityReservation(ctx, "testComputeCapacityReservation", &core.ComputeCapacityReservationArgs{
// 			AvailabilityDomain: pulumi.Any(_var.Compute_capacity_reservation_availability_domain),
// 			CompartmentId:      pulumi.Any(_var.Compartment_id),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			DisplayName: pulumi.Any(_var.Compute_capacity_reservation_display_name),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			InstanceReservationConfigs: core.ComputeCapacityReservationInstanceReservationConfigArray{
// 				&core.ComputeCapacityReservationInstanceReservationConfigArgs{
// 					InstanceShape: pulumi.Any(_var.Compute_capacity_reservation_instance_reservation_configs_instance_shape),
// 					ReservedCount: pulumi.Any(_var.Compute_capacity_reservation_instance_reservation_configs_reserved_count),
// 					FaultDomain:   pulumi.Any(_var.Compute_capacity_reservation_instance_reservation_configs_fault_domain),
// 					InstanceShapeConfig: &core.ComputeCapacityReservationInstanceReservationConfigInstanceShapeConfigArgs{
// 						MemoryInGbs: pulumi.Any(_var.Compute_capacity_reservation_instance_reservation_configs_instance_shape_config_memory_in_gbs),
// 						Ocpus:       pulumi.Any(_var.Compute_capacity_reservation_instance_reservation_configs_instance_shape_config_ocpus),
// 					},
// 				},
// 			},
// 			IsDefaultReservation: pulumi.Any(_var.Compute_capacity_reservation_is_default_reservation),
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
// ComputeCapacityReservations can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:core/computeCapacityReservation:ComputeCapacityReservation test_compute_capacity_reservation "id"
// ```
type ComputeCapacityReservation struct {
	pulumi.CustomResourceState

	// The availability domain of this compute capacity reservation.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringOutput `pulumi:"availabilityDomain"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the capacity reservation.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) A user-friendly name for the compute capacity reservation. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// (Updatable) The reservation configurations for the capacity reservation.
	InstanceReservationConfigs ComputeCapacityReservationInstanceReservationConfigArrayOutput `pulumi:"instanceReservationConfigs"`
	// (Updatable) Whether this capacity reservation is the default. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default).
	IsDefaultReservation pulumi.BoolOutput `pulumi:"isDefaultReservation"`
	// The number of instances for which capacity will be held with this compute capacity reservation. This number is the sum of the values of the `reservedCount` fields for all of the instance reservation configurations under this reservation. The purpose of this field is to calculate the percentage usage of the reservation.
	ReservedInstanceCount pulumi.StringOutput `pulumi:"reservedInstanceCount"`
	// The current state of the compute capacity reservation.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the compute capacity reservation was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The date and time the compute capacity reservation was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated pulumi.StringOutput `pulumi:"timeUpdated"`
	// The total number of instances currently consuming space in this compute capacity reservation. This number is the sum of the values of the `usedCount` fields for all of the instance reservation configurations under this reservation. The purpose of this field is to calculate the percentage usage of the reservation.
	UsedInstanceCount pulumi.StringOutput `pulumi:"usedInstanceCount"`
}

// NewComputeCapacityReservation registers a new resource with the given unique name, arguments, and options.
func NewComputeCapacityReservation(ctx *pulumi.Context,
	name string, args *ComputeCapacityReservationArgs, opts ...pulumi.ResourceOption) (*ComputeCapacityReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityDomain == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityDomain'")
	}
	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.InstanceReservationConfigs == nil {
		return nil, errors.New("invalid value for required argument 'InstanceReservationConfigs'")
	}
	var resource ComputeCapacityReservation
	err := ctx.RegisterResource("oci:core/computeCapacityReservation:ComputeCapacityReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeCapacityReservation gets an existing ComputeCapacityReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeCapacityReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeCapacityReservationState, opts ...pulumi.ResourceOption) (*ComputeCapacityReservation, error) {
	var resource ComputeCapacityReservation
	err := ctx.ReadResource("oci:core/computeCapacityReservation:ComputeCapacityReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeCapacityReservation resources.
type computeCapacityReservationState struct {
	// The availability domain of this compute capacity reservation.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the capacity reservation.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name for the compute capacity reservation. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The reservation configurations for the capacity reservation.
	InstanceReservationConfigs []ComputeCapacityReservationInstanceReservationConfig `pulumi:"instanceReservationConfigs"`
	// (Updatable) Whether this capacity reservation is the default. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default).
	IsDefaultReservation *bool `pulumi:"isDefaultReservation"`
	// The number of instances for which capacity will be held with this compute capacity reservation. This number is the sum of the values of the `reservedCount` fields for all of the instance reservation configurations under this reservation. The purpose of this field is to calculate the percentage usage of the reservation.
	ReservedInstanceCount *string `pulumi:"reservedInstanceCount"`
	// The current state of the compute capacity reservation.
	State *string `pulumi:"state"`
	// The date and time the compute capacity reservation was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *string `pulumi:"timeCreated"`
	// The date and time the compute capacity reservation was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *string `pulumi:"timeUpdated"`
	// The total number of instances currently consuming space in this compute capacity reservation. This number is the sum of the values of the `usedCount` fields for all of the instance reservation configurations under this reservation. The purpose of this field is to calculate the percentage usage of the reservation.
	UsedInstanceCount *string `pulumi:"usedInstanceCount"`
}

type ComputeCapacityReservationState struct {
	// The availability domain of this compute capacity reservation.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringPtrInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the capacity reservation.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name for the compute capacity reservation. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The reservation configurations for the capacity reservation.
	InstanceReservationConfigs ComputeCapacityReservationInstanceReservationConfigArrayInput
	// (Updatable) Whether this capacity reservation is the default. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default).
	IsDefaultReservation pulumi.BoolPtrInput
	// The number of instances for which capacity will be held with this compute capacity reservation. This number is the sum of the values of the `reservedCount` fields for all of the instance reservation configurations under this reservation. The purpose of this field is to calculate the percentage usage of the reservation.
	ReservedInstanceCount pulumi.StringPtrInput
	// The current state of the compute capacity reservation.
	State pulumi.StringPtrInput
	// The date and time the compute capacity reservation was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringPtrInput
	// The date and time the compute capacity reservation was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated pulumi.StringPtrInput
	// The total number of instances currently consuming space in this compute capacity reservation. This number is the sum of the values of the `usedCount` fields for all of the instance reservation configurations under this reservation. The purpose of this field is to calculate the percentage usage of the reservation.
	UsedInstanceCount pulumi.StringPtrInput
}

func (ComputeCapacityReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeCapacityReservationState)(nil)).Elem()
}

type computeCapacityReservationArgs struct {
	// The availability domain of this compute capacity reservation.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain string `pulumi:"availabilityDomain"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the capacity reservation.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name for the compute capacity reservation. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The reservation configurations for the capacity reservation.
	InstanceReservationConfigs []ComputeCapacityReservationInstanceReservationConfig `pulumi:"instanceReservationConfigs"`
	// (Updatable) Whether this capacity reservation is the default. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default).
	IsDefaultReservation *bool `pulumi:"isDefaultReservation"`
}

// The set of arguments for constructing a ComputeCapacityReservation resource.
type ComputeCapacityReservationArgs struct {
	// The availability domain of this compute capacity reservation.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the capacity reservation.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name for the compute capacity reservation. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The reservation configurations for the capacity reservation.
	InstanceReservationConfigs ComputeCapacityReservationInstanceReservationConfigArrayInput
	// (Updatable) Whether this capacity reservation is the default. For more information, see [Capacity Reservations](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default).
	IsDefaultReservation pulumi.BoolPtrInput
}

func (ComputeCapacityReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeCapacityReservationArgs)(nil)).Elem()
}

type ComputeCapacityReservationInput interface {
	pulumi.Input

	ToComputeCapacityReservationOutput() ComputeCapacityReservationOutput
	ToComputeCapacityReservationOutputWithContext(ctx context.Context) ComputeCapacityReservationOutput
}

func (*ComputeCapacityReservation) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeCapacityReservation)(nil))
}

func (i *ComputeCapacityReservation) ToComputeCapacityReservationOutput() ComputeCapacityReservationOutput {
	return i.ToComputeCapacityReservationOutputWithContext(context.Background())
}

func (i *ComputeCapacityReservation) ToComputeCapacityReservationOutputWithContext(ctx context.Context) ComputeCapacityReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeCapacityReservationOutput)
}

func (i *ComputeCapacityReservation) ToComputeCapacityReservationPtrOutput() ComputeCapacityReservationPtrOutput {
	return i.ToComputeCapacityReservationPtrOutputWithContext(context.Background())
}

func (i *ComputeCapacityReservation) ToComputeCapacityReservationPtrOutputWithContext(ctx context.Context) ComputeCapacityReservationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeCapacityReservationPtrOutput)
}

type ComputeCapacityReservationPtrInput interface {
	pulumi.Input

	ToComputeCapacityReservationPtrOutput() ComputeCapacityReservationPtrOutput
	ToComputeCapacityReservationPtrOutputWithContext(ctx context.Context) ComputeCapacityReservationPtrOutput
}

type computeCapacityReservationPtrType ComputeCapacityReservationArgs

func (*computeCapacityReservationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeCapacityReservation)(nil))
}

func (i *computeCapacityReservationPtrType) ToComputeCapacityReservationPtrOutput() ComputeCapacityReservationPtrOutput {
	return i.ToComputeCapacityReservationPtrOutputWithContext(context.Background())
}

func (i *computeCapacityReservationPtrType) ToComputeCapacityReservationPtrOutputWithContext(ctx context.Context) ComputeCapacityReservationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeCapacityReservationPtrOutput)
}

// ComputeCapacityReservationArrayInput is an input type that accepts ComputeCapacityReservationArray and ComputeCapacityReservationArrayOutput values.
// You can construct a concrete instance of `ComputeCapacityReservationArrayInput` via:
//
//          ComputeCapacityReservationArray{ ComputeCapacityReservationArgs{...} }
type ComputeCapacityReservationArrayInput interface {
	pulumi.Input

	ToComputeCapacityReservationArrayOutput() ComputeCapacityReservationArrayOutput
	ToComputeCapacityReservationArrayOutputWithContext(context.Context) ComputeCapacityReservationArrayOutput
}

type ComputeCapacityReservationArray []ComputeCapacityReservationInput

func (ComputeCapacityReservationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeCapacityReservation)(nil)).Elem()
}

func (i ComputeCapacityReservationArray) ToComputeCapacityReservationArrayOutput() ComputeCapacityReservationArrayOutput {
	return i.ToComputeCapacityReservationArrayOutputWithContext(context.Background())
}

func (i ComputeCapacityReservationArray) ToComputeCapacityReservationArrayOutputWithContext(ctx context.Context) ComputeCapacityReservationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeCapacityReservationArrayOutput)
}

// ComputeCapacityReservationMapInput is an input type that accepts ComputeCapacityReservationMap and ComputeCapacityReservationMapOutput values.
// You can construct a concrete instance of `ComputeCapacityReservationMapInput` via:
//
//          ComputeCapacityReservationMap{ "key": ComputeCapacityReservationArgs{...} }
type ComputeCapacityReservationMapInput interface {
	pulumi.Input

	ToComputeCapacityReservationMapOutput() ComputeCapacityReservationMapOutput
	ToComputeCapacityReservationMapOutputWithContext(context.Context) ComputeCapacityReservationMapOutput
}

type ComputeCapacityReservationMap map[string]ComputeCapacityReservationInput

func (ComputeCapacityReservationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeCapacityReservation)(nil)).Elem()
}

func (i ComputeCapacityReservationMap) ToComputeCapacityReservationMapOutput() ComputeCapacityReservationMapOutput {
	return i.ToComputeCapacityReservationMapOutputWithContext(context.Background())
}

func (i ComputeCapacityReservationMap) ToComputeCapacityReservationMapOutputWithContext(ctx context.Context) ComputeCapacityReservationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeCapacityReservationMapOutput)
}

type ComputeCapacityReservationOutput struct {
	*pulumi.OutputState
}

func (ComputeCapacityReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeCapacityReservation)(nil))
}

func (o ComputeCapacityReservationOutput) ToComputeCapacityReservationOutput() ComputeCapacityReservationOutput {
	return o
}

func (o ComputeCapacityReservationOutput) ToComputeCapacityReservationOutputWithContext(ctx context.Context) ComputeCapacityReservationOutput {
	return o
}

func (o ComputeCapacityReservationOutput) ToComputeCapacityReservationPtrOutput() ComputeCapacityReservationPtrOutput {
	return o.ToComputeCapacityReservationPtrOutputWithContext(context.Background())
}

func (o ComputeCapacityReservationOutput) ToComputeCapacityReservationPtrOutputWithContext(ctx context.Context) ComputeCapacityReservationPtrOutput {
	return o.ApplyT(func(v ComputeCapacityReservation) *ComputeCapacityReservation {
		return &v
	}).(ComputeCapacityReservationPtrOutput)
}

type ComputeCapacityReservationPtrOutput struct {
	*pulumi.OutputState
}

func (ComputeCapacityReservationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeCapacityReservation)(nil))
}

func (o ComputeCapacityReservationPtrOutput) ToComputeCapacityReservationPtrOutput() ComputeCapacityReservationPtrOutput {
	return o
}

func (o ComputeCapacityReservationPtrOutput) ToComputeCapacityReservationPtrOutputWithContext(ctx context.Context) ComputeCapacityReservationPtrOutput {
	return o
}

type ComputeCapacityReservationArrayOutput struct{ *pulumi.OutputState }

func (ComputeCapacityReservationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ComputeCapacityReservation)(nil))
}

func (o ComputeCapacityReservationArrayOutput) ToComputeCapacityReservationArrayOutput() ComputeCapacityReservationArrayOutput {
	return o
}

func (o ComputeCapacityReservationArrayOutput) ToComputeCapacityReservationArrayOutputWithContext(ctx context.Context) ComputeCapacityReservationArrayOutput {
	return o
}

func (o ComputeCapacityReservationArrayOutput) Index(i pulumi.IntInput) ComputeCapacityReservationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ComputeCapacityReservation {
		return vs[0].([]ComputeCapacityReservation)[vs[1].(int)]
	}).(ComputeCapacityReservationOutput)
}

type ComputeCapacityReservationMapOutput struct{ *pulumi.OutputState }

func (ComputeCapacityReservationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ComputeCapacityReservation)(nil))
}

func (o ComputeCapacityReservationMapOutput) ToComputeCapacityReservationMapOutput() ComputeCapacityReservationMapOutput {
	return o
}

func (o ComputeCapacityReservationMapOutput) ToComputeCapacityReservationMapOutputWithContext(ctx context.Context) ComputeCapacityReservationMapOutput {
	return o
}

func (o ComputeCapacityReservationMapOutput) MapIndex(k pulumi.StringInput) ComputeCapacityReservationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ComputeCapacityReservation {
		return vs[0].(map[string]ComputeCapacityReservation)[vs[1].(string)]
	}).(ComputeCapacityReservationOutput)
}

func init() {
	pulumi.RegisterOutputType(ComputeCapacityReservationOutput{})
	pulumi.RegisterOutputType(ComputeCapacityReservationPtrOutput{})
	pulumi.RegisterOutputType(ComputeCapacityReservationArrayOutput{})
	pulumi.RegisterOutputType(ComputeCapacityReservationMapOutput{})
}
