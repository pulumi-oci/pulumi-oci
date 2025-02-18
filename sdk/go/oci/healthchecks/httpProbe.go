// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthchecks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Http Probe resource in Oracle Cloud Infrastructure Health Checks service.
//
// Creates an on-demand HTTP probe. The location response header contains the URL for
// fetching the probe results.
//
// *Note:* On-demand probe configurations are not saved.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/healthchecks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := healthchecks.NewHttpProbe(ctx, "testHttpProbe", &healthchecks.HttpProbeArgs{
// 			CompartmentId:     pulumi.Any(_var.Compartment_id),
// 			Protocol:          pulumi.Any(_var.Http_probe_protocol),
// 			Targets:           pulumi.Any(_var.Http_probe_targets),
// 			Headers:           pulumi.Any(_var.Http_probe_headers),
// 			Method:            pulumi.Any(_var.Http_probe_method),
// 			Path:              pulumi.Any(_var.Http_probe_path),
// 			Port:              pulumi.Any(_var.Http_probe_port),
// 			TimeoutInSeconds:  pulumi.Any(_var.Http_probe_timeout_in_seconds),
// 			VantagePointNames: pulumi.Any(_var.Http_probe_vantage_point_names),
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
// HttpProbes can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:healthchecks/httpProbe:HttpProbe test_http_probe "id"
// ```
type HttpProbe struct {
	pulumi.CustomResourceState

	// The OCID of the compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// A dictionary of HTTP request headers.
	Headers pulumi.MapOutput `pulumi:"headers"`
	// The region where updates must be made and where results must be fetched from.
	HomeRegion pulumi.StringOutput `pulumi:"homeRegion"`
	// The supported HTTP methods available for probes.
	Method pulumi.StringOutput `pulumi:"method"`
	// The optional URL path to probe, including query parameters.
	Path pulumi.StringOutput `pulumi:"path"`
	// The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol.
	Port pulumi.IntOutput `pulumi:"port"`
	// The supported protocols available for HTTP probes.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// A URL for fetching the probe results.
	ResultsUrl pulumi.StringOutput `pulumi:"resultsUrl"`
	// A list of targets (hostnames or IP addresses) of the probe.
	Targets pulumi.StringArrayOutput `pulumi:"targets"`
	// The RFC 3339-formatted creation date and time of the probe.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds pulumi.IntOutput `pulumi:"timeoutInSeconds"`
	// A list of names of vantage points from which to execute the probe.
	VantagePointNames pulumi.StringArrayOutput `pulumi:"vantagePointNames"`
}

// NewHttpProbe registers a new resource with the given unique name, arguments, and options.
func NewHttpProbe(ctx *pulumi.Context,
	name string, args *HttpProbeArgs, opts ...pulumi.ResourceOption) (*HttpProbe, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.Targets == nil {
		return nil, errors.New("invalid value for required argument 'Targets'")
	}
	var resource HttpProbe
	err := ctx.RegisterResource("oci:healthchecks/httpProbe:HttpProbe", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHttpProbe gets an existing HttpProbe resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpProbe(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HttpProbeState, opts ...pulumi.ResourceOption) (*HttpProbe, error) {
	var resource HttpProbe
	err := ctx.ReadResource("oci:healthchecks/httpProbe:HttpProbe", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HttpProbe resources.
type httpProbeState struct {
	// The OCID of the compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// A dictionary of HTTP request headers.
	Headers map[string]interface{} `pulumi:"headers"`
	// The region where updates must be made and where results must be fetched from.
	HomeRegion *string `pulumi:"homeRegion"`
	// The supported HTTP methods available for probes.
	Method *string `pulumi:"method"`
	// The optional URL path to probe, including query parameters.
	Path *string `pulumi:"path"`
	// The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol.
	Port *int `pulumi:"port"`
	// The supported protocols available for HTTP probes.
	Protocol *string `pulumi:"protocol"`
	// A URL for fetching the probe results.
	ResultsUrl *string `pulumi:"resultsUrl"`
	// A list of targets (hostnames or IP addresses) of the probe.
	Targets []string `pulumi:"targets"`
	// The RFC 3339-formatted creation date and time of the probe.
	TimeCreated *string `pulumi:"timeCreated"`
	// The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
	// A list of names of vantage points from which to execute the probe.
	VantagePointNames []string `pulumi:"vantagePointNames"`
}

type HttpProbeState struct {
	// The OCID of the compartment.
	CompartmentId pulumi.StringPtrInput
	// A dictionary of HTTP request headers.
	Headers pulumi.MapInput
	// The region where updates must be made and where results must be fetched from.
	HomeRegion pulumi.StringPtrInput
	// The supported HTTP methods available for probes.
	Method pulumi.StringPtrInput
	// The optional URL path to probe, including query parameters.
	Path pulumi.StringPtrInput
	// The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol.
	Port pulumi.IntPtrInput
	// The supported protocols available for HTTP probes.
	Protocol pulumi.StringPtrInput
	// A URL for fetching the probe results.
	ResultsUrl pulumi.StringPtrInput
	// A list of targets (hostnames or IP addresses) of the probe.
	Targets pulumi.StringArrayInput
	// The RFC 3339-formatted creation date and time of the probe.
	TimeCreated pulumi.StringPtrInput
	// The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds pulumi.IntPtrInput
	// A list of names of vantage points from which to execute the probe.
	VantagePointNames pulumi.StringArrayInput
}

func (HttpProbeState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpProbeState)(nil)).Elem()
}

type httpProbeArgs struct {
	// The OCID of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// A dictionary of HTTP request headers.
	Headers map[string]interface{} `pulumi:"headers"`
	// The supported HTTP methods available for probes.
	Method *string `pulumi:"method"`
	// The optional URL path to probe, including query parameters.
	Path *string `pulumi:"path"`
	// The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol.
	Port *int `pulumi:"port"`
	// The supported protocols available for HTTP probes.
	Protocol string `pulumi:"protocol"`
	// A list of targets (hostnames or IP addresses) of the probe.
	Targets []string `pulumi:"targets"`
	// The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
	// A list of names of vantage points from which to execute the probe.
	VantagePointNames []string `pulumi:"vantagePointNames"`
}

// The set of arguments for constructing a HttpProbe resource.
type HttpProbeArgs struct {
	// The OCID of the compartment.
	CompartmentId pulumi.StringInput
	// A dictionary of HTTP request headers.
	Headers pulumi.MapInput
	// The supported HTTP methods available for probes.
	Method pulumi.StringPtrInput
	// The optional URL path to probe, including query parameters.
	Path pulumi.StringPtrInput
	// The port on which to probe endpoints. If unspecified, probes will use the default port of their protocol.
	Port pulumi.IntPtrInput
	// The supported protocols available for HTTP probes.
	Protocol pulumi.StringInput
	// A list of targets (hostnames or IP addresses) of the probe.
	Targets pulumi.StringArrayInput
	// The probe timeout in seconds. Valid values: 10, 20, 30, and 60. The probe timeout must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds pulumi.IntPtrInput
	// A list of names of vantage points from which to execute the probe.
	VantagePointNames pulumi.StringArrayInput
}

func (HttpProbeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*httpProbeArgs)(nil)).Elem()
}

type HttpProbeInput interface {
	pulumi.Input

	ToHttpProbeOutput() HttpProbeOutput
	ToHttpProbeOutputWithContext(ctx context.Context) HttpProbeOutput
}

func (*HttpProbe) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProbe)(nil))
}

func (i *HttpProbe) ToHttpProbeOutput() HttpProbeOutput {
	return i.ToHttpProbeOutputWithContext(context.Background())
}

func (i *HttpProbe) ToHttpProbeOutputWithContext(ctx context.Context) HttpProbeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProbeOutput)
}

func (i *HttpProbe) ToHttpProbePtrOutput() HttpProbePtrOutput {
	return i.ToHttpProbePtrOutputWithContext(context.Background())
}

func (i *HttpProbe) ToHttpProbePtrOutputWithContext(ctx context.Context) HttpProbePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProbePtrOutput)
}

type HttpProbePtrInput interface {
	pulumi.Input

	ToHttpProbePtrOutput() HttpProbePtrOutput
	ToHttpProbePtrOutputWithContext(ctx context.Context) HttpProbePtrOutput
}

type httpProbePtrType HttpProbeArgs

func (*httpProbePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProbe)(nil))
}

func (i *httpProbePtrType) ToHttpProbePtrOutput() HttpProbePtrOutput {
	return i.ToHttpProbePtrOutputWithContext(context.Background())
}

func (i *httpProbePtrType) ToHttpProbePtrOutputWithContext(ctx context.Context) HttpProbePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProbePtrOutput)
}

// HttpProbeArrayInput is an input type that accepts HttpProbeArray and HttpProbeArrayOutput values.
// You can construct a concrete instance of `HttpProbeArrayInput` via:
//
//          HttpProbeArray{ HttpProbeArgs{...} }
type HttpProbeArrayInput interface {
	pulumi.Input

	ToHttpProbeArrayOutput() HttpProbeArrayOutput
	ToHttpProbeArrayOutputWithContext(context.Context) HttpProbeArrayOutput
}

type HttpProbeArray []HttpProbeInput

func (HttpProbeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpProbe)(nil)).Elem()
}

func (i HttpProbeArray) ToHttpProbeArrayOutput() HttpProbeArrayOutput {
	return i.ToHttpProbeArrayOutputWithContext(context.Background())
}

func (i HttpProbeArray) ToHttpProbeArrayOutputWithContext(ctx context.Context) HttpProbeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProbeArrayOutput)
}

// HttpProbeMapInput is an input type that accepts HttpProbeMap and HttpProbeMapOutput values.
// You can construct a concrete instance of `HttpProbeMapInput` via:
//
//          HttpProbeMap{ "key": HttpProbeArgs{...} }
type HttpProbeMapInput interface {
	pulumi.Input

	ToHttpProbeMapOutput() HttpProbeMapOutput
	ToHttpProbeMapOutputWithContext(context.Context) HttpProbeMapOutput
}

type HttpProbeMap map[string]HttpProbeInput

func (HttpProbeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpProbe)(nil)).Elem()
}

func (i HttpProbeMap) ToHttpProbeMapOutput() HttpProbeMapOutput {
	return i.ToHttpProbeMapOutputWithContext(context.Background())
}

func (i HttpProbeMap) ToHttpProbeMapOutputWithContext(ctx context.Context) HttpProbeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProbeMapOutput)
}

type HttpProbeOutput struct {
	*pulumi.OutputState
}

func (HttpProbeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProbe)(nil))
}

func (o HttpProbeOutput) ToHttpProbeOutput() HttpProbeOutput {
	return o
}

func (o HttpProbeOutput) ToHttpProbeOutputWithContext(ctx context.Context) HttpProbeOutput {
	return o
}

func (o HttpProbeOutput) ToHttpProbePtrOutput() HttpProbePtrOutput {
	return o.ToHttpProbePtrOutputWithContext(context.Background())
}

func (o HttpProbeOutput) ToHttpProbePtrOutputWithContext(ctx context.Context) HttpProbePtrOutput {
	return o.ApplyT(func(v HttpProbe) *HttpProbe {
		return &v
	}).(HttpProbePtrOutput)
}

type HttpProbePtrOutput struct {
	*pulumi.OutputState
}

func (HttpProbePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProbe)(nil))
}

func (o HttpProbePtrOutput) ToHttpProbePtrOutput() HttpProbePtrOutput {
	return o
}

func (o HttpProbePtrOutput) ToHttpProbePtrOutputWithContext(ctx context.Context) HttpProbePtrOutput {
	return o
}

type HttpProbeArrayOutput struct{ *pulumi.OutputState }

func (HttpProbeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpProbe)(nil))
}

func (o HttpProbeArrayOutput) ToHttpProbeArrayOutput() HttpProbeArrayOutput {
	return o
}

func (o HttpProbeArrayOutput) ToHttpProbeArrayOutputWithContext(ctx context.Context) HttpProbeArrayOutput {
	return o
}

func (o HttpProbeArrayOutput) Index(i pulumi.IntInput) HttpProbeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpProbe {
		return vs[0].([]HttpProbe)[vs[1].(int)]
	}).(HttpProbeOutput)
}

type HttpProbeMapOutput struct{ *pulumi.OutputState }

func (HttpProbeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]HttpProbe)(nil))
}

func (o HttpProbeMapOutput) ToHttpProbeMapOutput() HttpProbeMapOutput {
	return o
}

func (o HttpProbeMapOutput) ToHttpProbeMapOutputWithContext(ctx context.Context) HttpProbeMapOutput {
	return o
}

func (o HttpProbeMapOutput) MapIndex(k pulumi.StringInput) HttpProbeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) HttpProbe {
		return vs[0].(map[string]HttpProbe)[vs[1].(string)]
	}).(HttpProbeOutput)
}

func init() {
	pulumi.RegisterOutputType(HttpProbeOutput{})
	pulumi.RegisterOutputType(HttpProbePtrOutput{})
	pulumi.RegisterOutputType(HttpProbeArrayOutput{})
	pulumi.RegisterOutputType(HttpProbeMapOutput{})
}
