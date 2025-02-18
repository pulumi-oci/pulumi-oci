// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ons

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Subscription resource in Oracle Cloud Infrastructure Notifications service.
//
// Gets the specified subscription's configuration information.
//
// Transactions Per Minute (TPM) per-tenancy limit for this operation: 60.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/ons"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ons.LookupSubscription(ctx, &ons.LookupSubscriptionArgs{
// 			SubscriptionId: oci_ons_subscription.Test_subscription.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSubscription(ctx *pulumi.Context, args *LookupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionResult, error) {
	var rv LookupSubscriptionResult
	err := ctx.Invoke("oci:ons/getSubscription:getSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubscription.
type LookupSubscriptionArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription to retrieve.
	SubscriptionId string `pulumi:"subscriptionId"`
}

// A collection of values returned by getSubscription.
type LookupSubscriptionResult struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for the subscription.
	CompartmentId string `pulumi:"compartmentId"`
	// The time when this suscription was created.
	CreatedTime string `pulumi:"createdTime"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The delivery policy of the subscription. Stored as a JSON string.
	DeliveryPolicy string `pulumi:"deliveryPolicy"`
	// A locator that corresponds to the subscription protocol. For example, an email address for a subscription that uses the `EMAIL` protocol, or a URL for a subscription that uses an HTTP-based protocol.
	Endpoint string `pulumi:"endpoint"`
	// For optimistic concurrency control. See `if-match`.
	Etag string `pulumi:"etag"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription.
	Id string `pulumi:"id"`
	// The protocol used for the subscription.
	Protocol string `pulumi:"protocol"`
	// The lifecycle state of the subscription. The status of a new subscription is PENDING; when confirmed, the subscription status changes to ACTIVE.
	State          string `pulumi:"state"`
	SubscriptionId string `pulumi:"subscriptionId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated topic.
	TopicId string `pulumi:"topicId"`
}
