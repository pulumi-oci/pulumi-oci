// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aianomalydetection

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Data Asset resource in Oracle Cloud Infrastructure Ai Anomaly Detection service.
//
// Creates a new DataAsset.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/aianomalydetection"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := aianomalydetection.NewDataAsset(ctx, "testDataAsset", &aianomalydetection.DataAssetArgs{
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			DataSourceDetails: &aianomalydetection.DataAssetDataSourceDetailsArgs{
// 				DataSourceType:         pulumi.Any(_var.Data_asset_data_source_details_data_source_type),
// 				AtpPasswordSecretId:    pulumi.Any(oci_vault_secret.Test_secret.Id),
// 				AtpUserName:            pulumi.Any(oci_identity_user.Test_user.Name),
// 				Bucket:                 pulumi.Any(_var.Data_asset_data_source_details_bucket),
// 				CwalletFileSecretId:    pulumi.Any(oci_vault_secret.Test_secret.Id),
// 				DatabaseName:           pulumi.Any(oci_database_database.Test_database.Name),
// 				EwalletFileSecretId:    pulumi.Any(oci_vault_secret.Test_secret.Id),
// 				KeyStoreFileSecretId:   pulumi.Any(oci_vault_secret.Test_secret.Id),
// 				MeasurementName:        pulumi.Any(_var.Data_asset_data_source_details_measurement_name),
// 				Namespace:              pulumi.Any(_var.Data_asset_data_source_details_namespace),
// 				Object:                 pulumi.Any(_var.Data_asset_data_source_details_object),
// 				OjdbcFileSecretId:      pulumi.Any(oci_vault_secret.Test_secret.Id),
// 				PasswordSecretId:       pulumi.Any(oci_vault_secret.Test_secret.Id),
// 				TableName:              pulumi.Any(oci_nosql_table.Test_table.Name),
// 				TnsnamesFileSecretId:   pulumi.Any(oci_vault_secret.Test_secret.Id),
// 				TruststoreFileSecretId: pulumi.Any(oci_vault_secret.Test_secret.Id),
// 				Url:                    pulumi.Any(_var.Data_asset_data_source_details_url),
// 				UserName:               pulumi.Any(oci_identity_user.Test_user.Name),
// 				VersionSpecificDetails: &aianomalydetection.DataAssetDataSourceDetailsVersionSpecificDetailsArgs{
// 					InfluxVersion:       pulumi.Any(_var.Data_asset_data_source_details_version_specific_details_influx_version),
// 					Bucket:              pulumi.Any(_var.Data_asset_data_source_details_version_specific_details_bucket),
// 					DatabaseName:        pulumi.Any(oci_database_database.Test_database.Name),
// 					OrganizationName:    pulumi.Any(_var.Data_asset_data_source_details_version_specific_details_organization_name),
// 					RetentionPolicyName: pulumi.Any(oci_identity_policy.Test_policy.Name),
// 				},
// 				WalletPasswordSecretId: pulumi.Any(oci_vault_secret.Test_secret.Id),
// 			},
// 			ProjectId: pulumi.Any(oci_ai_anomaly_detection_project.Test_project.Id),
// 			DefinedTags: pulumi.AnyMap{
// 				"foo-namespace.bar-key": pulumi.Any("value"),
// 			},
// 			Description: pulumi.Any(_var.Data_asset_description),
// 			DisplayName: pulumi.Any(_var.Data_asset_display_name),
// 			FreeformTags: pulumi.AnyMap{
// 				"bar-key": pulumi.Any("value"),
// 			},
// 			PrivateEndpointId: pulumi.Any(oci_dataflow_private_endpoint.Test_private_endpoint.Id),
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
// DataAssets can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:aianomalydetection/dataAsset:DataAsset test_data_asset "id"
// ```
type DataAsset struct {
	pulumi.CustomResourceState

	// (Updatable) The OCID for the data asset's compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// Possible data sources
	DataSourceDetails DataAssetDataSourceDetailsOutput `pulumi:"dataSourceDetails"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) A short description of the Ai data asset
	Description pulumi.StringOutput `pulumi:"description"`
	// (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// OCID of Private Endpoint.
	PrivateEndpointId pulumi.StringOutput `pulumi:"privateEndpointId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the data asset.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The lifecycle state of the Data Asset.
	State pulumi.StringOutput `pulumi:"state"`
	// Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags pulumi.MapOutput `pulumi:"systemTags"`
	// The time the the DataAsset was created. An RFC3339 formatted datetime string
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The time the the DataAsset was updated. An RFC3339 formatted datetime string
	TimeUpdated pulumi.StringOutput `pulumi:"timeUpdated"`
}

// NewDataAsset registers a new resource with the given unique name, arguments, and options.
func NewDataAsset(ctx *pulumi.Context,
	name string, args *DataAssetArgs, opts ...pulumi.ResourceOption) (*DataAsset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.DataSourceDetails == nil {
		return nil, errors.New("invalid value for required argument 'DataSourceDetails'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource DataAsset
	err := ctx.RegisterResource("oci:aianomalydetection/dataAsset:DataAsset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataAsset gets an existing DataAsset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataAsset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataAssetState, opts ...pulumi.ResourceOption) (*DataAsset, error) {
	var resource DataAsset
	err := ctx.ReadResource("oci:aianomalydetection/dataAsset:DataAsset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataAsset resources.
type dataAssetState struct {
	// (Updatable) The OCID for the data asset's compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// Possible data sources
	DataSourceDetails *DataAssetDataSourceDetails `pulumi:"dataSourceDetails"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A short description of the Ai data asset
	Description *string `pulumi:"description"`
	// (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// OCID of Private Endpoint.
	PrivateEndpointId *string `pulumi:"privateEndpointId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the data asset.
	ProjectId *string `pulumi:"projectId"`
	// The lifecycle state of the Data Asset.
	State *string `pulumi:"state"`
	// Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// The time the the DataAsset was created. An RFC3339 formatted datetime string
	TimeCreated *string `pulumi:"timeCreated"`
	// The time the the DataAsset was updated. An RFC3339 formatted datetime string
	TimeUpdated *string `pulumi:"timeUpdated"`
}

type DataAssetState struct {
	// (Updatable) The OCID for the data asset's compartment.
	CompartmentId pulumi.StringPtrInput
	// Possible data sources
	DataSourceDetails DataAssetDataSourceDetailsPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A short description of the Ai data asset
	Description pulumi.StringPtrInput
	// (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// OCID of Private Endpoint.
	PrivateEndpointId pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the data asset.
	ProjectId pulumi.StringPtrInput
	// The lifecycle state of the Data Asset.
	State pulumi.StringPtrInput
	// Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags pulumi.MapInput
	// The time the the DataAsset was created. An RFC3339 formatted datetime string
	TimeCreated pulumi.StringPtrInput
	// The time the the DataAsset was updated. An RFC3339 formatted datetime string
	TimeUpdated pulumi.StringPtrInput
}

func (DataAssetState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataAssetState)(nil)).Elem()
}

type dataAssetArgs struct {
	// (Updatable) The OCID for the data asset's compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// Possible data sources
	DataSourceDetails DataAssetDataSourceDetails `pulumi:"dataSourceDetails"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A short description of the Ai data asset
	Description *string `pulumi:"description"`
	// (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// OCID of Private Endpoint.
	PrivateEndpointId *string `pulumi:"privateEndpointId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the data asset.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a DataAsset resource.
type DataAssetArgs struct {
	// (Updatable) The OCID for the data asset's compartment.
	CompartmentId pulumi.StringInput
	// Possible data sources
	DataSourceDetails DataAssetDataSourceDetailsInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A short description of the Ai data asset
	Description pulumi.StringPtrInput
	// (Updatable) A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// OCID of Private Endpoint.
	PrivateEndpointId pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate with the data asset.
	ProjectId pulumi.StringInput
}

func (DataAssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataAssetArgs)(nil)).Elem()
}

type DataAssetInput interface {
	pulumi.Input

	ToDataAssetOutput() DataAssetOutput
	ToDataAssetOutputWithContext(ctx context.Context) DataAssetOutput
}

func (*DataAsset) ElementType() reflect.Type {
	return reflect.TypeOf((*DataAsset)(nil))
}

func (i *DataAsset) ToDataAssetOutput() DataAssetOutput {
	return i.ToDataAssetOutputWithContext(context.Background())
}

func (i *DataAsset) ToDataAssetOutputWithContext(ctx context.Context) DataAssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataAssetOutput)
}

func (i *DataAsset) ToDataAssetPtrOutput() DataAssetPtrOutput {
	return i.ToDataAssetPtrOutputWithContext(context.Background())
}

func (i *DataAsset) ToDataAssetPtrOutputWithContext(ctx context.Context) DataAssetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataAssetPtrOutput)
}

type DataAssetPtrInput interface {
	pulumi.Input

	ToDataAssetPtrOutput() DataAssetPtrOutput
	ToDataAssetPtrOutputWithContext(ctx context.Context) DataAssetPtrOutput
}

type dataAssetPtrType DataAssetArgs

func (*dataAssetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataAsset)(nil))
}

func (i *dataAssetPtrType) ToDataAssetPtrOutput() DataAssetPtrOutput {
	return i.ToDataAssetPtrOutputWithContext(context.Background())
}

func (i *dataAssetPtrType) ToDataAssetPtrOutputWithContext(ctx context.Context) DataAssetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataAssetPtrOutput)
}

// DataAssetArrayInput is an input type that accepts DataAssetArray and DataAssetArrayOutput values.
// You can construct a concrete instance of `DataAssetArrayInput` via:
//
//          DataAssetArray{ DataAssetArgs{...} }
type DataAssetArrayInput interface {
	pulumi.Input

	ToDataAssetArrayOutput() DataAssetArrayOutput
	ToDataAssetArrayOutputWithContext(context.Context) DataAssetArrayOutput
}

type DataAssetArray []DataAssetInput

func (DataAssetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataAsset)(nil)).Elem()
}

func (i DataAssetArray) ToDataAssetArrayOutput() DataAssetArrayOutput {
	return i.ToDataAssetArrayOutputWithContext(context.Background())
}

func (i DataAssetArray) ToDataAssetArrayOutputWithContext(ctx context.Context) DataAssetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataAssetArrayOutput)
}

// DataAssetMapInput is an input type that accepts DataAssetMap and DataAssetMapOutput values.
// You can construct a concrete instance of `DataAssetMapInput` via:
//
//          DataAssetMap{ "key": DataAssetArgs{...} }
type DataAssetMapInput interface {
	pulumi.Input

	ToDataAssetMapOutput() DataAssetMapOutput
	ToDataAssetMapOutputWithContext(context.Context) DataAssetMapOutput
}

type DataAssetMap map[string]DataAssetInput

func (DataAssetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataAsset)(nil)).Elem()
}

func (i DataAssetMap) ToDataAssetMapOutput() DataAssetMapOutput {
	return i.ToDataAssetMapOutputWithContext(context.Background())
}

func (i DataAssetMap) ToDataAssetMapOutputWithContext(ctx context.Context) DataAssetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataAssetMapOutput)
}

type DataAssetOutput struct {
	*pulumi.OutputState
}

func (DataAssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataAsset)(nil))
}

func (o DataAssetOutput) ToDataAssetOutput() DataAssetOutput {
	return o
}

func (o DataAssetOutput) ToDataAssetOutputWithContext(ctx context.Context) DataAssetOutput {
	return o
}

func (o DataAssetOutput) ToDataAssetPtrOutput() DataAssetPtrOutput {
	return o.ToDataAssetPtrOutputWithContext(context.Background())
}

func (o DataAssetOutput) ToDataAssetPtrOutputWithContext(ctx context.Context) DataAssetPtrOutput {
	return o.ApplyT(func(v DataAsset) *DataAsset {
		return &v
	}).(DataAssetPtrOutput)
}

type DataAssetPtrOutput struct {
	*pulumi.OutputState
}

func (DataAssetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataAsset)(nil))
}

func (o DataAssetPtrOutput) ToDataAssetPtrOutput() DataAssetPtrOutput {
	return o
}

func (o DataAssetPtrOutput) ToDataAssetPtrOutputWithContext(ctx context.Context) DataAssetPtrOutput {
	return o
}

type DataAssetArrayOutput struct{ *pulumi.OutputState }

func (DataAssetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataAsset)(nil))
}

func (o DataAssetArrayOutput) ToDataAssetArrayOutput() DataAssetArrayOutput {
	return o
}

func (o DataAssetArrayOutput) ToDataAssetArrayOutputWithContext(ctx context.Context) DataAssetArrayOutput {
	return o
}

func (o DataAssetArrayOutput) Index(i pulumi.IntInput) DataAssetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataAsset {
		return vs[0].([]DataAsset)[vs[1].(int)]
	}).(DataAssetOutput)
}

type DataAssetMapOutput struct{ *pulumi.OutputState }

func (DataAssetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DataAsset)(nil))
}

func (o DataAssetMapOutput) ToDataAssetMapOutput() DataAssetMapOutput {
	return o
}

func (o DataAssetMapOutput) ToDataAssetMapOutputWithContext(ctx context.Context) DataAssetMapOutput {
	return o
}

func (o DataAssetMapOutput) MapIndex(k pulumi.StringInput) DataAssetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DataAsset {
		return vs[0].(map[string]DataAsset)[vs[1].(string)]
	}).(DataAssetOutput)
}

func init() {
	pulumi.RegisterOutputType(DataAssetOutput{})
	pulumi.RegisterOutputType(DataAssetPtrOutput{})
	pulumi.RegisterOutputType(DataAssetArrayOutput{})
	pulumi.RegisterOutputType(DataAssetMapOutput{})
}