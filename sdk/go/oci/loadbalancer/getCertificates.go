// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Certificates in Oracle Cloud Infrastructure Load Balancer service.
//
// Lists all SSL certificates bundles associated with a given load balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/loadbalancer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := loadbalancer.GetCertificates(ctx, &loadbalancer.GetCertificatesArgs{
// 			LoadBalancerId: oci_load_balancer_load_balancer.Test_load_balancer.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCertificates(ctx *pulumi.Context, args *GetCertificatesArgs, opts ...pulumi.InvokeOption) (*GetCertificatesResult, error) {
	var rv GetCertificatesResult
	err := ctx.Invoke("oci:loadbalancer/getCertificates:getCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificates.
type GetCertificatesArgs struct {
	Filters []GetCertificatesFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the certificate bundles to be listed.
	LoadBalancerId string `pulumi:"loadBalancerId"`
}

// A collection of values returned by getCertificates.
type GetCertificatesResult struct {
	// The list of certificates.
	Certificates []GetCertificatesCertificate `pulumi:"certificates"`
	Filters      []GetCertificatesFilter      `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	LoadBalancerId string `pulumi:"loadBalancerId"`
}
