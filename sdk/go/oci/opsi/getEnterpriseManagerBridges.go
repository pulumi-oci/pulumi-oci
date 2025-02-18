// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsi

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Enterprise Manager Bridges in Oracle Cloud Infrastructure Opsi service.
//
// Gets a list of Operations Insights Enterprise Manager bridges. Either compartmentId or id must be specified.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/opsi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Compartment_id
// 		opt1 := _var.Enterprise_manager_bridge_display_name
// 		opt2 := _var.Enterprise_manager_bridge_id
// 		_, err := opsi.GetEnterpriseManagerBridges(ctx, &opsi.GetEnterpriseManagerBridgesArgs{
// 			CompartmentId: &opt0,
// 			DisplayName:   &opt1,
// 			Id:            &opt2,
// 			States:        _var.Enterprise_manager_bridge_state,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetEnterpriseManagerBridges(ctx *pulumi.Context, args *GetEnterpriseManagerBridgesArgs, opts ...pulumi.InvokeOption) (*GetEnterpriseManagerBridgesResult, error) {
	var rv GetEnterpriseManagerBridgesResult
	err := ctx.Invoke("oci:opsi/getEnterpriseManagerBridges:getEnterpriseManagerBridges", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnterpriseManagerBridges.
type GetEnterpriseManagerBridgesArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// A filter to return only resources that match the entire display name.
	DisplayName *string                             `pulumi:"displayName"`
	Filters     []GetEnterpriseManagerBridgesFilter `pulumi:"filters"`
	// Unique Enterprise Manager bridge identifier
	Id *string `pulumi:"id"`
	// Lifecycle states
	States []string `pulumi:"states"`
}

// A collection of values returned by getEnterpriseManagerBridges.
type GetEnterpriseManagerBridgesResult struct {
	// Compartment identifier of the Enterprise Manager bridge
	CompartmentId *string `pulumi:"compartmentId"`
	// User-friedly name of Enterprise Manager Bridge that does not have to be unique.
	DisplayName *string `pulumi:"displayName"`
	// The list of enterprise_manager_bridge_collection.
	EnterpriseManagerBridgeCollections []GetEnterpriseManagerBridgesEnterpriseManagerBridgeCollection `pulumi:"enterpriseManagerBridgeCollections"`
	Filters                            []GetEnterpriseManagerBridgesFilter                            `pulumi:"filters"`
	// Enterprise Manager bridge identifier
	Id *string `pulumi:"id"`
	// The current state of the Enterprise Manager bridge.
	States []string `pulumi:"states"`
}
