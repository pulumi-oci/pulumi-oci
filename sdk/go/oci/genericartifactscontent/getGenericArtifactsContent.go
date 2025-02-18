// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package genericartifactscontent

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Generic Artifacts Content resource in Oracle Cloud Infrastructure Generic Artifacts Content service.
//
// Gets the specified artifact's content.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/genericartifactscontent"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := genericartifactscontent.GetGenericArtifactsContent(ctx, &genericartifactscontent.GetGenericArtifactsContentArgs{
// 			ArtifactId: oci_generic_artifacts_content_artifact.Test_artifact.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetGenericArtifactsContent(ctx *pulumi.Context, args *GetGenericArtifactsContentArgs, opts ...pulumi.InvokeOption) (*GetGenericArtifactsContentResult, error) {
	var rv GetGenericArtifactsContentResult
	err := ctx.Invoke("oci:genericartifactscontent/getGenericArtifactsContent:getGenericArtifactsContent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGenericArtifactsContent.
type GetGenericArtifactsContentArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the artifact.  Example: `ocid1.genericartifact.oc1..exampleuniqueID`
	ArtifactId string `pulumi:"artifactId"`
}

// A collection of values returned by getGenericArtifactsContent.
type GetGenericArtifactsContentResult struct {
	ArtifactId string `pulumi:"artifactId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
