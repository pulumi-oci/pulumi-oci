// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Cross Connect Groups in Oracle Cloud Infrastructure Core service.
//
// Lists the cross-connect groups in the specified compartment.
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
// 		opt0 := _var.Cross_connect_group_display_name
// 		opt1 := _var.Cross_connect_group_state
// 		_, err := core.GetCrossConnectGroups(ctx, &core.GetCrossConnectGroupsArgs{
// 			CompartmentId: _var.Compartment_id,
// 			DisplayName:   &opt0,
// 			State:         &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCrossConnectGroups(ctx *pulumi.Context, args *GetCrossConnectGroupsArgs, opts ...pulumi.InvokeOption) (*GetCrossConnectGroupsResult, error) {
	var rv GetCrossConnectGroupsResult
	err := ctx.Invoke("oci:core/getCrossConnectGroups:getCrossConnectGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCrossConnectGroups.
type GetCrossConnectGroupsArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// A filter to return only resources that match the given display name exactly.
	DisplayName *string                       `pulumi:"displayName"`
	Filters     []GetCrossConnectGroupsFilter `pulumi:"filters"`
	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	State *string `pulumi:"state"`
}

// A collection of values returned by getCrossConnectGroups.
type GetCrossConnectGroupsResult struct {
	// The OCID of the compartment containing the cross-connect group.
	CompartmentId string `pulumi:"compartmentId"`
	// The list of cross_connect_groups.
	CrossConnectGroups []GetCrossConnectGroupsCrossConnectGroup `pulumi:"crossConnectGroups"`
	// The display name of a user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string                       `pulumi:"displayName"`
	Filters     []GetCrossConnectGroupsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The cross-connect group's current state.
	State *string `pulumi:"state"`
}
