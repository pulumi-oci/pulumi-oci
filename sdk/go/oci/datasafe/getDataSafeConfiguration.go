// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasafe

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Data Safe Configuration resource in Oracle Cloud Infrastructure Data Safe service.
//
// Gets the details of the Data Safe configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/datasafe"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datasafe.LookupDataSafeConfiguration(ctx, &datasafe.LookupDataSafeConfigurationArgs{
// 			CompartmentId: _var.Compartment_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDataSafeConfiguration(ctx *pulumi.Context, args *LookupDataSafeConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDataSafeConfigurationResult, error) {
	var rv LookupDataSafeConfigurationResult
	err := ctx.Invoke("oci:datasafe/getDataSafeConfiguration:getDataSafeConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataSafeConfiguration.
type LookupDataSafeConfigurationArgs struct {
	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId string `pulumi:"compartmentId"`
}

// A collection of values returned by getDataSafeConfiguration.
type LookupDataSafeConfigurationResult struct {
	// The OCID of the tenancy used to enable Data Safe.
	CompartmentId string `pulumi:"compartmentId"`
	Id            string `pulumi:"id"`
	// Indicates if Data Safe is enabled.
	IsEnabled bool `pulumi:"isEnabled"`
	// The current state of Data Safe.
	State string `pulumi:"state"`
	// The date and time Data Safe was enabled, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	TimeEnabled string `pulumi:"timeEnabled"`
	// The URL of the Data Safe service.
	Url string `pulumi:"url"`
}
