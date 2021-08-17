// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetCoreComputeCapacityReservations
    {
        /// <summary>
        /// This data source provides the list of Compute Capacity Reservations in Oracle Cloud Infrastructure Core service.
        /// 
        /// Lists the compute capacity reservations that match the specified criteria and compartment.
        /// 
        /// You can limit the list by specifying a compute capacity reservation display name 
        /// (the list will include all the identically-named compute capacity reservations in the compartment).
        /// 
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
        ///         var testComputeCapacityReservations = Output.Create(Oci.GetCoreComputeCapacityReservations.InvokeAsync(new Oci.GetCoreComputeCapacityReservationsArgs
        ///         {
        ///             CompartmentId = @var.Compartment_id,
        ///             AvailabilityDomain = @var.Compute_capacity_reservation_availability_domain,
        ///             DisplayName = @var.Compute_capacity_reservation_display_name,
        ///             State = @var.Compute_capacity_reservation_state,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCoreComputeCapacityReservationsResult> InvokeAsync(GetCoreComputeCapacityReservationsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCoreComputeCapacityReservationsResult>("oci:index/getCoreComputeCapacityReservations:GetCoreComputeCapacityReservations", args ?? new GetCoreComputeCapacityReservationsArgs(), options.WithVersion());
    }


    public sealed class GetCoreComputeCapacityReservationsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the availability domain.  Example: `Uocm:PHX-AD-1`
        /// </summary>
        [Input("availabilityDomain")]
        public string? AvailabilityDomain { get; set; }

        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
        /// </summary>
        [Input("compartmentId", required: true)]
        public string CompartmentId { get; set; } = null!;

        /// <summary>
        /// A filter to return only resources that match the given display name exactly.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("filters")]
        private List<Inputs.GetCoreComputeCapacityReservationsFilterArgs>? _filters;
        public List<Inputs.GetCoreComputeCapacityReservationsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetCoreComputeCapacityReservationsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// A filter to only return resources that match the given lifecycle state.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetCoreComputeCapacityReservationsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCoreComputeCapacityReservationsResult
    {
        /// <summary>
        /// The availability domain of the compute capacity reservation.  Example: `Uocm:PHX-AD-1`
        /// </summary>
        public readonly string? AvailabilityDomain;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the compute capacity reservation.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// The list of compute_capacity_reservations.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCoreComputeCapacityReservationsComputeCapacityReservationResult> ComputeCapacityReservations;
        /// <summary>
        /// A user-friendly name for the capacity reservation. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My Reservation`
        /// </summary>
        public readonly string? DisplayName;
        public readonly ImmutableArray<Outputs.GetCoreComputeCapacityReservationsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The current state of the compute capacity reservation.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private GetCoreComputeCapacityReservationsResult(
            string? availabilityDomain,

            string compartmentId,

            ImmutableArray<Outputs.GetCoreComputeCapacityReservationsComputeCapacityReservationResult> computeCapacityReservations,

            string? displayName,

            ImmutableArray<Outputs.GetCoreComputeCapacityReservationsFilterResult> filters,

            string id,

            string? state)
        {
            AvailabilityDomain = availabilityDomain;
            CompartmentId = compartmentId;
            ComputeCapacityReservations = computeCapacityReservations;
            DisplayName = displayName;
            Filters = filters;
            Id = id;
            State = state;
        }
    }
}