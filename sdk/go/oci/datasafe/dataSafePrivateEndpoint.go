// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasafe

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Data Safe Private Endpoint resource in Oracle Cloud Infrastructure Data Safe service.
//
// Creates a new Data Safe private endpoint.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/datasafe"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datasafe.NewDataSafePrivateEndpoint(ctx, "testDataSafePrivateEndpoint", &datasafe.DataSafePrivateEndpointArgs{
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			DisplayName:   pulumi.Any(_var.Data_safe_private_endpoint_display_name),
// 			SubnetId:      pulumi.Any(oci_core_subnet.Test_subnet.Id),
// 			VcnId:         pulumi.Any(oci_core_vcn.Test_vcn.Id),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			Description: pulumi.Any(_var.Data_safe_private_endpoint_description),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			NsgIds:            pulumi.Any(_var.Data_safe_private_endpoint_nsg_ids),
// 			PrivateEndpointIp: pulumi.Any(_var.Data_safe_private_endpoint_private_endpoint_ip),
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
// DataSafePrivateEndpoints can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:datasafe/dataSafePrivateEndpoint:DataSafePrivateEndpoint test_data_safe_private_endpoint "id"
// ```
type DataSafePrivateEndpoint struct {
	pulumi.CustomResourceState

	// (Updatable) The OCID of the compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) The description of the private endpoint.
	Description pulumi.StringOutput `pulumi:"description"`
	// (Updatable) The display name for the private endpoint. The name does not have to be unique, and it's changeable.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The three-label fully qualified domain name (FQDN) of the private endpoint. The customer VCN's DNS records are updated with this FQDN.
	EndpointFqdn pulumi.StringOutput `pulumi:"endpointFqdn"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// (Updatable) The OCIDs of the network security groups that the private endpoint belongs to.
	NsgIds pulumi.StringArrayOutput `pulumi:"nsgIds"`
	// The OCID of the underlying private endpoint.
	PrivateEndpointId pulumi.StringOutput `pulumi:"privateEndpointId"`
	// The private IP address of the private endpoint.
	PrivateEndpointIp pulumi.StringOutput `pulumi:"privateEndpointIp"`
	// The current state of the private endpoint.
	State pulumi.StringOutput `pulumi:"state"`
	// The OCID of the subnet.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags pulumi.MapOutput `pulumi:"systemTags"`
	// The date and time the private endpoint was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The OCID of the VCN.
	VcnId pulumi.StringOutput `pulumi:"vcnId"`
}

// NewDataSafePrivateEndpoint registers a new resource with the given unique name, arguments, and options.
func NewDataSafePrivateEndpoint(ctx *pulumi.Context,
	name string, args *DataSafePrivateEndpointArgs, opts ...pulumi.ResourceOption) (*DataSafePrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VcnId == nil {
		return nil, errors.New("invalid value for required argument 'VcnId'")
	}
	var resource DataSafePrivateEndpoint
	err := ctx.RegisterResource("oci:datasafe/dataSafePrivateEndpoint:DataSafePrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSafePrivateEndpoint gets an existing DataSafePrivateEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSafePrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSafePrivateEndpointState, opts ...pulumi.ResourceOption) (*DataSafePrivateEndpoint, error) {
	var resource DataSafePrivateEndpoint
	err := ctx.ReadResource("oci:datasafe/dataSafePrivateEndpoint:DataSafePrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSafePrivateEndpoint resources.
type dataSafePrivateEndpointState struct {
	// (Updatable) The OCID of the compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) The description of the private endpoint.
	Description *string `pulumi:"description"`
	// (Updatable) The display name for the private endpoint. The name does not have to be unique, and it's changeable.
	DisplayName *string `pulumi:"displayName"`
	// The three-label fully qualified domain name (FQDN) of the private endpoint. The customer VCN's DNS records are updated with this FQDN.
	EndpointFqdn *string `pulumi:"endpointFqdn"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The OCIDs of the network security groups that the private endpoint belongs to.
	NsgIds []string `pulumi:"nsgIds"`
	// The OCID of the underlying private endpoint.
	PrivateEndpointId *string `pulumi:"privateEndpointId"`
	// The private IP address of the private endpoint.
	PrivateEndpointIp *string `pulumi:"privateEndpointIp"`
	// The current state of the private endpoint.
	State *string `pulumi:"state"`
	// The OCID of the subnet.
	SubnetId *string `pulumi:"subnetId"`
	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// The date and time the private endpoint was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	TimeCreated *string `pulumi:"timeCreated"`
	// The OCID of the VCN.
	VcnId *string `pulumi:"vcnId"`
}

type DataSafePrivateEndpointState struct {
	// (Updatable) The OCID of the compartment.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) The description of the private endpoint.
	Description pulumi.StringPtrInput
	// (Updatable) The display name for the private endpoint. The name does not have to be unique, and it's changeable.
	DisplayName pulumi.StringPtrInput
	// The three-label fully qualified domain name (FQDN) of the private endpoint. The customer VCN's DNS records are updated with this FQDN.
	EndpointFqdn pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The OCIDs of the network security groups that the private endpoint belongs to.
	NsgIds pulumi.StringArrayInput
	// The OCID of the underlying private endpoint.
	PrivateEndpointId pulumi.StringPtrInput
	// The private IP address of the private endpoint.
	PrivateEndpointIp pulumi.StringPtrInput
	// The current state of the private endpoint.
	State pulumi.StringPtrInput
	// The OCID of the subnet.
	SubnetId pulumi.StringPtrInput
	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags pulumi.MapInput
	// The date and time the private endpoint was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	TimeCreated pulumi.StringPtrInput
	// The OCID of the VCN.
	VcnId pulumi.StringPtrInput
}

func (DataSafePrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSafePrivateEndpointState)(nil)).Elem()
}

type dataSafePrivateEndpointArgs struct {
	// (Updatable) The OCID of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) The description of the private endpoint.
	Description *string `pulumi:"description"`
	// (Updatable) The display name for the private endpoint. The name does not have to be unique, and it's changeable.
	DisplayName string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The OCIDs of the network security groups that the private endpoint belongs to.
	NsgIds []string `pulumi:"nsgIds"`
	// The private IP address of the private endpoint.
	PrivateEndpointIp *string `pulumi:"privateEndpointIp"`
	// The OCID of the subnet.
	SubnetId string `pulumi:"subnetId"`
	// The OCID of the VCN.
	VcnId string `pulumi:"vcnId"`
}

// The set of arguments for constructing a DataSafePrivateEndpoint resource.
type DataSafePrivateEndpointArgs struct {
	// (Updatable) The OCID of the compartment.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) The description of the private endpoint.
	Description pulumi.StringPtrInput
	// (Updatable) The display name for the private endpoint. The name does not have to be unique, and it's changeable.
	DisplayName pulumi.StringInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The OCIDs of the network security groups that the private endpoint belongs to.
	NsgIds pulumi.StringArrayInput
	// The private IP address of the private endpoint.
	PrivateEndpointIp pulumi.StringPtrInput
	// The OCID of the subnet.
	SubnetId pulumi.StringInput
	// The OCID of the VCN.
	VcnId pulumi.StringInput
}

func (DataSafePrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSafePrivateEndpointArgs)(nil)).Elem()
}

type DataSafePrivateEndpointInput interface {
	pulumi.Input

	ToDataSafePrivateEndpointOutput() DataSafePrivateEndpointOutput
	ToDataSafePrivateEndpointOutputWithContext(ctx context.Context) DataSafePrivateEndpointOutput
}

func (*DataSafePrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSafePrivateEndpoint)(nil))
}

func (i *DataSafePrivateEndpoint) ToDataSafePrivateEndpointOutput() DataSafePrivateEndpointOutput {
	return i.ToDataSafePrivateEndpointOutputWithContext(context.Background())
}

func (i *DataSafePrivateEndpoint) ToDataSafePrivateEndpointOutputWithContext(ctx context.Context) DataSafePrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSafePrivateEndpointOutput)
}

func (i *DataSafePrivateEndpoint) ToDataSafePrivateEndpointPtrOutput() DataSafePrivateEndpointPtrOutput {
	return i.ToDataSafePrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *DataSafePrivateEndpoint) ToDataSafePrivateEndpointPtrOutputWithContext(ctx context.Context) DataSafePrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSafePrivateEndpointPtrOutput)
}

type DataSafePrivateEndpointPtrInput interface {
	pulumi.Input

	ToDataSafePrivateEndpointPtrOutput() DataSafePrivateEndpointPtrOutput
	ToDataSafePrivateEndpointPtrOutputWithContext(ctx context.Context) DataSafePrivateEndpointPtrOutput
}

type dataSafePrivateEndpointPtrType DataSafePrivateEndpointArgs

func (*dataSafePrivateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSafePrivateEndpoint)(nil))
}

func (i *dataSafePrivateEndpointPtrType) ToDataSafePrivateEndpointPtrOutput() DataSafePrivateEndpointPtrOutput {
	return i.ToDataSafePrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *dataSafePrivateEndpointPtrType) ToDataSafePrivateEndpointPtrOutputWithContext(ctx context.Context) DataSafePrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSafePrivateEndpointPtrOutput)
}

// DataSafePrivateEndpointArrayInput is an input type that accepts DataSafePrivateEndpointArray and DataSafePrivateEndpointArrayOutput values.
// You can construct a concrete instance of `DataSafePrivateEndpointArrayInput` via:
//
//          DataSafePrivateEndpointArray{ DataSafePrivateEndpointArgs{...} }
type DataSafePrivateEndpointArrayInput interface {
	pulumi.Input

	ToDataSafePrivateEndpointArrayOutput() DataSafePrivateEndpointArrayOutput
	ToDataSafePrivateEndpointArrayOutputWithContext(context.Context) DataSafePrivateEndpointArrayOutput
}

type DataSafePrivateEndpointArray []DataSafePrivateEndpointInput

func (DataSafePrivateEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataSafePrivateEndpoint)(nil)).Elem()
}

func (i DataSafePrivateEndpointArray) ToDataSafePrivateEndpointArrayOutput() DataSafePrivateEndpointArrayOutput {
	return i.ToDataSafePrivateEndpointArrayOutputWithContext(context.Background())
}

func (i DataSafePrivateEndpointArray) ToDataSafePrivateEndpointArrayOutputWithContext(ctx context.Context) DataSafePrivateEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSafePrivateEndpointArrayOutput)
}

// DataSafePrivateEndpointMapInput is an input type that accepts DataSafePrivateEndpointMap and DataSafePrivateEndpointMapOutput values.
// You can construct a concrete instance of `DataSafePrivateEndpointMapInput` via:
//
//          DataSafePrivateEndpointMap{ "key": DataSafePrivateEndpointArgs{...} }
type DataSafePrivateEndpointMapInput interface {
	pulumi.Input

	ToDataSafePrivateEndpointMapOutput() DataSafePrivateEndpointMapOutput
	ToDataSafePrivateEndpointMapOutputWithContext(context.Context) DataSafePrivateEndpointMapOutput
}

type DataSafePrivateEndpointMap map[string]DataSafePrivateEndpointInput

func (DataSafePrivateEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataSafePrivateEndpoint)(nil)).Elem()
}

func (i DataSafePrivateEndpointMap) ToDataSafePrivateEndpointMapOutput() DataSafePrivateEndpointMapOutput {
	return i.ToDataSafePrivateEndpointMapOutputWithContext(context.Background())
}

func (i DataSafePrivateEndpointMap) ToDataSafePrivateEndpointMapOutputWithContext(ctx context.Context) DataSafePrivateEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSafePrivateEndpointMapOutput)
}

type DataSafePrivateEndpointOutput struct {
	*pulumi.OutputState
}

func (DataSafePrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSafePrivateEndpoint)(nil))
}

func (o DataSafePrivateEndpointOutput) ToDataSafePrivateEndpointOutput() DataSafePrivateEndpointOutput {
	return o
}

func (o DataSafePrivateEndpointOutput) ToDataSafePrivateEndpointOutputWithContext(ctx context.Context) DataSafePrivateEndpointOutput {
	return o
}

func (o DataSafePrivateEndpointOutput) ToDataSafePrivateEndpointPtrOutput() DataSafePrivateEndpointPtrOutput {
	return o.ToDataSafePrivateEndpointPtrOutputWithContext(context.Background())
}

func (o DataSafePrivateEndpointOutput) ToDataSafePrivateEndpointPtrOutputWithContext(ctx context.Context) DataSafePrivateEndpointPtrOutput {
	return o.ApplyT(func(v DataSafePrivateEndpoint) *DataSafePrivateEndpoint {
		return &v
	}).(DataSafePrivateEndpointPtrOutput)
}

type DataSafePrivateEndpointPtrOutput struct {
	*pulumi.OutputState
}

func (DataSafePrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSafePrivateEndpoint)(nil))
}

func (o DataSafePrivateEndpointPtrOutput) ToDataSafePrivateEndpointPtrOutput() DataSafePrivateEndpointPtrOutput {
	return o
}

func (o DataSafePrivateEndpointPtrOutput) ToDataSafePrivateEndpointPtrOutputWithContext(ctx context.Context) DataSafePrivateEndpointPtrOutput {
	return o
}

type DataSafePrivateEndpointArrayOutput struct{ *pulumi.OutputState }

func (DataSafePrivateEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSafePrivateEndpoint)(nil))
}

func (o DataSafePrivateEndpointArrayOutput) ToDataSafePrivateEndpointArrayOutput() DataSafePrivateEndpointArrayOutput {
	return o
}

func (o DataSafePrivateEndpointArrayOutput) ToDataSafePrivateEndpointArrayOutputWithContext(ctx context.Context) DataSafePrivateEndpointArrayOutput {
	return o
}

func (o DataSafePrivateEndpointArrayOutput) Index(i pulumi.IntInput) DataSafePrivateEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataSafePrivateEndpoint {
		return vs[0].([]DataSafePrivateEndpoint)[vs[1].(int)]
	}).(DataSafePrivateEndpointOutput)
}

type DataSafePrivateEndpointMapOutput struct{ *pulumi.OutputState }

func (DataSafePrivateEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DataSafePrivateEndpoint)(nil))
}

func (o DataSafePrivateEndpointMapOutput) ToDataSafePrivateEndpointMapOutput() DataSafePrivateEndpointMapOutput {
	return o
}

func (o DataSafePrivateEndpointMapOutput) ToDataSafePrivateEndpointMapOutputWithContext(ctx context.Context) DataSafePrivateEndpointMapOutput {
	return o
}

func (o DataSafePrivateEndpointMapOutput) MapIndex(k pulumi.StringInput) DataSafePrivateEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DataSafePrivateEndpoint {
		return vs[0].(map[string]DataSafePrivateEndpoint)[vs[1].(string)]
	}).(DataSafePrivateEndpointOutput)
}

func init() {
	pulumi.RegisterOutputType(DataSafePrivateEndpointOutput{})
	pulumi.RegisterOutputType(DataSafePrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(DataSafePrivateEndpointArrayOutput{})
	pulumi.RegisterOutputType(DataSafePrivateEndpointMapOutput{})
}
