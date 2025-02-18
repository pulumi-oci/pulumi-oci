// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Virtual Circuit resource in Oracle Cloud Infrastructure Core service.
//
// Gets the specified virtual circuit's information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.LookupVirtualCircuit(ctx, &core.LookupVirtualCircuitArgs{
// 			VirtualCircuitId: oci_core_virtual_circuit.Test_virtual_circuit.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVirtualCircuit(ctx *pulumi.Context, args *LookupVirtualCircuitArgs, opts ...pulumi.InvokeOption) (*LookupVirtualCircuitResult, error) {
	var rv LookupVirtualCircuitResult
	err := ctx.Invoke("oci:core/getVirtualCircuit:getVirtualCircuit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualCircuit.
type LookupVirtualCircuitArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the virtual circuit.
	VirtualCircuitId string `pulumi:"virtualCircuitId"`
}

// A collection of values returned by getVirtualCircuit.
type LookupVirtualCircuitResult struct {
	// The provisioned data rate of the connection. To get a list of the available bandwidth levels (that is, shapes), see [ListFastConnectProviderServiceVirtualCircuitBandwidthShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/FastConnectProviderService/ListFastConnectProviderVirtualCircuitBandwidthShapes).  Example: `10 Gbps`
	BandwidthShapeName string `pulumi:"bandwidthShapeName"`
	// The state of the Ipv6 BGP session associated with the virtual circuit.
	BgpIpv6sessionState string `pulumi:"bgpIpv6sessionState"`
	// Deprecated. Instead use the information in [FastConnectProviderService](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/FastConnectProviderService/).
	//
	// Deprecated: The 'bgp_management' field has been deprecated. Please use the 'oci_core_fast_connect_provider_service' data source instead.
	BgpManagement string `pulumi:"bgpManagement"`
	// The state of the Ipv4 BGP session associated with the virtual circuit.
	BgpSessionState string `pulumi:"bgpSessionState"`
	// The OCID of the compartment containing the virtual circuit.
	CompartmentId string `pulumi:"compartmentId"`
	// An array of mappings, each containing properties for a cross-connect or cross-connect group that is associated with this virtual circuit.
	CrossConnectMappings []GetVirtualCircuitCrossConnectMapping `pulumi:"crossConnectMappings"`
	// The BGP ASN of the network at the other end of the BGP session from Oracle. If the session is between the customer's edge router and Oracle, the value is the customer's ASN. If the BGP session is between the provider's edge router and Oracle, the value is the provider's ASN. Can be a 2-byte or 4-byte ASN. Uses "asplain" format.
	CustomerAsn string `pulumi:"customerAsn"`
	// Deprecated. Instead use `customerAsn`. If you specify values for both, the request will be rejected.
	//
	// Deprecated: The 'customer_bgp_asn' field has been deprecated. Please use 'customer_asn' instead.
	CustomerBgpAsn int `pulumi:"customerBgpAsn"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName string `pulumi:"displayName"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The OCID of the customer's [dynamic routing gateway (DRG)](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Drg) that this virtual circuit uses. Applicable only to private virtual circuits.
	GatewayId string `pulumi:"gatewayId"`
	// The virtual circuit's Oracle ID (OCID).
	Id string `pulumi:"id"`
	// The Oracle BGP ASN.
	OracleBgpAsn int `pulumi:"oracleBgpAsn"`
	// The OCID of the service offered by the provider (if the customer is connecting via a provider).
	ProviderServiceId string `pulumi:"providerServiceId"`
	// The service key name offered by the provider (if the customer is connecting via a provider).
	ProviderServiceKeyName string `pulumi:"providerServiceKeyName"`
	// The provider's state in relation to this virtual circuit (if the customer is connecting via a provider). ACTIVE means the provider has provisioned the virtual circuit from their end. INACTIVE means the provider has not yet provisioned the virtual circuit, or has de-provisioned it.
	ProviderState string `pulumi:"providerState"`
	// For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to advertise across the connection. All prefix sizes are allowed.
	PublicPrefixes []GetVirtualCircuitPublicPrefix `pulumi:"publicPrefixes"`
	// Provider-supplied reference information about this virtual circuit (if the customer is connecting via a provider).
	ReferenceComment string `pulumi:"referenceComment"`
	// The Oracle Cloud Infrastructure region where this virtual circuit is located.
	Region string `pulumi:"region"`
	// The routing policy sets how routing information about the Oracle cloud is shared over a public virtual circuit. Policies available are: `ORACLE_SERVICE_NETWORK`, `REGIONAL`, `MARKET_LEVEL`, and `GLOBAL`. See [Route Filtering](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/routingonprem.htm#route_filtering) for details. By default, routing information is shared for all routes in the same market.
	RoutingPolicies []string `pulumi:"routingPolicies"`
	// Provider service type.
	ServiceType string `pulumi:"serviceType"`
	// The virtual circuit's current state. For information about the different states, see [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).
	State string `pulumi:"state"`
	// The date and time the virtual circuit was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated string `pulumi:"timeCreated"`
	// Whether the virtual circuit supports private or public peering. For more information, see [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm).
	Type             string `pulumi:"type"`
	VirtualCircuitId string `pulumi:"virtualCircuitId"`
}
