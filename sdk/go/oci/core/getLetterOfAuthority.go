// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Letter Of Authority resource in Oracle Cloud Infrastructure Core service.
//
// Gets the Letter of Authority for the specified cross-connect.
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
// 		_, err := core.GetLetterOfAuthority(ctx, &core.GetLetterOfAuthorityArgs{
// 			CrossConnectId: oci_core_cross_connect.Test_cross_connect.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLetterOfAuthority(ctx *pulumi.Context, args *GetLetterOfAuthorityArgs, opts ...pulumi.InvokeOption) (*GetLetterOfAuthorityResult, error) {
	var rv GetLetterOfAuthorityResult
	err := ctx.Invoke("oci:core/getLetterOfAuthority:getLetterOfAuthority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLetterOfAuthority.
type GetLetterOfAuthorityArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect.
	CrossConnectId string `pulumi:"crossConnectId"`
}

// A collection of values returned by getLetterOfAuthority.
type GetLetterOfAuthorityResult struct {
	// The name of the entity authorized by this Letter of Authority.
	AuthorizedEntityName string `pulumi:"authorizedEntityName"`
	// The type of cross-connect fiber, termination, and optical specification.
	CircuitType string `pulumi:"circuitType"`
	// The OCID of the cross-connect.
	CrossConnectId string `pulumi:"crossConnectId"`
	// The address of the FastConnect location.
	FacilityLocation string `pulumi:"facilityLocation"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The meet-me room port for this cross-connect.
	PortName string `pulumi:"portName"`
	// The date and time when the Letter of Authority expires, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	TimeExpires string `pulumi:"timeExpires"`
	// The date and time the Letter of Authority was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeIssued string `pulumi:"timeIssued"`
}
