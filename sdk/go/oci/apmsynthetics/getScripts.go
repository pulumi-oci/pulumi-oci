// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apmsynthetics

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Scripts in Oracle Cloud Infrastructure Apm Synthetics service.
//
// Returns a list of scripts.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/apmsynthetics"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Script_content_type
// 		opt1 := _var.Script_display_name
// 		_, err := apmsynthetics.GetScripts(ctx, &apmsynthetics.GetScriptsArgs{
// 			ApmDomainId: oci_apm_synthetics_apm_domain.Test_apm_domain.Id,
// 			ContentType: &opt0,
// 			DisplayName: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetScripts(ctx *pulumi.Context, args *GetScriptsArgs, opts ...pulumi.InvokeOption) (*GetScriptsResult, error) {
	var rv GetScriptsResult
	err := ctx.Invoke("oci:apmsynthetics/getScripts:getScripts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScripts.
type GetScriptsArgs struct {
	// The APM domain ID the request is intended for.
	ApmDomainId string `pulumi:"apmDomainId"`
	// A filter to return only resources that match the content type given.
	ContentType *string `pulumi:"contentType"`
	// A filter to return only resources that match the entire display name given.
	DisplayName *string            `pulumi:"displayName"`
	Filters     []GetScriptsFilter `pulumi:"filters"`
}

// A collection of values returned by getScripts.
type GetScriptsResult struct {
	ApmDomainId string `pulumi:"apmDomainId"`
	// Content type of the script.
	ContentType *string `pulumi:"contentType"`
	// Unique name that can be edited. The name should not contain any confidential information.
	DisplayName *string            `pulumi:"displayName"`
	Filters     []GetScriptsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of script_collection.
	ScriptCollections []GetScriptsScriptCollection `pulumi:"scriptCollections"`
}
