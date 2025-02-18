// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package functions

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Application resource in Oracle Cloud Infrastructure Functions service.
//
// Retrieves an application.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/functions"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := functions.LookupApplication(ctx, &functions.LookupApplicationArgs{
// 			ApplicationId: oci_functions_application.Test_application.Id,
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
	err := ctx.Invoke("oci:functions/getApplication:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplication.
type LookupApplicationArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this application.
	ApplicationId string `pulumi:"applicationId"`
}

// A collection of values returned by getApplication.
type LookupApplicationResult struct {
	ApplicationId string `pulumi:"applicationId"`
	// The OCID of the compartment that contains the application.
	CompartmentId string `pulumi:"compartmentId"`
	// Application configuration for functions in this application (passed as environment variables). Can be overridden by function configuration. Keys must be ASCII strings consisting solely of letters, digits, and the '_' (underscore) character, and must not begin with a digit. Values should be limited to printable unicode characters.  Example: `{"MY_FUNCTION_CONFIG": "ConfVal"}`
	Config map[string]interface{} `pulumi:"config"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The display name of the application. The display name is unique within the compartment containing the application.
	DisplayName string `pulumi:"displayName"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application.
	Id string `pulumi:"id"`
	// The current state of the application.
	State string `pulumi:"state"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of the subnets in which to run functions in the application.
	SubnetIds []string `pulumi:"subnetIds"`
	// A syslog URL to which to send all function logs. Supports tcp, udp, and tcp+tls. The syslog URL must be reachable from all of the subnets configured for the application. Note: If you enable the Oracle Cloud Infrastructure Logging service for this application, the syslogUrl value is ignored. Function logs are sent to the Oracle Cloud Infrastructure Logging service, and not to the syslog URL.  Example: `tcp://logserver.myserver:1234`
	SyslogUrl string `pulumi:"syslogUrl"`
	// The time the application was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-09-12T22:47:12.613Z`
	TimeCreated string `pulumi:"timeCreated"`
	// The time the application was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2018-09-12T22:47:12.613Z`
	TimeUpdated string `pulumi:"timeUpdated"`
	// Define the tracing configuration for an application.
	TraceConfig GetApplicationTraceConfig `pulumi:"traceConfig"`
}
