// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Log Saved Searches in Oracle Cloud Infrastructure Logging service.
//
// Lists Logging Saved Searches for this compartment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/logging"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := oci_logging_log_saved_search.Test_log_saved_search.Id
// 		opt1 := _var.Log_saved_search_name
// 		_, err := logging.GetLogSavedSearches(ctx, &logging.GetLogSavedSearchesArgs{
// 			CompartmentId:    _var.Compartment_id,
// 			LogSavedSearchId: &opt0,
// 			Name:             &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLogSavedSearches(ctx *pulumi.Context, args *GetLogSavedSearchesArgs, opts ...pulumi.InvokeOption) (*GetLogSavedSearchesResult, error) {
	var rv GetLogSavedSearchesResult
	err := ctx.Invoke("oci:logging/getLogSavedSearches:getLogSavedSearches", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogSavedSearches.
type GetLogSavedSearchesArgs struct {
	// Compartment OCID to list resources in. See compartmentIdInSubtree for nested compartments traversal.
	CompartmentId string                      `pulumi:"compartmentId"`
	Filters       []GetLogSavedSearchesFilter `pulumi:"filters"`
	// OCID of the LogSavedSearch
	LogSavedSearchId *string `pulumi:"logSavedSearchId"`
	// Resource name
	Name *string `pulumi:"name"`
}

// A collection of values returned by getLogSavedSearches.
type GetLogSavedSearchesResult struct {
	// The OCID of the compartment that the resource belongs to.
	CompartmentId string                      `pulumi:"compartmentId"`
	Filters       []GetLogSavedSearchesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	LogSavedSearchId *string `pulumi:"logSavedSearchId"`
	// The list of log_saved_search_summary_collection.
	LogSavedSearchSummaryCollections []GetLogSavedSearchesLogSavedSearchSummaryCollection `pulumi:"logSavedSearchSummaryCollections"`
	// The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information.
	Name *string `pulumi:"name"`
}