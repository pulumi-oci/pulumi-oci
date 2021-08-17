// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Backends in Oracle Cloud Infrastructure Load Balancer service.
//
// Lists the backend servers for a given load balancer and backend set.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oci.GetLoadBalancerBackends(ctx, &GetLoadBalancerBackendsArgs{
// 			BackendsetName: oci_load_balancer_backend_set.Test_backend_set.Name,
// 			LoadBalancerId: oci_load_balancer_load_balancer.Test_load_balancer.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLoadBalancerBackends(ctx *pulumi.Context, args *GetLoadBalancerBackendsArgs, opts ...pulumi.InvokeOption) (*GetLoadBalancerBackendsResult, error) {
	var rv GetLoadBalancerBackendsResult
	err := ctx.Invoke("oci:index/getLoadBalancerBackends:GetLoadBalancerBackends", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetLoadBalancerBackends.
type GetLoadBalancerBackendsArgs struct {
	// The name of the backend set associated with the backend servers.  Example: `exampleBackendSet`
	BackendsetName string                          `pulumi:"backendsetName"`
	Filters        []GetLoadBalancerBackendsFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend set and servers.
	LoadBalancerId string `pulumi:"loadBalancerId"`
}

// A collection of values returned by GetLoadBalancerBackends.
type GetLoadBalancerBackendsResult struct {
	// The list of backends.
	Backends       []GetLoadBalancerBackendsBackend `pulumi:"backends"`
	BackendsetName string                           `pulumi:"backendsetName"`
	Filters        []GetLoadBalancerBackendsFilter  `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	LoadBalancerId string `pulumi:"loadBalancerId"`
}