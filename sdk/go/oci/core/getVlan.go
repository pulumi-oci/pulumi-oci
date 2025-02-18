// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Vlan resource in Oracle Cloud Infrastructure Core service.
//
// Gets the specified VLAN's information.
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
// 		_, err := core.LookupVlan(ctx, &core.LookupVlanArgs{
// 			VlanId: oci_core_vlan.Test_vlan.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVlan(ctx *pulumi.Context, args *LookupVlanArgs, opts ...pulumi.InvokeOption) (*LookupVlanResult, error) {
	var rv LookupVlanResult
	err := ctx.Invoke("oci:core/getVlan:getVlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVlan.
type LookupVlanArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN.
	VlanId string `pulumi:"vlanId"`
}

// A collection of values returned by getVlan.
type LookupVlanResult struct {
	// The VLAN's availability domain. This attribute will be null if this is a regional VLAN rather than an AD-specific VLAN.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain string `pulumi:"availabilityDomain"`
	// The range of IPv4 addresses that will be used for layer 3 communication with hosts outside the VLAN.  Example: `192.168.1.0/24`
	CidrBlock string `pulumi:"cidrBlock"`
	// The OCID of the compartment containing the VLAN.
	CompartmentId string `pulumi:"compartmentId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName string `pulumi:"displayName"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The VLAN's Oracle ID (OCID).
	Id string `pulumi:"id"`
	// A list of the OCIDs of the network security groups (NSGs) to use with this VLAN. All VNICs in the VLAN belong to these NSGs. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/).
	NsgIds []string `pulumi:"nsgIds"`
	// The OCID of the route table that the VLAN uses.
	RouteTableId string `pulumi:"routeTableId"`
	// The VLAN's current state.
	State string `pulumi:"state"`
	// The date and time the VLAN was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated string `pulumi:"timeCreated"`
	// The OCID of the VCN the VLAN is in.
	VcnId  string `pulumi:"vcnId"`
	VlanId string `pulumi:"vlanId"`
	// The IEEE 802.1Q VLAN tag of this VLAN.  Example: `100`
	VlanTag int `pulumi:"vlanTag"`
}
