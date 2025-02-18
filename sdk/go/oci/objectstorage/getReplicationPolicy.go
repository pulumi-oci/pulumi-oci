// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package objectstorage

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Replication Policy resource in Oracle Cloud Infrastructure Object Storage service.
//
// Get the replication policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/objectstorage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := objectstorage.GetReplicationPolicy(ctx, &objectstorage.GetReplicationPolicyArgs{
// 			Bucket:        _var.Replication_policy_bucket,
// 			Namespace:     _var.Replication_policy_namespace,
// 			ReplicationId: oci_objectstorage_replication.Test_replication.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetReplicationPolicy(ctx *pulumi.Context, args *GetReplicationPolicyArgs, opts ...pulumi.InvokeOption) (*GetReplicationPolicyResult, error) {
	var rv GetReplicationPolicyResult
	err := ctx.Invoke("oci:objectstorage/getReplicationPolicy:getReplicationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReplicationPolicy.
type GetReplicationPolicyArgs struct {
	// The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1`
	Bucket string `pulumi:"bucket"`
	// The Object Storage namespace used for the request.
	Namespace string `pulumi:"namespace"`
	// The ID of the replication policy.
	ReplicationId string `pulumi:"replicationId"`
}

// A collection of values returned by getReplicationPolicy.
type GetReplicationPolicyResult struct {
	Bucket string `pulumi:"bucket"`
	// Deprecated: The 'delete_object_in_destination_bucket' field has been deprecated. It is no longer supported.
	DeleteObjectInDestinationBucket string `pulumi:"deleteObjectInDestinationBucket"`
	// The bucket to replicate to in the destination region. Replication policy creation does not automatically create a destination bucket. Create the destination bucket before creating the policy.
	DestinationBucketName string `pulumi:"destinationBucketName"`
	// The destination region to replicate to, for example "us-ashburn-1".
	DestinationRegionName string `pulumi:"destinationRegionName"`
	// The id of the replication policy.
	Id string `pulumi:"id"`
	// The name of the policy.
	Name          string `pulumi:"name"`
	Namespace     string `pulumi:"namespace"`
	ReplicationId string `pulumi:"replicationId"`
	// The replication status of the policy. If the status is CLIENT_ERROR, once the user fixes the issue described in the status message, the status will become ACTIVE.
	Status string `pulumi:"status"`
	// A human-readable description of the status.
	StatusMessage string `pulumi:"statusMessage"`
	// The date when the replication policy was created as per [RFC 3339](https://tools.ietf.org/html/rfc3339).
	TimeCreated string `pulumi:"timeCreated"`
	// Changes made to the source bucket before this time has been replicated.
	TimeLastSync string `pulumi:"timeLastSync"`
}
