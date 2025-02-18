// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataflow

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Private Endpoint resource in Oracle Cloud Infrastructure Data Flow service.
//
// Retrieves an private endpoint using a `privateEndpointId`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/dataflow"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataflow.LookupPrivateEndpoint(ctx, &dataflow.LookupPrivateEndpointArgs{
// 			PrivateEndpointId: oci_dataflow_private_endpoint.Test_private_endpoint.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupPrivateEndpoint(ctx *pulumi.Context, args *LookupPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointResult, error) {
	var rv LookupPrivateEndpointResult
	err := ctx.Invoke("oci:dataflow/getPrivateEndpoint:getPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivateEndpoint.
type LookupPrivateEndpointArgs struct {
	// The unique ID for a private endpoint.
	PrivateEndpointId string `pulumi:"privateEndpointId"`
}

// A collection of values returned by getPrivateEndpoint.
type LookupPrivateEndpointResult struct {
	// The OCID of a compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// A user-friendly description. Avoid entering confidential information.
	Description string `pulumi:"description"`
	// A user-friendly name. It does not have to be unique. Avoid entering confidential information.
	DisplayName string `pulumi:"displayName"`
	// An array of DNS zone names. Example: `[ "app.examplecorp.com", "app.examplecorp2.com" ]`
	DnsZones []string `pulumi:"dnsZones"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The OCID of a private endpoint.
	Id string `pulumi:"id"`
	// The detailed messages about the lifecycle state.
	LifecycleDetails string `pulumi:"lifecycleDetails"`
	// The maximum number of hosts to be accessed through the private endpoint. This value is used to calculate the relevant CIDR block and should be a multiple of 256.  If the value is not a multiple of 256, it is rounded up to the next multiple of 256. For example, 300 is rounded up to 512.
	MaxHostCount int `pulumi:"maxHostCount"`
	// An array of network security group OCIDs.
	NsgIds []string `pulumi:"nsgIds"`
	// The OCID of the user who created the resource.
	OwnerPrincipalId string `pulumi:"ownerPrincipalId"`
	// The username of the user who created the resource.  If the username of the owner does not exist, `null` will be returned and the caller should refer to the ownerPrincipalId value instead.
	OwnerUserName     string `pulumi:"ownerUserName"`
	PrivateEndpointId string `pulumi:"privateEndpointId"`
	// The current state of this private endpoint.
	State string `pulumi:"state"`
	// The OCID of a subnet.
	SubnetId string `pulumi:"subnetId"`
	// The date and time a application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z`
	TimeCreated string `pulumi:"timeCreated"`
	// The date and time a application was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated string `pulumi:"timeUpdated"`
}
