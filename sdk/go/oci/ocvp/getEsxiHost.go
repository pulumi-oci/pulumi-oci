// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ocvp

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Esxi Host resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.
//
// Gets the specified ESXi host's information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/ocvp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ocvp.LookupEsxiHost(ctx, &ocvp.LookupEsxiHostArgs{
// 			EsxiHostId: oci_ocvp_esxi_host.Test_esxi_host.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupEsxiHost(ctx *pulumi.Context, args *LookupEsxiHostArgs, opts ...pulumi.InvokeOption) (*LookupEsxiHostResult, error) {
	var rv LookupEsxiHostResult
	err := ctx.Invoke("oci:ocvp/getEsxiHost:getEsxiHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEsxiHost.
type LookupEsxiHostArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host.
	EsxiHostId string `pulumi:"esxiHostId"`
}

// A collection of values returned by getEsxiHost.
type LookupEsxiHostResult struct {
	// Current billing cycle end date. If nextSku is different from existing SKU, then we switch to newSKu after this contractEndDate Example: `2016-08-25T21:10:29.600Z`
	BillingContractEndDate string `pulumi:"billingContractEndDate"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the SDDC.
	CompartmentId string `pulumi:"compartmentId"`
	// In terms of implementation, an ESXi host is a Compute instance that is configured with the chosen bundle of VMware software. The `computeInstanceId` is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of that Compute instance.
	ComputeInstanceId string `pulumi:"computeInstanceId"`
	// Billing option selected during SDDC creation. Oracle Cloud Infrastructure VMware Solution supports the following billing interval SKUs: HOUR, MONTH, ONE_YEAR, and THREE_YEARS. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus).
	CurrentSku string `pulumi:"currentSku"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// A descriptive name for the ESXi host. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName string `pulumi:"displayName"`
	EsxiHostId  string `pulumi:"esxiHostId"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ESXi host.
	Id string `pulumi:"id"`
	// Billing option to switch to once existing billing cycle ends. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus).
	NextSku string `pulumi:"nextSku"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC that the ESXi host belongs to.
	SddcId string `pulumi:"sddcId"`
	// The current state of the ESXi host.
	State string `pulumi:"state"`
	// The date and time the ESXi host was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated string `pulumi:"timeCreated"`
	// The date and time the ESXi host was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	TimeUpdated string `pulumi:"timeUpdated"`
}
