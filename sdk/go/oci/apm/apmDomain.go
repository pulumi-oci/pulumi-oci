// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Apm Domain resource in Oracle Cloud Infrastructure Apm service.
//
// Creates a new APM Domain.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/apm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apm.NewApmDomain(ctx, "testApmDomain", &apm.ApmDomainArgs{
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			DisplayName:   pulumi.Any(_var.Apm_domain_display_name),
// 			DefinedTags: pulumi.AnyMap{
// 				"foo-namespace.bar-key": pulumi.Any("value"),
// 			},
// 			Description: pulumi.Any(_var.Apm_domain_description),
// 			FreeformTags: pulumi.AnyMap{
// 				"bar-key": pulumi.Any("value"),
// 			},
// 			IsFreeTier: pulumi.Any(_var.Apm_domain_is_free_tier),
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
// ApmDomains can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:apm/apmDomain:ApmDomain test_apm_domain "id"
// ```
type ApmDomain struct {
	pulumi.CustomResourceState

	// (Updatable) The OCID of the compartment corresponding to the APM Domain.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// Where APM Agents upload their observations and metrics.
	DataUploadEndpoint pulumi.StringOutput `pulumi:"dataUploadEndpoint"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) Description of the APM Domain
	Description pulumi.StringOutput `pulumi:"description"`
	// (Updatable) Display name of the APM Domain
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// Indicates whether this is an "Always Free" resource. The default value is false.
	IsFreeTier pulumi.BoolOutput `pulumi:"isFreeTier"`
	// The current lifecycle state of the APM Domain.
	State pulumi.StringOutput `pulumi:"state"`
	// The time the the APM Domain was created. An RFC3339 formatted datetime string
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The time the APM Domain was updated. An RFC3339 formatted datetime string
	TimeUpdated pulumi.StringOutput `pulumi:"timeUpdated"`
}

// NewApmDomain registers a new resource with the given unique name, arguments, and options.
func NewApmDomain(ctx *pulumi.Context,
	name string, args *ApmDomainArgs, opts ...pulumi.ResourceOption) (*ApmDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource ApmDomain
	err := ctx.RegisterResource("oci:apm/apmDomain:ApmDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApmDomain gets an existing ApmDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApmDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApmDomainState, opts ...pulumi.ResourceOption) (*ApmDomain, error) {
	var resource ApmDomain
	err := ctx.ReadResource("oci:apm/apmDomain:ApmDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApmDomain resources.
type apmDomainState struct {
	// (Updatable) The OCID of the compartment corresponding to the APM Domain.
	CompartmentId *string `pulumi:"compartmentId"`
	// Where APM Agents upload their observations and metrics.
	DataUploadEndpoint *string `pulumi:"dataUploadEndpoint"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) Description of the APM Domain
	Description *string `pulumi:"description"`
	// (Updatable) Display name of the APM Domain
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Indicates whether this is an "Always Free" resource. The default value is false.
	IsFreeTier *bool `pulumi:"isFreeTier"`
	// The current lifecycle state of the APM Domain.
	State *string `pulumi:"state"`
	// The time the the APM Domain was created. An RFC3339 formatted datetime string
	TimeCreated *string `pulumi:"timeCreated"`
	// The time the APM Domain was updated. An RFC3339 formatted datetime string
	TimeUpdated *string `pulumi:"timeUpdated"`
}

type ApmDomainState struct {
	// (Updatable) The OCID of the compartment corresponding to the APM Domain.
	CompartmentId pulumi.StringPtrInput
	// Where APM Agents upload their observations and metrics.
	DataUploadEndpoint pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// (Updatable) Description of the APM Domain
	Description pulumi.StringPtrInput
	// (Updatable) Display name of the APM Domain
	DisplayName pulumi.StringPtrInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// Indicates whether this is an "Always Free" resource. The default value is false.
	IsFreeTier pulumi.BoolPtrInput
	// The current lifecycle state of the APM Domain.
	State pulumi.StringPtrInput
	// The time the the APM Domain was created. An RFC3339 formatted datetime string
	TimeCreated pulumi.StringPtrInput
	// The time the APM Domain was updated. An RFC3339 formatted datetime string
	TimeUpdated pulumi.StringPtrInput
}

func (ApmDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*apmDomainState)(nil)).Elem()
}

type apmDomainArgs struct {
	// (Updatable) The OCID of the compartment corresponding to the APM Domain.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) Description of the APM Domain
	Description *string `pulumi:"description"`
	// (Updatable) Display name of the APM Domain
	DisplayName string `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Indicates whether this is an "Always Free" resource. The default value is false.
	IsFreeTier *bool `pulumi:"isFreeTier"`
}

// The set of arguments for constructing a ApmDomain resource.
type ApmDomainArgs struct {
	// (Updatable) The OCID of the compartment corresponding to the APM Domain.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// (Updatable) Description of the APM Domain
	Description pulumi.StringPtrInput
	// (Updatable) Display name of the APM Domain
	DisplayName pulumi.StringInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// Indicates whether this is an "Always Free" resource. The default value is false.
	IsFreeTier pulumi.BoolPtrInput
}

func (ApmDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apmDomainArgs)(nil)).Elem()
}

type ApmDomainInput interface {
	pulumi.Input

	ToApmDomainOutput() ApmDomainOutput
	ToApmDomainOutputWithContext(ctx context.Context) ApmDomainOutput
}

func (*ApmDomain) ElementType() reflect.Type {
	return reflect.TypeOf((*ApmDomain)(nil))
}

func (i *ApmDomain) ToApmDomainOutput() ApmDomainOutput {
	return i.ToApmDomainOutputWithContext(context.Background())
}

func (i *ApmDomain) ToApmDomainOutputWithContext(ctx context.Context) ApmDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApmDomainOutput)
}

func (i *ApmDomain) ToApmDomainPtrOutput() ApmDomainPtrOutput {
	return i.ToApmDomainPtrOutputWithContext(context.Background())
}

func (i *ApmDomain) ToApmDomainPtrOutputWithContext(ctx context.Context) ApmDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApmDomainPtrOutput)
}

type ApmDomainPtrInput interface {
	pulumi.Input

	ToApmDomainPtrOutput() ApmDomainPtrOutput
	ToApmDomainPtrOutputWithContext(ctx context.Context) ApmDomainPtrOutput
}

type apmDomainPtrType ApmDomainArgs

func (*apmDomainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApmDomain)(nil))
}

func (i *apmDomainPtrType) ToApmDomainPtrOutput() ApmDomainPtrOutput {
	return i.ToApmDomainPtrOutputWithContext(context.Background())
}

func (i *apmDomainPtrType) ToApmDomainPtrOutputWithContext(ctx context.Context) ApmDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApmDomainPtrOutput)
}

// ApmDomainArrayInput is an input type that accepts ApmDomainArray and ApmDomainArrayOutput values.
// You can construct a concrete instance of `ApmDomainArrayInput` via:
//
//          ApmDomainArray{ ApmDomainArgs{...} }
type ApmDomainArrayInput interface {
	pulumi.Input

	ToApmDomainArrayOutput() ApmDomainArrayOutput
	ToApmDomainArrayOutputWithContext(context.Context) ApmDomainArrayOutput
}

type ApmDomainArray []ApmDomainInput

func (ApmDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApmDomain)(nil)).Elem()
}

func (i ApmDomainArray) ToApmDomainArrayOutput() ApmDomainArrayOutput {
	return i.ToApmDomainArrayOutputWithContext(context.Background())
}

func (i ApmDomainArray) ToApmDomainArrayOutputWithContext(ctx context.Context) ApmDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApmDomainArrayOutput)
}

// ApmDomainMapInput is an input type that accepts ApmDomainMap and ApmDomainMapOutput values.
// You can construct a concrete instance of `ApmDomainMapInput` via:
//
//          ApmDomainMap{ "key": ApmDomainArgs{...} }
type ApmDomainMapInput interface {
	pulumi.Input

	ToApmDomainMapOutput() ApmDomainMapOutput
	ToApmDomainMapOutputWithContext(context.Context) ApmDomainMapOutput
}

type ApmDomainMap map[string]ApmDomainInput

func (ApmDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApmDomain)(nil)).Elem()
}

func (i ApmDomainMap) ToApmDomainMapOutput() ApmDomainMapOutput {
	return i.ToApmDomainMapOutputWithContext(context.Background())
}

func (i ApmDomainMap) ToApmDomainMapOutputWithContext(ctx context.Context) ApmDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApmDomainMapOutput)
}

type ApmDomainOutput struct {
	*pulumi.OutputState
}

func (ApmDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApmDomain)(nil))
}

func (o ApmDomainOutput) ToApmDomainOutput() ApmDomainOutput {
	return o
}

func (o ApmDomainOutput) ToApmDomainOutputWithContext(ctx context.Context) ApmDomainOutput {
	return o
}

func (o ApmDomainOutput) ToApmDomainPtrOutput() ApmDomainPtrOutput {
	return o.ToApmDomainPtrOutputWithContext(context.Background())
}

func (o ApmDomainOutput) ToApmDomainPtrOutputWithContext(ctx context.Context) ApmDomainPtrOutput {
	return o.ApplyT(func(v ApmDomain) *ApmDomain {
		return &v
	}).(ApmDomainPtrOutput)
}

type ApmDomainPtrOutput struct {
	*pulumi.OutputState
}

func (ApmDomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApmDomain)(nil))
}

func (o ApmDomainPtrOutput) ToApmDomainPtrOutput() ApmDomainPtrOutput {
	return o
}

func (o ApmDomainPtrOutput) ToApmDomainPtrOutputWithContext(ctx context.Context) ApmDomainPtrOutput {
	return o
}

type ApmDomainArrayOutput struct{ *pulumi.OutputState }

func (ApmDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApmDomain)(nil))
}

func (o ApmDomainArrayOutput) ToApmDomainArrayOutput() ApmDomainArrayOutput {
	return o
}

func (o ApmDomainArrayOutput) ToApmDomainArrayOutputWithContext(ctx context.Context) ApmDomainArrayOutput {
	return o
}

func (o ApmDomainArrayOutput) Index(i pulumi.IntInput) ApmDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApmDomain {
		return vs[0].([]ApmDomain)[vs[1].(int)]
	}).(ApmDomainOutput)
}

type ApmDomainMapOutput struct{ *pulumi.OutputState }

func (ApmDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApmDomain)(nil))
}

func (o ApmDomainMapOutput) ToApmDomainMapOutput() ApmDomainMapOutput {
	return o
}

func (o ApmDomainMapOutput) ToApmDomainMapOutputWithContext(ctx context.Context) ApmDomainMapOutput {
	return o
}

func (o ApmDomainMapOutput) MapIndex(k pulumi.StringInput) ApmDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApmDomain {
		return vs[0].(map[string]ApmDomain)[vs[1].(string)]
	}).(ApmDomainOutput)
}

func init() {
	pulumi.RegisterOutputType(ApmDomainOutput{})
	pulumi.RegisterOutputType(ApmDomainPtrOutput{})
	pulumi.RegisterOutputType(ApmDomainArrayOutput{})
	pulumi.RegisterOutputType(ApmDomainMapOutput{})
}
