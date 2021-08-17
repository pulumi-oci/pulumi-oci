// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Exadata Iorm Config resource in Oracle Cloud Infrastructure Database service.
//
// Updates IORM settings for the specified Exadata DB system.
//
// **Note:** Deprecated for Exadata Cloud Service systems. Use the [new resource model APIs](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
//
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See [Switching an Exadata DB System to the New Resource Model and APIs](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
//
// The [UpdateCloudVmClusterIormConfig](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudVmCluster/UpdateCloudVmClusterIormConfig/) API is used for Exadata systems using the
// new resource model.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oci.NewDatabaseExadataIormConfig(ctx, "testExadataIormConfig", &oci.DatabaseExadataIormConfigArgs{
// 			DbPlans: DatabaseExadataIormConfigDbPlanArray{
// 				&DatabaseExadataIormConfigDbPlanArgs{
// 					DbName: pulumi.Any(_var.Exadata_iorm_config_db_plans_db_name),
// 					Share:  pulumi.Any(_var.Exadata_iorm_config_db_plans_share),
// 				},
// 			},
// 			DbSystemId: pulumi.Any(oci_database_db_system.Test_db_system.Id),
// 			Objective:  pulumi.String("AUTO"),
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
// Import is not supported for this resource.
type DatabaseExadataIormConfig struct {
	pulumi.CustomResourceState

	// (Updatable) Array of IORM Setting for all the database in this Exadata DB System
	DbPlans DatabaseExadataIormConfigDbPlanArrayOutput `pulumi:"dbPlans"`
	// (Updatable) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbSystemId pulumi.StringOutput `pulumi:"dbSystemId"`
	// Additional information about the current `lifecycleState`.
	LifecycleDetails pulumi.StringOutput `pulumi:"lifecycleDetails"`
	// (Updatable) Value for the IORM objective Default is "Auto"
	Objective pulumi.StringOutput `pulumi:"objective"`
	// The current state of IORM configuration for the Exadata DB system.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewDatabaseExadataIormConfig registers a new resource with the given unique name, arguments, and options.
func NewDatabaseExadataIormConfig(ctx *pulumi.Context,
	name string, args *DatabaseExadataIormConfigArgs, opts ...pulumi.ResourceOption) (*DatabaseExadataIormConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbPlans == nil {
		return nil, errors.New("invalid value for required argument 'DbPlans'")
	}
	if args.DbSystemId == nil {
		return nil, errors.New("invalid value for required argument 'DbSystemId'")
	}
	var resource DatabaseExadataIormConfig
	err := ctx.RegisterResource("oci:index/databaseExadataIormConfig:DatabaseExadataIormConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseExadataIormConfig gets an existing DatabaseExadataIormConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseExadataIormConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseExadataIormConfigState, opts ...pulumi.ResourceOption) (*DatabaseExadataIormConfig, error) {
	var resource DatabaseExadataIormConfig
	err := ctx.ReadResource("oci:index/databaseExadataIormConfig:DatabaseExadataIormConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseExadataIormConfig resources.
type databaseExadataIormConfigState struct {
	// (Updatable) Array of IORM Setting for all the database in this Exadata DB System
	DbPlans []DatabaseExadataIormConfigDbPlan `pulumi:"dbPlans"`
	// (Updatable) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `pulumi:"dbSystemId"`
	// Additional information about the current `lifecycleState`.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// (Updatable) Value for the IORM objective Default is "Auto"
	Objective *string `pulumi:"objective"`
	// The current state of IORM configuration for the Exadata DB system.
	State *string `pulumi:"state"`
}

type DatabaseExadataIormConfigState struct {
	// (Updatable) Array of IORM Setting for all the database in this Exadata DB System
	DbPlans DatabaseExadataIormConfigDbPlanArrayInput
	// (Updatable) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbSystemId pulumi.StringPtrInput
	// Additional information about the current `lifecycleState`.
	LifecycleDetails pulumi.StringPtrInput
	// (Updatable) Value for the IORM objective Default is "Auto"
	Objective pulumi.StringPtrInput
	// The current state of IORM configuration for the Exadata DB system.
	State pulumi.StringPtrInput
}

func (DatabaseExadataIormConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseExadataIormConfigState)(nil)).Elem()
}

type databaseExadataIormConfigArgs struct {
	// (Updatable) Array of IORM Setting for all the database in this Exadata DB System
	DbPlans []DatabaseExadataIormConfigDbPlan `pulumi:"dbPlans"`
	// (Updatable) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbSystemId string `pulumi:"dbSystemId"`
	// (Updatable) Value for the IORM objective Default is "Auto"
	Objective *string `pulumi:"objective"`
}

// The set of arguments for constructing a DatabaseExadataIormConfig resource.
type DatabaseExadataIormConfigArgs struct {
	// (Updatable) Array of IORM Setting for all the database in this Exadata DB System
	DbPlans DatabaseExadataIormConfigDbPlanArrayInput
	// (Updatable) The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbSystemId pulumi.StringInput
	// (Updatable) Value for the IORM objective Default is "Auto"
	Objective pulumi.StringPtrInput
}

func (DatabaseExadataIormConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseExadataIormConfigArgs)(nil)).Elem()
}

type DatabaseExadataIormConfigInput interface {
	pulumi.Input

	ToDatabaseExadataIormConfigOutput() DatabaseExadataIormConfigOutput
	ToDatabaseExadataIormConfigOutputWithContext(ctx context.Context) DatabaseExadataIormConfigOutput
}

func (*DatabaseExadataIormConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseExadataIormConfig)(nil))
}

func (i *DatabaseExadataIormConfig) ToDatabaseExadataIormConfigOutput() DatabaseExadataIormConfigOutput {
	return i.ToDatabaseExadataIormConfigOutputWithContext(context.Background())
}

func (i *DatabaseExadataIormConfig) ToDatabaseExadataIormConfigOutputWithContext(ctx context.Context) DatabaseExadataIormConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseExadataIormConfigOutput)
}

func (i *DatabaseExadataIormConfig) ToDatabaseExadataIormConfigPtrOutput() DatabaseExadataIormConfigPtrOutput {
	return i.ToDatabaseExadataIormConfigPtrOutputWithContext(context.Background())
}

func (i *DatabaseExadataIormConfig) ToDatabaseExadataIormConfigPtrOutputWithContext(ctx context.Context) DatabaseExadataIormConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseExadataIormConfigPtrOutput)
}

type DatabaseExadataIormConfigPtrInput interface {
	pulumi.Input

	ToDatabaseExadataIormConfigPtrOutput() DatabaseExadataIormConfigPtrOutput
	ToDatabaseExadataIormConfigPtrOutputWithContext(ctx context.Context) DatabaseExadataIormConfigPtrOutput
}

type databaseExadataIormConfigPtrType DatabaseExadataIormConfigArgs

func (*databaseExadataIormConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseExadataIormConfig)(nil))
}

func (i *databaseExadataIormConfigPtrType) ToDatabaseExadataIormConfigPtrOutput() DatabaseExadataIormConfigPtrOutput {
	return i.ToDatabaseExadataIormConfigPtrOutputWithContext(context.Background())
}

func (i *databaseExadataIormConfigPtrType) ToDatabaseExadataIormConfigPtrOutputWithContext(ctx context.Context) DatabaseExadataIormConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseExadataIormConfigPtrOutput)
}

// DatabaseExadataIormConfigArrayInput is an input type that accepts DatabaseExadataIormConfigArray and DatabaseExadataIormConfigArrayOutput values.
// You can construct a concrete instance of `DatabaseExadataIormConfigArrayInput` via:
//
//          DatabaseExadataIormConfigArray{ DatabaseExadataIormConfigArgs{...} }
type DatabaseExadataIormConfigArrayInput interface {
	pulumi.Input

	ToDatabaseExadataIormConfigArrayOutput() DatabaseExadataIormConfigArrayOutput
	ToDatabaseExadataIormConfigArrayOutputWithContext(context.Context) DatabaseExadataIormConfigArrayOutput
}

type DatabaseExadataIormConfigArray []DatabaseExadataIormConfigInput

func (DatabaseExadataIormConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseExadataIormConfig)(nil)).Elem()
}

func (i DatabaseExadataIormConfigArray) ToDatabaseExadataIormConfigArrayOutput() DatabaseExadataIormConfigArrayOutput {
	return i.ToDatabaseExadataIormConfigArrayOutputWithContext(context.Background())
}

func (i DatabaseExadataIormConfigArray) ToDatabaseExadataIormConfigArrayOutputWithContext(ctx context.Context) DatabaseExadataIormConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseExadataIormConfigArrayOutput)
}

// DatabaseExadataIormConfigMapInput is an input type that accepts DatabaseExadataIormConfigMap and DatabaseExadataIormConfigMapOutput values.
// You can construct a concrete instance of `DatabaseExadataIormConfigMapInput` via:
//
//          DatabaseExadataIormConfigMap{ "key": DatabaseExadataIormConfigArgs{...} }
type DatabaseExadataIormConfigMapInput interface {
	pulumi.Input

	ToDatabaseExadataIormConfigMapOutput() DatabaseExadataIormConfigMapOutput
	ToDatabaseExadataIormConfigMapOutputWithContext(context.Context) DatabaseExadataIormConfigMapOutput
}

type DatabaseExadataIormConfigMap map[string]DatabaseExadataIormConfigInput

func (DatabaseExadataIormConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseExadataIormConfig)(nil)).Elem()
}

func (i DatabaseExadataIormConfigMap) ToDatabaseExadataIormConfigMapOutput() DatabaseExadataIormConfigMapOutput {
	return i.ToDatabaseExadataIormConfigMapOutputWithContext(context.Background())
}

func (i DatabaseExadataIormConfigMap) ToDatabaseExadataIormConfigMapOutputWithContext(ctx context.Context) DatabaseExadataIormConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseExadataIormConfigMapOutput)
}

type DatabaseExadataIormConfigOutput struct {
	*pulumi.OutputState
}

func (DatabaseExadataIormConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseExadataIormConfig)(nil))
}

func (o DatabaseExadataIormConfigOutput) ToDatabaseExadataIormConfigOutput() DatabaseExadataIormConfigOutput {
	return o
}

func (o DatabaseExadataIormConfigOutput) ToDatabaseExadataIormConfigOutputWithContext(ctx context.Context) DatabaseExadataIormConfigOutput {
	return o
}

func (o DatabaseExadataIormConfigOutput) ToDatabaseExadataIormConfigPtrOutput() DatabaseExadataIormConfigPtrOutput {
	return o.ToDatabaseExadataIormConfigPtrOutputWithContext(context.Background())
}

func (o DatabaseExadataIormConfigOutput) ToDatabaseExadataIormConfigPtrOutputWithContext(ctx context.Context) DatabaseExadataIormConfigPtrOutput {
	return o.ApplyT(func(v DatabaseExadataIormConfig) *DatabaseExadataIormConfig {
		return &v
	}).(DatabaseExadataIormConfigPtrOutput)
}

type DatabaseExadataIormConfigPtrOutput struct {
	*pulumi.OutputState
}

func (DatabaseExadataIormConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseExadataIormConfig)(nil))
}

func (o DatabaseExadataIormConfigPtrOutput) ToDatabaseExadataIormConfigPtrOutput() DatabaseExadataIormConfigPtrOutput {
	return o
}

func (o DatabaseExadataIormConfigPtrOutput) ToDatabaseExadataIormConfigPtrOutputWithContext(ctx context.Context) DatabaseExadataIormConfigPtrOutput {
	return o
}

type DatabaseExadataIormConfigArrayOutput struct{ *pulumi.OutputState }

func (DatabaseExadataIormConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseExadataIormConfig)(nil))
}

func (o DatabaseExadataIormConfigArrayOutput) ToDatabaseExadataIormConfigArrayOutput() DatabaseExadataIormConfigArrayOutput {
	return o
}

func (o DatabaseExadataIormConfigArrayOutput) ToDatabaseExadataIormConfigArrayOutputWithContext(ctx context.Context) DatabaseExadataIormConfigArrayOutput {
	return o
}

func (o DatabaseExadataIormConfigArrayOutput) Index(i pulumi.IntInput) DatabaseExadataIormConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseExadataIormConfig {
		return vs[0].([]DatabaseExadataIormConfig)[vs[1].(int)]
	}).(DatabaseExadataIormConfigOutput)
}

type DatabaseExadataIormConfigMapOutput struct{ *pulumi.OutputState }

func (DatabaseExadataIormConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatabaseExadataIormConfig)(nil))
}

func (o DatabaseExadataIormConfigMapOutput) ToDatabaseExadataIormConfigMapOutput() DatabaseExadataIormConfigMapOutput {
	return o
}

func (o DatabaseExadataIormConfigMapOutput) ToDatabaseExadataIormConfigMapOutputWithContext(ctx context.Context) DatabaseExadataIormConfigMapOutput {
	return o
}

func (o DatabaseExadataIormConfigMapOutput) MapIndex(k pulumi.StringInput) DatabaseExadataIormConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatabaseExadataIormConfig {
		return vs[0].(map[string]DatabaseExadataIormConfig)[vs[1].(string)]
	}).(DatabaseExadataIormConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseExadataIormConfigOutput{})
	pulumi.RegisterOutputType(DatabaseExadataIormConfigPtrOutput{})
	pulumi.RegisterOutputType(DatabaseExadataIormConfigArrayOutput{})
	pulumi.RegisterOutputType(DatabaseExadataIormConfigMapOutput{})
}