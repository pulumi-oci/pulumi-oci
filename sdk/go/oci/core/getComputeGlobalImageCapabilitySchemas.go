// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Compute Global Image Capability Schemas in Oracle Cloud Infrastructure Core service.
//
// Lists Compute Global Image Capability Schema in the specified compartment.
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
// 		opt0 := _var.Compartment_id
// 		opt1 := _var.Compute_global_image_capability_schema_display_name
// 		_, err := core.GetComputeGlobalImageCapabilitySchemas(ctx, &core.GetComputeGlobalImageCapabilitySchemasArgs{
// 			CompartmentId: &opt0,
// 			DisplayName:   &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetComputeGlobalImageCapabilitySchemas(ctx *pulumi.Context, args *GetComputeGlobalImageCapabilitySchemasArgs, opts ...pulumi.InvokeOption) (*GetComputeGlobalImageCapabilitySchemasResult, error) {
	var rv GetComputeGlobalImageCapabilitySchemasResult
	err := ctx.Invoke("oci:core/getComputeGlobalImageCapabilitySchemas:getComputeGlobalImageCapabilitySchemas", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeGlobalImageCapabilitySchemas.
type GetComputeGlobalImageCapabilitySchemasArgs struct {
	// A filter to return only resources that match the given compartment OCID exactly.
	CompartmentId *string `pulumi:"compartmentId"`
	// A filter to return only resources that match the given display name exactly.
	DisplayName *string                                        `pulumi:"displayName"`
	Filters     []GetComputeGlobalImageCapabilitySchemasFilter `pulumi:"filters"`
}

// A collection of values returned by getComputeGlobalImageCapabilitySchemas.
type GetComputeGlobalImageCapabilitySchemasResult struct {
	// The OCID of the compartment containing the compute global image capability schema
	CompartmentId *string `pulumi:"compartmentId"`
	// The list of compute_global_image_capability_schemas.
	ComputeGlobalImageCapabilitySchemas []GetComputeGlobalImageCapabilitySchemasComputeGlobalImageCapabilitySchema `pulumi:"computeGlobalImageCapabilitySchemas"`
	// A user-friendly name for the compute global image capability schema.
	DisplayName *string                                        `pulumi:"displayName"`
	Filters     []GetComputeGlobalImageCapabilitySchemasFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}