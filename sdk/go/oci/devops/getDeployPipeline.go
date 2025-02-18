// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devops

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Deploy Pipeline resource in Oracle Cloud Infrastructure Devops service.
//
// Retrieves a deployment pipeline by identifier.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/devops"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := devops.LookupDeployPipeline(ctx, &devops.LookupDeployPipelineArgs{
// 			DeployPipelineId: oci_devops_deploy_pipeline.Test_deploy_pipeline.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDeployPipeline(ctx *pulumi.Context, args *LookupDeployPipelineArgs, opts ...pulumi.InvokeOption) (*LookupDeployPipelineResult, error) {
	var rv LookupDeployPipelineResult
	err := ctx.Invoke("oci:devops/getDeployPipeline:getDeployPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDeployPipeline.
type LookupDeployPipelineArgs struct {
	// Unique pipeline identifier.
	DeployPipelineId string `pulumi:"deployPipelineId"`
}

// A collection of values returned by getDeployPipeline.
type LookupDeployPipelineResult struct {
	// The OCID of the compartment where the pipeline is created.
	CompartmentId string `pulumi:"compartmentId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// List of all artifacts used in the pipeline.
	DeployPipelineArtifacts GetDeployPipelineDeployPipelineArtifacts `pulumi:"deployPipelineArtifacts"`
	// List of all environments used in the pipeline.
	DeployPipelineEnvironments GetDeployPipelineDeployPipelineEnvironments `pulumi:"deployPipelineEnvironments"`
	DeployPipelineId           string                                      `pulumi:"deployPipelineId"`
	// Specifies list of parameters present in the deployment pipeline. In case of Update operation, replaces existing parameters list. Merging with existing parameters is not supported.
	DeployPipelineParameters GetDeployPipelineDeployPipelineParameters `pulumi:"deployPipelineParameters"`
	// Optional description about the deployment pipeline.
	Description string `pulumi:"description"`
	// Deployment pipeline display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName string `pulumi:"displayName"`
	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Unique identifier that is immutable on creation.
	Id string `pulumi:"id"`
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails string `pulumi:"lifecycleDetails"`
	// The OCID of a project.
	ProjectId string `pulumi:"projectId"`
	// The current state of the deployment pipeline.
	State string `pulumi:"state"`
	// Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// Time the deployment pipeline was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated string `pulumi:"timeCreated"`
	// Time the deployment pipeline was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated string `pulumi:"timeUpdated"`
}
