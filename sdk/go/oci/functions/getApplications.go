// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package functions

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Applications in Oracle Cloud Infrastructure Functions service.
//
// Lists applications for a compartment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/functions"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Application_display_name
// 		opt1 := _var.Application_id
// 		opt2 := _var.Application_state
// 		_, err := functions.GetApplications(ctx, &functions.GetApplicationsArgs{
// 			CompartmentId: _var.Compartment_id,
// 			DisplayName:   &opt0,
// 			Id:            &opt1,
// 			State:         &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetApplications(ctx *pulumi.Context, args *GetApplicationsArgs, opts ...pulumi.InvokeOption) (*GetApplicationsResult, error) {
	var rv GetApplicationsResult
	err := ctx.Invoke("oci:functions/getApplications:getApplications", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplications.
type GetApplicationsArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this resource belongs.
	CompartmentId string `pulumi:"compartmentId"`
	// A filter to return only applications with display names that match the display name string. Matching is exact.
	DisplayName *string                 `pulumi:"displayName"`
	Filters     []GetApplicationsFilter `pulumi:"filters"`
	// A filter to return only applications with the specified OCID.
	Id *string `pulumi:"id"`
	// A filter to return only applications that match the lifecycle state in this parameter. Example: `Creating`
	State *string `pulumi:"state"`
}

// A collection of values returned by getApplications.
type GetApplicationsResult struct {
	// The list of applications.
	Applications []GetApplicationsApplication `pulumi:"applications"`
	// The OCID of the compartment that contains the application.
	CompartmentId string `pulumi:"compartmentId"`
	// The display name of the application. The display name is unique within the compartment containing the application.
	DisplayName *string                 `pulumi:"displayName"`
	Filters     []GetApplicationsFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.
	Id *string `pulumi:"id"`
	// The current state of the application.
	State *string `pulumi:"state"`
}
