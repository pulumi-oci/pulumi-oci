// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waas

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Http Redirects in Oracle Cloud Infrastructure Web Application Acceleration and Security service.
//
// Gets a list of HTTP Redirects.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/waas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Http_redirect_time_created_greater_than_or_equal_to
// 		opt1 := _var.Http_redirect_time_created_less_than
// 		_, err := waas.GetHttpRedirects(ctx, &waas.GetHttpRedirectsArgs{
// 			CompartmentId:                   _var.Compartment_id,
// 			DisplayNames:                    _var.Http_redirect_display_names,
// 			Ids:                             _var.Http_redirect_ids,
// 			States:                          _var.Http_redirect_states,
// 			TimeCreatedGreaterThanOrEqualTo: &opt0,
// 			TimeCreatedLessThan:             &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetHttpRedirects(ctx *pulumi.Context, args *GetHttpRedirectsArgs, opts ...pulumi.InvokeOption) (*GetHttpRedirectsResult, error) {
	var rv GetHttpRedirectsResult
	err := ctx.Invoke("oci:waas/getHttpRedirects:getHttpRedirects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHttpRedirects.
type GetHttpRedirectsArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
	CompartmentId string `pulumi:"compartmentId"`
	// Filter redirects using a display name.
	DisplayNames []string                 `pulumi:"displayNames"`
	Filters      []GetHttpRedirectsFilter `pulumi:"filters"`
	// Filter redirects using a list of redirect OCIDs.
	Ids []string `pulumi:"ids"`
	// Filter redirects using a list of lifecycle states.
	States []string `pulumi:"states"`
	// A filter that matches redirects created on or after the specified date and time.
	TimeCreatedGreaterThanOrEqualTo *string `pulumi:"timeCreatedGreaterThanOrEqualTo"`
	// A filter that matches redirects created before the specified date-time. Default to 1 day before now.
	TimeCreatedLessThan *string `pulumi:"timeCreatedLessThan"`
}

// A collection of values returned by getHttpRedirects.
type GetHttpRedirectsResult struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the HTTP Redirect's compartment.
	CompartmentId string                   `pulumi:"compartmentId"`
	DisplayNames  []string                 `pulumi:"displayNames"`
	Filters       []GetHttpRedirectsFilter `pulumi:"filters"`
	// The list of http_redirects.
	HttpRedirects []GetHttpRedirectsHttpRedirect `pulumi:"httpRedirects"`
	// The provider-assigned unique ID for this managed resource.
	Id                              string   `pulumi:"id"`
	Ids                             []string `pulumi:"ids"`
	States                          []string `pulumi:"states"`
	TimeCreatedGreaterThanOrEqualTo *string  `pulumi:"timeCreatedGreaterThanOrEqualTo"`
	TimeCreatedLessThan             *string  `pulumi:"timeCreatedLessThan"`
}
