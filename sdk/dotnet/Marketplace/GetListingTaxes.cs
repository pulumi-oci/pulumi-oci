// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Marketplace
{
    public static class GetListingTaxes
    {
        /// <summary>
        /// This data source provides the list of Listing Taxes in Oracle Cloud Infrastructure Marketplace service.
        /// 
        /// Returns list of all tax implications that current tenant may be liable to once they launch the listing.
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
        ///         var testListingTaxes = Output.Create(Oci.Marketplace.GetListingTaxes.InvokeAsync(new Oci.Marketplace.GetListingTaxesArgs
        ///         {
        ///             ListingId = oci_marketplace_listing.Test_listing.Id,
        ///             CompartmentId = @var.Compartment_id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetListingTaxesResult> InvokeAsync(GetListingTaxesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetListingTaxesResult>("oci:marketplace/getListingTaxes:getListingTaxes", args ?? new GetListingTaxesArgs(), options.WithVersion());
    }


    public sealed class GetListingTaxesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier for the compartment.
        /// </summary>
        [Input("compartmentId")]
        public string? CompartmentId { get; set; }

        [Input("filters")]
        private List<Inputs.GetListingTaxesFilterArgs>? _filters;
        public List<Inputs.GetListingTaxesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetListingTaxesFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The unique identifier for the listing.
        /// </summary>
        [Input("listingId", required: true)]
        public string ListingId { get; set; } = null!;

        public GetListingTaxesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetListingTaxesResult
    {
        public readonly string? CompartmentId;
        public readonly ImmutableArray<Outputs.GetListingTaxesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ListingId;
        /// <summary>
        /// The list of taxes.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListingTaxesTaxResult> Taxes;

        [OutputConstructor]
        private GetListingTaxesResult(
            string? compartmentId,

            ImmutableArray<Outputs.GetListingTaxesFilterResult> filters,

            string id,

            string listingId,

            ImmutableArray<Outputs.GetListingTaxesTaxResult> taxes)
        {
            CompartmentId = compartmentId;
            Filters = filters;
            Id = id;
            ListingId = listingId;
            Taxes = taxes;
        }
    }
}