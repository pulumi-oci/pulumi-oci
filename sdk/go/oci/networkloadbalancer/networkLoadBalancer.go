// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkloadbalancer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Network Load Balancer resource in Oracle Cloud Infrastructure Network Load Balancer service.
//
// Creates a network load balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/networkloadbalancer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkloadbalancer.NewNetworkLoadBalancer(ctx, "testNetworkLoadBalancer", &networkloadbalancer.NetworkLoadBalancerArgs{
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			DisplayName:   pulumi.Any(_var.Network_load_balancer_display_name),
// 			SubnetId:      pulumi.Any(oci_core_subnet.Test_subnet.Id),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			IsPreserveSourceDestination: pulumi.Any(_var.Network_load_balancer_is_preserve_source_destination),
// 			IsPrivate:                   pulumi.Any(_var.Network_load_balancer_is_private),
// 			NetworkSecurityGroupIds:     pulumi.Any(_var.Network_load_balancer_network_security_group_ids),
// 			ReservedIps: networkloadbalancer.NetworkLoadBalancerReservedIpArray{
// 				&networkloadbalancer.NetworkLoadBalancerReservedIpArgs{
// 					Id: pulumi.Any(_var.Network_load_balancer_reserved_ips_id),
// 				},
// 			},
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
// NetworkLoadBalancers can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:networkloadbalancer/networkLoadBalancer:NetworkLoadBalancer test_network_load_balancer "id"
// ```
type NetworkLoadBalancer struct {
	pulumi.CustomResourceState

	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) Network load balancer identifier, which can be renamed.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// An array of IP addresses.
	IpAddresses NetworkLoadBalancerIpAddressArrayOutput `pulumi:"ipAddresses"`
	// (Updatable) This parameter can be enabled only if backends are compute OCIDs. When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC, and packets are sent to the backend with the entire IP header intact.
	IsPreserveSourceDestination pulumi.BoolOutput `pulumi:"isPreserveSourceDestination"`
	// Whether the network load balancer has a virtual cloud network-local (private) IP address.
	IsPrivate pulumi.BoolOutput `pulumi:"isPrivate"`
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails pulumi.StringOutput `pulumi:"lifecycleDetails"`
	// (Updatable) An array of network security groups [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the network load balancer.
	NetworkSecurityGroupIds pulumi.StringArrayOutput `pulumi:"networkSecurityGroupIds"`
	// An array of reserved Ips.
	ReservedIps NetworkLoadBalancerReservedIpArrayOutput `pulumi:"reservedIps"`
	// The current state of the network load balancer.
	State pulumi.StringOutput `pulumi:"state"`
	// The subnet in which the network load balancer is spawned [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Key-value pair representing system tags' keys and values scoped to a namespace. Example: `{"bar-key": "value"}`
	SystemTags pulumi.MapOutput `pulumi:"systemTags"`
	// The date and time the network load balancer was created, in the format defined by RFC3339.  Example: `2020-05-01T21:10:29.600Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The time the network load balancer was updated. An RFC3339 formatted date-time string.  Example: `2020-05-01T22:10:29.600Z`
	TimeUpdated pulumi.StringOutput `pulumi:"timeUpdated"`
}

// NewNetworkLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewNetworkLoadBalancer(ctx *pulumi.Context,
	name string, args *NetworkLoadBalancerArgs, opts ...pulumi.ResourceOption) (*NetworkLoadBalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource NetworkLoadBalancer
	err := ctx.RegisterResource("oci:networkloadbalancer/networkLoadBalancer:NetworkLoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkLoadBalancer gets an existing NetworkLoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkLoadBalancerState, opts ...pulumi.ResourceOption) (*NetworkLoadBalancer, error) {
	var resource NetworkLoadBalancer
	err := ctx.ReadResource("oci:networkloadbalancer/networkLoadBalancer:NetworkLoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkLoadBalancer resources.
type networkLoadBalancerState struct {
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) Network load balancer identifier, which can be renamed.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// An array of IP addresses.
	IpAddresses []NetworkLoadBalancerIpAddress `pulumi:"ipAddresses"`
	// (Updatable) This parameter can be enabled only if backends are compute OCIDs. When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC, and packets are sent to the backend with the entire IP header intact.
	IsPreserveSourceDestination *bool `pulumi:"isPreserveSourceDestination"`
	// Whether the network load balancer has a virtual cloud network-local (private) IP address.
	IsPrivate *bool `pulumi:"isPrivate"`
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// (Updatable) An array of network security groups [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the network load balancer.
	NetworkSecurityGroupIds []string `pulumi:"networkSecurityGroupIds"`
	// An array of reserved Ips.
	ReservedIps []NetworkLoadBalancerReservedIp `pulumi:"reservedIps"`
	// The current state of the network load balancer.
	State *string `pulumi:"state"`
	// The subnet in which the network load balancer is spawned [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	SubnetId *string `pulumi:"subnetId"`
	// Key-value pair representing system tags' keys and values scoped to a namespace. Example: `{"bar-key": "value"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// The date and time the network load balancer was created, in the format defined by RFC3339.  Example: `2020-05-01T21:10:29.600Z`
	TimeCreated *string `pulumi:"timeCreated"`
	// The time the network load balancer was updated. An RFC3339 formatted date-time string.  Example: `2020-05-01T22:10:29.600Z`
	TimeUpdated *string `pulumi:"timeUpdated"`
}

type NetworkLoadBalancerState struct {
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// (Updatable) Network load balancer identifier, which can be renamed.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// An array of IP addresses.
	IpAddresses NetworkLoadBalancerIpAddressArrayInput
	// (Updatable) This parameter can be enabled only if backends are compute OCIDs. When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC, and packets are sent to the backend with the entire IP header intact.
	IsPreserveSourceDestination pulumi.BoolPtrInput
	// Whether the network load balancer has a virtual cloud network-local (private) IP address.
	IsPrivate pulumi.BoolPtrInput
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails pulumi.StringPtrInput
	// (Updatable) An array of network security groups [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the network load balancer.
	NetworkSecurityGroupIds pulumi.StringArrayInput
	// An array of reserved Ips.
	ReservedIps NetworkLoadBalancerReservedIpArrayInput
	// The current state of the network load balancer.
	State pulumi.StringPtrInput
	// The subnet in which the network load balancer is spawned [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	SubnetId pulumi.StringPtrInput
	// Key-value pair representing system tags' keys and values scoped to a namespace. Example: `{"bar-key": "value"}`
	SystemTags pulumi.MapInput
	// The date and time the network load balancer was created, in the format defined by RFC3339.  Example: `2020-05-01T21:10:29.600Z`
	TimeCreated pulumi.StringPtrInput
	// The time the network load balancer was updated. An RFC3339 formatted date-time string.  Example: `2020-05-01T22:10:29.600Z`
	TimeUpdated pulumi.StringPtrInput
}

func (NetworkLoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkLoadBalancerState)(nil)).Elem()
}

type networkLoadBalancerArgs struct {
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) Network load balancer identifier, which can be renamed.
	DisplayName string `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) This parameter can be enabled only if backends are compute OCIDs. When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC, and packets are sent to the backend with the entire IP header intact.
	IsPreserveSourceDestination *bool `pulumi:"isPreserveSourceDestination"`
	// Whether the network load balancer has a virtual cloud network-local (private) IP address.
	IsPrivate *bool `pulumi:"isPrivate"`
	// (Updatable) An array of network security groups [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the network load balancer.
	NetworkSecurityGroupIds []string `pulumi:"networkSecurityGroupIds"`
	// An array of reserved Ips.
	ReservedIps []NetworkLoadBalancerReservedIp `pulumi:"reservedIps"`
	// The subnet in which the network load balancer is spawned [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a NetworkLoadBalancer resource.
type NetworkLoadBalancerArgs struct {
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// (Updatable) Network load balancer identifier, which can be renamed.
	DisplayName pulumi.StringInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// (Updatable) This parameter can be enabled only if backends are compute OCIDs. When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC, and packets are sent to the backend with the entire IP header intact.
	IsPreserveSourceDestination pulumi.BoolPtrInput
	// Whether the network load balancer has a virtual cloud network-local (private) IP address.
	IsPrivate pulumi.BoolPtrInput
	// (Updatable) An array of network security groups [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the network load balancer.
	NetworkSecurityGroupIds pulumi.StringArrayInput
	// An array of reserved Ips.
	ReservedIps NetworkLoadBalancerReservedIpArrayInput
	// The subnet in which the network load balancer is spawned [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	SubnetId pulumi.StringInput
}

func (NetworkLoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkLoadBalancerArgs)(nil)).Elem()
}

type NetworkLoadBalancerInput interface {
	pulumi.Input

	ToNetworkLoadBalancerOutput() NetworkLoadBalancerOutput
	ToNetworkLoadBalancerOutputWithContext(ctx context.Context) NetworkLoadBalancerOutput
}

func (*NetworkLoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkLoadBalancer)(nil))
}

func (i *NetworkLoadBalancer) ToNetworkLoadBalancerOutput() NetworkLoadBalancerOutput {
	return i.ToNetworkLoadBalancerOutputWithContext(context.Background())
}

func (i *NetworkLoadBalancer) ToNetworkLoadBalancerOutputWithContext(ctx context.Context) NetworkLoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkLoadBalancerOutput)
}

func (i *NetworkLoadBalancer) ToNetworkLoadBalancerPtrOutput() NetworkLoadBalancerPtrOutput {
	return i.ToNetworkLoadBalancerPtrOutputWithContext(context.Background())
}

func (i *NetworkLoadBalancer) ToNetworkLoadBalancerPtrOutputWithContext(ctx context.Context) NetworkLoadBalancerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkLoadBalancerPtrOutput)
}

type NetworkLoadBalancerPtrInput interface {
	pulumi.Input

	ToNetworkLoadBalancerPtrOutput() NetworkLoadBalancerPtrOutput
	ToNetworkLoadBalancerPtrOutputWithContext(ctx context.Context) NetworkLoadBalancerPtrOutput
}

type networkLoadBalancerPtrType NetworkLoadBalancerArgs

func (*networkLoadBalancerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkLoadBalancer)(nil))
}

func (i *networkLoadBalancerPtrType) ToNetworkLoadBalancerPtrOutput() NetworkLoadBalancerPtrOutput {
	return i.ToNetworkLoadBalancerPtrOutputWithContext(context.Background())
}

func (i *networkLoadBalancerPtrType) ToNetworkLoadBalancerPtrOutputWithContext(ctx context.Context) NetworkLoadBalancerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkLoadBalancerPtrOutput)
}

// NetworkLoadBalancerArrayInput is an input type that accepts NetworkLoadBalancerArray and NetworkLoadBalancerArrayOutput values.
// You can construct a concrete instance of `NetworkLoadBalancerArrayInput` via:
//
//          NetworkLoadBalancerArray{ NetworkLoadBalancerArgs{...} }
type NetworkLoadBalancerArrayInput interface {
	pulumi.Input

	ToNetworkLoadBalancerArrayOutput() NetworkLoadBalancerArrayOutput
	ToNetworkLoadBalancerArrayOutputWithContext(context.Context) NetworkLoadBalancerArrayOutput
}

type NetworkLoadBalancerArray []NetworkLoadBalancerInput

func (NetworkLoadBalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkLoadBalancer)(nil)).Elem()
}

func (i NetworkLoadBalancerArray) ToNetworkLoadBalancerArrayOutput() NetworkLoadBalancerArrayOutput {
	return i.ToNetworkLoadBalancerArrayOutputWithContext(context.Background())
}

func (i NetworkLoadBalancerArray) ToNetworkLoadBalancerArrayOutputWithContext(ctx context.Context) NetworkLoadBalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkLoadBalancerArrayOutput)
}

// NetworkLoadBalancerMapInput is an input type that accepts NetworkLoadBalancerMap and NetworkLoadBalancerMapOutput values.
// You can construct a concrete instance of `NetworkLoadBalancerMapInput` via:
//
//          NetworkLoadBalancerMap{ "key": NetworkLoadBalancerArgs{...} }
type NetworkLoadBalancerMapInput interface {
	pulumi.Input

	ToNetworkLoadBalancerMapOutput() NetworkLoadBalancerMapOutput
	ToNetworkLoadBalancerMapOutputWithContext(context.Context) NetworkLoadBalancerMapOutput
}

type NetworkLoadBalancerMap map[string]NetworkLoadBalancerInput

func (NetworkLoadBalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkLoadBalancer)(nil)).Elem()
}

func (i NetworkLoadBalancerMap) ToNetworkLoadBalancerMapOutput() NetworkLoadBalancerMapOutput {
	return i.ToNetworkLoadBalancerMapOutputWithContext(context.Background())
}

func (i NetworkLoadBalancerMap) ToNetworkLoadBalancerMapOutputWithContext(ctx context.Context) NetworkLoadBalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkLoadBalancerMapOutput)
}

type NetworkLoadBalancerOutput struct {
	*pulumi.OutputState
}

func (NetworkLoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkLoadBalancer)(nil))
}

func (o NetworkLoadBalancerOutput) ToNetworkLoadBalancerOutput() NetworkLoadBalancerOutput {
	return o
}

func (o NetworkLoadBalancerOutput) ToNetworkLoadBalancerOutputWithContext(ctx context.Context) NetworkLoadBalancerOutput {
	return o
}

func (o NetworkLoadBalancerOutput) ToNetworkLoadBalancerPtrOutput() NetworkLoadBalancerPtrOutput {
	return o.ToNetworkLoadBalancerPtrOutputWithContext(context.Background())
}

func (o NetworkLoadBalancerOutput) ToNetworkLoadBalancerPtrOutputWithContext(ctx context.Context) NetworkLoadBalancerPtrOutput {
	return o.ApplyT(func(v NetworkLoadBalancer) *NetworkLoadBalancer {
		return &v
	}).(NetworkLoadBalancerPtrOutput)
}

type NetworkLoadBalancerPtrOutput struct {
	*pulumi.OutputState
}

func (NetworkLoadBalancerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkLoadBalancer)(nil))
}

func (o NetworkLoadBalancerPtrOutput) ToNetworkLoadBalancerPtrOutput() NetworkLoadBalancerPtrOutput {
	return o
}

func (o NetworkLoadBalancerPtrOutput) ToNetworkLoadBalancerPtrOutputWithContext(ctx context.Context) NetworkLoadBalancerPtrOutput {
	return o
}

type NetworkLoadBalancerArrayOutput struct{ *pulumi.OutputState }

func (NetworkLoadBalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkLoadBalancer)(nil))
}

func (o NetworkLoadBalancerArrayOutput) ToNetworkLoadBalancerArrayOutput() NetworkLoadBalancerArrayOutput {
	return o
}

func (o NetworkLoadBalancerArrayOutput) ToNetworkLoadBalancerArrayOutputWithContext(ctx context.Context) NetworkLoadBalancerArrayOutput {
	return o
}

func (o NetworkLoadBalancerArrayOutput) Index(i pulumi.IntInput) NetworkLoadBalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkLoadBalancer {
		return vs[0].([]NetworkLoadBalancer)[vs[1].(int)]
	}).(NetworkLoadBalancerOutput)
}

type NetworkLoadBalancerMapOutput struct{ *pulumi.OutputState }

func (NetworkLoadBalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkLoadBalancer)(nil))
}

func (o NetworkLoadBalancerMapOutput) ToNetworkLoadBalancerMapOutput() NetworkLoadBalancerMapOutput {
	return o
}

func (o NetworkLoadBalancerMapOutput) ToNetworkLoadBalancerMapOutputWithContext(ctx context.Context) NetworkLoadBalancerMapOutput {
	return o
}

func (o NetworkLoadBalancerMapOutput) MapIndex(k pulumi.StringInput) NetworkLoadBalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkLoadBalancer {
		return vs[0].(map[string]NetworkLoadBalancer)[vs[1].(string)]
	}).(NetworkLoadBalancerOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkLoadBalancerOutput{})
	pulumi.RegisterOutputType(NetworkLoadBalancerPtrOutput{})
	pulumi.RegisterOutputType(NetworkLoadBalancerArrayOutput{})
	pulumi.RegisterOutputType(NetworkLoadBalancerMapOutput{})
}
