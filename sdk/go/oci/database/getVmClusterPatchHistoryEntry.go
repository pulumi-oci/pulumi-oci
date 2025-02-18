// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Vm Cluster Patch History Entry resource in Oracle Cloud Infrastructure Database service.
//
// Gets the patch history details for the specified patch history entry.
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
// 		_, err := database.GetVmClusterPatchHistoryEntry(ctx, &database.GetVmClusterPatchHistoryEntryArgs{
// 			PatchHistoryEntryId: oci_database_patch_history_entry.Test_patch_history_entry.Id,
// 			VmClusterId:         oci_database_vm_cluster.Test_vm_cluster.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetVmClusterPatchHistoryEntry(ctx *pulumi.Context, args *GetVmClusterPatchHistoryEntryArgs, opts ...pulumi.InvokeOption) (*GetVmClusterPatchHistoryEntryResult, error) {
	var rv GetVmClusterPatchHistoryEntryResult
	err := ctx.Invoke("oci:database/getVmClusterPatchHistoryEntry:getVmClusterPatchHistoryEntry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVmClusterPatchHistoryEntry.
type GetVmClusterPatchHistoryEntryArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch history entry.
	PatchHistoryEntryId string `pulumi:"patchHistoryEntryId"`
	// The VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	VmClusterId string `pulumi:"vmClusterId"`
}

// A collection of values returned by getVmClusterPatchHistoryEntry.
type GetVmClusterPatchHistoryEntryResult struct {
	// The action being performed or was completed.
	Action string `pulumi:"action"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A descriptive text associated with the lifecycleState. Typically contains additional displayable text.
	LifecycleDetails    string `pulumi:"lifecycleDetails"`
	PatchHistoryEntryId string `pulumi:"patchHistoryEntryId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch.
	PatchId string `pulumi:"patchId"`
	// The current state of the action.
	State string `pulumi:"state"`
	// The date and time when the patch action completed
	TimeEnded string `pulumi:"timeEnded"`
	// The date and time when the patch action started.
	TimeStarted string `pulumi:"timeStarted"`
	VmClusterId string `pulumi:"vmClusterId"`
}
