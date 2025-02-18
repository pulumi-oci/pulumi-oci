// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsi

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Host Insight resource in Oracle Cloud Infrastructure Opsi service.
//
// Gets details of a host insight.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/opsi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := opsi.LookupHostInsight(ctx, &opsi.LookupHostInsightArgs{
// 			HostInsightId: oci_opsi_host_insight.Test_host_insight.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupHostInsight(ctx *pulumi.Context, args *LookupHostInsightArgs, opts ...pulumi.InvokeOption) (*LookupHostInsightResult, error) {
	var rv LookupHostInsightResult
	err := ctx.Invoke("oci:opsi/getHostInsight:getHostInsight", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHostInsight.
type LookupHostInsightArgs struct {
	// Unique host insight identifier
	HostInsightId string `pulumi:"hostInsightId"`
}

// A collection of values returned by getHostInsight.
type LookupHostInsightResult struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// Source of the host entity.
	EntitySource string `pulumi:"entitySource"`
	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The user-friendly name for the host. The name does not have to be unique.
	HostDisplayName string `pulumi:"hostDisplayName"`
	HostInsightId   string `pulumi:"hostInsightId"`
	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName string `pulumi:"hostName"`
	// Operations Insights internal representation of the host type. Possible value is EXTERNAL-HOST.
	HostType string `pulumi:"hostType"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id string `pulumi:"id"`
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails  string `pulumi:"lifecycleDetails"`
	ManagementAgentId string `pulumi:"managementAgentId"`
	PlatformName      string `pulumi:"platformName"`
	PlatformType      string `pulumi:"platformType"`
	PlatformVersion   string `pulumi:"platformVersion"`
	// Processor count.
	ProcessorCount int `pulumi:"processorCount"`
	// The current state of the host.
	State string `pulumi:"state"`
	// Indicates the status of a host insight in Operations Insights
	Status string `pulumi:"status"`
	// System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// The time the the host insight was first enabled. An RFC3339 formatted datetime string
	TimeCreated string `pulumi:"timeCreated"`
	// The time the host insight was updated. An RFC3339 formatted datetime string
	TimeUpdated string `pulumi:"timeUpdated"`
}
