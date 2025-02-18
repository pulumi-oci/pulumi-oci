// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerengine

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Work Request Errors in Oracle Cloud Infrastructure Container Engine service.
//
// Get the errors of a work request.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/containerengine"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := containerengine.GetWorkRequestErrors(ctx, &containerengine.GetWorkRequestErrorsArgs{
// 			CompartmentId: _var.Compartment_id,
// 			WorkRequestId: oci_containerengine_work_request.Test_work_request.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetWorkRequestErrors(ctx *pulumi.Context, args *GetWorkRequestErrorsArgs, opts ...pulumi.InvokeOption) (*GetWorkRequestErrorsResult, error) {
	var rv GetWorkRequestErrorsResult
	err := ctx.Invoke("oci:containerengine/getWorkRequestErrors:getWorkRequestErrors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWorkRequestErrors.
type GetWorkRequestErrorsArgs struct {
	// The OCID of the compartment.
	CompartmentId string                       `pulumi:"compartmentId"`
	Filters       []GetWorkRequestErrorsFilter `pulumi:"filters"`
	// The OCID of the work request.
	WorkRequestId string `pulumi:"workRequestId"`
}

// A collection of values returned by getWorkRequestErrors.
type GetWorkRequestErrorsResult struct {
	CompartmentId string                       `pulumi:"compartmentId"`
	Filters       []GetWorkRequestErrorsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of work_request_errors.
	WorkRequestErrors []GetWorkRequestErrorsWorkRequestError `pulumi:"workRequestErrors"`
	WorkRequestId     string                                 `pulumi:"workRequestId"`
}
