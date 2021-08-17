// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetApigatewayApiDeploymentSpecification
    {
        /// <summary>
        /// This data source provides details about a specific Api Deployment Specification resource in Oracle Cloud Infrastructure API Gateway service.
        /// 
        /// Gets an API Deployment specification by identifier.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Oci = Pulumi.Oci;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var testApiDeploymentSpecification = Output.Create(Oci.GetApigatewayApiDeploymentSpecification.InvokeAsync(new Oci.GetApigatewayApiDeploymentSpecificationArgs
        ///         {
        ///             ApiId = oci_apigateway_api.Test_api.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApigatewayApiDeploymentSpecificationResult> InvokeAsync(GetApigatewayApiDeploymentSpecificationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApigatewayApiDeploymentSpecificationResult>("oci:index/getApigatewayApiDeploymentSpecification:GetApigatewayApiDeploymentSpecification", args ?? new GetApigatewayApiDeploymentSpecificationArgs(), options.WithVersion());
    }


    public sealed class GetApigatewayApiDeploymentSpecificationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ocid of the API.
        /// </summary>
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        public GetApigatewayApiDeploymentSpecificationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApigatewayApiDeploymentSpecificationResult
    {
        public readonly string ApiId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Policies controlling the pushing of logs to Oracle Cloud Infrastructure Public Logging.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApigatewayApiDeploymentSpecificationLoggingPolicyResult> LoggingPolicies;
        /// <summary>
        /// Behavior applied to any requests received by the API on this route.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApigatewayApiDeploymentSpecificationRequestPolicyResult> RequestPolicies;
        /// <summary>
        /// A list of routes that this API exposes.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApigatewayApiDeploymentSpecificationRouteResult> Routes;

        [OutputConstructor]
        private GetApigatewayApiDeploymentSpecificationResult(
            string apiId,

            string id,

            ImmutableArray<Outputs.GetApigatewayApiDeploymentSpecificationLoggingPolicyResult> loggingPolicies,

            ImmutableArray<Outputs.GetApigatewayApiDeploymentSpecificationRequestPolicyResult> requestPolicies,

            ImmutableArray<Outputs.GetApigatewayApiDeploymentSpecificationRouteResult> routes)
        {
            ApiId = apiId;
            Id = id;
            LoggingPolicies = loggingPolicies;
            RequestPolicies = requestPolicies;
            Routes = routes;
        }
    }
}