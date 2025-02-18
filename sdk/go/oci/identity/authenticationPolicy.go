// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Authentication Policy resource in Oracle Cloud Infrastructure Identity service.
//
// Updates authentication policy for the specified tenancy
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
// 		_, err := identity.NewAuthenticationPolicy(ctx, "testAuthenticationPolicy", &identity.AuthenticationPolicyArgs{
// 			CompartmentId: pulumi.Any(_var.Tenancy_ocid),
// 			NetworkPolicy: &identity.AuthenticationPolicyNetworkPolicyArgs{
// 				NetworkSourceIds: pulumi.Any(_var.Authentication_policy_network_policy_network_source_ids),
// 			},
// 			PasswordPolicy: &identity.AuthenticationPolicyPasswordPolicyArgs{
// 				IsLowercaseCharactersRequired: pulumi.Any(_var.Authentication_policy_password_policy_is_lowercase_characters_required),
// 				IsNumericCharactersRequired:   pulumi.Any(_var.Authentication_policy_password_policy_is_numeric_characters_required),
// 				IsSpecialCharactersRequired:   pulumi.Any(_var.Authentication_policy_password_policy_is_special_characters_required),
// 				IsUppercaseCharactersRequired: pulumi.Any(_var.Authentication_policy_password_policy_is_uppercase_characters_required),
// 				IsUsernameContainmentAllowed:  pulumi.Any(_var.Authentication_policy_password_policy_is_username_containment_allowed),
// 				MinimumPasswordLength:         pulumi.Any(_var.Authentication_policy_password_policy_minimum_password_length),
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
// AuthenticationPolicies can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:identity/authenticationPolicy:AuthenticationPolicy test_authentication_policy "authenticationPolicies/{compartmentId}"
// ```
type AuthenticationPolicy struct {
	pulumi.CustomResourceState

	// The OCID of the compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Network policy, Consists of a list of Network Source ids.
	NetworkPolicy AuthenticationPolicyNetworkPolicyOutput `pulumi:"networkPolicy"`
	// (Updatable) Password policy, currently set for the given compartment.
	PasswordPolicy AuthenticationPolicyPasswordPolicyOutput `pulumi:"passwordPolicy"`
}

// NewAuthenticationPolicy registers a new resource with the given unique name, arguments, and options.
func NewAuthenticationPolicy(ctx *pulumi.Context,
	name string, args *AuthenticationPolicyArgs, opts ...pulumi.ResourceOption) (*AuthenticationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	var resource AuthenticationPolicy
	err := ctx.RegisterResource("oci:identity/authenticationPolicy:AuthenticationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthenticationPolicy gets an existing AuthenticationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthenticationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthenticationPolicyState, opts ...pulumi.ResourceOption) (*AuthenticationPolicy, error) {
	var resource AuthenticationPolicy
	err := ctx.ReadResource("oci:identity/authenticationPolicy:AuthenticationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthenticationPolicy resources.
type authenticationPolicyState struct {
	// The OCID of the compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Network policy, Consists of a list of Network Source ids.
	NetworkPolicy *AuthenticationPolicyNetworkPolicy `pulumi:"networkPolicy"`
	// (Updatable) Password policy, currently set for the given compartment.
	PasswordPolicy *AuthenticationPolicyPasswordPolicy `pulumi:"passwordPolicy"`
}

type AuthenticationPolicyState struct {
	// The OCID of the compartment.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Network policy, Consists of a list of Network Source ids.
	NetworkPolicy AuthenticationPolicyNetworkPolicyPtrInput
	// (Updatable) Password policy, currently set for the given compartment.
	PasswordPolicy AuthenticationPolicyPasswordPolicyPtrInput
}

func (AuthenticationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationPolicyState)(nil)).Elem()
}

type authenticationPolicyArgs struct {
	// The OCID of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Network policy, Consists of a list of Network Source ids.
	NetworkPolicy *AuthenticationPolicyNetworkPolicy `pulumi:"networkPolicy"`
	// (Updatable) Password policy, currently set for the given compartment.
	PasswordPolicy *AuthenticationPolicyPasswordPolicy `pulumi:"passwordPolicy"`
}

// The set of arguments for constructing a AuthenticationPolicy resource.
type AuthenticationPolicyArgs struct {
	// The OCID of the compartment.
	CompartmentId pulumi.StringInput
	// (Updatable) Network policy, Consists of a list of Network Source ids.
	NetworkPolicy AuthenticationPolicyNetworkPolicyPtrInput
	// (Updatable) Password policy, currently set for the given compartment.
	PasswordPolicy AuthenticationPolicyPasswordPolicyPtrInput
}

func (AuthenticationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationPolicyArgs)(nil)).Elem()
}

type AuthenticationPolicyInput interface {
	pulumi.Input

	ToAuthenticationPolicyOutput() AuthenticationPolicyOutput
	ToAuthenticationPolicyOutputWithContext(ctx context.Context) AuthenticationPolicyOutput
}

func (*AuthenticationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationPolicy)(nil))
}

func (i *AuthenticationPolicy) ToAuthenticationPolicyOutput() AuthenticationPolicyOutput {
	return i.ToAuthenticationPolicyOutputWithContext(context.Background())
}

func (i *AuthenticationPolicy) ToAuthenticationPolicyOutputWithContext(ctx context.Context) AuthenticationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationPolicyOutput)
}

func (i *AuthenticationPolicy) ToAuthenticationPolicyPtrOutput() AuthenticationPolicyPtrOutput {
	return i.ToAuthenticationPolicyPtrOutputWithContext(context.Background())
}

func (i *AuthenticationPolicy) ToAuthenticationPolicyPtrOutputWithContext(ctx context.Context) AuthenticationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationPolicyPtrOutput)
}

type AuthenticationPolicyPtrInput interface {
	pulumi.Input

	ToAuthenticationPolicyPtrOutput() AuthenticationPolicyPtrOutput
	ToAuthenticationPolicyPtrOutputWithContext(ctx context.Context) AuthenticationPolicyPtrOutput
}

type authenticationPolicyPtrType AuthenticationPolicyArgs

func (*authenticationPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationPolicy)(nil))
}

func (i *authenticationPolicyPtrType) ToAuthenticationPolicyPtrOutput() AuthenticationPolicyPtrOutput {
	return i.ToAuthenticationPolicyPtrOutputWithContext(context.Background())
}

func (i *authenticationPolicyPtrType) ToAuthenticationPolicyPtrOutputWithContext(ctx context.Context) AuthenticationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationPolicyPtrOutput)
}

// AuthenticationPolicyArrayInput is an input type that accepts AuthenticationPolicyArray and AuthenticationPolicyArrayOutput values.
// You can construct a concrete instance of `AuthenticationPolicyArrayInput` via:
//
//          AuthenticationPolicyArray{ AuthenticationPolicyArgs{...} }
type AuthenticationPolicyArrayInput interface {
	pulumi.Input

	ToAuthenticationPolicyArrayOutput() AuthenticationPolicyArrayOutput
	ToAuthenticationPolicyArrayOutputWithContext(context.Context) AuthenticationPolicyArrayOutput
}

type AuthenticationPolicyArray []AuthenticationPolicyInput

func (AuthenticationPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthenticationPolicy)(nil)).Elem()
}

func (i AuthenticationPolicyArray) ToAuthenticationPolicyArrayOutput() AuthenticationPolicyArrayOutput {
	return i.ToAuthenticationPolicyArrayOutputWithContext(context.Background())
}

func (i AuthenticationPolicyArray) ToAuthenticationPolicyArrayOutputWithContext(ctx context.Context) AuthenticationPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationPolicyArrayOutput)
}

// AuthenticationPolicyMapInput is an input type that accepts AuthenticationPolicyMap and AuthenticationPolicyMapOutput values.
// You can construct a concrete instance of `AuthenticationPolicyMapInput` via:
//
//          AuthenticationPolicyMap{ "key": AuthenticationPolicyArgs{...} }
type AuthenticationPolicyMapInput interface {
	pulumi.Input

	ToAuthenticationPolicyMapOutput() AuthenticationPolicyMapOutput
	ToAuthenticationPolicyMapOutputWithContext(context.Context) AuthenticationPolicyMapOutput
}

type AuthenticationPolicyMap map[string]AuthenticationPolicyInput

func (AuthenticationPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthenticationPolicy)(nil)).Elem()
}

func (i AuthenticationPolicyMap) ToAuthenticationPolicyMapOutput() AuthenticationPolicyMapOutput {
	return i.ToAuthenticationPolicyMapOutputWithContext(context.Background())
}

func (i AuthenticationPolicyMap) ToAuthenticationPolicyMapOutputWithContext(ctx context.Context) AuthenticationPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationPolicyMapOutput)
}

type AuthenticationPolicyOutput struct {
	*pulumi.OutputState
}

func (AuthenticationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationPolicy)(nil))
}

func (o AuthenticationPolicyOutput) ToAuthenticationPolicyOutput() AuthenticationPolicyOutput {
	return o
}

func (o AuthenticationPolicyOutput) ToAuthenticationPolicyOutputWithContext(ctx context.Context) AuthenticationPolicyOutput {
	return o
}

func (o AuthenticationPolicyOutput) ToAuthenticationPolicyPtrOutput() AuthenticationPolicyPtrOutput {
	return o.ToAuthenticationPolicyPtrOutputWithContext(context.Background())
}

func (o AuthenticationPolicyOutput) ToAuthenticationPolicyPtrOutputWithContext(ctx context.Context) AuthenticationPolicyPtrOutput {
	return o.ApplyT(func(v AuthenticationPolicy) *AuthenticationPolicy {
		return &v
	}).(AuthenticationPolicyPtrOutput)
}

type AuthenticationPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (AuthenticationPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationPolicy)(nil))
}

func (o AuthenticationPolicyPtrOutput) ToAuthenticationPolicyPtrOutput() AuthenticationPolicyPtrOutput {
	return o
}

func (o AuthenticationPolicyPtrOutput) ToAuthenticationPolicyPtrOutputWithContext(ctx context.Context) AuthenticationPolicyPtrOutput {
	return o
}

type AuthenticationPolicyArrayOutput struct{ *pulumi.OutputState }

func (AuthenticationPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthenticationPolicy)(nil))
}

func (o AuthenticationPolicyArrayOutput) ToAuthenticationPolicyArrayOutput() AuthenticationPolicyArrayOutput {
	return o
}

func (o AuthenticationPolicyArrayOutput) ToAuthenticationPolicyArrayOutputWithContext(ctx context.Context) AuthenticationPolicyArrayOutput {
	return o
}

func (o AuthenticationPolicyArrayOutput) Index(i pulumi.IntInput) AuthenticationPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthenticationPolicy {
		return vs[0].([]AuthenticationPolicy)[vs[1].(int)]
	}).(AuthenticationPolicyOutput)
}

type AuthenticationPolicyMapOutput struct{ *pulumi.OutputState }

func (AuthenticationPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AuthenticationPolicy)(nil))
}

func (o AuthenticationPolicyMapOutput) ToAuthenticationPolicyMapOutput() AuthenticationPolicyMapOutput {
	return o
}

func (o AuthenticationPolicyMapOutput) ToAuthenticationPolicyMapOutputWithContext(ctx context.Context) AuthenticationPolicyMapOutput {
	return o
}

func (o AuthenticationPolicyMapOutput) MapIndex(k pulumi.StringInput) AuthenticationPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AuthenticationPolicy {
		return vs[0].(map[string]AuthenticationPolicy)[vs[1].(string)]
	}).(AuthenticationPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthenticationPolicyOutput{})
	pulumi.RegisterOutputType(AuthenticationPolicyPtrOutput{})
	pulumi.RegisterOutputType(AuthenticationPolicyArrayOutput{})
	pulumi.RegisterOutputType(AuthenticationPolicyMapOutput{})
}
