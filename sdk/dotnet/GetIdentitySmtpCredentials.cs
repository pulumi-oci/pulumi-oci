// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetIdentitySmtpCredentials
    {
        /// <summary>
        /// This data source provides the list of Smtp Credentials in Oracle Cloud Infrastructure Identity service.
        /// 
        /// Lists the SMTP credentials for the specified user. The returned object contains the credential's OCID,
        /// the SMTP user name but not the SMTP password. The SMTP password is returned only upon creation.
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
        ///         var testSmtpCredentials = Output.Create(Oci.GetIdentitySmtpCredentials.InvokeAsync(new Oci.GetIdentitySmtpCredentialsArgs
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
        public static Task<GetIdentitySmtpCredentialsResult> InvokeAsync(GetIdentitySmtpCredentialsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIdentitySmtpCredentialsResult>("oci:index/getIdentitySmtpCredentials:GetIdentitySmtpCredentials", args ?? new GetIdentitySmtpCredentialsArgs(), options.WithVersion());
    }


    public sealed class GetIdentitySmtpCredentialsArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetIdentitySmtpCredentialsFilterArgs>? _filters;
        public List<Inputs.GetIdentitySmtpCredentialsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetIdentitySmtpCredentialsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The OCID of the user.
        /// </summary>
        [Input("userId", required: true)]
        public string UserId { get; set; } = null!;

        public GetIdentitySmtpCredentialsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIdentitySmtpCredentialsResult
    {
        public readonly ImmutableArray<Outputs.GetIdentitySmtpCredentialsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of smtp_credentials.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIdentitySmtpCredentialsSmtpCredentialResult> SmtpCredentials;
        /// <summary>
        /// The OCID of the user the SMTP credential belongs to.
        /// </summary>
        public readonly string UserId;

        [OutputConstructor]
        private GetIdentitySmtpCredentialsResult(
            ImmutableArray<Outputs.GetIdentitySmtpCredentialsFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetIdentitySmtpCredentialsSmtpCredentialResult> smtpCredentials,

            string userId)
        {
            Filters = filters;
            Id = id;
            SmtpCredentials = smtpCredentials;
            UserId = userId;
        }
    }
}