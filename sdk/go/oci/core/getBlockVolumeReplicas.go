// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Block Volume Replicas in Oracle Cloud Infrastructure Core service.
//
// Lists the block volume replicas in the specified compartment and availability domain.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Block_volume_replica_display_name
// 		opt1 := _var.Block_volume_replica_state
// 		_, err := core.GetBlockVolumeReplicas(ctx, &core.GetBlockVolumeReplicasArgs{
// 			AvailabilityDomain: _var.Block_volume_replica_availability_domain,
// 			CompartmentId:      _var.Compartment_id,
// 			DisplayName:        &opt0,
// 			State:              &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetBlockVolumeReplicas(ctx *pulumi.Context, args *GetBlockVolumeReplicasArgs, opts ...pulumi.InvokeOption) (*GetBlockVolumeReplicasResult, error) {
	var rv GetBlockVolumeReplicasResult
	err := ctx.Invoke("oci:core/getBlockVolumeReplicas:getBlockVolumeReplicas", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBlockVolumeReplicas.
type GetBlockVolumeReplicasArgs struct {
	// The name of the availability domain.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain string `pulumi:"availabilityDomain"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// A filter to return only resources that match the given display name exactly.
	DisplayName *string                        `pulumi:"displayName"`
	Filters     []GetBlockVolumeReplicasFilter `pulumi:"filters"`
	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	State *string `pulumi:"state"`
}

// A collection of values returned by getBlockVolumeReplicas.
type GetBlockVolumeReplicasResult struct {
	// The availability domain of the block volume replica.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain string `pulumi:"availabilityDomain"`
	// The list of block_volume_replicas.
	BlockVolumeReplicas []GetBlockVolumeReplicasBlockVolumeReplica `pulumi:"blockVolumeReplicas"`
	// The OCID of the compartment that contains the block volume replica.
	CompartmentId string `pulumi:"compartmentId"`
	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string                        `pulumi:"displayName"`
	Filters     []GetBlockVolumeReplicasFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current state of a block volume replica.
	State *string `pulumi:"state"`
}
