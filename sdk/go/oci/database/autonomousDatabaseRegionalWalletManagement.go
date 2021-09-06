// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Autonomous Database Regional Wallet Management resource in Oracle Cloud Infrastructure Database service.
//
// Updates the Autonomous Database regional wallet.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/database"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := database.NewAutonomousDatabaseRegionalWalletManagement(ctx, "testAutonomousDatabaseRegionalWalletManagement", &database.AutonomousDatabaseRegionalWalletManagementArgs{
// 			ShouldRotate: pulumi.Any(_var.Autonomous_database_regional_wallet_management_should_rotate),
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
// Import is not supported for this resource.
type AutonomousDatabaseRegionalWalletManagement struct {
	pulumi.CustomResourceState

	// (Updatable) Indicates whether to rotate the wallet or not. If `false`, the wallet will not be rotated. The default is `false`.
	ShouldRotate pulumi.BoolPtrOutput `pulumi:"shouldRotate"`
	// The current lifecycle state of the Autonomous Database wallet.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the wallet was last rotated.
	TimeRotated pulumi.StringOutput `pulumi:"timeRotated"`
}

// NewAutonomousDatabaseRegionalWalletManagement registers a new resource with the given unique name, arguments, and options.
func NewAutonomousDatabaseRegionalWalletManagement(ctx *pulumi.Context,
	name string, args *AutonomousDatabaseRegionalWalletManagementArgs, opts ...pulumi.ResourceOption) (*AutonomousDatabaseRegionalWalletManagement, error) {
	if args == nil {
		args = &AutonomousDatabaseRegionalWalletManagementArgs{}
	}

	var resource AutonomousDatabaseRegionalWalletManagement
	err := ctx.RegisterResource("oci:database/autonomousDatabaseRegionalWalletManagement:AutonomousDatabaseRegionalWalletManagement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutonomousDatabaseRegionalWalletManagement gets an existing AutonomousDatabaseRegionalWalletManagement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutonomousDatabaseRegionalWalletManagement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutonomousDatabaseRegionalWalletManagementState, opts ...pulumi.ResourceOption) (*AutonomousDatabaseRegionalWalletManagement, error) {
	var resource AutonomousDatabaseRegionalWalletManagement
	err := ctx.ReadResource("oci:database/autonomousDatabaseRegionalWalletManagement:AutonomousDatabaseRegionalWalletManagement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutonomousDatabaseRegionalWalletManagement resources.
type autonomousDatabaseRegionalWalletManagementState struct {
	// (Updatable) Indicates whether to rotate the wallet or not. If `false`, the wallet will not be rotated. The default is `false`.
	ShouldRotate *bool `pulumi:"shouldRotate"`
	// The current lifecycle state of the Autonomous Database wallet.
	State *string `pulumi:"state"`
	// The date and time the wallet was last rotated.
	TimeRotated *string `pulumi:"timeRotated"`
}

type AutonomousDatabaseRegionalWalletManagementState struct {
	// (Updatable) Indicates whether to rotate the wallet or not. If `false`, the wallet will not be rotated. The default is `false`.
	ShouldRotate pulumi.BoolPtrInput
	// The current lifecycle state of the Autonomous Database wallet.
	State pulumi.StringPtrInput
	// The date and time the wallet was last rotated.
	TimeRotated pulumi.StringPtrInput
}

func (AutonomousDatabaseRegionalWalletManagementState) ElementType() reflect.Type {
	return reflect.TypeOf((*autonomousDatabaseRegionalWalletManagementState)(nil)).Elem()
}

type autonomousDatabaseRegionalWalletManagementArgs struct {
	// (Updatable) Indicates whether to rotate the wallet or not. If `false`, the wallet will not be rotated. The default is `false`.
	ShouldRotate *bool `pulumi:"shouldRotate"`
}

// The set of arguments for constructing a AutonomousDatabaseRegionalWalletManagement resource.
type AutonomousDatabaseRegionalWalletManagementArgs struct {
	// (Updatable) Indicates whether to rotate the wallet or not. If `false`, the wallet will not be rotated. The default is `false`.
	ShouldRotate pulumi.BoolPtrInput
}

func (AutonomousDatabaseRegionalWalletManagementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autonomousDatabaseRegionalWalletManagementArgs)(nil)).Elem()
}

type AutonomousDatabaseRegionalWalletManagementInput interface {
	pulumi.Input

	ToAutonomousDatabaseRegionalWalletManagementOutput() AutonomousDatabaseRegionalWalletManagementOutput
	ToAutonomousDatabaseRegionalWalletManagementOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementOutput
}

func (*AutonomousDatabaseRegionalWalletManagement) ElementType() reflect.Type {
	return reflect.TypeOf((*AutonomousDatabaseRegionalWalletManagement)(nil))
}

func (i *AutonomousDatabaseRegionalWalletManagement) ToAutonomousDatabaseRegionalWalletManagementOutput() AutonomousDatabaseRegionalWalletManagementOutput {
	return i.ToAutonomousDatabaseRegionalWalletManagementOutputWithContext(context.Background())
}

func (i *AutonomousDatabaseRegionalWalletManagement) ToAutonomousDatabaseRegionalWalletManagementOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutonomousDatabaseRegionalWalletManagementOutput)
}

func (i *AutonomousDatabaseRegionalWalletManagement) ToAutonomousDatabaseRegionalWalletManagementPtrOutput() AutonomousDatabaseRegionalWalletManagementPtrOutput {
	return i.ToAutonomousDatabaseRegionalWalletManagementPtrOutputWithContext(context.Background())
}

func (i *AutonomousDatabaseRegionalWalletManagement) ToAutonomousDatabaseRegionalWalletManagementPtrOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutonomousDatabaseRegionalWalletManagementPtrOutput)
}

type AutonomousDatabaseRegionalWalletManagementPtrInput interface {
	pulumi.Input

	ToAutonomousDatabaseRegionalWalletManagementPtrOutput() AutonomousDatabaseRegionalWalletManagementPtrOutput
	ToAutonomousDatabaseRegionalWalletManagementPtrOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementPtrOutput
}

type autonomousDatabaseRegionalWalletManagementPtrType AutonomousDatabaseRegionalWalletManagementArgs

func (*autonomousDatabaseRegionalWalletManagementPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutonomousDatabaseRegionalWalletManagement)(nil))
}

func (i *autonomousDatabaseRegionalWalletManagementPtrType) ToAutonomousDatabaseRegionalWalletManagementPtrOutput() AutonomousDatabaseRegionalWalletManagementPtrOutput {
	return i.ToAutonomousDatabaseRegionalWalletManagementPtrOutputWithContext(context.Background())
}

func (i *autonomousDatabaseRegionalWalletManagementPtrType) ToAutonomousDatabaseRegionalWalletManagementPtrOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutonomousDatabaseRegionalWalletManagementPtrOutput)
}

// AutonomousDatabaseRegionalWalletManagementArrayInput is an input type that accepts AutonomousDatabaseRegionalWalletManagementArray and AutonomousDatabaseRegionalWalletManagementArrayOutput values.
// You can construct a concrete instance of `AutonomousDatabaseRegionalWalletManagementArrayInput` via:
//
//          AutonomousDatabaseRegionalWalletManagementArray{ AutonomousDatabaseRegionalWalletManagementArgs{...} }
type AutonomousDatabaseRegionalWalletManagementArrayInput interface {
	pulumi.Input

	ToAutonomousDatabaseRegionalWalletManagementArrayOutput() AutonomousDatabaseRegionalWalletManagementArrayOutput
	ToAutonomousDatabaseRegionalWalletManagementArrayOutputWithContext(context.Context) AutonomousDatabaseRegionalWalletManagementArrayOutput
}

type AutonomousDatabaseRegionalWalletManagementArray []AutonomousDatabaseRegionalWalletManagementInput

func (AutonomousDatabaseRegionalWalletManagementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutonomousDatabaseRegionalWalletManagement)(nil)).Elem()
}

func (i AutonomousDatabaseRegionalWalletManagementArray) ToAutonomousDatabaseRegionalWalletManagementArrayOutput() AutonomousDatabaseRegionalWalletManagementArrayOutput {
	return i.ToAutonomousDatabaseRegionalWalletManagementArrayOutputWithContext(context.Background())
}

func (i AutonomousDatabaseRegionalWalletManagementArray) ToAutonomousDatabaseRegionalWalletManagementArrayOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutonomousDatabaseRegionalWalletManagementArrayOutput)
}

// AutonomousDatabaseRegionalWalletManagementMapInput is an input type that accepts AutonomousDatabaseRegionalWalletManagementMap and AutonomousDatabaseRegionalWalletManagementMapOutput values.
// You can construct a concrete instance of `AutonomousDatabaseRegionalWalletManagementMapInput` via:
//
//          AutonomousDatabaseRegionalWalletManagementMap{ "key": AutonomousDatabaseRegionalWalletManagementArgs{...} }
type AutonomousDatabaseRegionalWalletManagementMapInput interface {
	pulumi.Input

	ToAutonomousDatabaseRegionalWalletManagementMapOutput() AutonomousDatabaseRegionalWalletManagementMapOutput
	ToAutonomousDatabaseRegionalWalletManagementMapOutputWithContext(context.Context) AutonomousDatabaseRegionalWalletManagementMapOutput
}

type AutonomousDatabaseRegionalWalletManagementMap map[string]AutonomousDatabaseRegionalWalletManagementInput

func (AutonomousDatabaseRegionalWalletManagementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutonomousDatabaseRegionalWalletManagement)(nil)).Elem()
}

func (i AutonomousDatabaseRegionalWalletManagementMap) ToAutonomousDatabaseRegionalWalletManagementMapOutput() AutonomousDatabaseRegionalWalletManagementMapOutput {
	return i.ToAutonomousDatabaseRegionalWalletManagementMapOutputWithContext(context.Background())
}

func (i AutonomousDatabaseRegionalWalletManagementMap) ToAutonomousDatabaseRegionalWalletManagementMapOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutonomousDatabaseRegionalWalletManagementMapOutput)
}

type AutonomousDatabaseRegionalWalletManagementOutput struct {
	*pulumi.OutputState
}

func (AutonomousDatabaseRegionalWalletManagementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutonomousDatabaseRegionalWalletManagement)(nil))
}

func (o AutonomousDatabaseRegionalWalletManagementOutput) ToAutonomousDatabaseRegionalWalletManagementOutput() AutonomousDatabaseRegionalWalletManagementOutput {
	return o
}

func (o AutonomousDatabaseRegionalWalletManagementOutput) ToAutonomousDatabaseRegionalWalletManagementOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementOutput {
	return o
}

func (o AutonomousDatabaseRegionalWalletManagementOutput) ToAutonomousDatabaseRegionalWalletManagementPtrOutput() AutonomousDatabaseRegionalWalletManagementPtrOutput {
	return o.ToAutonomousDatabaseRegionalWalletManagementPtrOutputWithContext(context.Background())
}

func (o AutonomousDatabaseRegionalWalletManagementOutput) ToAutonomousDatabaseRegionalWalletManagementPtrOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementPtrOutput {
	return o.ApplyT(func(v AutonomousDatabaseRegionalWalletManagement) *AutonomousDatabaseRegionalWalletManagement {
		return &v
	}).(AutonomousDatabaseRegionalWalletManagementPtrOutput)
}

type AutonomousDatabaseRegionalWalletManagementPtrOutput struct {
	*pulumi.OutputState
}

func (AutonomousDatabaseRegionalWalletManagementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutonomousDatabaseRegionalWalletManagement)(nil))
}

func (o AutonomousDatabaseRegionalWalletManagementPtrOutput) ToAutonomousDatabaseRegionalWalletManagementPtrOutput() AutonomousDatabaseRegionalWalletManagementPtrOutput {
	return o
}

func (o AutonomousDatabaseRegionalWalletManagementPtrOutput) ToAutonomousDatabaseRegionalWalletManagementPtrOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementPtrOutput {
	return o
}

type AutonomousDatabaseRegionalWalletManagementArrayOutput struct{ *pulumi.OutputState }

func (AutonomousDatabaseRegionalWalletManagementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutonomousDatabaseRegionalWalletManagement)(nil))
}

func (o AutonomousDatabaseRegionalWalletManagementArrayOutput) ToAutonomousDatabaseRegionalWalletManagementArrayOutput() AutonomousDatabaseRegionalWalletManagementArrayOutput {
	return o
}

func (o AutonomousDatabaseRegionalWalletManagementArrayOutput) ToAutonomousDatabaseRegionalWalletManagementArrayOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementArrayOutput {
	return o
}

func (o AutonomousDatabaseRegionalWalletManagementArrayOutput) Index(i pulumi.IntInput) AutonomousDatabaseRegionalWalletManagementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutonomousDatabaseRegionalWalletManagement {
		return vs[0].([]AutonomousDatabaseRegionalWalletManagement)[vs[1].(int)]
	}).(AutonomousDatabaseRegionalWalletManagementOutput)
}

type AutonomousDatabaseRegionalWalletManagementMapOutput struct{ *pulumi.OutputState }

func (AutonomousDatabaseRegionalWalletManagementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AutonomousDatabaseRegionalWalletManagement)(nil))
}

func (o AutonomousDatabaseRegionalWalletManagementMapOutput) ToAutonomousDatabaseRegionalWalletManagementMapOutput() AutonomousDatabaseRegionalWalletManagementMapOutput {
	return o
}

func (o AutonomousDatabaseRegionalWalletManagementMapOutput) ToAutonomousDatabaseRegionalWalletManagementMapOutputWithContext(ctx context.Context) AutonomousDatabaseRegionalWalletManagementMapOutput {
	return o
}

func (o AutonomousDatabaseRegionalWalletManagementMapOutput) MapIndex(k pulumi.StringInput) AutonomousDatabaseRegionalWalletManagementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AutonomousDatabaseRegionalWalletManagement {
		return vs[0].(map[string]AutonomousDatabaseRegionalWalletManagement)[vs[1].(string)]
	}).(AutonomousDatabaseRegionalWalletManagementOutput)
}

func init() {
	pulumi.RegisterOutputType(AutonomousDatabaseRegionalWalletManagementOutput{})
	pulumi.RegisterOutputType(AutonomousDatabaseRegionalWalletManagementPtrOutput{})
	pulumi.RegisterOutputType(AutonomousDatabaseRegionalWalletManagementArrayOutput{})
	pulumi.RegisterOutputType(AutonomousDatabaseRegionalWalletManagementMapOutput{})
}