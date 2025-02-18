// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Fast Connect Provider Services in Oracle Cloud Infrastructure Core service.
//
// Lists the service offerings from supported providers. You need this
// information so you can specify your desired provider and service
// offering when you create a virtual circuit.
//
// For the compartment ID, provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of your tenancy (the root compartment).
//
// For more information, see [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.GetFastConnectProviderServices(ctx, &core.GetFastConnectProviderServicesArgs{
// 			CompartmentId: _var.Compartment_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetFastConnectProviderServices(ctx *pulumi.Context, args *GetFastConnectProviderServicesArgs, opts ...pulumi.InvokeOption) (*GetFastConnectProviderServicesResult, error) {
	var rv GetFastConnectProviderServicesResult
	err := ctx.Invoke("oci:core/getFastConnectProviderServices:getFastConnectProviderServices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFastConnectProviderServices.
type GetFastConnectProviderServicesArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string                                 `pulumi:"compartmentId"`
	Filters       []GetFastConnectProviderServicesFilter `pulumi:"filters"`
}

// A collection of values returned by getFastConnectProviderServices.
type GetFastConnectProviderServicesResult struct {
	CompartmentId string `pulumi:"compartmentId"`
	// The list of fast_connect_provider_services.
	FastConnectProviderServices []GetFastConnectProviderServicesFastConnectProviderService `pulumi:"fastConnectProviderServices"`
	Filters                     []GetFastConnectProviderServicesFilter                     `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
