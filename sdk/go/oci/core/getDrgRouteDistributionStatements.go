// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Drg Route Distribution Statements in Oracle Cloud Infrastructure Core service.
//
// Lists the statements for the specified route distribution.
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
// 		_, err := core.GetDrgRouteDistributionStatements(ctx, &core.GetDrgRouteDistributionStatementsArgs{
// 			DrgRouteDistributionId: oci_core_drg_route_distribution.Test_drg_route_distribution.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDrgRouteDistributionStatements(ctx *pulumi.Context, args *GetDrgRouteDistributionStatementsArgs, opts ...pulumi.InvokeOption) (*GetDrgRouteDistributionStatementsResult, error) {
	var rv GetDrgRouteDistributionStatementsResult
	err := ctx.Invoke("oci:core/getDrgRouteDistributionStatements:getDrgRouteDistributionStatements", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDrgRouteDistributionStatements.
type GetDrgRouteDistributionStatementsArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route distribution.
	DrgRouteDistributionId string                                    `pulumi:"drgRouteDistributionId"`
	Filters                []GetDrgRouteDistributionStatementsFilter `pulumi:"filters"`
}

// A collection of values returned by getDrgRouteDistributionStatements.
type GetDrgRouteDistributionStatementsResult struct {
	DrgRouteDistributionId string `pulumi:"drgRouteDistributionId"`
	// The list of drg_route_distribution_statements.
	DrgRouteDistributionStatements []GetDrgRouteDistributionStatementsDrgRouteDistributionStatement `pulumi:"drgRouteDistributionStatements"`
	Filters                        []GetDrgRouteDistributionStatementsFilter                        `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
