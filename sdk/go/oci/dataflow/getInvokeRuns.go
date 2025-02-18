// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataflow

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Invoke Runs in Oracle Cloud Infrastructure Data Flow service.
//
// Lists all runs of an application in the specified compartment.  Only one parameter other than compartmentId may also be included in a query. The query must include compartmentId. If the query does not include compartmentId, or includes compartmentId but two or more other parameters an error is returned.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/dataflow"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := oci_dataflow_application.Test_application.Id
// 		opt1 := _var.Invoke_run_display_name
// 		opt2 := _var.Invoke_run_display_name_starts_with
// 		opt3 := oci_dataflow_owner_principal.Test_owner_principal.Id
// 		opt4 := _var.Invoke_run_state
// 		opt5 := _var.Invoke_run_time_created_greater_than
// 		_, err := dataflow.GetInvokeRuns(ctx, &dataflow.GetInvokeRunsArgs{
// 			CompartmentId:          _var.Compartment_id,
// 			ApplicationId:          &opt0,
// 			DisplayName:            &opt1,
// 			DisplayNameStartsWith:  &opt2,
// 			OwnerPrincipalId:       &opt3,
// 			State:                  &opt4,
// 			TimeCreatedGreaterThan: &opt5,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetInvokeRuns(ctx *pulumi.Context, args *GetInvokeRunsArgs, opts ...pulumi.InvokeOption) (*GetInvokeRunsResult, error) {
	var rv GetInvokeRunsResult
	err := ctx.Invoke("oci:dataflow/getInvokeRuns:getInvokeRuns", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInvokeRuns.
type GetInvokeRunsArgs struct {
	// The ID of the application.
	ApplicationId *string `pulumi:"applicationId"`
	// The OCID of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// The query parameter for the Spark application name.
	DisplayName *string `pulumi:"displayName"`
	// The displayName prefix.
	DisplayNameStartsWith *string               `pulumi:"displayNameStartsWith"`
	Filters               []GetInvokeRunsFilter `pulumi:"filters"`
	// The OCID of the user who created the resource.
	OwnerPrincipalId *string `pulumi:"ownerPrincipalId"`
	// The LifecycleState of the run.
	State *string `pulumi:"state"`
	// The epoch time that the resource was created.
	TimeCreatedGreaterThan *string `pulumi:"timeCreatedGreaterThan"`
}

// A collection of values returned by getInvokeRuns.
type GetInvokeRunsResult struct {
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// The OCID of a compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// A user-friendly name. This name is not necessarily unique.
	DisplayName           *string               `pulumi:"displayName"`
	DisplayNameStartsWith *string               `pulumi:"displayNameStartsWith"`
	Filters               []GetInvokeRunsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The OCID of the user who created the resource.
	OwnerPrincipalId *string `pulumi:"ownerPrincipalId"`
	// The list of runs.
	Runs []GetInvokeRunsRun `pulumi:"runs"`
	// The current state of this run.
	State                  *string `pulumi:"state"`
	TimeCreatedGreaterThan *string `pulumi:"timeCreatedGreaterThan"`
}
