// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

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
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/database"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := database.GetDbSystemPatchHistoryEntries(ctx, &database.GetDbSystemPatchHistoryEntriesArgs{
// 			DbSystemId: oci_database_db_system.Test_db_system.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDbSystemPatchHistoryEntries(ctx *pulumi.Context, args *GetDbSystemPatchHistoryEntriesArgs, opts ...pulumi.InvokeOption) (*GetDbSystemPatchHistoryEntriesResult, error) {
	var rv GetDbSystemPatchHistoryEntriesResult
	err := ctx.Invoke("oci:database/getDbSystemPatchHistoryEntries:getDbSystemPatchHistoryEntries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDbSystemPatchHistoryEntries.
type GetDbSystemPatchHistoryEntriesArgs struct {
	// The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbSystemId string                                 `pulumi:"dbSystemId"`
	Filters    []GetDbSystemPatchHistoryEntriesFilter `pulumi:"filters"`
}

// A collection of values returned by getDbSystemPatchHistoryEntries.
type GetDbSystemPatchHistoryEntriesResult struct {
	DbSystemId string                                 `pulumi:"dbSystemId"`
	Filters    []GetDbSystemPatchHistoryEntriesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of patch_history_entries.
	PatchHistoryEntries []GetDbSystemPatchHistoryEntriesPatchHistoryEntry `pulumi:"patchHistoryEntries"`
}
