// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetMonitoringAlarmHistoryCollection
    {
        /// <summary>
        /// This data source provides details about a specific Alarm History Collection resource in Oracle Cloud Infrastructure Monitoring service.
        /// 
        /// Get the history of the specified alarm.
        /// For important limits information, see [Limits on Monitoring](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#Limits).
        /// 
        /// This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations. 
        /// Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests, 
        /// or transactions, per second (TPS) for a given tenancy.
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
        ///         var testAlarmHistoryCollection = Output.Create(Oci.GetMonitoringAlarmHistoryCollection.InvokeAsync(new Oci.GetMonitoringAlarmHistoryCollectionArgs
        ///         {
        ///             AlarmId = oci_monitoring_alarm.Test_alarm.Id,
        ///             AlarmHistorytype = @var.Alarm_history_collection_alarm_historytype,
        ///             TimestampGreaterThanOrEqualTo = @var.Alarm_history_collection_timestamp_greater_than_or_equal_to,
        ///             TimestampLessThan = @var.Alarm_history_collection_timestamp_less_than,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMonitoringAlarmHistoryCollectionResult> InvokeAsync(GetMonitoringAlarmHistoryCollectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMonitoringAlarmHistoryCollectionResult>("oci:index/getMonitoringAlarmHistoryCollection:GetMonitoringAlarmHistoryCollection", args ?? new GetMonitoringAlarmHistoryCollectionArgs(), options.WithVersion());
    }


    public sealed class GetMonitoringAlarmHistoryCollectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of history entries to retrieve. State history (STATE_HISTORY) or state transition history (STATE_TRANSITION_HISTORY). If not specified, entries of both types are retrieved.  Example: `STATE_HISTORY`
        /// </summary>
        [Input("alarmHistorytype")]
        public string? AlarmHistorytype { get; set; }

        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of an alarm.
        /// </summary>
        [Input("alarmId", required: true)]
        public string AlarmId { get; set; } = null!;

        /// <summary>
        /// A filter to return only alarm history entries with timestamps occurring on or after the specified date and time. Format defined by RFC3339.  Example: `2019-01-01T01:00:00.789Z`
        /// </summary>
        [Input("timestampGreaterThanOrEqualTo")]
        public string? TimestampGreaterThanOrEqualTo { get; set; }

        /// <summary>
        /// A filter to return only alarm history entries with timestamps occurring before the specified date and time. Format defined by RFC3339.  Example: `2019-01-02T01:00:00.789Z`
        /// </summary>
        [Input("timestampLessThan")]
        public string? TimestampLessThan { get; set; }

        public GetMonitoringAlarmHistoryCollectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMonitoringAlarmHistoryCollectionResult
    {
        public readonly string? AlarmHistorytype;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm for which to retrieve history.
        /// </summary>
        public readonly string AlarmId;
        /// <summary>
        /// The set of history entries retrieved for the alarm.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMonitoringAlarmHistoryCollectionEntryResult> Entries;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether the alarm is enabled.  Example: `true`
        /// </summary>
        public readonly bool IsEnabled;
        public readonly string? TimestampGreaterThanOrEqualTo;
        public readonly string? TimestampLessThan;

        [OutputConstructor]
        private GetMonitoringAlarmHistoryCollectionResult(
            string? alarmHistorytype,

            string alarmId,

            ImmutableArray<Outputs.GetMonitoringAlarmHistoryCollectionEntryResult> entries,

            string id,

            bool isEnabled,

            string? timestampGreaterThanOrEqualTo,

            string? timestampLessThan)
        {
            AlarmHistorytype = alarmHistorytype;
            AlarmId = alarmId;
            Entries = entries;
            Id = id;
            IsEnabled = isEnabled;
            TimestampGreaterThanOrEqualTo = timestampGreaterThanOrEqualTo;
            TimestampLessThan = timestampLessThan;
        }
    }
}