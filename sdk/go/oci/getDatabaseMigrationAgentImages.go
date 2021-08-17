// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Agent Images in Oracle Cloud Infrastructure Database Migration service.
//
// Get details of the ODMS Agent Images available to install on-premises.
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
// 		_, err := oci.GetDatabaseMigrationAgentImages(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDatabaseMigrationAgentImages(ctx *pulumi.Context, args *GetDatabaseMigrationAgentImagesArgs, opts ...pulumi.InvokeOption) (*GetDatabaseMigrationAgentImagesResult, error) {
	var rv GetDatabaseMigrationAgentImagesResult
	err := ctx.Invoke("oci:index/getDatabaseMigrationAgentImages:GetDatabaseMigrationAgentImages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetDatabaseMigrationAgentImages.
type GetDatabaseMigrationAgentImagesArgs struct {
	Filters []GetDatabaseMigrationAgentImagesFilter `pulumi:"filters"`
}

// A collection of values returned by GetDatabaseMigrationAgentImages.
type GetDatabaseMigrationAgentImagesResult struct {
	// The list of agent_image_collection.
	AgentImageCollections []GetDatabaseMigrationAgentImagesAgentImageCollection `pulumi:"agentImageCollections"`
	Filters               []GetDatabaseMigrationAgentImagesFilter               `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}