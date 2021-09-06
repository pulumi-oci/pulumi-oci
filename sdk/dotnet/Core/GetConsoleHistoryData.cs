// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Core
{
    public static class GetConsoleHistoryData
    {
        /// <summary>
        /// This data source provides details about a specific Console History Content resource in Oracle Cloud Infrastructure Core service.
        /// 
        /// Gets the actual console history data (not the metadata).
        /// See [CaptureConsoleHistory](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ConsoleHistory/CaptureConsoleHistory)
        /// for details about using the console history operations.
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
        ///         var testConsoleHistoryData = Output.Create(Oci.Core.GetConsoleHistoryData.InvokeAsync(new Oci.Core.GetConsoleHistoryDataArgs
        ///         {
        ///             ConsoleHistoryId = oci_core_console_history.Test_console_history.Id,
        ///             Length = @var.Console_history_content_length,
        ///             Offset = @var.Console_history_content_offset,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetConsoleHistoryDataResult> InvokeAsync(GetConsoleHistoryDataArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConsoleHistoryDataResult>("oci:core/getConsoleHistoryData:getConsoleHistoryData", args ?? new GetConsoleHistoryDataArgs(), options.WithVersion());
    }


    public sealed class GetConsoleHistoryDataArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The OCID of the console history.
        /// </summary>
        [Input("consoleHistoryId", required: true)]
        public string ConsoleHistoryId { get; set; } = null!;

        /// <summary>
        /// Length of the snapshot data to retrieve. Cannot be less than 10240.
        /// </summary>
        [Input("length")]
        public int? Length { get; set; }

        /// <summary>
        /// Offset of the snapshot data to retrieve.
        /// </summary>
        [Input("offset")]
        public int? Offset { get; set; }

        public GetConsoleHistoryDataArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConsoleHistoryDataResult
    {
        public readonly string ConsoleHistoryId;
        /// <summary>
        /// The console history data.
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int? Length;
        public readonly int? Offset;

        [OutputConstructor]
        private GetConsoleHistoryDataResult(
            string consoleHistoryId,

            string data,

            string id,

            int? length,

            int? offset)
        {
            ConsoleHistoryId = consoleHistoryId;
            Data = data;
            Id = id;
            Length = length;
            Offset = offset;
        }
    }
}