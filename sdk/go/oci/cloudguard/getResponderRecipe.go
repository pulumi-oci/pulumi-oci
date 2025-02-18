// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudguard

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Responder Recipe resource in Oracle Cloud Infrastructure Cloud Guard service.
//
// Get a ResponderRecipe by identifier
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/cloudguard"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudguard.LookupResponderRecipe(ctx, &cloudguard.LookupResponderRecipeArgs{
// 			ResponderRecipeId: oci_cloud_guard_responder_recipe.Test_responder_recipe.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupResponderRecipe(ctx *pulumi.Context, args *LookupResponderRecipeArgs, opts ...pulumi.InvokeOption) (*LookupResponderRecipeResult, error) {
	var rv LookupResponderRecipeResult
	err := ctx.Invoke("oci:cloudguard/getResponderRecipe:getResponderRecipe", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResponderRecipe.
type LookupResponderRecipeArgs struct {
	// OCID of ResponderRecipe
	ResponderRecipeId string `pulumi:"responderRecipeId"`
}

// A collection of values returned by getResponderRecipe.
type LookupResponderRecipeResult struct {
	// Compartment Identifier
	CompartmentId string `pulumi:"compartmentId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// ResponderRule Description
	Description string `pulumi:"description"`
	// ResponderRule Display Name
	DisplayName string `pulumi:"displayName"`
	// List of responder rules associated with the recipe
	EffectiveResponderRules []GetResponderRecipeEffectiveResponderRule `pulumi:"effectiveResponderRules"`
	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Identifier for ResponderRecipe.
	Id string `pulumi:"id"`
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails string `pulumi:"lifecycleDetails"`
	// Owner of ResponderRecipe
	Owner             string `pulumi:"owner"`
	ResponderRecipeId string `pulumi:"responderRecipeId"`
	// List of responder rules associated with the recipe
	ResponderRules []GetResponderRecipeResponderRule `pulumi:"responderRules"`
	// The id of the source responder recipe.
	SourceResponderRecipeId string `pulumi:"sourceResponderRecipeId"`
	// The current state of the Example.
	State string `pulumi:"state"`
	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// The date and time the responder recipe was created. Format defined by RFC3339.
	TimeCreated string `pulumi:"timeCreated"`
	// The date and time the responder recipe was updated. Format defined by RFC3339.
	TimeUpdated string `pulumi:"timeUpdated"`
}
