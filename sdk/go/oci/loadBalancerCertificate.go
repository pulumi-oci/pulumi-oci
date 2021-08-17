// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Certificates can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:index/loadBalancerCertificate:LoadBalancerCertificate test_certificate "loadBalancers/{loadBalancerId}/certificates/{certificateName}"
// ```
type LoadBalancerCertificate struct {
	pulumi.CustomResourceState

	// The Certificate Authority certificate, or any interim certificate, that you received from your SSL certificate provider.
	CaCertificate pulumi.StringOutput `pulumi:"caCertificate"`
	// A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `exampleCertificateBundle`
	CertificateName pulumi.StringOutput `pulumi:"certificateName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer on which to add the certificate bundle.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// A passphrase for encrypted private keys. This is needed only if you created your certificate with a passphrase.
	Passphrase pulumi.StringPtrOutput `pulumi:"passphrase"`
	// The SSL private key for your certificate, in PEM format.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The public certificate, in PEM format, that you received from your SSL certificate provider.
	PublicCertificate pulumi.StringOutput `pulumi:"publicCertificate"`
	State             pulumi.StringOutput `pulumi:"state"`
}

// NewLoadBalancerCertificate registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerCertificate(ctx *pulumi.Context,
	name string, args *LoadBalancerCertificateArgs, opts ...pulumi.ResourceOption) (*LoadBalancerCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateName == nil {
		return nil, errors.New("invalid value for required argument 'CertificateName'")
	}
	if args.LoadBalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerId'")
	}
	var resource LoadBalancerCertificate
	err := ctx.RegisterResource("oci:index/loadBalancerCertificate:LoadBalancerCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancerCertificate gets an existing LoadBalancerCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerCertificateState, opts ...pulumi.ResourceOption) (*LoadBalancerCertificate, error) {
	var resource LoadBalancerCertificate
	err := ctx.ReadResource("oci:index/loadBalancerCertificate:LoadBalancerCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancerCertificate resources.
type loadBalancerCertificateState struct {
	// The Certificate Authority certificate, or any interim certificate, that you received from your SSL certificate provider.
	CaCertificate *string `pulumi:"caCertificate"`
	// A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `exampleCertificateBundle`
	CertificateName *string `pulumi:"certificateName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer on which to add the certificate bundle.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// A passphrase for encrypted private keys. This is needed only if you created your certificate with a passphrase.
	Passphrase *string `pulumi:"passphrase"`
	// The SSL private key for your certificate, in PEM format.
	PrivateKey *string `pulumi:"privateKey"`
	// The public certificate, in PEM format, that you received from your SSL certificate provider.
	PublicCertificate *string `pulumi:"publicCertificate"`
	State             *string `pulumi:"state"`
}

type LoadBalancerCertificateState struct {
	// The Certificate Authority certificate, or any interim certificate, that you received from your SSL certificate provider.
	CaCertificate pulumi.StringPtrInput
	// A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `exampleCertificateBundle`
	CertificateName pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer on which to add the certificate bundle.
	LoadBalancerId pulumi.StringPtrInput
	// A passphrase for encrypted private keys. This is needed only if you created your certificate with a passphrase.
	Passphrase pulumi.StringPtrInput
	// The SSL private key for your certificate, in PEM format.
	PrivateKey pulumi.StringPtrInput
	// The public certificate, in PEM format, that you received from your SSL certificate provider.
	PublicCertificate pulumi.StringPtrInput
	State             pulumi.StringPtrInput
}

func (LoadBalancerCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerCertificateState)(nil)).Elem()
}

type loadBalancerCertificateArgs struct {
	// The Certificate Authority certificate, or any interim certificate, that you received from your SSL certificate provider.
	CaCertificate *string `pulumi:"caCertificate"`
	// A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `exampleCertificateBundle`
	CertificateName string `pulumi:"certificateName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer on which to add the certificate bundle.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// A passphrase for encrypted private keys. This is needed only if you created your certificate with a passphrase.
	Passphrase *string `pulumi:"passphrase"`
	// The SSL private key for your certificate, in PEM format.
	PrivateKey *string `pulumi:"privateKey"`
	// The public certificate, in PEM format, that you received from your SSL certificate provider.
	PublicCertificate *string `pulumi:"publicCertificate"`
}

// The set of arguments for constructing a LoadBalancerCertificate resource.
type LoadBalancerCertificateArgs struct {
	// The Certificate Authority certificate, or any interim certificate, that you received from your SSL certificate provider.
	CaCertificate pulumi.StringPtrInput
	// A friendly name for the certificate bundle. It must be unique and it cannot be changed. Valid certificate bundle names include only alphanumeric characters, dashes, and underscores. Certificate bundle names cannot contain spaces. Avoid entering confidential information.  Example: `exampleCertificateBundle`
	CertificateName pulumi.StringInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer on which to add the certificate bundle.
	LoadBalancerId pulumi.StringInput
	// A passphrase for encrypted private keys. This is needed only if you created your certificate with a passphrase.
	Passphrase pulumi.StringPtrInput
	// The SSL private key for your certificate, in PEM format.
	PrivateKey pulumi.StringPtrInput
	// The public certificate, in PEM format, that you received from your SSL certificate provider.
	PublicCertificate pulumi.StringPtrInput
}

func (LoadBalancerCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerCertificateArgs)(nil)).Elem()
}

type LoadBalancerCertificateInput interface {
	pulumi.Input

	ToLoadBalancerCertificateOutput() LoadBalancerCertificateOutput
	ToLoadBalancerCertificateOutputWithContext(ctx context.Context) LoadBalancerCertificateOutput
}

func (*LoadBalancerCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerCertificate)(nil))
}

func (i *LoadBalancerCertificate) ToLoadBalancerCertificateOutput() LoadBalancerCertificateOutput {
	return i.ToLoadBalancerCertificateOutputWithContext(context.Background())
}

func (i *LoadBalancerCertificate) ToLoadBalancerCertificateOutputWithContext(ctx context.Context) LoadBalancerCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerCertificateOutput)
}

func (i *LoadBalancerCertificate) ToLoadBalancerCertificatePtrOutput() LoadBalancerCertificatePtrOutput {
	return i.ToLoadBalancerCertificatePtrOutputWithContext(context.Background())
}

func (i *LoadBalancerCertificate) ToLoadBalancerCertificatePtrOutputWithContext(ctx context.Context) LoadBalancerCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerCertificatePtrOutput)
}

type LoadBalancerCertificatePtrInput interface {
	pulumi.Input

	ToLoadBalancerCertificatePtrOutput() LoadBalancerCertificatePtrOutput
	ToLoadBalancerCertificatePtrOutputWithContext(ctx context.Context) LoadBalancerCertificatePtrOutput
}

type loadBalancerCertificatePtrType LoadBalancerCertificateArgs

func (*loadBalancerCertificatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerCertificate)(nil))
}

func (i *loadBalancerCertificatePtrType) ToLoadBalancerCertificatePtrOutput() LoadBalancerCertificatePtrOutput {
	return i.ToLoadBalancerCertificatePtrOutputWithContext(context.Background())
}

func (i *loadBalancerCertificatePtrType) ToLoadBalancerCertificatePtrOutputWithContext(ctx context.Context) LoadBalancerCertificatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerCertificatePtrOutput)
}

// LoadBalancerCertificateArrayInput is an input type that accepts LoadBalancerCertificateArray and LoadBalancerCertificateArrayOutput values.
// You can construct a concrete instance of `LoadBalancerCertificateArrayInput` via:
//
//          LoadBalancerCertificateArray{ LoadBalancerCertificateArgs{...} }
type LoadBalancerCertificateArrayInput interface {
	pulumi.Input

	ToLoadBalancerCertificateArrayOutput() LoadBalancerCertificateArrayOutput
	ToLoadBalancerCertificateArrayOutputWithContext(context.Context) LoadBalancerCertificateArrayOutput
}

type LoadBalancerCertificateArray []LoadBalancerCertificateInput

func (LoadBalancerCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancerCertificate)(nil)).Elem()
}

func (i LoadBalancerCertificateArray) ToLoadBalancerCertificateArrayOutput() LoadBalancerCertificateArrayOutput {
	return i.ToLoadBalancerCertificateArrayOutputWithContext(context.Background())
}

func (i LoadBalancerCertificateArray) ToLoadBalancerCertificateArrayOutputWithContext(ctx context.Context) LoadBalancerCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerCertificateArrayOutput)
}

// LoadBalancerCertificateMapInput is an input type that accepts LoadBalancerCertificateMap and LoadBalancerCertificateMapOutput values.
// You can construct a concrete instance of `LoadBalancerCertificateMapInput` via:
//
//          LoadBalancerCertificateMap{ "key": LoadBalancerCertificateArgs{...} }
type LoadBalancerCertificateMapInput interface {
	pulumi.Input

	ToLoadBalancerCertificateMapOutput() LoadBalancerCertificateMapOutput
	ToLoadBalancerCertificateMapOutputWithContext(context.Context) LoadBalancerCertificateMapOutput
}

type LoadBalancerCertificateMap map[string]LoadBalancerCertificateInput

func (LoadBalancerCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancerCertificate)(nil)).Elem()
}

func (i LoadBalancerCertificateMap) ToLoadBalancerCertificateMapOutput() LoadBalancerCertificateMapOutput {
	return i.ToLoadBalancerCertificateMapOutputWithContext(context.Background())
}

func (i LoadBalancerCertificateMap) ToLoadBalancerCertificateMapOutputWithContext(ctx context.Context) LoadBalancerCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerCertificateMapOutput)
}

type LoadBalancerCertificateOutput struct {
	*pulumi.OutputState
}

func (LoadBalancerCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerCertificate)(nil))
}

func (o LoadBalancerCertificateOutput) ToLoadBalancerCertificateOutput() LoadBalancerCertificateOutput {
	return o
}

func (o LoadBalancerCertificateOutput) ToLoadBalancerCertificateOutputWithContext(ctx context.Context) LoadBalancerCertificateOutput {
	return o
}

func (o LoadBalancerCertificateOutput) ToLoadBalancerCertificatePtrOutput() LoadBalancerCertificatePtrOutput {
	return o.ToLoadBalancerCertificatePtrOutputWithContext(context.Background())
}

func (o LoadBalancerCertificateOutput) ToLoadBalancerCertificatePtrOutputWithContext(ctx context.Context) LoadBalancerCertificatePtrOutput {
	return o.ApplyT(func(v LoadBalancerCertificate) *LoadBalancerCertificate {
		return &v
	}).(LoadBalancerCertificatePtrOutput)
}

type LoadBalancerCertificatePtrOutput struct {
	*pulumi.OutputState
}

func (LoadBalancerCertificatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerCertificate)(nil))
}

func (o LoadBalancerCertificatePtrOutput) ToLoadBalancerCertificatePtrOutput() LoadBalancerCertificatePtrOutput {
	return o
}

func (o LoadBalancerCertificatePtrOutput) ToLoadBalancerCertificatePtrOutputWithContext(ctx context.Context) LoadBalancerCertificatePtrOutput {
	return o
}

type LoadBalancerCertificateArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerCertificate)(nil))
}

func (o LoadBalancerCertificateArrayOutput) ToLoadBalancerCertificateArrayOutput() LoadBalancerCertificateArrayOutput {
	return o
}

func (o LoadBalancerCertificateArrayOutput) ToLoadBalancerCertificateArrayOutputWithContext(ctx context.Context) LoadBalancerCertificateArrayOutput {
	return o
}

func (o LoadBalancerCertificateArrayOutput) Index(i pulumi.IntInput) LoadBalancerCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerCertificate {
		return vs[0].([]LoadBalancerCertificate)[vs[1].(int)]
	}).(LoadBalancerCertificateOutput)
}

type LoadBalancerCertificateMapOutput struct{ *pulumi.OutputState }

func (LoadBalancerCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LoadBalancerCertificate)(nil))
}

func (o LoadBalancerCertificateMapOutput) ToLoadBalancerCertificateMapOutput() LoadBalancerCertificateMapOutput {
	return o
}

func (o LoadBalancerCertificateMapOutput) ToLoadBalancerCertificateMapOutputWithContext(ctx context.Context) LoadBalancerCertificateMapOutput {
	return o
}

func (o LoadBalancerCertificateMapOutput) MapIndex(k pulumi.StringInput) LoadBalancerCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LoadBalancerCertificate {
		return vs[0].(map[string]LoadBalancerCertificate)[vs[1].(string)]
	}).(LoadBalancerCertificateOutput)
}

func init() {
	pulumi.RegisterOutputType(LoadBalancerCertificateOutput{})
	pulumi.RegisterOutputType(LoadBalancerCertificatePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerCertificateArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerCertificateMapOutput{})
}