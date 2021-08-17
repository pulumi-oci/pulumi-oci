// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetWaasCertificate
    {
        /// <summary>
        /// This data source provides details about a specific Certificate resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service.
        /// 
        /// Gets the details of an SSL certificate.
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
        ///         var testCertificate = Output.Create(Oci.GetWaasCertificate.InvokeAsync(new Oci.GetWaasCertificateArgs
        ///         {
        ///             CertificateId = oci_waas_certificate.Test_certificate.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetWaasCertificateResult> InvokeAsync(GetWaasCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWaasCertificateResult>("oci:index/getWaasCertificate:GetWaasCertificate", args ?? new GetWaasCertificateArgs(), options.WithVersion());
    }


    public sealed class GetWaasCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SSL certificate used in the WAAS policy. This number is generated when the certificate is added to the policy.
        /// </summary>
        [Input("certificateId", required: true)]
        public string CertificateId { get; set; } = null!;

        public GetWaasCertificateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWaasCertificateResult
    {
        /// <summary>
        /// The data of the SSL certificate.
        /// </summary>
        public readonly string CertificateData;
        public readonly string CertificateId;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SSL certificate's compartment.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// The user-friendly name of the SSL certificate.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Additional attributes associated with users or public keys for managing relationships between Certificate Authorities.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetWaasCertificateExtensionResult> Extensions;
        /// <summary>
        /// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SSL certificate.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// This indicates whether trust verification was disabled during the creation of SSL certificate. If `true` SSL certificate trust verification was disabled and this SSL certificate is most likely self-signed.
        /// </summary>
        public readonly bool IsTrustVerificationDisabled;
        public readonly string IssuedBy;
        /// <summary>
        /// The issuer of the certificate.
        /// </summary>
        public readonly Outputs.GetWaasCertificateIssuerNameResult IssuerName;
        public readonly string PrivateKeyData;
        /// <summary>
        /// Information about the public key and the algorithm used by the public key.
        /// </summary>
        public readonly Outputs.GetWaasCertificatePublicKeyInfoResult PublicKeyInfo;
        /// <summary>
        /// A unique, positive integer assigned by the Certificate Authority (CA). The issuer name and serial number identify a unique certificate.
        /// </summary>
        public readonly string SerialNumber;
        /// <summary>
        /// The identifier for the cryptographic algorithm used by the Certificate Authority (CA) to sign this certificate.
        /// </summary>
        public readonly string SignatureAlgorithm;
        /// <summary>
        /// The current lifecycle state of the SSL certificate.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The entity to be secured by the certificate.
        /// </summary>
        public readonly Outputs.GetWaasCertificateSubjectNameResult SubjectName;
        /// <summary>
        /// The date and time the certificate was created, expressed in RFC 3339 timestamp format.
        /// </summary>
        public readonly string TimeCreated;
        /// <summary>
        /// The date and time the certificate will expire, expressed in RFC 3339 timestamp format.
        /// </summary>
        public readonly string TimeNotValidAfter;
        /// <summary>
        /// The date and time the certificate will become valid, expressed in RFC 3339 timestamp format.
        /// </summary>
        public readonly string TimeNotValidBefore;
        /// <summary>
        /// The version of the encoded certificate.
        /// </summary>
        public readonly int Version;

        [OutputConstructor]
        private GetWaasCertificateResult(
            string certificateData,

            string certificateId,

            string compartmentId,

            ImmutableDictionary<string, object> definedTags,

            string displayName,

            ImmutableArray<Outputs.GetWaasCertificateExtensionResult> extensions,

            ImmutableDictionary<string, object> freeformTags,

            string id,

            bool isTrustVerificationDisabled,

            string issuedBy,

            Outputs.GetWaasCertificateIssuerNameResult issuerName,

            string privateKeyData,

            Outputs.GetWaasCertificatePublicKeyInfoResult publicKeyInfo,

            string serialNumber,

            string signatureAlgorithm,

            string state,

            Outputs.GetWaasCertificateSubjectNameResult subjectName,

            string timeCreated,

            string timeNotValidAfter,

            string timeNotValidBefore,

            int version)
        {
            CertificateData = certificateData;
            CertificateId = certificateId;
            CompartmentId = compartmentId;
            DefinedTags = definedTags;
            DisplayName = displayName;
            Extensions = extensions;
            FreeformTags = freeformTags;
            Id = id;
            IsTrustVerificationDisabled = isTrustVerificationDisabled;
            IssuedBy = issuedBy;
            IssuerName = issuerName;
            PrivateKeyData = privateKeyData;
            PublicKeyInfo = publicKeyInfo;
            SerialNumber = serialNumber;
            SignatureAlgorithm = signatureAlgorithm;
            State = state;
            SubjectName = subjectName;
            TimeCreated = timeCreated;
            TimeNotValidAfter = timeNotValidAfter;
            TimeNotValidBefore = timeNotValidBefore;
            Version = version;
        }
    }
}