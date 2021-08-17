// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Run Log resource in Oracle Cloud Infrastructure Data Flow service.
//
// Retrieves the content of an run log.
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
// 		_, err := oci.GetDataflowRunLog(ctx, &GetDataflowRunLogArgs{
// 			Name:  _var.Run_log_name,
// 			RunId: oci_dataflow_run.Test_run.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDataflowRunLog(ctx *pulumi.Context, args *GetDataflowRunLogArgs, opts ...pulumi.InvokeOption) (*GetDataflowRunLogResult, error) {
	var rv GetDataflowRunLogResult
	err := ctx.Invoke("oci:index/getDataflowRunLog:GetDataflowRunLog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetDataflowRunLog.
type GetDataflowRunLogArgs struct {
	Base64EncodeContent *bool `pulumi:"base64EncodeContent"`
	// The name of the log. Avoid entering confidential information.
	Name string `pulumi:"name"`
	// The unique ID for the run
	RunId string `pulumi:"runId"`
}

// A collection of values returned by GetDataflowRunLog.
type GetDataflowRunLogResult struct {
	Base64EncodeContent *bool `pulumi:"base64EncodeContent"`
	// The content of the run log.
	Content string `pulumi:"content"`
	// The content type of the run log.
	ContentType string `pulumi:"contentType"`
	// The provider-assigned unique ID for this managed resource.
	Id    string `pulumi:"id"`
	Name  string `pulumi:"name"`
	RunId string `pulumi:"runId"`
}