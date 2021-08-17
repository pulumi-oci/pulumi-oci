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
    public sealed class GetHealthChecksVantagePointsHealthChecksVantagePointResult
    {
        /// <summary>
        /// Filters results that exactly match the `displayName` field.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Geographic information about a vantage point.
        /// </summary>
        public readonly Outputs.GetHealthChecksVantagePointsHealthChecksVantagePointGeoResult Geo;
        /// <summary>
        /// Filters results that exactly match the `name` field.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The organization on whose infrastructure this vantage point resides. Provider names are not unique, as Oracle Cloud Infrastructure maintains many vantage points in each major provider.
        /// </summary>
        public readonly string ProviderName;
        /// <summary>
        /// An array of objects that describe how traffic to this vantage point is routed, including which prefixes and ASNs connect it to the internet.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHealthChecksVantagePointsHealthChecksVantagePointRoutingResult> Routings;

        [OutputConstructor]
        private GetHealthChecksVantagePointsHealthChecksVantagePointResult(
            string displayName,

            Outputs.GetHealthChecksVantagePointsHealthChecksVantagePointGeoResult geo,

            string name,

            string providerName,

            ImmutableArray<Outputs.GetHealthChecksVantagePointsHealthChecksVantagePointRoutingResult> routings)
        {
            DisplayName = displayName;
            Geo = geo;
            Name = name;
            ProviderName = providerName;
            Routings = routings;
        }
    }
}