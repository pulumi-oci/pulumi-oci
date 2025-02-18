// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkloadbalancer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Backend Set resource in Oracle Cloud Infrastructure Network Load Balancer service.
//
// Adds a backend set to a network load balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/networkloadbalancer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkloadbalancer.NewBackendSet(ctx, "testBackendSet", &networkloadbalancer.BackendSetArgs{
// 			HealthChecker: &networkloadbalancer.BackendSetHealthCheckerArgs{
// 				Protocol:          pulumi.Any(_var.Backend_set_health_checker_protocol),
// 				IntervalInMillis:  pulumi.Any(_var.Backend_set_health_checker_interval_in_millis),
// 				Port:              pulumi.Any(_var.Backend_set_health_checker_port),
// 				RequestData:       pulumi.Any(_var.Backend_set_health_checker_request_data),
// 				ResponseBodyRegex: pulumi.Any(_var.Backend_set_health_checker_response_body_regex),
// 				ResponseData:      pulumi.Any(_var.Backend_set_health_checker_response_data),
// 				Retries:           pulumi.Any(_var.Backend_set_health_checker_retries),
// 				ReturnCode:        pulumi.Any(_var.Backend_set_health_checker_return_code),
// 				TimeoutInMillis:   pulumi.Any(_var.Backend_set_health_checker_timeout_in_millis),
// 				UrlPath:           pulumi.Any(_var.Backend_set_health_checker_url_path),
// 			},
// 			NetworkLoadBalancerId: pulumi.Any(oci_network_load_balancer_network_load_balancer.Test_network_load_balancer.Id),
// 			Policy:                pulumi.Any(_var.Backend_set_policy),
// 			IsPreserveSource:      pulumi.Any(_var.Backend_set_is_preserve_source),
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
// BackendSets can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:networkloadbalancer/backendSet:BackendSet test_backend_set "networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}"
// ```
type BackendSet struct {
	pulumi.CustomResourceState

	// Array of backends.
	Backends BackendSetBackendArrayOutput `pulumi:"backends"`
	// (Updatable) The health check policy configuration. For more information, see [Editing Health Check Policies](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/editinghealthcheck.htm).
	HealthChecker BackendSetHealthCheckerOutput `pulumi:"healthChecker"`
	// (Updatable) If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends. Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled. The value is true by default.
	IsPreserveSource pulumi.BoolOutput `pulumi:"isPreserveSource"`
	// A user-friendly name for the backend set that must be unique and cannot be changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId pulumi.StringOutput `pulumi:"networkLoadBalancerId"`
	// (Updatable) The network load balancer policy for the backend set.  Example: `FIVE_TUPLE``
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewBackendSet registers a new resource with the given unique name, arguments, and options.
func NewBackendSet(ctx *pulumi.Context,
	name string, args *BackendSetArgs, opts ...pulumi.ResourceOption) (*BackendSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HealthChecker == nil {
		return nil, errors.New("invalid value for required argument 'HealthChecker'")
	}
	if args.NetworkLoadBalancerId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkLoadBalancerId'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource BackendSet
	err := ctx.RegisterResource("oci:networkloadbalancer/backendSet:BackendSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackendSet gets an existing BackendSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendSetState, opts ...pulumi.ResourceOption) (*BackendSet, error) {
	var resource BackendSet
	err := ctx.ReadResource("oci:networkloadbalancer/backendSet:BackendSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackendSet resources.
type backendSetState struct {
	// Array of backends.
	Backends []BackendSetBackend `pulumi:"backends"`
	// (Updatable) The health check policy configuration. For more information, see [Editing Health Check Policies](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/editinghealthcheck.htm).
	HealthChecker *BackendSetHealthChecker `pulumi:"healthChecker"`
	// (Updatable) If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends. Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled. The value is true by default.
	IsPreserveSource *bool `pulumi:"isPreserveSource"`
	// A user-friendly name for the backend set that must be unique and cannot be changed.
	Name *string `pulumi:"name"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId *string `pulumi:"networkLoadBalancerId"`
	// (Updatable) The network load balancer policy for the backend set.  Example: `FIVE_TUPLE``
	Policy *string `pulumi:"policy"`
}

type BackendSetState struct {
	// Array of backends.
	Backends BackendSetBackendArrayInput
	// (Updatable) The health check policy configuration. For more information, see [Editing Health Check Policies](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/editinghealthcheck.htm).
	HealthChecker BackendSetHealthCheckerPtrInput
	// (Updatable) If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends. Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled. The value is true by default.
	IsPreserveSource pulumi.BoolPtrInput
	// A user-friendly name for the backend set that must be unique and cannot be changed.
	Name pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId pulumi.StringPtrInput
	// (Updatable) The network load balancer policy for the backend set.  Example: `FIVE_TUPLE``
	Policy pulumi.StringPtrInput
}

func (BackendSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendSetState)(nil)).Elem()
}

type backendSetArgs struct {
	// (Updatable) The health check policy configuration. For more information, see [Editing Health Check Policies](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/editinghealthcheck.htm).
	HealthChecker BackendSetHealthChecker `pulumi:"healthChecker"`
	// (Updatable) If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends. Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled. The value is true by default.
	IsPreserveSource *bool `pulumi:"isPreserveSource"`
	// A user-friendly name for the backend set that must be unique and cannot be changed.
	Name *string `pulumi:"name"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId string `pulumi:"networkLoadBalancerId"`
	// (Updatable) The network load balancer policy for the backend set.  Example: `FIVE_TUPLE``
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a BackendSet resource.
type BackendSetArgs struct {
	// (Updatable) The health check policy configuration. For more information, see [Editing Health Check Policies](https://docs.cloud.oracle.com/iaas/Content/Balance/Tasks/editinghealthcheck.htm).
	HealthChecker BackendSetHealthCheckerInput
	// (Updatable) If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends. Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled. The value is true by default.
	IsPreserveSource pulumi.BoolPtrInput
	// A user-friendly name for the backend set that must be unique and cannot be changed.
	Name pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId pulumi.StringInput
	// (Updatable) The network load balancer policy for the backend set.  Example: `FIVE_TUPLE``
	Policy pulumi.StringInput
}

func (BackendSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendSetArgs)(nil)).Elem()
}

type BackendSetInput interface {
	pulumi.Input

	ToBackendSetOutput() BackendSetOutput
	ToBackendSetOutputWithContext(ctx context.Context) BackendSetOutput
}

func (*BackendSet) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendSet)(nil))
}

func (i *BackendSet) ToBackendSetOutput() BackendSetOutput {
	return i.ToBackendSetOutputWithContext(context.Background())
}

func (i *BackendSet) ToBackendSetOutputWithContext(ctx context.Context) BackendSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendSetOutput)
}

func (i *BackendSet) ToBackendSetPtrOutput() BackendSetPtrOutput {
	return i.ToBackendSetPtrOutputWithContext(context.Background())
}

func (i *BackendSet) ToBackendSetPtrOutputWithContext(ctx context.Context) BackendSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendSetPtrOutput)
}

type BackendSetPtrInput interface {
	pulumi.Input

	ToBackendSetPtrOutput() BackendSetPtrOutput
	ToBackendSetPtrOutputWithContext(ctx context.Context) BackendSetPtrOutput
}

type backendSetPtrType BackendSetArgs

func (*backendSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendSet)(nil))
}

func (i *backendSetPtrType) ToBackendSetPtrOutput() BackendSetPtrOutput {
	return i.ToBackendSetPtrOutputWithContext(context.Background())
}

func (i *backendSetPtrType) ToBackendSetPtrOutputWithContext(ctx context.Context) BackendSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendSetPtrOutput)
}

// BackendSetArrayInput is an input type that accepts BackendSetArray and BackendSetArrayOutput values.
// You can construct a concrete instance of `BackendSetArrayInput` via:
//
//          BackendSetArray{ BackendSetArgs{...} }
type BackendSetArrayInput interface {
	pulumi.Input

	ToBackendSetArrayOutput() BackendSetArrayOutput
	ToBackendSetArrayOutputWithContext(context.Context) BackendSetArrayOutput
}

type BackendSetArray []BackendSetInput

func (BackendSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackendSet)(nil)).Elem()
}

func (i BackendSetArray) ToBackendSetArrayOutput() BackendSetArrayOutput {
	return i.ToBackendSetArrayOutputWithContext(context.Background())
}

func (i BackendSetArray) ToBackendSetArrayOutputWithContext(ctx context.Context) BackendSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendSetArrayOutput)
}

// BackendSetMapInput is an input type that accepts BackendSetMap and BackendSetMapOutput values.
// You can construct a concrete instance of `BackendSetMapInput` via:
//
//          BackendSetMap{ "key": BackendSetArgs{...} }
type BackendSetMapInput interface {
	pulumi.Input

	ToBackendSetMapOutput() BackendSetMapOutput
	ToBackendSetMapOutputWithContext(context.Context) BackendSetMapOutput
}

type BackendSetMap map[string]BackendSetInput

func (BackendSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackendSet)(nil)).Elem()
}

func (i BackendSetMap) ToBackendSetMapOutput() BackendSetMapOutput {
	return i.ToBackendSetMapOutputWithContext(context.Background())
}

func (i BackendSetMap) ToBackendSetMapOutputWithContext(ctx context.Context) BackendSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendSetMapOutput)
}

type BackendSetOutput struct {
	*pulumi.OutputState
}

func (BackendSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendSet)(nil))
}

func (o BackendSetOutput) ToBackendSetOutput() BackendSetOutput {
	return o
}

func (o BackendSetOutput) ToBackendSetOutputWithContext(ctx context.Context) BackendSetOutput {
	return o
}

func (o BackendSetOutput) ToBackendSetPtrOutput() BackendSetPtrOutput {
	return o.ToBackendSetPtrOutputWithContext(context.Background())
}

func (o BackendSetOutput) ToBackendSetPtrOutputWithContext(ctx context.Context) BackendSetPtrOutput {
	return o.ApplyT(func(v BackendSet) *BackendSet {
		return &v
	}).(BackendSetPtrOutput)
}

type BackendSetPtrOutput struct {
	*pulumi.OutputState
}

func (BackendSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendSet)(nil))
}

func (o BackendSetPtrOutput) ToBackendSetPtrOutput() BackendSetPtrOutput {
	return o
}

func (o BackendSetPtrOutput) ToBackendSetPtrOutputWithContext(ctx context.Context) BackendSetPtrOutput {
	return o
}

type BackendSetArrayOutput struct{ *pulumi.OutputState }

func (BackendSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendSet)(nil))
}

func (o BackendSetArrayOutput) ToBackendSetArrayOutput() BackendSetArrayOutput {
	return o
}

func (o BackendSetArrayOutput) ToBackendSetArrayOutputWithContext(ctx context.Context) BackendSetArrayOutput {
	return o
}

func (o BackendSetArrayOutput) Index(i pulumi.IntInput) BackendSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendSet {
		return vs[0].([]BackendSet)[vs[1].(int)]
	}).(BackendSetOutput)
}

type BackendSetMapOutput struct{ *pulumi.OutputState }

func (BackendSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BackendSet)(nil))
}

func (o BackendSetMapOutput) ToBackendSetMapOutput() BackendSetMapOutput {
	return o
}

func (o BackendSetMapOutput) ToBackendSetMapOutputWithContext(ctx context.Context) BackendSetMapOutput {
	return o
}

func (o BackendSetMapOutput) MapIndex(k pulumi.StringInput) BackendSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BackendSet {
		return vs[0].(map[string]BackendSet)[vs[1].(string)]
	}).(BackendSetOutput)
}

func init() {
	pulumi.RegisterOutputType(BackendSetOutput{})
	pulumi.RegisterOutputType(BackendSetPtrOutput{})
	pulumi.RegisterOutputType(BackendSetArrayOutput{})
	pulumi.RegisterOutputType(BackendSetMapOutput{})
}
