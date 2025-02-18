// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package meteringcomputation

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Custom Table resource in Oracle Cloud Infrastructure Metering Computation service.
//
// Returns the saved custom table.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/meteringcomputation"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := meteringcomputation.LookupCustomTable(ctx, &meteringcomputation.LookupCustomTableArgs{
// 			CustomTableId: oci_metering_computation_custom_table.Test_custom_table.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCustomTable(ctx *pulumi.Context, args *LookupCustomTableArgs, opts ...pulumi.InvokeOption) (*LookupCustomTableResult, error) {
	var rv LookupCustomTableResult
	err := ctx.Invoke("oci:meteringcomputation/getCustomTable:getCustomTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomTable.
type LookupCustomTableArgs struct {
	// The custom table unique OCID.
	CustomTableId string `pulumi:"customTableId"`
}

// A collection of values returned by getCustomTable.
type LookupCustomTableResult struct {
	// The custom table compartment OCID.
	CompartmentId string `pulumi:"compartmentId"`
	CustomTableId string `pulumi:"customTableId"`
	// The custom table OCID.
	Id string `pulumi:"id"`
	// The custom table for Cost Analysis UI rendering.
	SavedCustomTable GetCustomTableSavedCustomTable `pulumi:"savedCustomTable"`
	// The custom table associated saved report OCID.
	SavedReportId string `pulumi:"savedReportId"`
}
