// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package email

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Suppression resource in Oracle Cloud Infrastructure Email service.
//
// Adds recipient email addresses to the suppression list for a tenancy.
// Addresses added to the suppression list via the API are denoted as
// "MANUAL" in the `reason` field. *Note:* All email addresses added to the
// suppression list are normalized to include only lowercase letters.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/email"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := email.NewSuppression(ctx, "testSuppression", &email.SuppressionArgs{
// 			CompartmentId: pulumi.Any(_var.Tenancy_ocid),
// 			EmailAddress:  pulumi.Any(_var.Suppression_email_address),
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
// Suppressions can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:email/suppression:Suppression test_suppression "id"
// ```
type Suppression struct {
	pulumi.CustomResourceState

	// The OCID of the compartment to contain the suppression. Since suppressions are at the customer level, this must be the tenancy OCID.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// The recipient email address of the suppression.
	EmailAddress pulumi.StringOutput `pulumi:"emailAddress"`
	// The specific error message returned by a system that resulted in the suppression. This message is usually an SMTP error code with additional descriptive text. Not provided for all types of suppressions.
	ErrorDetail pulumi.StringOutput `pulumi:"errorDetail"`
	// DNS name of the source of the error that caused the suppression. Will be set to either the remote-mta or reporting-mta field from a delivery status notification (RFC 3464) when available. Not provided for all types of suppressions, and not always known.
	ErrorSource pulumi.StringOutput `pulumi:"errorSource"`
	// The value of the Message-ID header from the email that triggered a suppression. This value is as defined in RFC 5322 section 3.6.4, excluding angle-brackets. Not provided for all types of suppressions.
	MessageId pulumi.StringOutput `pulumi:"messageId"`
	// The reason that the email address was suppressed. For more information on the types of bounces, see [Suppression List](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm#components).
	Reason pulumi.StringOutput `pulumi:"reason"`
	// The date and time a recipient's email address was added to the suppression list, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The last date and time the suppression prevented submission in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.
	TimeLastSuppressed pulumi.StringOutput `pulumi:"timeLastSuppressed"`
}

// NewSuppression registers a new resource with the given unique name, arguments, and options.
func NewSuppression(ctx *pulumi.Context,
	name string, args *SuppressionArgs, opts ...pulumi.ResourceOption) (*Suppression, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.EmailAddress == nil {
		return nil, errors.New("invalid value for required argument 'EmailAddress'")
	}
	var resource Suppression
	err := ctx.RegisterResource("oci:email/suppression:Suppression", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSuppression gets an existing Suppression resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSuppression(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SuppressionState, opts ...pulumi.ResourceOption) (*Suppression, error) {
	var resource Suppression
	err := ctx.ReadResource("oci:email/suppression:Suppression", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Suppression resources.
type suppressionState struct {
	// The OCID of the compartment to contain the suppression. Since suppressions are at the customer level, this must be the tenancy OCID.
	CompartmentId *string `pulumi:"compartmentId"`
	// The recipient email address of the suppression.
	EmailAddress *string `pulumi:"emailAddress"`
	// The specific error message returned by a system that resulted in the suppression. This message is usually an SMTP error code with additional descriptive text. Not provided for all types of suppressions.
	ErrorDetail *string `pulumi:"errorDetail"`
	// DNS name of the source of the error that caused the suppression. Will be set to either the remote-mta or reporting-mta field from a delivery status notification (RFC 3464) when available. Not provided for all types of suppressions, and not always known.
	ErrorSource *string `pulumi:"errorSource"`
	// The value of the Message-ID header from the email that triggered a suppression. This value is as defined in RFC 5322 section 3.6.4, excluding angle-brackets. Not provided for all types of suppressions.
	MessageId *string `pulumi:"messageId"`
	// The reason that the email address was suppressed. For more information on the types of bounces, see [Suppression List](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm#components).
	Reason *string `pulumi:"reason"`
	// The date and time a recipient's email address was added to the suppression list, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.
	TimeCreated *string `pulumi:"timeCreated"`
	// The last date and time the suppression prevented submission in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.
	TimeLastSuppressed *string `pulumi:"timeLastSuppressed"`
}

type SuppressionState struct {
	// The OCID of the compartment to contain the suppression. Since suppressions are at the customer level, this must be the tenancy OCID.
	CompartmentId pulumi.StringPtrInput
	// The recipient email address of the suppression.
	EmailAddress pulumi.StringPtrInput
	// The specific error message returned by a system that resulted in the suppression. This message is usually an SMTP error code with additional descriptive text. Not provided for all types of suppressions.
	ErrorDetail pulumi.StringPtrInput
	// DNS name of the source of the error that caused the suppression. Will be set to either the remote-mta or reporting-mta field from a delivery status notification (RFC 3464) when available. Not provided for all types of suppressions, and not always known.
	ErrorSource pulumi.StringPtrInput
	// The value of the Message-ID header from the email that triggered a suppression. This value is as defined in RFC 5322 section 3.6.4, excluding angle-brackets. Not provided for all types of suppressions.
	MessageId pulumi.StringPtrInput
	// The reason that the email address was suppressed. For more information on the types of bounces, see [Suppression List](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm#components).
	Reason pulumi.StringPtrInput
	// The date and time a recipient's email address was added to the suppression list, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.
	TimeCreated pulumi.StringPtrInput
	// The last date and time the suppression prevented submission in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.
	TimeLastSuppressed pulumi.StringPtrInput
}

func (SuppressionState) ElementType() reflect.Type {
	return reflect.TypeOf((*suppressionState)(nil)).Elem()
}

type suppressionArgs struct {
	// The OCID of the compartment to contain the suppression. Since suppressions are at the customer level, this must be the tenancy OCID.
	CompartmentId string `pulumi:"compartmentId"`
	// The recipient email address of the suppression.
	EmailAddress string `pulumi:"emailAddress"`
}

// The set of arguments for constructing a Suppression resource.
type SuppressionArgs struct {
	// The OCID of the compartment to contain the suppression. Since suppressions are at the customer level, this must be the tenancy OCID.
	CompartmentId pulumi.StringInput
	// The recipient email address of the suppression.
	EmailAddress pulumi.StringInput
}

func (SuppressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*suppressionArgs)(nil)).Elem()
}

type SuppressionInput interface {
	pulumi.Input

	ToSuppressionOutput() SuppressionOutput
	ToSuppressionOutputWithContext(ctx context.Context) SuppressionOutput
}

func (*Suppression) ElementType() reflect.Type {
	return reflect.TypeOf((*Suppression)(nil))
}

func (i *Suppression) ToSuppressionOutput() SuppressionOutput {
	return i.ToSuppressionOutputWithContext(context.Background())
}

func (i *Suppression) ToSuppressionOutputWithContext(ctx context.Context) SuppressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionOutput)
}

func (i *Suppression) ToSuppressionPtrOutput() SuppressionPtrOutput {
	return i.ToSuppressionPtrOutputWithContext(context.Background())
}

func (i *Suppression) ToSuppressionPtrOutputWithContext(ctx context.Context) SuppressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionPtrOutput)
}

type SuppressionPtrInput interface {
	pulumi.Input

	ToSuppressionPtrOutput() SuppressionPtrOutput
	ToSuppressionPtrOutputWithContext(ctx context.Context) SuppressionPtrOutput
}

type suppressionPtrType SuppressionArgs

func (*suppressionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Suppression)(nil))
}

func (i *suppressionPtrType) ToSuppressionPtrOutput() SuppressionPtrOutput {
	return i.ToSuppressionPtrOutputWithContext(context.Background())
}

func (i *suppressionPtrType) ToSuppressionPtrOutputWithContext(ctx context.Context) SuppressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionPtrOutput)
}

// SuppressionArrayInput is an input type that accepts SuppressionArray and SuppressionArrayOutput values.
// You can construct a concrete instance of `SuppressionArrayInput` via:
//
//          SuppressionArray{ SuppressionArgs{...} }
type SuppressionArrayInput interface {
	pulumi.Input

	ToSuppressionArrayOutput() SuppressionArrayOutput
	ToSuppressionArrayOutputWithContext(context.Context) SuppressionArrayOutput
}

type SuppressionArray []SuppressionInput

func (SuppressionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Suppression)(nil)).Elem()
}

func (i SuppressionArray) ToSuppressionArrayOutput() SuppressionArrayOutput {
	return i.ToSuppressionArrayOutputWithContext(context.Background())
}

func (i SuppressionArray) ToSuppressionArrayOutputWithContext(ctx context.Context) SuppressionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionArrayOutput)
}

// SuppressionMapInput is an input type that accepts SuppressionMap and SuppressionMapOutput values.
// You can construct a concrete instance of `SuppressionMapInput` via:
//
//          SuppressionMap{ "key": SuppressionArgs{...} }
type SuppressionMapInput interface {
	pulumi.Input

	ToSuppressionMapOutput() SuppressionMapOutput
	ToSuppressionMapOutputWithContext(context.Context) SuppressionMapOutput
}

type SuppressionMap map[string]SuppressionInput

func (SuppressionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Suppression)(nil)).Elem()
}

func (i SuppressionMap) ToSuppressionMapOutput() SuppressionMapOutput {
	return i.ToSuppressionMapOutputWithContext(context.Background())
}

func (i SuppressionMap) ToSuppressionMapOutputWithContext(ctx context.Context) SuppressionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionMapOutput)
}

type SuppressionOutput struct {
	*pulumi.OutputState
}

func (SuppressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Suppression)(nil))
}

func (o SuppressionOutput) ToSuppressionOutput() SuppressionOutput {
	return o
}

func (o SuppressionOutput) ToSuppressionOutputWithContext(ctx context.Context) SuppressionOutput {
	return o
}

func (o SuppressionOutput) ToSuppressionPtrOutput() SuppressionPtrOutput {
	return o.ToSuppressionPtrOutputWithContext(context.Background())
}

func (o SuppressionOutput) ToSuppressionPtrOutputWithContext(ctx context.Context) SuppressionPtrOutput {
	return o.ApplyT(func(v Suppression) *Suppression {
		return &v
	}).(SuppressionPtrOutput)
}

type SuppressionPtrOutput struct {
	*pulumi.OutputState
}

func (SuppressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Suppression)(nil))
}

func (o SuppressionPtrOutput) ToSuppressionPtrOutput() SuppressionPtrOutput {
	return o
}

func (o SuppressionPtrOutput) ToSuppressionPtrOutputWithContext(ctx context.Context) SuppressionPtrOutput {
	return o
}

type SuppressionArrayOutput struct{ *pulumi.OutputState }

func (SuppressionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Suppression)(nil))
}

func (o SuppressionArrayOutput) ToSuppressionArrayOutput() SuppressionArrayOutput {
	return o
}

func (o SuppressionArrayOutput) ToSuppressionArrayOutputWithContext(ctx context.Context) SuppressionArrayOutput {
	return o
}

func (o SuppressionArrayOutput) Index(i pulumi.IntInput) SuppressionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Suppression {
		return vs[0].([]Suppression)[vs[1].(int)]
	}).(SuppressionOutput)
}

type SuppressionMapOutput struct{ *pulumi.OutputState }

func (SuppressionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Suppression)(nil))
}

func (o SuppressionMapOutput) ToSuppressionMapOutput() SuppressionMapOutput {
	return o
}

func (o SuppressionMapOutput) ToSuppressionMapOutputWithContext(ctx context.Context) SuppressionMapOutput {
	return o
}

func (o SuppressionMapOutput) MapIndex(k pulumi.StringInput) SuppressionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Suppression {
		return vs[0].(map[string]Suppression)[vs[1].(string)]
	}).(SuppressionOutput)
}

func init() {
	pulumi.RegisterOutputType(SuppressionOutput{})
	pulumi.RegisterOutputType(SuppressionPtrOutput{})
	pulumi.RegisterOutputType(SuppressionArrayOutput{})
	pulumi.RegisterOutputType(SuppressionMapOutput{})
}
