// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ons

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Notification Topic resource in Oracle Cloud Infrastructure Notifications service.
//
// Creates a topic in the specified compartment. For general information about topics, see
// [Managing Topics and Subscriptions](https://docs.cloud.oracle.com/iaas/Content/Notification/Tasks/managingtopicsandsubscriptions.htm).
//
// For the purposes of access control, you must provide the OCID of the compartment where you want the topic to reside.
// For information about access control and compartments, see [Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
//
// You must specify a display name for the topic.
//
// All Oracle Cloud Infrastructure resources, including topics, get an Oracle-assigned, unique ID called an
// Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID in the response. You can also
// retrieve a resource's OCID by using a List API operation on that resource type, or by viewing the resource in the
// Console. For more information, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
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
// 		_, err := ons.NewNotificationTopic(ctx, "testNotificationTopic", &ons.NotificationTopicArgs{
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			Description: pulumi.Any(_var.Notification_topic_description),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// NotificationTopics can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:ons/notificationTopic:NotificationTopic test_notification_topic "id"
// ```
type NotificationTopic struct {
	pulumi.CustomResourceState

	// The endpoint for managing subscriptions or publishing messages to the topic.
	ApiEndpoint pulumi.StringOutput `pulumi:"apiEndpoint"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the topic in.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) The description of the topic being created. Avoid entering confidential information.
	Description pulumi.StringOutput `pulumi:"description"`
	// For optimistic concurrency control. See `if-match`.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// The name of the topic being created. The topic name must be unique across the tenancy. Avoid entering confidential information.
	Name pulumi.StringOutput `pulumi:"name"`
	// A unique short topic Id. This is used only for SMS subscriptions.
	ShortTopicId pulumi.StringOutput `pulumi:"shortTopicId"`
	// The lifecycle state of the topic.
	State pulumi.StringOutput `pulumi:"state"`
	// The time the topic was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic.
	TopicId pulumi.StringOutput `pulumi:"topicId"`
}

// NewNotificationTopic registers a new resource with the given unique name, arguments, and options.
func NewNotificationTopic(ctx *pulumi.Context,
	name string, args *NotificationTopicArgs, opts ...pulumi.ResourceOption) (*NotificationTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	var resource NotificationTopic
	err := ctx.RegisterResource("oci:ons/notificationTopic:NotificationTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationTopic gets an existing NotificationTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationTopicState, opts ...pulumi.ResourceOption) (*NotificationTopic, error) {
	var resource NotificationTopic
	err := ctx.ReadResource("oci:ons/notificationTopic:NotificationTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationTopic resources.
type notificationTopicState struct {
	// The endpoint for managing subscriptions or publishing messages to the topic.
	ApiEndpoint *string `pulumi:"apiEndpoint"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the topic in.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) The description of the topic being created. Avoid entering confidential information.
	Description *string `pulumi:"description"`
	// For optimistic concurrency control. See `if-match`.
	Etag *string `pulumi:"etag"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The name of the topic being created. The topic name must be unique across the tenancy. Avoid entering confidential information.
	Name *string `pulumi:"name"`
	// A unique short topic Id. This is used only for SMS subscriptions.
	ShortTopicId *string `pulumi:"shortTopicId"`
	// The lifecycle state of the topic.
	State *string `pulumi:"state"`
	// The time the topic was created.
	TimeCreated *string `pulumi:"timeCreated"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic.
	TopicId *string `pulumi:"topicId"`
}

type NotificationTopicState struct {
	// The endpoint for managing subscriptions or publishing messages to the topic.
	ApiEndpoint pulumi.StringPtrInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the topic in.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) The description of the topic being created. Avoid entering confidential information.
	Description pulumi.StringPtrInput
	// For optimistic concurrency control. See `if-match`.
	Etag pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// The name of the topic being created. The topic name must be unique across the tenancy. Avoid entering confidential information.
	Name pulumi.StringPtrInput
	// A unique short topic Id. This is used only for SMS subscriptions.
	ShortTopicId pulumi.StringPtrInput
	// The lifecycle state of the topic.
	State pulumi.StringPtrInput
	// The time the topic was created.
	TimeCreated pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic.
	TopicId pulumi.StringPtrInput
}

func (NotificationTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationTopicState)(nil)).Elem()
}

type notificationTopicArgs struct {
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the topic in.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) The description of the topic being created. Avoid entering confidential information.
	Description *string `pulumi:"description"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The name of the topic being created. The topic name must be unique across the tenancy. Avoid entering confidential information.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a NotificationTopic resource.
type NotificationTopicArgs struct {
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the topic in.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) The description of the topic being created. Avoid entering confidential information.
	Description pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// The name of the topic being created. The topic name must be unique across the tenancy. Avoid entering confidential information.
	Name pulumi.StringPtrInput
}

func (NotificationTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationTopicArgs)(nil)).Elem()
}

type NotificationTopicInput interface {
	pulumi.Input

	ToNotificationTopicOutput() NotificationTopicOutput
	ToNotificationTopicOutputWithContext(ctx context.Context) NotificationTopicOutput
}

func (*NotificationTopic) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationTopic)(nil))
}

func (i *NotificationTopic) ToNotificationTopicOutput() NotificationTopicOutput {
	return i.ToNotificationTopicOutputWithContext(context.Background())
}

func (i *NotificationTopic) ToNotificationTopicOutputWithContext(ctx context.Context) NotificationTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationTopicOutput)
}

func (i *NotificationTopic) ToNotificationTopicPtrOutput() NotificationTopicPtrOutput {
	return i.ToNotificationTopicPtrOutputWithContext(context.Background())
}

func (i *NotificationTopic) ToNotificationTopicPtrOutputWithContext(ctx context.Context) NotificationTopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationTopicPtrOutput)
}

type NotificationTopicPtrInput interface {
	pulumi.Input

	ToNotificationTopicPtrOutput() NotificationTopicPtrOutput
	ToNotificationTopicPtrOutputWithContext(ctx context.Context) NotificationTopicPtrOutput
}

type notificationTopicPtrType NotificationTopicArgs

func (*notificationTopicPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationTopic)(nil))
}

func (i *notificationTopicPtrType) ToNotificationTopicPtrOutput() NotificationTopicPtrOutput {
	return i.ToNotificationTopicPtrOutputWithContext(context.Background())
}

func (i *notificationTopicPtrType) ToNotificationTopicPtrOutputWithContext(ctx context.Context) NotificationTopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationTopicPtrOutput)
}

// NotificationTopicArrayInput is an input type that accepts NotificationTopicArray and NotificationTopicArrayOutput values.
// You can construct a concrete instance of `NotificationTopicArrayInput` via:
//
//          NotificationTopicArray{ NotificationTopicArgs{...} }
type NotificationTopicArrayInput interface {
	pulumi.Input

	ToNotificationTopicArrayOutput() NotificationTopicArrayOutput
	ToNotificationTopicArrayOutputWithContext(context.Context) NotificationTopicArrayOutput
}

type NotificationTopicArray []NotificationTopicInput

func (NotificationTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationTopic)(nil)).Elem()
}

func (i NotificationTopicArray) ToNotificationTopicArrayOutput() NotificationTopicArrayOutput {
	return i.ToNotificationTopicArrayOutputWithContext(context.Background())
}

func (i NotificationTopicArray) ToNotificationTopicArrayOutputWithContext(ctx context.Context) NotificationTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationTopicArrayOutput)
}

// NotificationTopicMapInput is an input type that accepts NotificationTopicMap and NotificationTopicMapOutput values.
// You can construct a concrete instance of `NotificationTopicMapInput` via:
//
//          NotificationTopicMap{ "key": NotificationTopicArgs{...} }
type NotificationTopicMapInput interface {
	pulumi.Input

	ToNotificationTopicMapOutput() NotificationTopicMapOutput
	ToNotificationTopicMapOutputWithContext(context.Context) NotificationTopicMapOutput
}

type NotificationTopicMap map[string]NotificationTopicInput

func (NotificationTopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationTopic)(nil)).Elem()
}

func (i NotificationTopicMap) ToNotificationTopicMapOutput() NotificationTopicMapOutput {
	return i.ToNotificationTopicMapOutputWithContext(context.Background())
}

func (i NotificationTopicMap) ToNotificationTopicMapOutputWithContext(ctx context.Context) NotificationTopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationTopicMapOutput)
}

type NotificationTopicOutput struct {
	*pulumi.OutputState
}

func (NotificationTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationTopic)(nil))
}

func (o NotificationTopicOutput) ToNotificationTopicOutput() NotificationTopicOutput {
	return o
}

func (o NotificationTopicOutput) ToNotificationTopicOutputWithContext(ctx context.Context) NotificationTopicOutput {
	return o
}

func (o NotificationTopicOutput) ToNotificationTopicPtrOutput() NotificationTopicPtrOutput {
	return o.ToNotificationTopicPtrOutputWithContext(context.Background())
}

func (o NotificationTopicOutput) ToNotificationTopicPtrOutputWithContext(ctx context.Context) NotificationTopicPtrOutput {
	return o.ApplyT(func(v NotificationTopic) *NotificationTopic {
		return &v
	}).(NotificationTopicPtrOutput)
}

type NotificationTopicPtrOutput struct {
	*pulumi.OutputState
}

func (NotificationTopicPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationTopic)(nil))
}

func (o NotificationTopicPtrOutput) ToNotificationTopicPtrOutput() NotificationTopicPtrOutput {
	return o
}

func (o NotificationTopicPtrOutput) ToNotificationTopicPtrOutputWithContext(ctx context.Context) NotificationTopicPtrOutput {
	return o
}

type NotificationTopicArrayOutput struct{ *pulumi.OutputState }

func (NotificationTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationTopic)(nil))
}

func (o NotificationTopicArrayOutput) ToNotificationTopicArrayOutput() NotificationTopicArrayOutput {
	return o
}

func (o NotificationTopicArrayOutput) ToNotificationTopicArrayOutputWithContext(ctx context.Context) NotificationTopicArrayOutput {
	return o
}

func (o NotificationTopicArrayOutput) Index(i pulumi.IntInput) NotificationTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationTopic {
		return vs[0].([]NotificationTopic)[vs[1].(int)]
	}).(NotificationTopicOutput)
}

type NotificationTopicMapOutput struct{ *pulumi.OutputState }

func (NotificationTopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NotificationTopic)(nil))
}

func (o NotificationTopicMapOutput) ToNotificationTopicMapOutput() NotificationTopicMapOutput {
	return o
}

func (o NotificationTopicMapOutput) ToNotificationTopicMapOutputWithContext(ctx context.Context) NotificationTopicMapOutput {
	return o
}

func (o NotificationTopicMapOutput) MapIndex(k pulumi.StringInput) NotificationTopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NotificationTopic {
		return vs[0].(map[string]NotificationTopic)[vs[1].(string)]
	}).(NotificationTopicOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationTopicOutput{})
	pulumi.RegisterOutputType(NotificationTopicPtrOutput{})
	pulumi.RegisterOutputType(NotificationTopicArrayOutput{})
	pulumi.RegisterOutputType(NotificationTopicMapOutput{})
}
