// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Ssl Cipher Suite resource in Oracle Cloud Infrastructure Load Balancer service.
//
// Creates a custom SSL cipher suite.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/loadbalancer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := loadbalancer.NewSslCipherSuite(ctx, "testSslCipherSuite", &loadbalancer.SslCipherSuiteArgs{
// 			Ciphers:        pulumi.Any(_var.Ssl_cipher_suite_ciphers),
// 			LoadBalancerId: pulumi.Any(oci_load_balancer_load_balancer.Test_load_balancer.Id),
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
// SslCipherSuites can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:loadbalancer/sslCipherSuite:SslCipherSuite test_ssl_cipher_suite "loadBalancers/{loadBalancerId}/sslCipherSuites/{name}"
// ```
type SslCipherSuite struct {
	pulumi.CustomResourceState

	// A list of SSL ciphers the load balancer must support for HTTPS or SSL connections.
	Ciphers pulumi.StringArrayOutput `pulumi:"ciphers"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated load balancer.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.
	Name  pulumi.StringOutput `pulumi:"name"`
	State pulumi.StringOutput `pulumi:"state"`
}

// NewSslCipherSuite registers a new resource with the given unique name, arguments, and options.
func NewSslCipherSuite(ctx *pulumi.Context,
	name string, args *SslCipherSuiteArgs, opts ...pulumi.ResourceOption) (*SslCipherSuite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ciphers == nil {
		return nil, errors.New("invalid value for required argument 'Ciphers'")
	}
	var resource SslCipherSuite
	err := ctx.RegisterResource("oci:loadbalancer/sslCipherSuite:SslCipherSuite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslCipherSuite gets an existing SslCipherSuite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslCipherSuite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslCipherSuiteState, opts ...pulumi.ResourceOption) (*SslCipherSuite, error) {
	var resource SslCipherSuite
	err := ctx.ReadResource("oci:loadbalancer/sslCipherSuite:SslCipherSuite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslCipherSuite resources.
type sslCipherSuiteState struct {
	// A list of SSL ciphers the load balancer must support for HTTPS or SSL connections.
	Ciphers []string `pulumi:"ciphers"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated load balancer.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.
	Name  *string `pulumi:"name"`
	State *string `pulumi:"state"`
}

type SslCipherSuiteState struct {
	// A list of SSL ciphers the load balancer must support for HTTPS or SSL connections.
	Ciphers pulumi.StringArrayInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated load balancer.
	LoadBalancerId pulumi.StringPtrInput
	// A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.
	Name  pulumi.StringPtrInput
	State pulumi.StringPtrInput
}

func (SslCipherSuiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslCipherSuiteState)(nil)).Elem()
}

type sslCipherSuiteArgs struct {
	// A list of SSL ciphers the load balancer must support for HTTPS or SSL connections.
	Ciphers []string `pulumi:"ciphers"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated load balancer.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a SslCipherSuite resource.
type SslCipherSuiteArgs struct {
	// A list of SSL ciphers the load balancer must support for HTTPS or SSL connections.
	Ciphers pulumi.StringArrayInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated load balancer.
	LoadBalancerId pulumi.StringPtrInput
	// A friendly name for the SSL cipher suite. It must be unique and it cannot be changed.
	Name pulumi.StringPtrInput
}

func (SslCipherSuiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslCipherSuiteArgs)(nil)).Elem()
}

type SslCipherSuiteInput interface {
	pulumi.Input

	ToSslCipherSuiteOutput() SslCipherSuiteOutput
	ToSslCipherSuiteOutputWithContext(ctx context.Context) SslCipherSuiteOutput
}

func (*SslCipherSuite) ElementType() reflect.Type {
	return reflect.TypeOf((*SslCipherSuite)(nil))
}

func (i *SslCipherSuite) ToSslCipherSuiteOutput() SslCipherSuiteOutput {
	return i.ToSslCipherSuiteOutputWithContext(context.Background())
}

func (i *SslCipherSuite) ToSslCipherSuiteOutputWithContext(ctx context.Context) SslCipherSuiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslCipherSuiteOutput)
}

func (i *SslCipherSuite) ToSslCipherSuitePtrOutput() SslCipherSuitePtrOutput {
	return i.ToSslCipherSuitePtrOutputWithContext(context.Background())
}

func (i *SslCipherSuite) ToSslCipherSuitePtrOutputWithContext(ctx context.Context) SslCipherSuitePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslCipherSuitePtrOutput)
}

type SslCipherSuitePtrInput interface {
	pulumi.Input

	ToSslCipherSuitePtrOutput() SslCipherSuitePtrOutput
	ToSslCipherSuitePtrOutputWithContext(ctx context.Context) SslCipherSuitePtrOutput
}

type sslCipherSuitePtrType SslCipherSuiteArgs

func (*sslCipherSuitePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SslCipherSuite)(nil))
}

func (i *sslCipherSuitePtrType) ToSslCipherSuitePtrOutput() SslCipherSuitePtrOutput {
	return i.ToSslCipherSuitePtrOutputWithContext(context.Background())
}

func (i *sslCipherSuitePtrType) ToSslCipherSuitePtrOutputWithContext(ctx context.Context) SslCipherSuitePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslCipherSuitePtrOutput)
}

// SslCipherSuiteArrayInput is an input type that accepts SslCipherSuiteArray and SslCipherSuiteArrayOutput values.
// You can construct a concrete instance of `SslCipherSuiteArrayInput` via:
//
//          SslCipherSuiteArray{ SslCipherSuiteArgs{...} }
type SslCipherSuiteArrayInput interface {
	pulumi.Input

	ToSslCipherSuiteArrayOutput() SslCipherSuiteArrayOutput
	ToSslCipherSuiteArrayOutputWithContext(context.Context) SslCipherSuiteArrayOutput
}

type SslCipherSuiteArray []SslCipherSuiteInput

func (SslCipherSuiteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SslCipherSuite)(nil)).Elem()
}

func (i SslCipherSuiteArray) ToSslCipherSuiteArrayOutput() SslCipherSuiteArrayOutput {
	return i.ToSslCipherSuiteArrayOutputWithContext(context.Background())
}

func (i SslCipherSuiteArray) ToSslCipherSuiteArrayOutputWithContext(ctx context.Context) SslCipherSuiteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslCipherSuiteArrayOutput)
}

// SslCipherSuiteMapInput is an input type that accepts SslCipherSuiteMap and SslCipherSuiteMapOutput values.
// You can construct a concrete instance of `SslCipherSuiteMapInput` via:
//
//          SslCipherSuiteMap{ "key": SslCipherSuiteArgs{...} }
type SslCipherSuiteMapInput interface {
	pulumi.Input

	ToSslCipherSuiteMapOutput() SslCipherSuiteMapOutput
	ToSslCipherSuiteMapOutputWithContext(context.Context) SslCipherSuiteMapOutput
}

type SslCipherSuiteMap map[string]SslCipherSuiteInput

func (SslCipherSuiteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SslCipherSuite)(nil)).Elem()
}

func (i SslCipherSuiteMap) ToSslCipherSuiteMapOutput() SslCipherSuiteMapOutput {
	return i.ToSslCipherSuiteMapOutputWithContext(context.Background())
}

func (i SslCipherSuiteMap) ToSslCipherSuiteMapOutputWithContext(ctx context.Context) SslCipherSuiteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslCipherSuiteMapOutput)
}

type SslCipherSuiteOutput struct {
	*pulumi.OutputState
}

func (SslCipherSuiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslCipherSuite)(nil))
}

func (o SslCipherSuiteOutput) ToSslCipherSuiteOutput() SslCipherSuiteOutput {
	return o
}

func (o SslCipherSuiteOutput) ToSslCipherSuiteOutputWithContext(ctx context.Context) SslCipherSuiteOutput {
	return o
}

func (o SslCipherSuiteOutput) ToSslCipherSuitePtrOutput() SslCipherSuitePtrOutput {
	return o.ToSslCipherSuitePtrOutputWithContext(context.Background())
}

func (o SslCipherSuiteOutput) ToSslCipherSuitePtrOutputWithContext(ctx context.Context) SslCipherSuitePtrOutput {
	return o.ApplyT(func(v SslCipherSuite) *SslCipherSuite {
		return &v
	}).(SslCipherSuitePtrOutput)
}

type SslCipherSuitePtrOutput struct {
	*pulumi.OutputState
}

func (SslCipherSuitePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslCipherSuite)(nil))
}

func (o SslCipherSuitePtrOutput) ToSslCipherSuitePtrOutput() SslCipherSuitePtrOutput {
	return o
}

func (o SslCipherSuitePtrOutput) ToSslCipherSuitePtrOutputWithContext(ctx context.Context) SslCipherSuitePtrOutput {
	return o
}

type SslCipherSuiteArrayOutput struct{ *pulumi.OutputState }

func (SslCipherSuiteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SslCipherSuite)(nil))
}

func (o SslCipherSuiteArrayOutput) ToSslCipherSuiteArrayOutput() SslCipherSuiteArrayOutput {
	return o
}

func (o SslCipherSuiteArrayOutput) ToSslCipherSuiteArrayOutputWithContext(ctx context.Context) SslCipherSuiteArrayOutput {
	return o
}

func (o SslCipherSuiteArrayOutput) Index(i pulumi.IntInput) SslCipherSuiteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SslCipherSuite {
		return vs[0].([]SslCipherSuite)[vs[1].(int)]
	}).(SslCipherSuiteOutput)
}

type SslCipherSuiteMapOutput struct{ *pulumi.OutputState }

func (SslCipherSuiteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SslCipherSuite)(nil))
}

func (o SslCipherSuiteMapOutput) ToSslCipherSuiteMapOutput() SslCipherSuiteMapOutput {
	return o
}

func (o SslCipherSuiteMapOutput) ToSslCipherSuiteMapOutputWithContext(ctx context.Context) SslCipherSuiteMapOutput {
	return o
}

func (o SslCipherSuiteMapOutput) MapIndex(k pulumi.StringInput) SslCipherSuiteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SslCipherSuite {
		return vs[0].(map[string]SslCipherSuite)[vs[1].(string)]
	}).(SslCipherSuiteOutput)
}

func init() {
	pulumi.RegisterOutputType(SslCipherSuiteOutput{})
	pulumi.RegisterOutputType(SslCipherSuitePtrOutput{})
	pulumi.RegisterOutputType(SslCipherSuiteArrayOutput{})
	pulumi.RegisterOutputType(SslCipherSuiteMapOutput{})
}