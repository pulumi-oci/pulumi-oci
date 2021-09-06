// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Core
{
    public static class GetAppCatalogListing
    {
        /// <summary>
        /// This data source provides details about a specific App Catalog Listing resource in Oracle Cloud Infrastructure Core service.
        /// 
        /// Gets the specified listing.
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
        ///         var testAppCatalogListing = Output.Create(Oci.Core.GetAppCatalogListing.InvokeAsync(new Oci.Core.GetAppCatalogListingArgs
        ///         {
        ///             ListingId = data.Oci_core_app_catalog_listing.Test_listing.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppCatalogListingResult> InvokeAsync(GetAppCatalogListingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppCatalogListingResult>("oci:core/getAppCatalogListing:getAppCatalogListing", args ?? new GetAppCatalogListingArgs(), options.WithVersion());
    }


    public sealed class GetAppCatalogListingArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The OCID of the listing.
        /// </summary>
        [Input("listingId", required: true)]
        public string ListingId { get; set; } = null!;

        public GetAppCatalogListingArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppCatalogListingResult
    {
        /// <summary>
        /// Listing's contact URL.
        /// </summary>
        public readonly string ContactUrl;
        /// <summary>
        /// Description of the listing.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name of the listing.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// the region free ocid of the listing resource.
        /// </summary>
        public readonly string ListingId;
        /// <summary>
        /// Publisher's logo URL.
        /// </summary>
        public readonly string PublisherLogoUrl;
        /// <summary>
        /// The name of the publisher who published this listing.
        /// </summary>
        public readonly string PublisherName;
        /// <summary>
        /// The short summary for the listing.
        /// </summary>
        public readonly string Summary;
        /// <summary>
        /// Date and time the listing was published, in [RFC3339](https://tools.ietf.org/html/rfc3339) format. Example: `2018-03-20T12:32:53.532Z`
        /// </summary>
        public readonly string TimePublished;

        [OutputConstructor]
        private GetAppCatalogListingResult(
            string contactUrl,

            string description,

            string displayName,

            string id,

            string listingId,

            string publisherLogoUrl,

            string publisherName,

            string summary,

            string timePublished)
        {
            ContactUrl = contactUrl;
            Description = description;
            DisplayName = displayName;
            Id = id;
            ListingId = listingId;
            PublisherLogoUrl = publisherLogoUrl;
            PublisherName = publisherName;
            Summary = summary;
            TimePublished = timePublished;
        }
    }
}