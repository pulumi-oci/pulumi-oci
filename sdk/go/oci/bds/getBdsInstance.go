// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bds

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Bds Instance resource in Oracle Cloud Infrastructure Big Data Service service.
//
// Returns information about the Big Data Service cluster identified by the given ID.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/bds"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bds.LookupBdsInstance(ctx, &bds.LookupBdsInstanceArgs{
// 			BdsInstanceId: oci_bds_bds_instance.Test_bds_instance.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupBdsInstance(ctx *pulumi.Context, args *LookupBdsInstanceArgs, opts ...pulumi.InvokeOption) (*LookupBdsInstanceResult, error) {
	var rv LookupBdsInstanceResult
	err := ctx.Invoke("oci:bds/getBdsInstance:getBdsInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBdsInstance.
type LookupBdsInstanceArgs struct {
	// The OCID of the cluster.
	BdsInstanceId string `pulumi:"bdsInstanceId"`
}

// A collection of values returned by getBdsInstance.
type LookupBdsInstanceResult struct {
	BdsInstanceId string `pulumi:"bdsInstanceId"`
	// The information about added Cloud SQL capability
	CloudSqlDetails      GetBdsInstanceCloudSqlDetails `pulumi:"cloudSqlDetails"`
	ClusterAdminPassword string                        `pulumi:"clusterAdminPassword"`
	// Specific info about a Hadoop cluster
	ClusterDetails   GetBdsInstanceClusterDetails `pulumi:"clusterDetails"`
	ClusterPublicKey string                       `pulumi:"clusterPublicKey"`
	// Version of the Hadoop distribution.
	ClusterVersion string `pulumi:"clusterVersion"`
	// The OCID of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// The user who created the cluster.
	CreatedBy string `pulumi:"createdBy"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For example, `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The name of the node.
	DisplayName string `pulumi:"displayName"`
	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. For example, `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The OCID of the Big Data Service resource.
	Id string `pulumi:"id"`
	// Boolean flag specifying whether or not Cloud SQL should be configured.
	IsCloudSqlConfigured bool `pulumi:"isCloudSqlConfigured"`
	// Boolean flag specifying whether or not the cluster is highly available (HA)
	IsHighAvailability bool `pulumi:"isHighAvailability"`
	// Boolean flag specifying whether or not the cluster should be set up as secure.
	IsSecure   bool                     `pulumi:"isSecure"`
	MasterNode GetBdsInstanceMasterNode `pulumi:"masterNode"`
	// Additional configuration of the user's network.
	NetworkConfig GetBdsInstanceNetworkConfig `pulumi:"networkConfig"`
	// The list of nodes in the cluster.
	Nodes []GetBdsInstanceNode `pulumi:"nodes"`
	// The number of nodes that form the cluster.
	NumberOfNodes int `pulumi:"numberOfNodes"`
	// The state of the cluster.
	State string `pulumi:"state"`
	// The time the cluster was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated string `pulumi:"timeCreated"`
	// The time the cluster was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated string                   `pulumi:"timeUpdated"`
	UtilNode    GetBdsInstanceUtilNode   `pulumi:"utilNode"`
	WorkerNode  GetBdsInstanceWorkerNode `pulumi:"workerNode"`
}
