// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Compute Global Image Capability Schemas Version resource in Oracle Cloud Infrastructure Core service.
//
// Gets the specified Compute Global Image Capability Schema Version
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
// 		_, err := oci.GetCoreComputeGlobalImageCapabilitySchemasVersion(ctx, &GetCoreComputeGlobalImageCapabilitySchemasVersionArgs{
// 			ComputeGlobalImageCapabilitySchemaId:          oci_core_compute_global_image_capability_schema.Test_compute_global_image_capability_schema.Id,
// 			ComputeGlobalImageCapabilitySchemaVersionName: _var.Compute_global_image_capability_schemas_version_compute_global_image_capability_schema_version_name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCoreComputeGlobalImageCapabilitySchemasVersion(ctx *pulumi.Context, args *GetCoreComputeGlobalImageCapabilitySchemasVersionArgs, opts ...pulumi.InvokeOption) (*GetCoreComputeGlobalImageCapabilitySchemasVersionResult, error) {
	var rv GetCoreComputeGlobalImageCapabilitySchemasVersionResult
	err := ctx.Invoke("oci:index/getCoreComputeGlobalImageCapabilitySchemasVersion:GetCoreComputeGlobalImageCapabilitySchemasVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetCoreComputeGlobalImageCapabilitySchemasVersion.
type GetCoreComputeGlobalImageCapabilitySchemasVersionArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute global image capability schema
	ComputeGlobalImageCapabilitySchemaId string `pulumi:"computeGlobalImageCapabilitySchemaId"`
	// The name of the compute global image capability schema version
	ComputeGlobalImageCapabilitySchemaVersionName string `pulumi:"computeGlobalImageCapabilitySchemaVersionName"`
}

// A collection of values returned by GetCoreComputeGlobalImageCapabilitySchemasVersion.
type GetCoreComputeGlobalImageCapabilitySchemasVersionResult struct {
	// The ocid of the compute global image capability schema
	ComputeGlobalImageCapabilitySchemaId          string `pulumi:"computeGlobalImageCapabilitySchemaId"`
	ComputeGlobalImageCapabilitySchemaVersionName string `pulumi:"computeGlobalImageCapabilitySchemaVersionName"`
	// A user-friendly name for the compute global image capability schema
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the compute global image capability schema version
	Name string `pulumi:"name"`
	// The map of each capability name to its ImageCapabilityDescriptor.
	SchemaData map[string]interface{} `pulumi:"schemaData"`
	// The date and time the compute global image capability schema version was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated string `pulumi:"timeCreated"`
}