// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databasemigration

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Agent resource in Oracle Cloud Infrastructure Database Migration service.
//
// Display the ODMS Agent configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/databasemigration"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := databasemigration.LookupAgent(ctx, &databasemigration.LookupAgentArgs{
// 			AgentId: oci_database_migration_agent.Test_agent.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupAgent(ctx *pulumi.Context, args *LookupAgentArgs, opts ...pulumi.InvokeOption) (*LookupAgentResult, error) {
	var rv LookupAgentResult
	err := ctx.Invoke("oci:databasemigration/getAgent:getAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAgent.
type LookupAgentArgs struct {
	// The OCID of the agent
	AgentId string `pulumi:"agentId"`
}

// A collection of values returned by getAgent.
type LookupAgentResult struct {
	AgentId string `pulumi:"agentId"`
	// OCID of the compartment
	CompartmentId string `pulumi:"compartmentId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// ODMS Agent name
	DisplayName string `pulumi:"displayName"`
	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The OCID of the resource
	Id string `pulumi:"id"`
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails string `pulumi:"lifecycleDetails"`
	// ODMS Agent public key.
	PublicKey string `pulumi:"publicKey"`
	// The current state of the ODMS On Prem Agent.
	State string `pulumi:"state"`
	// The OCID of the Stream
	StreamId string `pulumi:"streamId"`
	// Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// The time the Agent was created. An RFC3339 formatted datetime string.
	TimeCreated string `pulumi:"timeCreated"`
	// The time of the last Agent details update. An RFC3339 formatted datetime string.
	TimeUpdated string `pulumi:"timeUpdated"`
	// ODMS Agent version
	Version string `pulumi:"version"`
}
