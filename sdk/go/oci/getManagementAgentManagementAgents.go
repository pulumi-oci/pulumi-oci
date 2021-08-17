// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Management Agents in Oracle Cloud Infrastructure Management Agent service.
//
// Returns a list of Management Agent.
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
// 		opt0 := _var.Management_agent_display_name
// 		opt1 := _var.Management_agent_platform_type
// 		opt2 := _var.Management_agent_plugin_name
// 		opt3 := _var.Management_agent_state
// 		opt4 := _var.Management_agent_version
// 		_, err := oci.GetManagementAgentManagementAgents(ctx, &GetManagementAgentManagementAgentsArgs{
// 			CompartmentId: _var.Compartment_id,
// 			DisplayName:   &opt0,
// 			PlatformType:  &opt1,
// 			PluginName:    &opt2,
// 			State:         &opt3,
// 			Version:       &opt4,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetManagementAgentManagementAgents(ctx *pulumi.Context, args *GetManagementAgentManagementAgentsArgs, opts ...pulumi.InvokeOption) (*GetManagementAgentManagementAgentsResult, error) {
	var rv GetManagementAgentManagementAgentsResult
	err := ctx.Invoke("oci:index/getManagementAgentManagementAgents:GetManagementAgentManagementAgents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetManagementAgentManagementAgents.
type GetManagementAgentManagementAgentsArgs struct {
	// The ID of the compartment from which the Management Agents to be listed.
	CompartmentId string `pulumi:"compartmentId"`
	// Filter to return only Management Agents having the particular display name.
	DisplayName *string                                    `pulumi:"displayName"`
	Filters     []GetManagementAgentManagementAgentsFilter `pulumi:"filters"`
	// Filter to return only Management Agents having the particular platform type.
	PlatformType *string `pulumi:"platformType"`
	// Filter to return only Management Agents having the particular Plugin installed.
	PluginName *string `pulumi:"pluginName"`
	// Filter to return only Management Agents in the particular lifecycle state.
	State *string `pulumi:"state"`
	// Filter to return only Management Agents having the particular agent version.
	Version *string `pulumi:"version"`
}

// A collection of values returned by GetManagementAgentManagementAgents.
type GetManagementAgentManagementAgentsResult struct {
	// Compartment Identifier
	CompartmentId string `pulumi:"compartmentId"`
	// Management Agent Name
	DisplayName *string                                    `pulumi:"displayName"`
	Filters     []GetManagementAgentManagementAgentsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of management_agents.
	ManagementAgents []GetManagementAgentManagementAgentsManagementAgent `pulumi:"managementAgents"`
	// Platform Type
	PlatformType *string `pulumi:"platformType"`
	// Management Agent Plugin Name
	PluginName *string `pulumi:"pluginName"`
	// The current state of managementAgent
	State *string `pulumi:"state"`
	// Management Agent Version
	Version *string `pulumi:"version"`
}