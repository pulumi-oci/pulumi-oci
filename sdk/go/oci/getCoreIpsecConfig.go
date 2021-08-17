// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Ip Sec Connection Device Config resource in Oracle Cloud Infrastructure Core service.
//
// Deprecated. To get tunnel information, instead use:
//
// * [GetIPSecConnectionTunnel](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnel/GetIPSecConnectionTunnel)
// * [GetIPSecConnectionTunnelSharedSecret](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnectionTunnelSharedSecret/GetIPSecConnectionTunnelSharedSecret)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oci.GetCoreIpsecConfig(ctx, &GetCoreIpsecConfigArgs{
// 			IpsecId: oci_core_ipsec.Test_ipsec.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCoreIpsecConfig(ctx *pulumi.Context, args *GetCoreIpsecConfigArgs, opts ...pulumi.InvokeOption) (*GetCoreIpsecConfigResult, error) {
	var rv GetCoreIpsecConfigResult
	err := ctx.Invoke("oci:index/getCoreIpsecConfig:GetCoreIpsecConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetCoreIpsecConfig.
type GetCoreIpsecConfigArgs struct {
	Filters []GetCoreIpsecConfigFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IPSec connection.
	IpsecId string `pulumi:"ipsecId"`
}

// A collection of values returned by GetCoreIpsecConfig.
type GetCoreIpsecConfigResult struct {
	// The OCID of the compartment containing the IPSec connection.
	CompartmentId string                     `pulumi:"compartmentId"`
	Filters       []GetCoreIpsecConfigFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	IpsecId string `pulumi:"ipsecId"`
	// The date and time the IPSec connection was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated string `pulumi:"timeCreated"`
	// Two [TunnelConfig](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/TunnelConfig/) objects.
	Tunnels []GetCoreIpsecConfigTunnel `pulumi:"tunnels"`
}