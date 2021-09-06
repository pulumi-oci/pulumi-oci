// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Identity
{
    public static class GetUiPassword
    {
        /// <summary>
        /// This data source provides details about a specific Ui Password resource in Oracle Cloud Infrastructure Identity service.
        /// 
        /// Gets the specified user's console password information. The returned object contains the user's OCID,
        /// but not the password itself. The actual password is returned only when created or reset.
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
        ///         var testUiPassword = Output.Create(Oci.Identity.GetUiPassword.InvokeAsync(new Oci.Identity.GetUiPasswordArgs
        ///         {
        ///             UserId = oci_identity_user.Test_user.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUiPasswordResult> InvokeAsync(GetUiPasswordArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUiPasswordResult>("oci:identity/getUiPassword:getUiPassword", args ?? new GetUiPasswordArgs(), options.WithVersion());
    }


    public sealed class GetUiPasswordArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The OCID of the user.
        /// </summary>
        [Input("userId", required: true)]
        public string UserId { get; set; } = null!;

        public GetUiPasswordArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUiPasswordResult
    {
        public readonly string Id;
        public readonly string InactiveStatus;
        public readonly string Password;
        /// <summary>
        /// The password's current state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Date and time the password was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The OCID of the user.
        /// </summary>
        public readonly string UserId;

        [OutputConstructor]
        private GetUiPasswordResult(
            string id,

            string inactiveStatus,

            string password,

            string state,

            string timeCreated,

            string userId)
        {
            Id = id;
            InactiveStatus = inactiveStatus;
            Password = password;
            State = state;
            TimeCreated = timeCreated;
            UserId = userId;
        }
    }
}