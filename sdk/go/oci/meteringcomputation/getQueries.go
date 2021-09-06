// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package meteringcomputation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Queries in Oracle Cloud Infrastructure Metering Computation service.
//
// Returns the saved query list.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/meteringcomputation"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := meteringcomputation.GetQueries(ctx, &meteringcomputation.GetQueriesArgs{
// 			CompartmentId: _var.Compartment_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetQueries(ctx *pulumi.Context, args *GetQueriesArgs, opts ...pulumi.InvokeOption) (*GetQueriesResult, error) {
	var rv GetQueriesResult
	err := ctx.Invoke("oci:meteringcomputation/getQueries:getQueries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueries.
type GetQueriesArgs struct {
	// The compartment ID in which to list resources.
	CompartmentId string             `pulumi:"compartmentId"`
	Filters       []GetQueriesFilter `pulumi:"filters"`
}

// A collection of values returned by getQueries.
type GetQueriesResult struct {
	// The compartment OCID.
	CompartmentId string             `pulumi:"compartmentId"`
	Filters       []GetQueriesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of query_collection.
	QueryCollections []GetQueriesQueryCollection `pulumi:"queryCollections"`
}