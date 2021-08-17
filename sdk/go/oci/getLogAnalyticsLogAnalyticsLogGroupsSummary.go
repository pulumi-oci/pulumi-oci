// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Log Analytics Log Groups Summary resource in Oracle Cloud Infrastructure Log Analytics service.
//
// Returns the count of log groups in a compartment.
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
// 		_, err := oci.GetLogAnalyticsLogAnalyticsLogGroupsSummary(ctx, &GetLogAnalyticsLogAnalyticsLogGroupsSummaryArgs{
// 			CompartmentId: _var.Compartment_id,
// 			Namespace:     _var.Log_analytics_log_groups_summary_namespace,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLogAnalyticsLogAnalyticsLogGroupsSummary(ctx *pulumi.Context, args *GetLogAnalyticsLogAnalyticsLogGroupsSummaryArgs, opts ...pulumi.InvokeOption) (*GetLogAnalyticsLogAnalyticsLogGroupsSummaryResult, error) {
	var rv GetLogAnalyticsLogAnalyticsLogGroupsSummaryResult
	err := ctx.Invoke("oci:index/getLogAnalyticsLogAnalyticsLogGroupsSummary:GetLogAnalyticsLogAnalyticsLogGroupsSummary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetLogAnalyticsLogAnalyticsLogGroupsSummary.
type GetLogAnalyticsLogAnalyticsLogGroupsSummaryArgs struct {
	// The ID of the compartment in which to list resources.
	CompartmentId string `pulumi:"compartmentId"`
	// The Logging Analytics namespace used for the request.
	Namespace string `pulumi:"namespace"`
}

// A collection of values returned by GetLogAnalyticsLogAnalyticsLogGroupsSummary.
type GetLogAnalyticsLogAnalyticsLogGroupsSummaryResult struct {
	CompartmentId string `pulumi:"compartmentId"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	LogGroupCount int    `pulumi:"logGroupCount"`
	Namespace     string `pulumi:"namespace"`
}