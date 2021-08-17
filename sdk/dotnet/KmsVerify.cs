// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    /// <summary>
    /// This resource provides the Verify resource in Oracle Cloud Infrastructure Kms service.
    /// 
    /// Verifies a digital signature that was generated by the [Sign](https://docs.cloud.oracle.com/iaas/api/#/en/key/latest/SignedData/Sign) operation
    /// by using the public key of the same asymmetric key that was used to sign the data. If you want to validate the\
    /// digital signature outside of the service, you can do so by using the public key of the asymmetric key.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Oci = Pulumi.Oci;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testVerify = new Oci.KmsVerify("testVerify", new Oci.KmsVerifyArgs
    ///         {
    ///             CryptoEndpoint = @var.Verify_message_crypto_endpoint,
    ///             KeyId = oci_kms_key.Test_key.Id,
    ///             KeyVersionId = oci_kms_key_version.Test_key_version.Id,
    ///             Message = @var.Verify_message,
    ///             Signature = @var.Verify_signature,
    ///             SigningAlgorithm = @var.Verify_signing_algorithm,
    ///             MessageType = @var.Verify_message_type,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Verify can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import oci:index/kmsVerify:KmsVerify test_verify "id"
    /// ```
    /// </summary>
    [OciResourceType("oci:index/kmsVerify:KmsVerify")]
    public partial class KmsVerify : Pulumi.CustomResource
    {
        /// <summary>
        /// The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,', 'GenerateDataEncryptionKey', 'Sign' and 'Verify' operations. see Vault Crypto endpoint.
        /// </summary>
        [Output("cryptoEndpoint")]
        public Output<string> CryptoEndpoint { get; private set; } = null!;

        /// <summary>
        /// A Boolean value that indicates whether the signature was verified.
        /// </summary>
        [Output("isSignatureValid")]
        public Output<bool> IsSignatureValid { get; private set; } = null!;

        /// <summary>
        /// The OCID of the key used to sign the message.
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        /// <summary>
        /// The OCID of the key version used to sign the message.
        /// </summary>
        [Output("keyVersionId")]
        public Output<string> KeyVersionId { get; private set; } = null!;

        /// <summary>
        /// The base64-encoded binary data object denoting the message or message digest to sign. You can have a message up to 4096 bytes in size. To sign a larger message, provide the message digest.
        /// </summary>
        [Output("message")]
        public Output<string> Message { get; private set; } = null!;

        /// <summary>
        /// Denotes whether the value of the message parameter is a raw message or a message digest.  The default value, `RAW`, indicates a message. To indicate a message digest, use `DIGEST`.
        /// </summary>
        [Output("messageType")]
        public Output<string> MessageType { get; private set; } = null!;

        /// <summary>
        /// The base64-encoded binary data object denoting the cryptographic signature generated for the message.
        /// </summary>
        [Output("signature")]
        public Output<string> Signature { get; private set; } = null!;

        /// <summary>
        /// The algorithm to use to sign the message or message digest. For RSA keys, supported signature schemes include PKCS #1 and RSASSA-PSS, along with  different hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. When you pass a message digest for signing, ensure that you specify the same hashing algorithm  as used when creating the message digest.
        /// </summary>
        [Output("signingAlgorithm")]
        public Output<string> SigningAlgorithm { get; private set; } = null!;


        /// <summary>
        /// Create a KmsVerify resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KmsVerify(string name, KmsVerifyArgs args, CustomResourceOptions? options = null)
            : base("oci:index/kmsVerify:KmsVerify", name, args ?? new KmsVerifyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KmsVerify(string name, Input<string> id, KmsVerifyState? state = null, CustomResourceOptions? options = null)
            : base("oci:index/kmsVerify:KmsVerify", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KmsVerify resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KmsVerify Get(string name, Input<string> id, KmsVerifyState? state = null, CustomResourceOptions? options = null)
        {
            return new KmsVerify(name, id, state, options);
        }
    }

    public sealed class KmsVerifyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,', 'GenerateDataEncryptionKey', 'Sign' and 'Verify' operations. see Vault Crypto endpoint.
        /// </summary>
        [Input("cryptoEndpoint", required: true)]
        public Input<string> CryptoEndpoint { get; set; } = null!;

        /// <summary>
        /// The OCID of the key used to sign the message.
        /// </summary>
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        /// <summary>
        /// The OCID of the key version used to sign the message.
        /// </summary>
        [Input("keyVersionId", required: true)]
        public Input<string> KeyVersionId { get; set; } = null!;

        /// <summary>
        /// The base64-encoded binary data object denoting the message or message digest to sign. You can have a message up to 4096 bytes in size. To sign a larger message, provide the message digest.
        /// </summary>
        [Input("message", required: true)]
        public Input<string> Message { get; set; } = null!;

        /// <summary>
        /// Denotes whether the value of the message parameter is a raw message or a message digest.  The default value, `RAW`, indicates a message. To indicate a message digest, use `DIGEST`.
        /// </summary>
        [Input("messageType")]
        public Input<string>? MessageType { get; set; }

        /// <summary>
        /// The base64-encoded binary data object denoting the cryptographic signature generated for the message.
        /// </summary>
        [Input("signature", required: true)]
        public Input<string> Signature { get; set; } = null!;

        /// <summary>
        /// The algorithm to use to sign the message or message digest. For RSA keys, supported signature schemes include PKCS #1 and RSASSA-PSS, along with  different hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. When you pass a message digest for signing, ensure that you specify the same hashing algorithm  as used when creating the message digest.
        /// </summary>
        [Input("signingAlgorithm", required: true)]
        public Input<string> SigningAlgorithm { get; set; } = null!;

        public KmsVerifyArgs()
        {
        }
    }

    public sealed class KmsVerifyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,', 'GenerateDataEncryptionKey', 'Sign' and 'Verify' operations. see Vault Crypto endpoint.
        /// </summary>
        [Input("cryptoEndpoint")]
        public Input<string>? CryptoEndpoint { get; set; }

        /// <summary>
        /// A Boolean value that indicates whether the signature was verified.
        /// </summary>
        [Input("isSignatureValid")]
        public Input<bool>? IsSignatureValid { get; set; }

        /// <summary>
        /// The OCID of the key used to sign the message.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The OCID of the key version used to sign the message.
        /// </summary>
        [Input("keyVersionId")]
        public Input<string>? KeyVersionId { get; set; }

        /// <summary>
        /// The base64-encoded binary data object denoting the message or message digest to sign. You can have a message up to 4096 bytes in size. To sign a larger message, provide the message digest.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Denotes whether the value of the message parameter is a raw message or a message digest.  The default value, `RAW`, indicates a message. To indicate a message digest, use `DIGEST`.
        /// </summary>
        [Input("messageType")]
        public Input<string>? MessageType { get; set; }

        /// <summary>
        /// The base64-encoded binary data object denoting the cryptographic signature generated for the message.
        /// </summary>
        [Input("signature")]
        public Input<string>? Signature { get; set; }

        /// <summary>
        /// The algorithm to use to sign the message or message digest. For RSA keys, supported signature schemes include PKCS #1 and RSASSA-PSS, along with  different hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. When you pass a message digest for signing, ensure that you specify the same hashing algorithm  as used when creating the message digest.
        /// </summary>
        [Input("signingAlgorithm")]
        public Input<string>? SigningAlgorithm { get; set; }

        public KmsVerifyState()
        {
        }
    }
}