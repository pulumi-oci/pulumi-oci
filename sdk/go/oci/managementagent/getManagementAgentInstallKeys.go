// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managementagent

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Management Agent Install Keys in Oracle Cloud Infrastructure Management Agent service.
//
// Returns a list of Management Agent installed Keys.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/managementagent"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Management_agent_install_key_access_level
// 		opt1 := _var.Management_agent_install_key_compartment_id_in_subtree
// 		opt2 := _var.Management_agent_install_key_display_name
// 		opt3 := _var.Management_agent_install_key_state
// 		_, err := managementagent.GetManagementAgentInstallKeys(ctx, &managementagent.GetManagementAgentInstallKeysArgs{
// 			CompartmentId:          _var.Compartment_id,
// 			AccessLevel:            &opt0,
// 			CompartmentIdInSubtree: &opt1,
// 			DisplayName:            &opt2,
// 			State:                  &opt3,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetManagementAgentInstallKeys(ctx *pulumi.Context, args *GetManagementAgentInstallKeysArgs, opts ...pulumi.InvokeOption) (*GetManagementAgentInstallKeysResult, error) {
	var rv GetManagementAgentInstallKeysResult
	err := ctx.Invoke("oci:managementagent/getManagementAgentInstallKeys:getManagementAgentInstallKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getManagementAgentInstallKeys.
type GetManagementAgentInstallKeysArgs struct {
	// Value of this is always "ACCESSIBLE" and any other value is not supported.
	AccessLevel *string `pulumi:"accessLevel"`
	// The ID of the compartment from which the Management Agents to be listed.
	CompartmentId string `pulumi:"compartmentId"`
	// if set to true then it fetches install key for all compartments where user has access to else only on the compartment specified.
	CompartmentIdInSubtree *bool `pulumi:"compartmentIdInSubtree"`
	// The display name for which the Key needs to be listed.
	DisplayName *string                               `pulumi:"displayName"`
	Filters     []GetManagementAgentInstallKeysFilter `pulumi:"filters"`
	// Filter to return only Management Agents in the particular lifecycle state.
	State *string `pulumi:"state"`
}

// A collection of values returned by getManagementAgentInstallKeys.
type GetManagementAgentInstallKeysResult struct {
	AccessLevel *string `pulumi:"accessLevel"`
	// Compartment Identifier
	CompartmentId          string `pulumi:"compartmentId"`
	CompartmentIdInSubtree *bool  `pulumi:"compartmentIdInSubtree"`
	// Management Agent Install Key Name
	DisplayName *string                               `pulumi:"displayName"`
	Filters     []GetManagementAgentInstallKeysFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of management_agent_install_keys.
	ManagementAgentInstallKeys []GetManagementAgentInstallKeysManagementAgentInstallKey `pulumi:"managementAgentInstallKeys"`
	// Status of Key
	State *string `pulumi:"state"`
}