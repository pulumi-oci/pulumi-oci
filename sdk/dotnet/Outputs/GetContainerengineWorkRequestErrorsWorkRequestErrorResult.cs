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
    public sealed class GetContainerengineWorkRequestErrorsWorkRequestErrorResult
    {
        /// <summary>
        /// A short error code that defines the error, meant for programmatic parsing. See [API Errors](https://docs.cloud.oracle.com/iaas/Content/API/References/apierrors.htm).
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// A human-readable error string.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// The date and time the error occurred.
        /// </summary>
        public readonly string Timestamp;

        [OutputConstructor]
        private GetContainerengineWorkRequestErrorsWorkRequestErrorResult(
            string code,

            string message,

            string timestamp)
        {
            Code = code;
            Message = message;
            Timestamp = timestamp;
        }
    }
}