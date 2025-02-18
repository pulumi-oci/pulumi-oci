// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Address List resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service.
//
// Creates an address list in a set compartment and allows it to be used in a WAAS policy and referenced by access rules. Addresses can be IP addresses and CIDR notations.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/waas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := waas.NewAddressList(ctx, "testAddressList", &waas.AddressListArgs{
// 			Addresses:     pulumi.Any(_var.Address_list_addresses),
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			DisplayName:   pulumi.Any(_var.Address_list_display_name),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
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
// AddressLists can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:waas/addressList:AddressList test_address_list "id"
// ```
type AddressList struct {
	pulumi.CustomResourceState

	// The total number of unique IP addresses in the address list.
	AddressCount pulumi.Float64Output `pulumi:"addressCount"`
	// (Updatable) A list of IP addresses or CIDR notations.
	Addresses pulumi.StringArrayOutput `pulumi:"addresses"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the address list.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) A user-friendly name for the address list.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// The current lifecycle state of the address list.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the address list was created, expressed in RFC 3339 timestamp format.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
}

// NewAddressList registers a new resource with the given unique name, arguments, and options.
func NewAddressList(ctx *pulumi.Context,
	name string, args *AddressListArgs, opts ...pulumi.ResourceOption) (*AddressList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Addresses == nil {
		return nil, errors.New("invalid value for required argument 'Addresses'")
	}
	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource AddressList
	err := ctx.RegisterResource("oci:waas/addressList:AddressList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddressList gets an existing AddressList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddressList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressListState, opts ...pulumi.ResourceOption) (*AddressList, error) {
	var resource AddressList
	err := ctx.ReadResource("oci:waas/addressList:AddressList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AddressList resources.
type addressListState struct {
	// The total number of unique IP addresses in the address list.
	AddressCount *float64 `pulumi:"addressCount"`
	// (Updatable) A list of IP addresses or CIDR notations.
	Addresses []string `pulumi:"addresses"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the address list.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name for the address list.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The current lifecycle state of the address list.
	State *string `pulumi:"state"`
	// The date and time the address list was created, expressed in RFC 3339 timestamp format.
	TimeCreated *string `pulumi:"timeCreated"`
}

type AddressListState struct {
	// The total number of unique IP addresses in the address list.
	AddressCount pulumi.Float64PtrInput
	// (Updatable) A list of IP addresses or CIDR notations.
	Addresses pulumi.StringArrayInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the address list.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name for the address list.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// The current lifecycle state of the address list.
	State pulumi.StringPtrInput
	// The date and time the address list was created, expressed in RFC 3339 timestamp format.
	TimeCreated pulumi.StringPtrInput
}

func (AddressListState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressListState)(nil)).Elem()
}

type addressListArgs struct {
	// (Updatable) A list of IP addresses or CIDR notations.
	Addresses []string `pulumi:"addresses"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the address list.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name for the address list.
	DisplayName string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
}

// The set of arguments for constructing a AddressList resource.
type AddressListArgs struct {
	// (Updatable) A list of IP addresses or CIDR notations.
	Addresses pulumi.StringArrayInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the address list.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name for the address list.
	DisplayName pulumi.StringInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
}

func (AddressListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressListArgs)(nil)).Elem()
}

type AddressListInput interface {
	pulumi.Input

	ToAddressListOutput() AddressListOutput
	ToAddressListOutputWithContext(ctx context.Context) AddressListOutput
}

func (*AddressList) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressList)(nil))
}

func (i *AddressList) ToAddressListOutput() AddressListOutput {
	return i.ToAddressListOutputWithContext(context.Background())
}

func (i *AddressList) ToAddressListOutputWithContext(ctx context.Context) AddressListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressListOutput)
}

func (i *AddressList) ToAddressListPtrOutput() AddressListPtrOutput {
	return i.ToAddressListPtrOutputWithContext(context.Background())
}

func (i *AddressList) ToAddressListPtrOutputWithContext(ctx context.Context) AddressListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressListPtrOutput)
}

type AddressListPtrInput interface {
	pulumi.Input

	ToAddressListPtrOutput() AddressListPtrOutput
	ToAddressListPtrOutputWithContext(ctx context.Context) AddressListPtrOutput
}

type addressListPtrType AddressListArgs

func (*addressListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressList)(nil))
}

func (i *addressListPtrType) ToAddressListPtrOutput() AddressListPtrOutput {
	return i.ToAddressListPtrOutputWithContext(context.Background())
}

func (i *addressListPtrType) ToAddressListPtrOutputWithContext(ctx context.Context) AddressListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressListPtrOutput)
}

// AddressListArrayInput is an input type that accepts AddressListArray and AddressListArrayOutput values.
// You can construct a concrete instance of `AddressListArrayInput` via:
//
//          AddressListArray{ AddressListArgs{...} }
type AddressListArrayInput interface {
	pulumi.Input

	ToAddressListArrayOutput() AddressListArrayOutput
	ToAddressListArrayOutputWithContext(context.Context) AddressListArrayOutput
}

type AddressListArray []AddressListInput

func (AddressListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AddressList)(nil)).Elem()
}

func (i AddressListArray) ToAddressListArrayOutput() AddressListArrayOutput {
	return i.ToAddressListArrayOutputWithContext(context.Background())
}

func (i AddressListArray) ToAddressListArrayOutputWithContext(ctx context.Context) AddressListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressListArrayOutput)
}

// AddressListMapInput is an input type that accepts AddressListMap and AddressListMapOutput values.
// You can construct a concrete instance of `AddressListMapInput` via:
//
//          AddressListMap{ "key": AddressListArgs{...} }
type AddressListMapInput interface {
	pulumi.Input

	ToAddressListMapOutput() AddressListMapOutput
	ToAddressListMapOutputWithContext(context.Context) AddressListMapOutput
}

type AddressListMap map[string]AddressListInput

func (AddressListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AddressList)(nil)).Elem()
}

func (i AddressListMap) ToAddressListMapOutput() AddressListMapOutput {
	return i.ToAddressListMapOutputWithContext(context.Background())
}

func (i AddressListMap) ToAddressListMapOutputWithContext(ctx context.Context) AddressListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressListMapOutput)
}

type AddressListOutput struct {
	*pulumi.OutputState
}

func (AddressListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressList)(nil))
}

func (o AddressListOutput) ToAddressListOutput() AddressListOutput {
	return o
}

func (o AddressListOutput) ToAddressListOutputWithContext(ctx context.Context) AddressListOutput {
	return o
}

func (o AddressListOutput) ToAddressListPtrOutput() AddressListPtrOutput {
	return o.ToAddressListPtrOutputWithContext(context.Background())
}

func (o AddressListOutput) ToAddressListPtrOutputWithContext(ctx context.Context) AddressListPtrOutput {
	return o.ApplyT(func(v AddressList) *AddressList {
		return &v
	}).(AddressListPtrOutput)
}

type AddressListPtrOutput struct {
	*pulumi.OutputState
}

func (AddressListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressList)(nil))
}

func (o AddressListPtrOutput) ToAddressListPtrOutput() AddressListPtrOutput {
	return o
}

func (o AddressListPtrOutput) ToAddressListPtrOutputWithContext(ctx context.Context) AddressListPtrOutput {
	return o
}

type AddressListArrayOutput struct{ *pulumi.OutputState }

func (AddressListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AddressList)(nil))
}

func (o AddressListArrayOutput) ToAddressListArrayOutput() AddressListArrayOutput {
	return o
}

func (o AddressListArrayOutput) ToAddressListArrayOutputWithContext(ctx context.Context) AddressListArrayOutput {
	return o
}

func (o AddressListArrayOutput) Index(i pulumi.IntInput) AddressListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AddressList {
		return vs[0].([]AddressList)[vs[1].(int)]
	}).(AddressListOutput)
}

type AddressListMapOutput struct{ *pulumi.OutputState }

func (AddressListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AddressList)(nil))
}

func (o AddressListMapOutput) ToAddressListMapOutput() AddressListMapOutput {
	return o
}

func (o AddressListMapOutput) ToAddressListMapOutputWithContext(ctx context.Context) AddressListMapOutput {
	return o
}

func (o AddressListMapOutput) MapIndex(k pulumi.StringInput) AddressListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AddressList {
		return vs[0].(map[string]AddressList)[vs[1].(string)]
	}).(AddressListOutput)
}

func init() {
	pulumi.RegisterOutputType(AddressListOutput{})
	pulumi.RegisterOutputType(AddressListPtrOutput{})
	pulumi.RegisterOutputType(AddressListArrayOutput{})
	pulumi.RegisterOutputType(AddressListMapOutput{})
}
