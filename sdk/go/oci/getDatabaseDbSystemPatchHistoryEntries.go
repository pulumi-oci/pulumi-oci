// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Db System Patch History Entries in Oracle Cloud Infrastructure Database service.
//
// Gets the history of the patch actions performed on the specified DB system.
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
// 		_, err := oci.GetDatabaseDbSystemPatchHistoryEntries(ctx, &GetDatabaseDbSystemPatchHistoryEntriesArgs{
// 			DbSystemId: oci_database_db_system.Test_db_system.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDatabaseDbSystemPatchHistoryEntries(ctx *pulumi.Context, args *GetDatabaseDbSystemPatchHistoryEntriesArgs, opts ...pulumi.InvokeOption) (*GetDatabaseDbSystemPatchHistoryEntriesResult, error) {
	var rv GetDatabaseDbSystemPatchHistoryEntriesResult
	err := ctx.Invoke("oci:index/getDatabaseDbSystemPatchHistoryEntries:GetDatabaseDbSystemPatchHistoryEntries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetDatabaseDbSystemPatchHistoryEntries.
type GetDatabaseDbSystemPatchHistoryEntriesArgs struct {
	// The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbSystemId string                                         `pulumi:"dbSystemId"`
	Filters    []GetDatabaseDbSystemPatchHistoryEntriesFilter `pulumi:"filters"`
}

// A collection of values returned by GetDatabaseDbSystemPatchHistoryEntries.
type GetDatabaseDbSystemPatchHistoryEntriesResult struct {
	DbSystemId string                                         `pulumi:"dbSystemId"`
	Filters    []GetDatabaseDbSystemPatchHistoryEntriesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of patch_history_entries.
	PatchHistoryEntries []GetDatabaseDbSystemPatchHistoryEntriesPatchHistoryEntry `pulumi:"patchHistoryEntries"`
}