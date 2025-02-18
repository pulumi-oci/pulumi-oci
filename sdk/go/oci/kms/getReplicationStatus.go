// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Replication Status resource in Oracle Cloud Infrastructure Kms service.
//
// When a vault has a replica, each operation on the vault or its resources, such as
// keys, is replicated and has an associated replicationId. Replication status provides
// details about whether the operation associated with the given replicationId has been
// successfully applied across replicas.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/kms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.GetReplicationStatus(ctx, &kms.GetReplicationStatusArgs{
// 			ReplicationId:      oci_kms_replication.Test_replication.Id,
// 			ManagementEndpoint: _var.Replication_status_management_endpoint,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetReplicationStatus(ctx *pulumi.Context, args *GetReplicationStatusArgs, opts ...pulumi.InvokeOption) (*GetReplicationStatusResult, error) {
	var rv GetReplicationStatusResult
	err := ctx.Invoke("oci:kms/getReplicationStatus:getReplicationStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReplicationStatus.
type GetReplicationStatusArgs struct {
	// The service endpoint to perform management operations against. See Vault Management endpoint.
	ManagementEndpoint string `pulumi:"managementEndpoint"`
	// replicationId associated with an operation on a resource
	ReplicationId string `pulumi:"replicationId"`
}

// A collection of values returned by getReplicationStatus.
type GetReplicationStatusResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string                              `pulumi:"id"`
	ManagementEndpoint string                              `pulumi:"managementEndpoint"`
	ReplicaDetails     []GetReplicationStatusReplicaDetail `pulumi:"replicaDetails"`
	ReplicationId      string                              `pulumi:"replicationId"`
}
