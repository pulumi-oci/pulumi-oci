// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Outputs
{

    [OutputType]
    public sealed class LoadBalancerPathRouteSetPathRoute
    {
        /// <summary>
        /// (Updatable) The name of the target backend set for requests where the incoming URI matches the specified path.  Example: `example_backend_set`
        /// </summary>
        public readonly string BackendSetName;
        /// <summary>
        /// (Updatable) The path string to match against the incoming URI path.
        /// *  Path strings are case-insensitive.
        /// *  Asterisk (*) wildcards are not supported.
        /// *  Regular expressions are not supported.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// (Updatable) The type of matching to apply to incoming URIs.
        /// </summary>
        public readonly Outputs.LoadBalancerPathRouteSetPathRoutePathMatchType PathMatchType;

        [OutputConstructor]
        private LoadBalancerPathRouteSetPathRoute(
            string backendSetName,

            string path,

            Outputs.LoadBalancerPathRouteSetPathRoutePathMatchType pathMatchType)
        {
            BackendSetName = backendSetName;
            Path = path;
            PathMatchType = pathMatchType;
        }
    }
}