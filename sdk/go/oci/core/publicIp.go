// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Public Ip resource in Oracle Cloud Infrastructure Core service.
//
// Creates a public IP. Use the `lifetime` property to specify whether it's an ephemeral or
// reserved public IP. For information about limits on how many you can create, see
// [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
//
// * **For an ephemeral public IP assigned to a private IP:** You must also specify a `privateIpId`
// with the OCID of the primary private IP you want to assign the public IP to. The public IP is
// created in the same availability domain as the private IP. An ephemeral public IP must always be
// assigned to a private IP, and only to the *primary* private IP on a VNIC, not a secondary
// private IP. Exception: If you create a [NatGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NatGateway/), Oracle
// automatically assigns the NAT gateway a regional ephemeral public IP that you cannot remove.
//
// * **For a reserved public IP:** You may also optionally assign the public IP to a private
// IP by specifying `privateIpId`. Or you can later assign the public IP with
// [UpdatePublicIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/UpdatePublicIp).
//
// **Note:** When assigning a public IP to a private IP, the private IP must not already have
// a public IP with `lifecycleState` = ASSIGNING or ASSIGNED. If it does, an error is returned.
//
// Also, for reserved public IPs, the optional assignment part of this operation is
// asynchronous. Poll the public IP's `lifecycleState` to determine if the assignment
// succeeded.
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
// 		_, err := core.NewPublicIp(ctx, "testPublicIp", &core.PublicIpArgs{
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			Lifetime:      pulumi.Any(_var.Public_ip_lifetime),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			DisplayName: pulumi.Any(_var.Public_ip_display_name),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			PrivateIpId:    pulumi.Any(oci_core_private_ip.Test_private_ip.Id),
// 			PublicIpPoolId: pulumi.Any(oci_core_public_ip_pool.Test_public_ip_pool.Id),
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
// PublicIps can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:core/publicIp:PublicIp test_public_ip "id"
// ```
type PublicIp struct {
	pulumi.CustomResourceState

	// The OCID of the entity the public IP is assigned to, or in the process of being assigned to.
	AssignedEntityId pulumi.StringOutput `pulumi:"assignedEntityId"`
	// The type of entity the public IP is assigned to, or in the process of being assigned to.
	AssignedEntityType pulumi.StringOutput `pulumi:"assignedEntityType"`
	// The public IP's availability domain. This property is set only for ephemeral public IPs that are assigned to a private IP (that is, when the `scope` of the public IP is set to AVAILABILITY_DOMAIN). The value is the availability domain of the assigned private IP.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringOutput `pulumi:"availabilityDomain"`
	// (Updatable) The OCID of the compartment to contain the public IP. For ephemeral public IPs, you must set this to the private IP's compartment OCID.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// The public IP address of the `publicIp` object.  Example: `203.0.113.2`
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Defines when the public IP is deleted and released back to the Oracle Cloud Infrastructure public IP pool. For more information, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
	Lifetime pulumi.StringOutput `pulumi:"lifetime"`
	// (Updatable) The OCID of the private IP to assign the public IP to.
	PrivateIpId pulumi.StringPtrOutput `pulumi:"privateIpId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
	PublicIpPoolId pulumi.StringOutput `pulumi:"publicIpPoolId"`
	// Whether the public IP is regional or specific to a particular availability domain.
	// * `REGION`: The public IP exists within a region and is assigned to a regional entity (such as a [NatGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NatGateway/)), or can be assigned to a private IP in any availability domain in the region. Reserved public IPs and ephemeral public IPs assigned to a regional entity have `scope` = `REGION`.
	// * `AVAILABILITY_DOMAIN`: The public IP exists within the availability domain of the entity it's assigned to, which is specified by the `availabilityDomain` property of the public IP object. Ephemeral public IPs that are assigned to private IPs have `scope` = `AVAILABILITY_DOMAIN`.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// The public IP's current state.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the public IP was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
}

// NewPublicIp registers a new resource with the given unique name, arguments, and options.
func NewPublicIp(ctx *pulumi.Context,
	name string, args *PublicIpArgs, opts ...pulumi.ResourceOption) (*PublicIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.Lifetime == nil {
		return nil, errors.New("invalid value for required argument 'Lifetime'")
	}
	var resource PublicIp
	err := ctx.RegisterResource("oci:core/publicIp:PublicIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIp gets an existing PublicIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIpState, opts ...pulumi.ResourceOption) (*PublicIp, error) {
	var resource PublicIp
	err := ctx.ReadResource("oci:core/publicIp:PublicIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIp resources.
type publicIpState struct {
	// The OCID of the entity the public IP is assigned to, or in the process of being assigned to.
	AssignedEntityId *string `pulumi:"assignedEntityId"`
	// The type of entity the public IP is assigned to, or in the process of being assigned to.
	AssignedEntityType *string `pulumi:"assignedEntityType"`
	// The public IP's availability domain. This property is set only for ephemeral public IPs that are assigned to a private IP (that is, when the `scope` of the public IP is set to AVAILABILITY_DOMAIN). The value is the availability domain of the assigned private IP.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	// (Updatable) The OCID of the compartment to contain the public IP. For ephemeral public IPs, you must set this to the private IP's compartment OCID.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The public IP address of the `publicIp` object.  Example: `203.0.113.2`
	IpAddress *string `pulumi:"ipAddress"`
	// Defines when the public IP is deleted and released back to the Oracle Cloud Infrastructure public IP pool. For more information, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
	Lifetime *string `pulumi:"lifetime"`
	// (Updatable) The OCID of the private IP to assign the public IP to.
	PrivateIpId *string `pulumi:"privateIpId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
	PublicIpPoolId *string `pulumi:"publicIpPoolId"`
	// Whether the public IP is regional or specific to a particular availability domain.
	// * `REGION`: The public IP exists within a region and is assigned to a regional entity (such as a [NatGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NatGateway/)), or can be assigned to a private IP in any availability domain in the region. Reserved public IPs and ephemeral public IPs assigned to a regional entity have `scope` = `REGION`.
	// * `AVAILABILITY_DOMAIN`: The public IP exists within the availability domain of the entity it's assigned to, which is specified by the `availabilityDomain` property of the public IP object. Ephemeral public IPs that are assigned to private IPs have `scope` = `AVAILABILITY_DOMAIN`.
	Scope *string `pulumi:"scope"`
	// The public IP's current state.
	State *string `pulumi:"state"`
	// The date and time the public IP was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *string `pulumi:"timeCreated"`
}

type PublicIpState struct {
	// The OCID of the entity the public IP is assigned to, or in the process of being assigned to.
	AssignedEntityId pulumi.StringPtrInput
	// The type of entity the public IP is assigned to, or in the process of being assigned to.
	AssignedEntityType pulumi.StringPtrInput
	// The public IP's availability domain. This property is set only for ephemeral public IPs that are assigned to a private IP (that is, when the `scope` of the public IP is set to AVAILABILITY_DOMAIN). The value is the availability domain of the assigned private IP.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringPtrInput
	// (Updatable) The OCID of the compartment to contain the public IP. For ephemeral public IPs, you must set this to the private IP's compartment OCID.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// The public IP address of the `publicIp` object.  Example: `203.0.113.2`
	IpAddress pulumi.StringPtrInput
	// Defines when the public IP is deleted and released back to the Oracle Cloud Infrastructure public IP pool. For more information, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
	Lifetime pulumi.StringPtrInput
	// (Updatable) The OCID of the private IP to assign the public IP to.
	PrivateIpId pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
	PublicIpPoolId pulumi.StringPtrInput
	// Whether the public IP is regional or specific to a particular availability domain.
	// * `REGION`: The public IP exists within a region and is assigned to a regional entity (such as a [NatGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NatGateway/)), or can be assigned to a private IP in any availability domain in the region. Reserved public IPs and ephemeral public IPs assigned to a regional entity have `scope` = `REGION`.
	// * `AVAILABILITY_DOMAIN`: The public IP exists within the availability domain of the entity it's assigned to, which is specified by the `availabilityDomain` property of the public IP object. Ephemeral public IPs that are assigned to private IPs have `scope` = `AVAILABILITY_DOMAIN`.
	Scope pulumi.StringPtrInput
	// The public IP's current state.
	State pulumi.StringPtrInput
	// The date and time the public IP was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringPtrInput
}

func (PublicIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpState)(nil)).Elem()
}

type publicIpArgs struct {
	// (Updatable) The OCID of the compartment to contain the public IP. For ephemeral public IPs, you must set this to the private IP's compartment OCID.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Defines when the public IP is deleted and released back to the Oracle Cloud Infrastructure public IP pool. For more information, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
	Lifetime string `pulumi:"lifetime"`
	// (Updatable) The OCID of the private IP to assign the public IP to.
	PrivateIpId *string `pulumi:"privateIpId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
	PublicIpPoolId *string `pulumi:"publicIpPoolId"`
}

// The set of arguments for constructing a PublicIp resource.
type PublicIpArgs struct {
	// (Updatable) The OCID of the compartment to contain the public IP. For ephemeral public IPs, you must set this to the private IP's compartment OCID.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// Defines when the public IP is deleted and released back to the Oracle Cloud Infrastructure public IP pool. For more information, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
	Lifetime pulumi.StringInput
	// (Updatable) The OCID of the private IP to assign the public IP to.
	PrivateIpId pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
	PublicIpPoolId pulumi.StringPtrInput
}

func (PublicIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpArgs)(nil)).Elem()
}

type PublicIpInput interface {
	pulumi.Input

	ToPublicIpOutput() PublicIpOutput
	ToPublicIpOutputWithContext(ctx context.Context) PublicIpOutput
}

func (*PublicIp) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIp)(nil))
}

func (i *PublicIp) ToPublicIpOutput() PublicIpOutput {
	return i.ToPublicIpOutputWithContext(context.Background())
}

func (i *PublicIp) ToPublicIpOutputWithContext(ctx context.Context) PublicIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpOutput)
}

func (i *PublicIp) ToPublicIpPtrOutput() PublicIpPtrOutput {
	return i.ToPublicIpPtrOutputWithContext(context.Background())
}

func (i *PublicIp) ToPublicIpPtrOutputWithContext(ctx context.Context) PublicIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpPtrOutput)
}

type PublicIpPtrInput interface {
	pulumi.Input

	ToPublicIpPtrOutput() PublicIpPtrOutput
	ToPublicIpPtrOutputWithContext(ctx context.Context) PublicIpPtrOutput
}

type publicIpPtrType PublicIpArgs

func (*publicIpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIp)(nil))
}

func (i *publicIpPtrType) ToPublicIpPtrOutput() PublicIpPtrOutput {
	return i.ToPublicIpPtrOutputWithContext(context.Background())
}

func (i *publicIpPtrType) ToPublicIpPtrOutputWithContext(ctx context.Context) PublicIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpPtrOutput)
}

// PublicIpArrayInput is an input type that accepts PublicIpArray and PublicIpArrayOutput values.
// You can construct a concrete instance of `PublicIpArrayInput` via:
//
//          PublicIpArray{ PublicIpArgs{...} }
type PublicIpArrayInput interface {
	pulumi.Input

	ToPublicIpArrayOutput() PublicIpArrayOutput
	ToPublicIpArrayOutputWithContext(context.Context) PublicIpArrayOutput
}

type PublicIpArray []PublicIpInput

func (PublicIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicIp)(nil)).Elem()
}

func (i PublicIpArray) ToPublicIpArrayOutput() PublicIpArrayOutput {
	return i.ToPublicIpArrayOutputWithContext(context.Background())
}

func (i PublicIpArray) ToPublicIpArrayOutputWithContext(ctx context.Context) PublicIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpArrayOutput)
}

// PublicIpMapInput is an input type that accepts PublicIpMap and PublicIpMapOutput values.
// You can construct a concrete instance of `PublicIpMapInput` via:
//
//          PublicIpMap{ "key": PublicIpArgs{...} }
type PublicIpMapInput interface {
	pulumi.Input

	ToPublicIpMapOutput() PublicIpMapOutput
	ToPublicIpMapOutputWithContext(context.Context) PublicIpMapOutput
}

type PublicIpMap map[string]PublicIpInput

func (PublicIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicIp)(nil)).Elem()
}

func (i PublicIpMap) ToPublicIpMapOutput() PublicIpMapOutput {
	return i.ToPublicIpMapOutputWithContext(context.Background())
}

func (i PublicIpMap) ToPublicIpMapOutputWithContext(ctx context.Context) PublicIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpMapOutput)
}

type PublicIpOutput struct {
	*pulumi.OutputState
}

func (PublicIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIp)(nil))
}

func (o PublicIpOutput) ToPublicIpOutput() PublicIpOutput {
	return o
}

func (o PublicIpOutput) ToPublicIpOutputWithContext(ctx context.Context) PublicIpOutput {
	return o
}

func (o PublicIpOutput) ToPublicIpPtrOutput() PublicIpPtrOutput {
	return o.ToPublicIpPtrOutputWithContext(context.Background())
}

func (o PublicIpOutput) ToPublicIpPtrOutputWithContext(ctx context.Context) PublicIpPtrOutput {
	return o.ApplyT(func(v PublicIp) *PublicIp {
		return &v
	}).(PublicIpPtrOutput)
}

type PublicIpPtrOutput struct {
	*pulumi.OutputState
}

func (PublicIpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIp)(nil))
}

func (o PublicIpPtrOutput) ToPublicIpPtrOutput() PublicIpPtrOutput {
	return o
}

func (o PublicIpPtrOutput) ToPublicIpPtrOutputWithContext(ctx context.Context) PublicIpPtrOutput {
	return o
}

type PublicIpArrayOutput struct{ *pulumi.OutputState }

func (PublicIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PublicIp)(nil))
}

func (o PublicIpArrayOutput) ToPublicIpArrayOutput() PublicIpArrayOutput {
	return o
}

func (o PublicIpArrayOutput) ToPublicIpArrayOutputWithContext(ctx context.Context) PublicIpArrayOutput {
	return o
}

func (o PublicIpArrayOutput) Index(i pulumi.IntInput) PublicIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PublicIp {
		return vs[0].([]PublicIp)[vs[1].(int)]
	}).(PublicIpOutput)
}

type PublicIpMapOutput struct{ *pulumi.OutputState }

func (PublicIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PublicIp)(nil))
}

func (o PublicIpMapOutput) ToPublicIpMapOutput() PublicIpMapOutput {
	return o
}

func (o PublicIpMapOutput) ToPublicIpMapOutputWithContext(ctx context.Context) PublicIpMapOutput {
	return o
}

func (o PublicIpMapOutput) MapIndex(k pulumi.StringInput) PublicIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PublicIp {
		return vs[0].(map[string]PublicIp)[vs[1].(string)]
	}).(PublicIpOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicIpOutput{})
	pulumi.RegisterOutputType(PublicIpPtrOutput{})
	pulumi.RegisterOutputType(PublicIpArrayOutput{})
	pulumi.RegisterOutputType(PublicIpMapOutput{})
}
