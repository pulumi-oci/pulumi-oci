// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Cloud Guard Configuration resource in Oracle Cloud Infrastructure Cloud Guard service.
//
// GET Cloud Guard Configuration Details for a Tenancy.
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
// 		_, err := oci.GetCloudGuardCloudGuardConfiguration(ctx, &GetCloudGuardCloudGuardConfigurationArgs{
// 			CompartmentId: _var.Compartment_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCloudGuardCloudGuardConfiguration(ctx *pulumi.Context, args *LookupCloudGuardCloudGuardConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupCloudGuardCloudGuardConfigurationResult, error) {
	var rv LookupCloudGuardCloudGuardConfigurationResult
	err := ctx.Invoke("oci:index/getCloudGuardCloudGuardConfiguration:GetCloudGuardCloudGuardConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetCloudGuardCloudGuardConfiguration.
type LookupCloudGuardCloudGuardConfigurationArgs struct {
	// The ID of the compartment in which to list resources.
	CompartmentId string `pulumi:"compartmentId"`
}

// A collection of values returned by GetCloudGuardCloudGuardConfiguration.
type LookupCloudGuardCloudGuardConfigurationResult struct {
	CompartmentId string `pulumi:"compartmentId"`
	Id            string `pulumi:"id"`
	// The reporting region value
	ReportingRegion string `pulumi:"reportingRegion"`
	// Identifies if Oracle managed resources were created by customers
	SelfManageResources bool `pulumi:"selfManageResources"`
	// Status of Cloud Guard Tenant
	Status string `pulumi:"status"`
}