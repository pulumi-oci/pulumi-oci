// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Log Groups in Oracle Cloud Infrastructure Logging service.
//
// Lists all log groups for the specified compartment or tenancy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/logging"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Log_group_display_name
// 		opt1 := _var.Log_group_is_compartment_id_in_subtree
// 		_, err := logging.GetLogGroups(ctx, &logging.GetLogGroupsArgs{
// 			CompartmentId:            _var.Compartment_id,
// 			DisplayName:              &opt0,
// 			IsCompartmentIdInSubtree: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLogGroups(ctx *pulumi.Context, args *GetLogGroupsArgs, opts ...pulumi.InvokeOption) (*GetLogGroupsResult, error) {
	var rv GetLogGroupsResult
	err := ctx.Invoke("oci:logging/getLogGroups:getLogGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogGroups.
type GetLogGroupsArgs struct {
	// Compartment OCID to list resources in. See compartmentIdInSubtree for nested compartments traversal.
	CompartmentId string `pulumi:"compartmentId"`
	// Resource name
	DisplayName *string              `pulumi:"displayName"`
	Filters     []GetLogGroupsFilter `pulumi:"filters"`
	// Specifies whether or not nested compartments should be traversed. Defaults to false.
	IsCompartmentIdInSubtree *bool `pulumi:"isCompartmentIdInSubtree"`
}

// A collection of values returned by getLogGroups.
type GetLogGroupsResult struct {
	// The OCID of the compartment that the resource belongs to.
	CompartmentId string `pulumi:"compartmentId"`
	// The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information.
	DisplayName *string              `pulumi:"displayName"`
	Filters     []GetLogGroupsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string `pulumi:"id"`
	IsCompartmentIdInSubtree *bool  `pulumi:"isCompartmentIdInSubtree"`
	// The list of log_groups.
	LogGroups []GetLogGroupsLogGroup `pulumi:"logGroups"`
}
