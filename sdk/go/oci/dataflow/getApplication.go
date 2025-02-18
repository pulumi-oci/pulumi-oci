// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataflow

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Application resource in Oracle Cloud Infrastructure Data Flow service.
//
// Retrieves an application using an `applicationId`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/dataflow"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataflow.LookupApplication(ctx, &dataflow.LookupApplicationArgs{
// 			ApplicationId: oci_dataflow_application.Test_application.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("oci:dataflow/getApplication:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplication.
type LookupApplicationArgs struct {
	// The unique ID for an application.
	ApplicationId string `pulumi:"applicationId"`
}

// A collection of values returned by getApplication.
type LookupApplicationResult struct {
	ApplicationId string `pulumi:"applicationId"`
	// An Oracle Cloud Infrastructure URI of an archive.zip file containing custom dependencies that may be used to support the execution a Python, Java, or Scala application. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	ArchiveUri string `pulumi:"archiveUri"`
	// The arguments passed to the running application as command line arguments.  An argument is either a plain text or a placeholder. Placeholders are replaced using values from the parameters map.  Each placeholder specified must be represented in the parameters map else the request (POST or PUT) will fail with a HTTP 400 status code.  Placeholders are specified as `Service Api Spec`, where `name` is the name of the parameter. Example:  `[ "--input", "${input_file}", "--name", "John Doe" ]` If "inputFile" has a value of "mydata.xml", then the value above will be translated to `--input mydata.xml --name "John Doe"`
	Arguments []string `pulumi:"arguments"`
	// The class for the application.
	ClassName string `pulumi:"className"`
	// The OCID of a compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// The Spark configuration passed to the running process. See https://spark.apache.org/docs/latest/configuration.html#available-properties. Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" } Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is not allowed to be overwritten will cause a 400 status to be returned.
	Configuration map[string]interface{} `pulumi:"configuration"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// A user-friendly description.
	Description string `pulumi:"description"`
	// A user-friendly name. This name is not necessarily unique.
	DisplayName string `pulumi:"displayName"`
	// The VM shape for the driver. Sets the driver cores and memory.
	DriverShape string `pulumi:"driverShape"`
	// The input used for spark-submit command. For more details see https://spark.apache.org/docs/latest/submitting-applications.html#launching-applications-with-spark-submit. Supported options include ``--class``, ``--file``, ``--jars``, ``--conf``, ``--py-files``, and main application file with arguments. Example: ``--jars oci://path/to/a.jar,oci://path/to/b.jar --files oci://path/to/a.json,oci://path/to/b.csv --py-files oci://path/to/a.py,oci://path/to/b.py --conf spark.sql.crossJoin.enabled=true --class org.apache.spark.examples.SparkPi oci://path/to/main.jar 10`` Note: If execute is specified together with applicationId, className, configuration, fileUri, language, arguments, parameters during application create/update, or run create/submit, Data Flow service will use derived information from execute input only.
	Execute string `pulumi:"execute"`
	// The VM shape for the executors. Sets the executor cores and memory.
	ExecutorShape string `pulumi:"executorShape"`
	// An Oracle Cloud Infrastructure URI of the file containing the application to execute. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	FileUri string `pulumi:"fileUri"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The application ID.
	Id string `pulumi:"id"`
	// The Spark language.
	Language string `pulumi:"language"`
	// An Oracle Cloud Infrastructure URI of the bucket where the Spark job logs are to be uploaded. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	LogsBucketUri string `pulumi:"logsBucketUri"`
	// The OCID of Oracle Cloud Infrastructure Hive Metastore.
	MetastoreId string `pulumi:"metastoreId"`
	// The number of executor VMs requested.
	NumExecutors int `pulumi:"numExecutors"`
	// The OCID of the user who created the resource.
	OwnerPrincipalId string `pulumi:"ownerPrincipalId"`
	// The username of the user who created the resource.  If the username of the owner does not exist, `null` will be returned and the caller should refer to the ownerPrincipalId value instead.
	OwnerUserName string `pulumi:"ownerUserName"`
	// An array of name/value pairs used to fill placeholders found in properties like `Application.arguments`.  The name must be a string of one or more word characters (a-z, A-Z, 0-9, _).  The value can be a string of 0 or more characters of any kind. Example:  [ { name: "iterations", value: "10"}, { name: "inputFile", value: "mydata.xml" }, { name: "variableX", value: "${x}"} ]
	Parameters []GetApplicationParameter `pulumi:"parameters"`
	// The OCID of a private endpoint.
	PrivateEndpointId string `pulumi:"privateEndpointId"`
	// The Spark version utilized to run the application.
	SparkVersion string `pulumi:"sparkVersion"`
	// The current state of this application.
	State string `pulumi:"state"`
	// The date and time a application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z`
	TimeCreated string `pulumi:"timeCreated"`
	// The date and time a application was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated string `pulumi:"timeUpdated"`
	// An Oracle Cloud Infrastructure URI of the bucket to be used as default warehouse directory for BATCH SQL runs. See https://docs.cloud.oracle.com/iaas/Content/API/SDKDocs/hdfsconnector.htm#uriformat.
	WarehouseBucketUri string `pulumi:"warehouseBucketUri"`
}
