// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datascience

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Notebook Sessions in Oracle Cloud Infrastructure Data Science service.
//
// Lists the notebook sessions in the specified compartment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/datascience"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Notebook_session_created_by
// 		opt1 := _var.Notebook_session_display_name
// 		opt2 := _var.Notebook_session_id
// 		opt3 := oci_datascience_project.Test_project.Id
// 		opt4 := _var.Notebook_session_state
// 		_, err := datascience.GetNotebookSessions(ctx, &datascience.GetNotebookSessionsArgs{
// 			CompartmentId: _var.Compartment_id,
// 			CreatedBy:     &opt0,
// 			DisplayName:   &opt1,
// 			Id:            &opt2,
// 			ProjectId:     &opt3,
// 			State:         &opt4,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetNotebookSessions(ctx *pulumi.Context, args *GetNotebookSessionsArgs, opts ...pulumi.InvokeOption) (*GetNotebookSessionsResult, error) {
	var rv GetNotebookSessionsResult
	err := ctx.Invoke("oci:datascience/getNotebookSessions:getNotebookSessions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNotebookSessions.
type GetNotebookSessionsArgs struct {
	// <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string                     `pulumi:"displayName"`
	Filters     []GetNotebookSessionsFilter `pulumi:"filters"`
	// <b>Filter</b> results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `pulumi:"id"`
	// <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
	ProjectId *string `pulumi:"projectId"`
	// <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type.
	State *string `pulumi:"state"`
}

// A collection of values returned by getNotebookSessions.
type GetNotebookSessionsResult struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the notebook session's compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the notebook session.
	CreatedBy *string `pulumi:"createdBy"`
	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information. Example: `My NotebookSession`
	DisplayName *string                     `pulumi:"displayName"`
	Filters     []GetNotebookSessionsFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the notebook session.
	Id *string `pulumi:"id"`
	// The list of notebook_sessions.
	NotebookSessions []GetNotebookSessionsNotebookSession `pulumi:"notebookSessions"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the notebook session.
	ProjectId *string `pulumi:"projectId"`
	// The state of the notebook session.
	State *string `pulumi:"state"`
}
