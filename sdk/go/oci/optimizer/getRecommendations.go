// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package optimizer

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Recommendations in Oracle Cloud Infrastructure Optimizer service.
//
// Lists the Cloud Advisor recommendations that are currently supported in the specified category.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/optimizer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Recommendation_name
// 		opt1 := _var.Recommendation_state
// 		opt2 := _var.Recommendation_status
// 		_, err := optimizer.GetRecommendations(ctx, &optimizer.GetRecommendationsArgs{
// 			CategoryId:             oci_optimizer_category.Test_category.Id,
// 			CompartmentId:          _var.Compartment_id,
// 			CompartmentIdInSubtree: _var.Recommendation_compartment_id_in_subtree,
// 			Name:                   &opt0,
// 			State:                  &opt1,
// 			Status:                 &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetRecommendations(ctx *pulumi.Context, args *GetRecommendationsArgs, opts ...pulumi.InvokeOption) (*GetRecommendationsResult, error) {
	var rv GetRecommendationsResult
	err := ctx.Invoke("oci:optimizer/getRecommendations:getRecommendations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRecommendations.
type GetRecommendationsArgs struct {
	// The unique OCID associated with the category.
	CategoryId string `pulumi:"categoryId"`
	// The OCID of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.
	CompartmentIdInSubtree bool                       `pulumi:"compartmentIdInSubtree"`
	Filters                []GetRecommendationsFilter `pulumi:"filters"`
	// Optional. A filter that returns results that match the name specified.
	Name *string `pulumi:"name"`
	// A filter that returns results that match the lifecycle state specified.
	State *string `pulumi:"state"`
	// A filter that returns recommendations that match the status specified.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getRecommendations.
type GetRecommendationsResult struct {
	// The unique OCID associated with the category.
	CategoryId string `pulumi:"categoryId"`
	// The OCID of the tenancy. The tenancy is the root compartment.
	CompartmentId          string                     `pulumi:"compartmentId"`
	CompartmentIdInSubtree bool                       `pulumi:"compartmentIdInSubtree"`
	Filters                []GetRecommendationsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the profile level.
	Name *string `pulumi:"name"`
	// The list of recommendation_collection.
	RecommendationCollections []GetRecommendationsRecommendationCollection `pulumi:"recommendationCollections"`
	// The recommendation's current state.
	State *string `pulumi:"state"`
	// The current status of the recommendation.
	Status *string `pulumi:"status"`
}
