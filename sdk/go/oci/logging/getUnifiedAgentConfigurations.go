// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Unified Agent Configurations in Oracle Cloud Infrastructure Logging service.
//
// Lists all unified agent configurations in the specified compartment.
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
// 		opt0 := _var.Unified_agent_configuration_display_name
// 		opt1 := oci_identity_group.Test_group.Id
// 		opt2 := _var.Unified_agent_configuration_is_compartment_id_in_subtree
// 		opt3 := oci_logging_log.Test_log.Id
// 		opt4 := _var.Unified_agent_configuration_state
// 		_, err := logging.GetUnifiedAgentConfigurations(ctx, &logging.GetUnifiedAgentConfigurationsArgs{
// 			CompartmentId:            _var.Compartment_id,
// 			DisplayName:              &opt0,
// 			GroupId:                  &opt1,
// 			IsCompartmentIdInSubtree: &opt2,
// 			LogId:                    &opt3,
// 			State:                    &opt4,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetUnifiedAgentConfigurations(ctx *pulumi.Context, args *GetUnifiedAgentConfigurationsArgs, opts ...pulumi.InvokeOption) (*GetUnifiedAgentConfigurationsResult, error) {
	var rv GetUnifiedAgentConfigurationsResult
	err := ctx.Invoke("oci:logging/getUnifiedAgentConfigurations:getUnifiedAgentConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUnifiedAgentConfigurations.
type GetUnifiedAgentConfigurationsArgs struct {
	// Compartment OCID to list resources in. See compartmentIdInSubtree for nested compartments traversal.
	CompartmentId string `pulumi:"compartmentId"`
	// Resource name
	DisplayName *string                               `pulumi:"displayName"`
	Filters     []GetUnifiedAgentConfigurationsFilter `pulumi:"filters"`
	// The OCID of a group or a dynamic group.
	GroupId *string `pulumi:"groupId"`
	// Specifies whether or not nested compartments should be traversed. Defaults to false.
	IsCompartmentIdInSubtree *bool `pulumi:"isCompartmentIdInSubtree"`
	// Custom log OCID to list resources with the log as destination.
	LogId *string `pulumi:"logId"`
	// Lifecycle state of the log object
	State *string `pulumi:"state"`
}

// A collection of values returned by getUnifiedAgentConfigurations.
type GetUnifiedAgentConfigurationsResult struct {
	// The OCID of the compartment that the resource belongs to.
	CompartmentId string `pulumi:"compartmentId"`
	// The user-friendly display name. This must be unique within the enclosing resource, and it's changeable. Avoid entering confidential information.
	DisplayName *string                               `pulumi:"displayName"`
	Filters     []GetUnifiedAgentConfigurationsFilter `pulumi:"filters"`
	GroupId     *string                               `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string  `pulumi:"id"`
	IsCompartmentIdInSubtree *bool   `pulumi:"isCompartmentIdInSubtree"`
	LogId                    *string `pulumi:"logId"`
	// The pipeline state.
	State *string `pulumi:"state"`
	// The list of unified_agent_configuration_collection.
	UnifiedAgentConfigurationCollections []GetUnifiedAgentConfigurationsUnifiedAgentConfigurationCollection `pulumi:"unifiedAgentConfigurationCollections"`
}
