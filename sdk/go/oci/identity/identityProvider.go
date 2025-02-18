// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Identity Provider resource in Oracle Cloud Infrastructure Identity service.
//
// Creates a new identity provider in your tenancy. For more information, see
// [Identity Providers and Federation](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/federation.htm).
//
// You must specify your tenancy's OCID as the compartment ID in the request object.
// Remember that the tenancy is simply the root compartment. For information about
// OCIDs, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
//
// You must also specify a *name* for the `IdentityProvider`, which must be unique
// across all `IdentityProvider` objects in your tenancy and cannot be changed.
//
// You must also specify a *description* for the `IdentityProvider` (although
// it can be an empty string). It does not have to be unique, and you can change
// it anytime with
// [UpdateIdentityProvider](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/IdentityProvider/UpdateIdentityProvider).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/identity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := identity.NewIdentityProvider(ctx, "testIdentityProvider", &identity.IdentityProviderArgs{
// 			CompartmentId: pulumi.Any(_var.Tenancy_ocid),
// 			Description:   pulumi.Any(_var.Identity_provider_description),
// 			Metadata:      pulumi.Any(_var.Identity_provider_metadata),
// 			MetadataUrl:   pulumi.Any(_var.Identity_provider_metadata_url),
// 			ProductType:   pulumi.Any(_var.Identity_provider_product_type),
// 			Protocol:      pulumi.Any(_var.Identity_provider_protocol),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			FreeformAttributes: pulumi.Any(_var.Identity_provider_freeform_attributes),
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
// IdentityProviders can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:identity/identityProvider:IdentityProvider test_identity_provider "id"
// ```
type IdentityProvider struct {
	pulumi.CustomResourceState

	// The OCID of your tenancy.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) The description you assign to the `IdentityProvider` during creation. Does not have to be unique, and it's changeable.
	Description pulumi.StringOutput `pulumi:"description"`
	// (Updatable) Extra name value pairs associated with this identity provider. Example: `{"clientId": "appSf3kdjf3"}`
	FreeformAttributes pulumi.MapOutput `pulumi:"freeformAttributes"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// The detailed status of INACTIVE lifecycleState.
	InactiveState pulumi.StringOutput `pulumi:"inactiveState"`
	// (Updatable) The XML that contains the information required for federating.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// (Updatable) The URL for retrieving the identity provider's metadata, which contains information required for federating.
	MetadataUrl pulumi.StringOutput `pulumi:"metadataUrl"`
	// The name you assign to the `IdentityProvider` during creation. The name must be unique across all `IdentityProvider` objects in the tenancy and cannot be changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The identity provider service or product. Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft Active Directory Federation Services (ADFS).  Example: `IDCS`
	ProductType pulumi.StringOutput `pulumi:"productType"`
	// (Updatable) The protocol used for federation.  Example: `SAML2`
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The URL to redirect federated users to for authentication with the identity provider.
	RedirectUrl pulumi.StringOutput `pulumi:"redirectUrl"`
	// The identity provider's signing certificate used by the IAM Service to validate the SAML2 token.
	SigningCertificate pulumi.StringOutput `pulumi:"signingCertificate"`
	// The current state.
	State pulumi.StringOutput `pulumi:"state"`
	// Date and time the `IdentityProvider` was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
}

// NewIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewIdentityProvider(ctx *pulumi.Context,
	name string, args *IdentityProviderArgs, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Metadata == nil {
		return nil, errors.New("invalid value for required argument 'Metadata'")
	}
	if args.MetadataUrl == nil {
		return nil, errors.New("invalid value for required argument 'MetadataUrl'")
	}
	if args.ProductType == nil {
		return nil, errors.New("invalid value for required argument 'ProductType'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	var resource IdentityProvider
	err := ctx.RegisterResource("oci:identity/identityProvider:IdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProvider gets an existing IdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderState, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	var resource IdentityProvider
	err := ctx.ReadResource("oci:identity/identityProvider:IdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProvider resources.
type identityProviderState struct {
	// The OCID of your tenancy.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) The description you assign to the `IdentityProvider` during creation. Does not have to be unique, and it's changeable.
	Description *string `pulumi:"description"`
	// (Updatable) Extra name value pairs associated with this identity provider. Example: `{"clientId": "appSf3kdjf3"}`
	FreeformAttributes map[string]interface{} `pulumi:"freeformAttributes"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The detailed status of INACTIVE lifecycleState.
	InactiveState *string `pulumi:"inactiveState"`
	// (Updatable) The XML that contains the information required for federating.
	Metadata *string `pulumi:"metadata"`
	// (Updatable) The URL for retrieving the identity provider's metadata, which contains information required for federating.
	MetadataUrl *string `pulumi:"metadataUrl"`
	// The name you assign to the `IdentityProvider` during creation. The name must be unique across all `IdentityProvider` objects in the tenancy and cannot be changed.
	Name *string `pulumi:"name"`
	// The identity provider service or product. Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft Active Directory Federation Services (ADFS).  Example: `IDCS`
	ProductType *string `pulumi:"productType"`
	// (Updatable) The protocol used for federation.  Example: `SAML2`
	Protocol *string `pulumi:"protocol"`
	// The URL to redirect federated users to for authentication with the identity provider.
	RedirectUrl *string `pulumi:"redirectUrl"`
	// The identity provider's signing certificate used by the IAM Service to validate the SAML2 token.
	SigningCertificate *string `pulumi:"signingCertificate"`
	// The current state.
	State *string `pulumi:"state"`
	// Date and time the `IdentityProvider` was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *string `pulumi:"timeCreated"`
}

type IdentityProviderState struct {
	// The OCID of your tenancy.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) The description you assign to the `IdentityProvider` during creation. Does not have to be unique, and it's changeable.
	Description pulumi.StringPtrInput
	// (Updatable) Extra name value pairs associated with this identity provider. Example: `{"clientId": "appSf3kdjf3"}`
	FreeformAttributes pulumi.MapInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// The detailed status of INACTIVE lifecycleState.
	InactiveState pulumi.StringPtrInput
	// (Updatable) The XML that contains the information required for federating.
	Metadata pulumi.StringPtrInput
	// (Updatable) The URL for retrieving the identity provider's metadata, which contains information required for federating.
	MetadataUrl pulumi.StringPtrInput
	// The name you assign to the `IdentityProvider` during creation. The name must be unique across all `IdentityProvider` objects in the tenancy and cannot be changed.
	Name pulumi.StringPtrInput
	// The identity provider service or product. Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft Active Directory Federation Services (ADFS).  Example: `IDCS`
	ProductType pulumi.StringPtrInput
	// (Updatable) The protocol used for federation.  Example: `SAML2`
	Protocol pulumi.StringPtrInput
	// The URL to redirect federated users to for authentication with the identity provider.
	RedirectUrl pulumi.StringPtrInput
	// The identity provider's signing certificate used by the IAM Service to validate the SAML2 token.
	SigningCertificate pulumi.StringPtrInput
	// The current state.
	State pulumi.StringPtrInput
	// Date and time the `IdentityProvider` was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringPtrInput
}

func (IdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderState)(nil)).Elem()
}

type identityProviderArgs struct {
	// The OCID of your tenancy.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) The description you assign to the `IdentityProvider` during creation. Does not have to be unique, and it's changeable.
	Description string `pulumi:"description"`
	// (Updatable) Extra name value pairs associated with this identity provider. Example: `{"clientId": "appSf3kdjf3"}`
	FreeformAttributes map[string]interface{} `pulumi:"freeformAttributes"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The XML that contains the information required for federating.
	Metadata string `pulumi:"metadata"`
	// (Updatable) The URL for retrieving the identity provider's metadata, which contains information required for federating.
	MetadataUrl string `pulumi:"metadataUrl"`
	// The name you assign to the `IdentityProvider` during creation. The name must be unique across all `IdentityProvider` objects in the tenancy and cannot be changed.
	Name *string `pulumi:"name"`
	// The identity provider service or product. Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft Active Directory Federation Services (ADFS).  Example: `IDCS`
	ProductType string `pulumi:"productType"`
	// (Updatable) The protocol used for federation.  Example: `SAML2`
	Protocol string `pulumi:"protocol"`
}

// The set of arguments for constructing a IdentityProvider resource.
type IdentityProviderArgs struct {
	// The OCID of your tenancy.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) The description you assign to the `IdentityProvider` during creation. Does not have to be unique, and it's changeable.
	Description pulumi.StringInput
	// (Updatable) Extra name value pairs associated with this identity provider. Example: `{"clientId": "appSf3kdjf3"}`
	FreeformAttributes pulumi.MapInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The XML that contains the information required for federating.
	Metadata pulumi.StringInput
	// (Updatable) The URL for retrieving the identity provider's metadata, which contains information required for federating.
	MetadataUrl pulumi.StringInput
	// The name you assign to the `IdentityProvider` during creation. The name must be unique across all `IdentityProvider` objects in the tenancy and cannot be changed.
	Name pulumi.StringPtrInput
	// The identity provider service or product. Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft Active Directory Federation Services (ADFS).  Example: `IDCS`
	ProductType pulumi.StringInput
	// (Updatable) The protocol used for federation.  Example: `SAML2`
	Protocol pulumi.StringInput
}

func (IdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderArgs)(nil)).Elem()
}

type IdentityProviderInput interface {
	pulumi.Input

	ToIdentityProviderOutput() IdentityProviderOutput
	ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput
}

func (*IdentityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProvider)(nil))
}

func (i *IdentityProvider) ToIdentityProviderOutput() IdentityProviderOutput {
	return i.ToIdentityProviderOutputWithContext(context.Background())
}

func (i *IdentityProvider) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderOutput)
}

func (i *IdentityProvider) ToIdentityProviderPtrOutput() IdentityProviderPtrOutput {
	return i.ToIdentityProviderPtrOutputWithContext(context.Background())
}

func (i *IdentityProvider) ToIdentityProviderPtrOutputWithContext(ctx context.Context) IdentityProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderPtrOutput)
}

type IdentityProviderPtrInput interface {
	pulumi.Input

	ToIdentityProviderPtrOutput() IdentityProviderPtrOutput
	ToIdentityProviderPtrOutputWithContext(ctx context.Context) IdentityProviderPtrOutput
}

type identityProviderPtrType IdentityProviderArgs

func (*identityProviderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil))
}

func (i *identityProviderPtrType) ToIdentityProviderPtrOutput() IdentityProviderPtrOutput {
	return i.ToIdentityProviderPtrOutputWithContext(context.Background())
}

func (i *identityProviderPtrType) ToIdentityProviderPtrOutputWithContext(ctx context.Context) IdentityProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderPtrOutput)
}

// IdentityProviderArrayInput is an input type that accepts IdentityProviderArray and IdentityProviderArrayOutput values.
// You can construct a concrete instance of `IdentityProviderArrayInput` via:
//
//          IdentityProviderArray{ IdentityProviderArgs{...} }
type IdentityProviderArrayInput interface {
	pulumi.Input

	ToIdentityProviderArrayOutput() IdentityProviderArrayOutput
	ToIdentityProviderArrayOutputWithContext(context.Context) IdentityProviderArrayOutput
}

type IdentityProviderArray []IdentityProviderInput

func (IdentityProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityProvider)(nil)).Elem()
}

func (i IdentityProviderArray) ToIdentityProviderArrayOutput() IdentityProviderArrayOutput {
	return i.ToIdentityProviderArrayOutputWithContext(context.Background())
}

func (i IdentityProviderArray) ToIdentityProviderArrayOutputWithContext(ctx context.Context) IdentityProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderArrayOutput)
}

// IdentityProviderMapInput is an input type that accepts IdentityProviderMap and IdentityProviderMapOutput values.
// You can construct a concrete instance of `IdentityProviderMapInput` via:
//
//          IdentityProviderMap{ "key": IdentityProviderArgs{...} }
type IdentityProviderMapInput interface {
	pulumi.Input

	ToIdentityProviderMapOutput() IdentityProviderMapOutput
	ToIdentityProviderMapOutputWithContext(context.Context) IdentityProviderMapOutput
}

type IdentityProviderMap map[string]IdentityProviderInput

func (IdentityProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityProvider)(nil)).Elem()
}

func (i IdentityProviderMap) ToIdentityProviderMapOutput() IdentityProviderMapOutput {
	return i.ToIdentityProviderMapOutputWithContext(context.Background())
}

func (i IdentityProviderMap) ToIdentityProviderMapOutputWithContext(ctx context.Context) IdentityProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderMapOutput)
}

type IdentityProviderOutput struct {
	*pulumi.OutputState
}

func (IdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProvider)(nil))
}

func (o IdentityProviderOutput) ToIdentityProviderOutput() IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) ToIdentityProviderPtrOutput() IdentityProviderPtrOutput {
	return o.ToIdentityProviderPtrOutputWithContext(context.Background())
}

func (o IdentityProviderOutput) ToIdentityProviderPtrOutputWithContext(ctx context.Context) IdentityProviderPtrOutput {
	return o.ApplyT(func(v IdentityProvider) *IdentityProvider {
		return &v
	}).(IdentityProviderPtrOutput)
}

type IdentityProviderPtrOutput struct {
	*pulumi.OutputState
}

func (IdentityProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil))
}

func (o IdentityProviderPtrOutput) ToIdentityProviderPtrOutput() IdentityProviderPtrOutput {
	return o
}

func (o IdentityProviderPtrOutput) ToIdentityProviderPtrOutputWithContext(ctx context.Context) IdentityProviderPtrOutput {
	return o
}

type IdentityProviderArrayOutput struct{ *pulumi.OutputState }

func (IdentityProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentityProvider)(nil))
}

func (o IdentityProviderArrayOutput) ToIdentityProviderArrayOutput() IdentityProviderArrayOutput {
	return o
}

func (o IdentityProviderArrayOutput) ToIdentityProviderArrayOutputWithContext(ctx context.Context) IdentityProviderArrayOutput {
	return o
}

func (o IdentityProviderArrayOutput) Index(i pulumi.IntInput) IdentityProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentityProvider {
		return vs[0].([]IdentityProvider)[vs[1].(int)]
	}).(IdentityProviderOutput)
}

type IdentityProviderMapOutput struct{ *pulumi.OutputState }

func (IdentityProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityProvider)(nil))
}

func (o IdentityProviderMapOutput) ToIdentityProviderMapOutput() IdentityProviderMapOutput {
	return o
}

func (o IdentityProviderMapOutput) ToIdentityProviderMapOutputWithContext(ctx context.Context) IdentityProviderMapOutput {
	return o
}

func (o IdentityProviderMapOutput) MapIndex(k pulumi.StringInput) IdentityProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityProvider {
		return vs[0].(map[string]IdentityProvider)[vs[1].(string)]
	}).(IdentityProviderOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityProviderOutput{})
	pulumi.RegisterOutputType(IdentityProviderPtrOutput{})
	pulumi.RegisterOutputType(IdentityProviderArrayOutput{})
	pulumi.RegisterOutputType(IdentityProviderMapOutput{})
}
