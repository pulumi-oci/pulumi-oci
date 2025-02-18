// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waas

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Address Lists in Oracle Cloud Infrastructure Web Application Acceleration and Security service.
//
// Gets a list of address lists that can be used in a WAAS policy.
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
// 		opt0 := _var.Address_list_time_created_greater_than_or_equal_to
// 		opt1 := _var.Address_list_time_created_less_than
// 		_, err := waas.GetAddressLists(ctx, &waas.GetAddressListsArgs{
// 			CompartmentId:                   _var.Compartment_id,
// 			Ids:                             _var.Address_list_ids,
// 			Names:                           _var.Address_list_names,
// 			States:                          _var.Address_list_states,
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
func GetAddressLists(ctx *pulumi.Context, args *GetAddressListsArgs, opts ...pulumi.InvokeOption) (*GetAddressListsResult, error) {
	var rv GetAddressListsResult
	err := ctx.Invoke("oci:waas/getAddressLists:getAddressLists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddressLists.
type GetAddressListsArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
	CompartmentId string                  `pulumi:"compartmentId"`
	Filters       []GetAddressListsFilter `pulumi:"filters"`
	// Filter address lists using a list of address lists OCIDs.
	Ids []string `pulumi:"ids"`
	// Filter address lists using a list of names.
	Names []string `pulumi:"names"`
	// Filter address lists using a list of lifecycle states.
	States []string `pulumi:"states"`
	// A filter that matches address lists created on or after the specified date-time.
	TimeCreatedGreaterThanOrEqualTo *string `pulumi:"timeCreatedGreaterThanOrEqualTo"`
	// A filter that matches address lists created before the specified date-time.
	TimeCreatedLessThan *string `pulumi:"timeCreatedLessThan"`
}

// A collection of values returned by getAddressLists.
type GetAddressListsResult struct {
	// The list of address_lists.
	AddressLists []GetAddressListsAddressList `pulumi:"addressLists"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the address list's compartment.
	CompartmentId string                  `pulumi:"compartmentId"`
	Filters       []GetAddressListsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                              string   `pulumi:"id"`
	Ids                             []string `pulumi:"ids"`
	Names                           []string `pulumi:"names"`
	States                          []string `pulumi:"states"`
	TimeCreatedGreaterThanOrEqualTo *string  `pulumi:"timeCreatedGreaterThanOrEqualTo"`
	TimeCreatedLessThan             *string  `pulumi:"timeCreatedLessThan"`
}
