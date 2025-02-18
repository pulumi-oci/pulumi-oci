// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Customer Secret Key resource in Oracle Cloud Infrastructure Identity service.
//
// Creates a new secret key for the specified user. Secret keys are used for authentication with the Object Storage Service's Amazon S3
// compatible API. The secret key consists of an Access Key/Secret Key pair. For information, see
// [Managing User Credentials](https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcredentials.htm).
//
// You must specify a *description* for the secret key (although it can be an empty string). It does not
// have to be unique, and you can change it anytime with
// [UpdateCustomerSecretKey](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/CustomerSecretKeySummary/UpdateCustomerSecretKey).
//
// Every user has permission to create a secret key for *their own user ID*. An administrator in your organization
// does not need to write a policy to give users this ability. To compare, administrators who have permission to the
// tenancy can use this operation to create a secret key for any user, including themselves.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/identity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := identity.NewCustomerSecretKey(ctx, "testCustomerSecretKey", &identity.CustomerSecretKeyArgs{
// 			DisplayName: pulumi.Any(_var.Customer_secret_key_display_name),
// 			UserId:      pulumi.Any(oci_identity_user.Test_user.Id),
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
// CustomerSecretKeys can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:identity/customerSecretKey:CustomerSecretKey test_customer_secret_key "users/{userId}/customerSecretKeys/{customerSecretKeyId}"
// ```
type CustomerSecretKey struct {
	pulumi.CustomResourceState

	// (Updatable) The name you assign to the secret key during creation. Does not have to be unique, and it's changeable.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The detailed status of INACTIVE lifecycleState.
	InactiveState pulumi.StringOutput `pulumi:"inactiveState"`
	// The secret key.
	Key pulumi.StringOutput `pulumi:"key"`
	// The secret key's current state.
	State pulumi.StringOutput `pulumi:"state"`
	// Date and time the `CustomerSecretKey` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z`
	TimeExpires pulumi.StringOutput `pulumi:"timeExpires"`
	// The OCID of the user.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewCustomerSecretKey registers a new resource with the given unique name, arguments, and options.
func NewCustomerSecretKey(ctx *pulumi.Context,
	name string, args *CustomerSecretKeyArgs, opts ...pulumi.ResourceOption) (*CustomerSecretKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource CustomerSecretKey
	err := ctx.RegisterResource("oci:identity/customerSecretKey:CustomerSecretKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerSecretKey gets an existing CustomerSecretKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerSecretKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerSecretKeyState, opts ...pulumi.ResourceOption) (*CustomerSecretKey, error) {
	var resource CustomerSecretKey
	err := ctx.ReadResource("oci:identity/customerSecretKey:CustomerSecretKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerSecretKey resources.
type customerSecretKeyState struct {
	// (Updatable) The name you assign to the secret key during creation. Does not have to be unique, and it's changeable.
	DisplayName *string `pulumi:"displayName"`
	// The detailed status of INACTIVE lifecycleState.
	InactiveState *string `pulumi:"inactiveState"`
	// The secret key.
	Key *string `pulumi:"key"`
	// The secret key's current state.
	State *string `pulumi:"state"`
	// Date and time the `CustomerSecretKey` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *string `pulumi:"timeCreated"`
	// Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z`
	TimeExpires *string `pulumi:"timeExpires"`
	// The OCID of the user.
	UserId *string `pulumi:"userId"`
}

type CustomerSecretKeyState struct {
	// (Updatable) The name you assign to the secret key during creation. Does not have to be unique, and it's changeable.
	DisplayName pulumi.StringPtrInput
	// The detailed status of INACTIVE lifecycleState.
	InactiveState pulumi.StringPtrInput
	// The secret key.
	Key pulumi.StringPtrInput
	// The secret key's current state.
	State pulumi.StringPtrInput
	// Date and time the `CustomerSecretKey` object was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringPtrInput
	// Date and time when this password will expire, in the format defined by RFC3339. Null if it never expires.  Example: `2016-08-25T21:10:29.600Z`
	TimeExpires pulumi.StringPtrInput
	// The OCID of the user.
	UserId pulumi.StringPtrInput
}

func (CustomerSecretKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerSecretKeyState)(nil)).Elem()
}

type customerSecretKeyArgs struct {
	// (Updatable) The name you assign to the secret key during creation. Does not have to be unique, and it's changeable.
	DisplayName string `pulumi:"displayName"`
	// The OCID of the user.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a CustomerSecretKey resource.
type CustomerSecretKeyArgs struct {
	// (Updatable) The name you assign to the secret key during creation. Does not have to be unique, and it's changeable.
	DisplayName pulumi.StringInput
	// The OCID of the user.
	UserId pulumi.StringInput
}

func (CustomerSecretKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerSecretKeyArgs)(nil)).Elem()
}

type CustomerSecretKeyInput interface {
	pulumi.Input

	ToCustomerSecretKeyOutput() CustomerSecretKeyOutput
	ToCustomerSecretKeyOutputWithContext(ctx context.Context) CustomerSecretKeyOutput
}

func (*CustomerSecretKey) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSecretKey)(nil))
}

func (i *CustomerSecretKey) ToCustomerSecretKeyOutput() CustomerSecretKeyOutput {
	return i.ToCustomerSecretKeyOutputWithContext(context.Background())
}

func (i *CustomerSecretKey) ToCustomerSecretKeyOutputWithContext(ctx context.Context) CustomerSecretKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSecretKeyOutput)
}

func (i *CustomerSecretKey) ToCustomerSecretKeyPtrOutput() CustomerSecretKeyPtrOutput {
	return i.ToCustomerSecretKeyPtrOutputWithContext(context.Background())
}

func (i *CustomerSecretKey) ToCustomerSecretKeyPtrOutputWithContext(ctx context.Context) CustomerSecretKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSecretKeyPtrOutput)
}

type CustomerSecretKeyPtrInput interface {
	pulumi.Input

	ToCustomerSecretKeyPtrOutput() CustomerSecretKeyPtrOutput
	ToCustomerSecretKeyPtrOutputWithContext(ctx context.Context) CustomerSecretKeyPtrOutput
}

type customerSecretKeyPtrType CustomerSecretKeyArgs

func (*customerSecretKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerSecretKey)(nil))
}

func (i *customerSecretKeyPtrType) ToCustomerSecretKeyPtrOutput() CustomerSecretKeyPtrOutput {
	return i.ToCustomerSecretKeyPtrOutputWithContext(context.Background())
}

func (i *customerSecretKeyPtrType) ToCustomerSecretKeyPtrOutputWithContext(ctx context.Context) CustomerSecretKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSecretKeyPtrOutput)
}

// CustomerSecretKeyArrayInput is an input type that accepts CustomerSecretKeyArray and CustomerSecretKeyArrayOutput values.
// You can construct a concrete instance of `CustomerSecretKeyArrayInput` via:
//
//          CustomerSecretKeyArray{ CustomerSecretKeyArgs{...} }
type CustomerSecretKeyArrayInput interface {
	pulumi.Input

	ToCustomerSecretKeyArrayOutput() CustomerSecretKeyArrayOutput
	ToCustomerSecretKeyArrayOutputWithContext(context.Context) CustomerSecretKeyArrayOutput
}

type CustomerSecretKeyArray []CustomerSecretKeyInput

func (CustomerSecretKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomerSecretKey)(nil)).Elem()
}

func (i CustomerSecretKeyArray) ToCustomerSecretKeyArrayOutput() CustomerSecretKeyArrayOutput {
	return i.ToCustomerSecretKeyArrayOutputWithContext(context.Background())
}

func (i CustomerSecretKeyArray) ToCustomerSecretKeyArrayOutputWithContext(ctx context.Context) CustomerSecretKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSecretKeyArrayOutput)
}

// CustomerSecretKeyMapInput is an input type that accepts CustomerSecretKeyMap and CustomerSecretKeyMapOutput values.
// You can construct a concrete instance of `CustomerSecretKeyMapInput` via:
//
//          CustomerSecretKeyMap{ "key": CustomerSecretKeyArgs{...} }
type CustomerSecretKeyMapInput interface {
	pulumi.Input

	ToCustomerSecretKeyMapOutput() CustomerSecretKeyMapOutput
	ToCustomerSecretKeyMapOutputWithContext(context.Context) CustomerSecretKeyMapOutput
}

type CustomerSecretKeyMap map[string]CustomerSecretKeyInput

func (CustomerSecretKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomerSecretKey)(nil)).Elem()
}

func (i CustomerSecretKeyMap) ToCustomerSecretKeyMapOutput() CustomerSecretKeyMapOutput {
	return i.ToCustomerSecretKeyMapOutputWithContext(context.Background())
}

func (i CustomerSecretKeyMap) ToCustomerSecretKeyMapOutputWithContext(ctx context.Context) CustomerSecretKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSecretKeyMapOutput)
}

type CustomerSecretKeyOutput struct {
	*pulumi.OutputState
}

func (CustomerSecretKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSecretKey)(nil))
}

func (o CustomerSecretKeyOutput) ToCustomerSecretKeyOutput() CustomerSecretKeyOutput {
	return o
}

func (o CustomerSecretKeyOutput) ToCustomerSecretKeyOutputWithContext(ctx context.Context) CustomerSecretKeyOutput {
	return o
}

func (o CustomerSecretKeyOutput) ToCustomerSecretKeyPtrOutput() CustomerSecretKeyPtrOutput {
	return o.ToCustomerSecretKeyPtrOutputWithContext(context.Background())
}

func (o CustomerSecretKeyOutput) ToCustomerSecretKeyPtrOutputWithContext(ctx context.Context) CustomerSecretKeyPtrOutput {
	return o.ApplyT(func(v CustomerSecretKey) *CustomerSecretKey {
		return &v
	}).(CustomerSecretKeyPtrOutput)
}

type CustomerSecretKeyPtrOutput struct {
	*pulumi.OutputState
}

func (CustomerSecretKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerSecretKey)(nil))
}

func (o CustomerSecretKeyPtrOutput) ToCustomerSecretKeyPtrOutput() CustomerSecretKeyPtrOutput {
	return o
}

func (o CustomerSecretKeyPtrOutput) ToCustomerSecretKeyPtrOutputWithContext(ctx context.Context) CustomerSecretKeyPtrOutput {
	return o
}

type CustomerSecretKeyArrayOutput struct{ *pulumi.OutputState }

func (CustomerSecretKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomerSecretKey)(nil))
}

func (o CustomerSecretKeyArrayOutput) ToCustomerSecretKeyArrayOutput() CustomerSecretKeyArrayOutput {
	return o
}

func (o CustomerSecretKeyArrayOutput) ToCustomerSecretKeyArrayOutputWithContext(ctx context.Context) CustomerSecretKeyArrayOutput {
	return o
}

func (o CustomerSecretKeyArrayOutput) Index(i pulumi.IntInput) CustomerSecretKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomerSecretKey {
		return vs[0].([]CustomerSecretKey)[vs[1].(int)]
	}).(CustomerSecretKeyOutput)
}

type CustomerSecretKeyMapOutput struct{ *pulumi.OutputState }

func (CustomerSecretKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomerSecretKey)(nil))
}

func (o CustomerSecretKeyMapOutput) ToCustomerSecretKeyMapOutput() CustomerSecretKeyMapOutput {
	return o
}

func (o CustomerSecretKeyMapOutput) ToCustomerSecretKeyMapOutputWithContext(ctx context.Context) CustomerSecretKeyMapOutput {
	return o
}

func (o CustomerSecretKeyMapOutput) MapIndex(k pulumi.StringInput) CustomerSecretKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomerSecretKey {
		return vs[0].(map[string]CustomerSecretKey)[vs[1].(string)]
	}).(CustomerSecretKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomerSecretKeyOutput{})
	pulumi.RegisterOutputType(CustomerSecretKeyPtrOutput{})
	pulumi.RegisterOutputType(CustomerSecretKeyArrayOutput{})
	pulumi.RegisterOutputType(CustomerSecretKeyMapOutput{})
}
