// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure Audit service.
//
// Get the configuration
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
// 		_, err := oci.GetAuditConfiguration(ctx, &GetAuditConfigurationArgs{
// 			CompartmentId: _var.Tenancy_ocid,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupAuditConfiguration(ctx *pulumi.Context, args *LookupAuditConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupAuditConfigurationResult, error) {
	var rv LookupAuditConfigurationResult
	err := ctx.Invoke("oci:index/getAuditConfiguration:GetAuditConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetAuditConfiguration.
type LookupAuditConfigurationArgs struct {
	// ID of the root compartment (tenancy)
	CompartmentId string `pulumi:"compartmentId"`
}

// A collection of values returned by GetAuditConfiguration.
type LookupAuditConfigurationResult struct {
	CompartmentId string `pulumi:"compartmentId"`
	Id            string `pulumi:"id"`
	// The retention period setting, specified in days. The minimum is 90, the maximum 365.  Example: `90`
	RetentionPeriodDays int `pulumi:"retentionPeriodDays"`
}