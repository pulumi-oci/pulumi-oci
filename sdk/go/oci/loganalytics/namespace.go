// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loganalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Namespace resource in Oracle Cloud Infrastructure Log Analytics service.
//
// Onboards a tenancy with Log Analytics or Offboards a tenancy from Log Analytics functionality.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/loganalytics"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := loganalytics.NewNamespace(ctx, "testNamespace", &loganalytics.NamespaceArgs{
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			IsOnboarded:   pulumi.Any(_var.Is_onboarded),
// 			Namespace:     pulumi.Any(_var.Namespace_namespace),
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
// Namespace can be imported using the `namespace`, e.g.
//
// ```sh
//  $ pulumi import oci:loganalytics/namespace:Namespace test_namespace "namespace"
// ```
type Namespace struct {
	pulumi.CustomResourceState

	// The OCID of the root compartment i.e. OCID of the tenancy.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// Use `true` if tenancy is to be onboarded to logging analytics and `false` if tenancy is to be offboarded
	IsOnboarded pulumi.BoolOutput `pulumi:"isOnboarded"`
	// The Log Analytics namespace used for the request.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.IsOnboarded == nil {
		return nil, errors.New("invalid value for required argument 'IsOnboarded'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	var resource Namespace
	err := ctx.RegisterResource("oci:loganalytics/namespace:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("oci:loganalytics/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// The OCID of the root compartment i.e. OCID of the tenancy.
	CompartmentId *string `pulumi:"compartmentId"`
	// Use `true` if tenancy is to be onboarded to logging analytics and `false` if tenancy is to be offboarded
	IsOnboarded *bool `pulumi:"isOnboarded"`
	// The Log Analytics namespace used for the request.
	Namespace *string `pulumi:"namespace"`
}

type NamespaceState struct {
	// The OCID of the root compartment i.e. OCID of the tenancy.
	CompartmentId pulumi.StringPtrInput
	// Use `true` if tenancy is to be onboarded to logging analytics and `false` if tenancy is to be offboarded
	IsOnboarded pulumi.BoolPtrInput
	// The Log Analytics namespace used for the request.
	Namespace pulumi.StringPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// The OCID of the root compartment i.e. OCID of the tenancy.
	CompartmentId string `pulumi:"compartmentId"`
	// Use `true` if tenancy is to be onboarded to logging analytics and `false` if tenancy is to be offboarded
	IsOnboarded bool `pulumi:"isOnboarded"`
	// The Log Analytics namespace used for the request.
	Namespace string `pulumi:"namespace"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// The OCID of the root compartment i.e. OCID of the tenancy.
	CompartmentId pulumi.StringInput
	// Use `true` if tenancy is to be onboarded to logging analytics and `false` if tenancy is to be offboarded
	IsOnboarded pulumi.BoolInput
	// The Log Analytics namespace used for the request.
	Namespace pulumi.StringInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil))
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

func (i *Namespace) ToNamespacePtrOutput() NamespacePtrOutput {
	return i.ToNamespacePtrOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePtrOutput)
}

type NamespacePtrInput interface {
	pulumi.Input

	ToNamespacePtrOutput() NamespacePtrOutput
	ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput
}

type namespacePtrType NamespaceArgs

func (*namespacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil))
}

func (i *namespacePtrType) ToNamespacePtrOutput() NamespacePtrOutput {
	return i.ToNamespacePtrOutputWithContext(context.Background())
}

func (i *namespacePtrType) ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespacePtrOutput)
}

// NamespaceArrayInput is an input type that accepts NamespaceArray and NamespaceArrayOutput values.
// You can construct a concrete instance of `NamespaceArrayInput` via:
//
//          NamespaceArray{ NamespaceArgs{...} }
type NamespaceArrayInput interface {
	pulumi.Input

	ToNamespaceArrayOutput() NamespaceArrayOutput
	ToNamespaceArrayOutputWithContext(context.Context) NamespaceArrayOutput
}

type NamespaceArray []NamespaceInput

func (NamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Namespace)(nil)).Elem()
}

func (i NamespaceArray) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return i.ToNamespaceArrayOutputWithContext(context.Background())
}

func (i NamespaceArray) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceArrayOutput)
}

// NamespaceMapInput is an input type that accepts NamespaceMap and NamespaceMapOutput values.
// You can construct a concrete instance of `NamespaceMapInput` via:
//
//          NamespaceMap{ "key": NamespaceArgs{...} }
type NamespaceMapInput interface {
	pulumi.Input

	ToNamespaceMapOutput() NamespaceMapOutput
	ToNamespaceMapOutputWithContext(context.Context) NamespaceMapOutput
}

type NamespaceMap map[string]NamespaceInput

func (NamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Namespace)(nil)).Elem()
}

func (i NamespaceMap) ToNamespaceMapOutput() NamespaceMapOutput {
	return i.ToNamespaceMapOutputWithContext(context.Background())
}

func (i NamespaceMap) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceMapOutput)
}

type NamespaceOutput struct {
	*pulumi.OutputState
}

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Namespace)(nil))
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespacePtrOutput() NamespacePtrOutput {
	return o.ToNamespacePtrOutputWithContext(context.Background())
}

func (o NamespaceOutput) ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput {
	return o.ApplyT(func(v Namespace) *Namespace {
		return &v
	}).(NamespacePtrOutput)
}

type NamespacePtrOutput struct {
	*pulumi.OutputState
}

func (NamespacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil))
}

func (o NamespacePtrOutput) ToNamespacePtrOutput() NamespacePtrOutput {
	return o
}

func (o NamespacePtrOutput) ToNamespacePtrOutputWithContext(ctx context.Context) NamespacePtrOutput {
	return o
}

type NamespaceArrayOutput struct{ *pulumi.OutputState }

func (NamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Namespace)(nil))
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutput() NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) ToNamespaceArrayOutputWithContext(ctx context.Context) NamespaceArrayOutput {
	return o
}

func (o NamespaceArrayOutput) Index(i pulumi.IntInput) NamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Namespace {
		return vs[0].([]Namespace)[vs[1].(int)]
	}).(NamespaceOutput)
}

type NamespaceMapOutput struct{ *pulumi.OutputState }

func (NamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Namespace)(nil))
}

func (o NamespaceMapOutput) ToNamespaceMapOutput() NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) ToNamespaceMapOutputWithContext(ctx context.Context) NamespaceMapOutput {
	return o
}

func (o NamespaceMapOutput) MapIndex(k pulumi.StringInput) NamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Namespace {
		return vs[0].(map[string]Namespace)[vs[1].(string)]
	}).(NamespaceOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
	pulumi.RegisterOutputType(NamespacePtrOutput{})
	pulumi.RegisterOutputType(NamespaceArrayOutput{})
	pulumi.RegisterOutputType(NamespaceMapOutput{})
}
