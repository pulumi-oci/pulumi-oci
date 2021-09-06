// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Nat Gateway resource in Oracle Cloud Infrastructure Core service.
//
// Gets the specified NAT gateway's information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.LookupNatGateway(ctx, &core.LookupNatGatewayArgs{
// 			NatGatewayId: oci_core_nat_gateway.Test_nat_gateway.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupNatGateway(ctx *pulumi.Context, args *LookupNatGatewayArgs, opts ...pulumi.InvokeOption) (*LookupNatGatewayResult, error) {
	var rv LookupNatGatewayResult
	err := ctx.Invoke("oci:core/getNatGateway:getNatGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNatGateway.
type LookupNatGatewayArgs struct {
	// The NAT gateway's [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	NatGatewayId string `pulumi:"natGatewayId"`
}

// A collection of values returned by getNatGateway.
type LookupNatGatewayResult struct {
	// Whether the NAT gateway blocks traffic through it. The default is `false`.  Example: `true`
	BlockTraffic bool `pulumi:"blockTraffic"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the NAT gateway.
	CompartmentId string `pulumi:"compartmentId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName string `pulumi:"displayName"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NAT gateway.
	Id           string `pulumi:"id"`
	NatGatewayId string `pulumi:"natGatewayId"`
	// The IP address associated with the NAT gateway.
	NatIp string `pulumi:"natIp"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP address associated with the NAT gateway.
	PublicIpId string `pulumi:"publicIpId"`
	// The NAT gateway's current state.
	State string `pulumi:"state"`
	// The date and time the NAT gateway was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated string `pulumi:"timeCreated"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the NAT gateway belongs to.
	VcnId string `pulumi:"vcnId"`
}