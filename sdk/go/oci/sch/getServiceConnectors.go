// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sch

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Service Connectors in Oracle Cloud Infrastructure Service Connector Hub service.
//
// Lists service connectors in the specified compartment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/sch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Service_connector_display_name
// 		opt1 := _var.Service_connector_state
// 		_, err := sch.GetServiceConnectors(ctx, &sch.GetServiceConnectorsArgs{
// 			CompartmentId: _var.Compartment_id,
// 			DisplayName:   &opt0,
// 			State:         &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetServiceConnectors(ctx *pulumi.Context, args *GetServiceConnectorsArgs, opts ...pulumi.InvokeOption) (*GetServiceConnectorsResult, error) {
	var rv GetServiceConnectorsResult
	err := ctx.Invoke("oci:sch/getServiceConnectors:getServiceConnectors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceConnectors.
type GetServiceConnectorsArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for this request.
	CompartmentId string `pulumi:"compartmentId"`
	// A filter to return only resources that match the given display name exactly.  Example: `exampleServiceConnector`
	DisplayName *string                      `pulumi:"displayName"`
	Filters     []GetServiceConnectorsFilter `pulumi:"filters"`
	// A filter to return only resources that match the given lifecycle state.  Example: `ACTIVE`
	State *string `pulumi:"state"`
}

// A collection of values returned by getServiceConnectors.
type GetServiceConnectorsResult struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric.
	CompartmentId string `pulumi:"compartmentId"`
	// A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.
	DisplayName *string                      `pulumi:"displayName"`
	Filters     []GetServiceConnectorsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of service_connector_collection.
	ServiceConnectorCollections []GetServiceConnectorsServiceConnectorCollection `pulumi:"serviceConnectorCollections"`
	// The current state of the service connector.
	State *string `pulumi:"state"`
}