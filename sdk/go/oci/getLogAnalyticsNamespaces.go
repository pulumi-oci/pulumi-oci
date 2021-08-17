// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Namespaces in Oracle Cloud Infrastructure Log Analytics service.
//
// Given a tenancy OCID, this API returns the namespace of the tenancy if it is valid and subscribed to the region.  The
// result also indicates if the tenancy is onboarded with Logging Analytics.
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
// 		_, err := oci.GetLogAnalyticsNamespaces(ctx, &GetLogAnalyticsNamespacesArgs{
// 			CompartmentId: _var.Compartment_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLogAnalyticsNamespaces(ctx *pulumi.Context, args *GetLogAnalyticsNamespacesArgs, opts ...pulumi.InvokeOption) (*GetLogAnalyticsNamespacesResult, error) {
	var rv GetLogAnalyticsNamespacesResult
	err := ctx.Invoke("oci:index/getLogAnalyticsNamespaces:GetLogAnalyticsNamespaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetLogAnalyticsNamespaces.
type GetLogAnalyticsNamespacesArgs struct {
	// The ID of the compartment in which to list resources.
	CompartmentId string                            `pulumi:"compartmentId"`
	Filters       []GetLogAnalyticsNamespacesFilter `pulumi:"filters"`
}

// A collection of values returned by GetLogAnalyticsNamespaces.
type GetLogAnalyticsNamespacesResult struct {
	// The is the tenancy ID
	CompartmentId string                            `pulumi:"compartmentId"`
	Filters       []GetLogAnalyticsNamespacesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of namespace_collection.
	NamespaceCollections []GetLogAnalyticsNamespacesNamespaceCollection `pulumi:"namespaceCollections"`
}