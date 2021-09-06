# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SignArgs', 'Sign']

@pulumi.input_type
class SignArgs:
    def __init__(__self__, *,
                 crypto_endpoint: pulumi.Input[str],
                 key_id: pulumi.Input[str],
                 message: pulumi.Input[str],
                 signing_algorithm: pulumi.Input[str],
                 key_version_id: Optional[pulumi.Input[str]] = None,
                 message_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Sign resource.
        :param pulumi.Input[str] crypto_endpoint: The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,', 'GenerateDataEncryptionKey', 'Sign' and 'Verify' operations. see Vault Crypto endpoint.
        :param pulumi.Input[str] key_id: The OCID of the key used to sign the message.
        :param pulumi.Input[str] message: The base64-encoded binary data object denoting the message or message digest to sign. You can have a message up to 4096 bytes in size. To sign a larger message, provide the message digest.
        :param pulumi.Input[str] signing_algorithm: The algorithm to use to sign the message or message digest. For RSA keys, supported signature schemes include PKCS #1 and RSASSA-PSS, along with  different hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. When you pass a message digest for signing, ensure that you specify the same hashing algorithm  as used when creating the message digest.
        :param pulumi.Input[str] key_version_id: The OCID of the key version used to sign the message.
        :param pulumi.Input[str] message_type: Denotes whether the value of the message parameter is a raw message or a message digest.  The default value, `RAW`, indicates a message. To indicate a message digest, use `DIGEST`.
        """
        pulumi.set(__self__, "crypto_endpoint", crypto_endpoint)
        pulumi.set(__self__, "key_id", key_id)
        pulumi.set(__self__, "message", message)
        pulumi.set(__self__, "signing_algorithm", signing_algorithm)
        if key_version_id is not None:
            pulumi.set(__self__, "key_version_id", key_version_id)
        if message_type is not None:
            pulumi.set(__self__, "message_type", message_type)

    @property
    @pulumi.getter(name="cryptoEndpoint")
    def crypto_endpoint(self) -> pulumi.Input[str]:
        """
        The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,', 'GenerateDataEncryptionKey', 'Sign' and 'Verify' operations. see Vault Crypto endpoint.
        """
        return pulumi.get(self, "crypto_endpoint")

    @crypto_endpoint.setter
    def crypto_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "crypto_endpoint", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Input[str]:
        """
        The OCID of the key used to sign the message.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter
    def message(self) -> pulumi.Input[str]:
        """
        The base64-encoded binary data object denoting the message or message digest to sign. You can have a message up to 4096 bytes in size. To sign a larger message, provide the message digest.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: pulumi.Input[str]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter(name="signingAlgorithm")
    def signing_algorithm(self) -> pulumi.Input[str]:
        """
        The algorithm to use to sign the message or message digest. For RSA keys, supported signature schemes include PKCS #1 and RSASSA-PSS, along with  different hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. When you pass a message digest for signing, ensure that you specify the same hashing algorithm  as used when creating the message digest.
        """
        return pulumi.get(self, "signing_algorithm")

    @signing_algorithm.setter
    def signing_algorithm(self, value: pulumi.Input[str]):
        pulumi.set(self, "signing_algorithm", value)

    @property
    @pulumi.getter(name="keyVersionId")
    def key_version_id(self) -> Optional[pulumi.Input[str]]:
        """
        The OCID of the key version used to sign the message.
        """
        return pulumi.get(self, "key_version_id")

    @key_version_id.setter
    def key_version_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_version_id", value)

    @property
    @pulumi.getter(name="messageType")
    def message_type(self) -> Optional[pulumi.Input[str]]:
        """
        Denotes whether the value of the message parameter is a raw message or a message digest.  The default value, `RAW`, indicates a message. To indicate a message digest, use `DIGEST`.
        """
        return pulumi.get(self, "message_type")

    @message_type.setter
    def message_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message_type", value)


@pulumi.input_type
class _SignState:
    def __init__(__self__, *,
                 crypto_endpoint: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 key_version_id: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 message_type: Optional[pulumi.Input[str]] = None,
                 signature: Optional[pulumi.Input[str]] = None,
                 signing_algorithm: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Sign resources.
        :param pulumi.Input[str] crypto_endpoint: The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,', 'GenerateDataEncryptionKey', 'Sign' and 'Verify' operations. see Vault Crypto endpoint.
        :param pulumi.Input[str] key_id: The OCID of the key used to sign the message.
        :param pulumi.Input[str] key_version_id: The OCID of the key version used to sign the message.
        :param pulumi.Input[str] message: The base64-encoded binary data object denoting the message or message digest to sign. You can have a message up to 4096 bytes in size. To sign a larger message, provide the message digest.
        :param pulumi.Input[str] message_type: Denotes whether the value of the message parameter is a raw message or a message digest.  The default value, `RAW`, indicates a message. To indicate a message digest, use `DIGEST`.
        :param pulumi.Input[str] signature: The base64-encoded binary data object denoting the cryptographic signature generated for the message or message digest.
        :param pulumi.Input[str] signing_algorithm: The algorithm to use to sign the message or message digest. For RSA keys, supported signature schemes include PKCS #1 and RSASSA-PSS, along with  different hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. When you pass a message digest for signing, ensure that you specify the same hashing algorithm  as used when creating the message digest.
        """
        if crypto_endpoint is not None:
            pulumi.set(__self__, "crypto_endpoint", crypto_endpoint)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if key_version_id is not None:
            pulumi.set(__self__, "key_version_id", key_version_id)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if message_type is not None:
            pulumi.set(__self__, "message_type", message_type)
        if signature is not None:
            pulumi.set(__self__, "signature", signature)
        if signing_algorithm is not None:
            pulumi.set(__self__, "signing_algorithm", signing_algorithm)

    @property
    @pulumi.getter(name="cryptoEndpoint")
    def crypto_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,', 'GenerateDataEncryptionKey', 'Sign' and 'Verify' operations. see Vault Crypto endpoint.
        """
        return pulumi.get(self, "crypto_endpoint")

    @crypto_endpoint.setter
    def crypto_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "crypto_endpoint", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The OCID of the key used to sign the message.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="keyVersionId")
    def key_version_id(self) -> Optional[pulumi.Input[str]]:
        """
        The OCID of the key version used to sign the message.
        """
        return pulumi.get(self, "key_version_id")

    @key_version_id.setter
    def key_version_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_version_id", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        """
        The base64-encoded binary data object denoting the message or message digest to sign. You can have a message up to 4096 bytes in size. To sign a larger message, provide the message digest.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter(name="messageType")
    def message_type(self) -> Optional[pulumi.Input[str]]:
        """
        Denotes whether the value of the message parameter is a raw message or a message digest.  The default value, `RAW`, indicates a message. To indicate a message digest, use `DIGEST`.
        """
        return pulumi.get(self, "message_type")

    @message_type.setter
    def message_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message_type", value)

    @property
    @pulumi.getter
    def signature(self) -> Optional[pulumi.Input[str]]:
        """
        The base64-encoded binary data object denoting the cryptographic signature generated for the message or message digest.
        """
        return pulumi.get(self, "signature")

    @signature.setter
    def signature(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signature", value)

    @property
    @pulumi.getter(name="signingAlgorithm")
    def signing_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The algorithm to use to sign the message or message digest. For RSA keys, supported signature schemes include PKCS #1 and RSASSA-PSS, along with  different hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. When you pass a message digest for signing, ensure that you specify the same hashing algorithm  as used when creating the message digest.
        """
        return pulumi.get(self, "signing_algorithm")

    @signing_algorithm.setter
    def signing_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signing_algorithm", value)


class Sign(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 crypto_endpoint: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 key_version_id: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 message_type: Optional[pulumi.Input[str]] = None,
                 signing_algorithm: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource provides the Sign resource in Oracle Cloud Infrastructure Kms service.

        Creates a digital signature for a message or message digest by using the private key of a public-private key pair,
        also known as an asymmetric key. To verify the generated signature, you can use the [Verify](https://docs.cloud.oracle.com/iaas/api/#/en/key/latest/VerifiedData/Verify)
        operation. Or, if you want to validate the signature outside of the service, you can do so by using the public key of the same asymmetric key.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_sign = oci.kms.Sign("testSign",
            crypto_endpoint=var["sign_message_crypto_endpoint"],
            key_id=oci_kms_key["test_key"]["id"],
            message=var["sign_message"],
            signing_algorithm=var["sign_signing_algorithm"],
            key_version_id=oci_kms_key_version["test_key_version"]["id"],
            message_type=var["sign_message_type"])
        ```

        ## Import

        Sign can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:kms/sign:Sign test_sign "id"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] crypto_endpoint: The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,', 'GenerateDataEncryptionKey', 'Sign' and 'Verify' operations. see Vault Crypto endpoint.
        :param pulumi.Input[str] key_id: The OCID of the key used to sign the message.
        :param pulumi.Input[str] key_version_id: The OCID of the key version used to sign the message.
        :param pulumi.Input[str] message: The base64-encoded binary data object denoting the message or message digest to sign. You can have a message up to 4096 bytes in size. To sign a larger message, provide the message digest.
        :param pulumi.Input[str] message_type: Denotes whether the value of the message parameter is a raw message or a message digest.  The default value, `RAW`, indicates a message. To indicate a message digest, use `DIGEST`.
        :param pulumi.Input[str] signing_algorithm: The algorithm to use to sign the message or message digest. For RSA keys, supported signature schemes include PKCS #1 and RSASSA-PSS, along with  different hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. When you pass a message digest for signing, ensure that you specify the same hashing algorithm  as used when creating the message digest.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SignArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource provides the Sign resource in Oracle Cloud Infrastructure Kms service.

        Creates a digital signature for a message or message digest by using the private key of a public-private key pair,
        also known as an asymmetric key. To verify the generated signature, you can use the [Verify](https://docs.cloud.oracle.com/iaas/api/#/en/key/latest/VerifiedData/Verify)
        operation. Or, if you want to validate the signature outside of the service, you can do so by using the public key of the same asymmetric key.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_sign = oci.kms.Sign("testSign",
            crypto_endpoint=var["sign_message_crypto_endpoint"],
            key_id=oci_kms_key["test_key"]["id"],
            message=var["sign_message"],
            signing_algorithm=var["sign_signing_algorithm"],
            key_version_id=oci_kms_key_version["test_key_version"]["id"],
            message_type=var["sign_message_type"])
        ```

        ## Import

        Sign can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:kms/sign:Sign test_sign "id"
        ```

        :param str resource_name: The name of the resource.
        :param SignArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SignArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 crypto_endpoint: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 key_version_id: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None,
                 message_type: Optional[pulumi.Input[str]] = None,
                 signing_algorithm: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SignArgs.__new__(SignArgs)

            if crypto_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'crypto_endpoint'")
            __props__.__dict__["crypto_endpoint"] = crypto_endpoint
            if key_id is None and not opts.urn:
                raise TypeError("Missing required property 'key_id'")
            __props__.__dict__["key_id"] = key_id
            __props__.__dict__["key_version_id"] = key_version_id
            if message is None and not opts.urn:
                raise TypeError("Missing required property 'message'")
            __props__.__dict__["message"] = message
            __props__.__dict__["message_type"] = message_type
            if signing_algorithm is None and not opts.urn:
                raise TypeError("Missing required property 'signing_algorithm'")
            __props__.__dict__["signing_algorithm"] = signing_algorithm
            __props__.__dict__["signature"] = None
        super(Sign, __self__).__init__(
            'oci:kms/sign:Sign',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            crypto_endpoint: Optional[pulumi.Input[str]] = None,
            key_id: Optional[pulumi.Input[str]] = None,
            key_version_id: Optional[pulumi.Input[str]] = None,
            message: Optional[pulumi.Input[str]] = None,
            message_type: Optional[pulumi.Input[str]] = None,
            signature: Optional[pulumi.Input[str]] = None,
            signing_algorithm: Optional[pulumi.Input[str]] = None) -> 'Sign':
        """
        Get an existing Sign resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] crypto_endpoint: The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,', 'GenerateDataEncryptionKey', 'Sign' and 'Verify' operations. see Vault Crypto endpoint.
        :param pulumi.Input[str] key_id: The OCID of the key used to sign the message.
        :param pulumi.Input[str] key_version_id: The OCID of the key version used to sign the message.
        :param pulumi.Input[str] message: The base64-encoded binary data object denoting the message or message digest to sign. You can have a message up to 4096 bytes in size. To sign a larger message, provide the message digest.
        :param pulumi.Input[str] message_type: Denotes whether the value of the message parameter is a raw message or a message digest.  The default value, `RAW`, indicates a message. To indicate a message digest, use `DIGEST`.
        :param pulumi.Input[str] signature: The base64-encoded binary data object denoting the cryptographic signature generated for the message or message digest.
        :param pulumi.Input[str] signing_algorithm: The algorithm to use to sign the message or message digest. For RSA keys, supported signature schemes include PKCS #1 and RSASSA-PSS, along with  different hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. When you pass a message digest for signing, ensure that you specify the same hashing algorithm  as used when creating the message digest.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SignState.__new__(_SignState)

        __props__.__dict__["crypto_endpoint"] = crypto_endpoint
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["key_version_id"] = key_version_id
        __props__.__dict__["message"] = message
        __props__.__dict__["message_type"] = message_type
        __props__.__dict__["signature"] = signature
        __props__.__dict__["signing_algorithm"] = signing_algorithm
        return Sign(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cryptoEndpoint")
    def crypto_endpoint(self) -> pulumi.Output[str]:
        """
        The service endpoint to perform cryptographic operations against. Cryptographic operations include 'Encrypt,' 'Decrypt,', 'GenerateDataEncryptionKey', 'Sign' and 'Verify' operations. see Vault Crypto endpoint.
        """
        return pulumi.get(self, "crypto_endpoint")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        """
        The OCID of the key used to sign the message.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="keyVersionId")
    def key_version_id(self) -> pulumi.Output[str]:
        """
        The OCID of the key version used to sign the message.
        """
        return pulumi.get(self, "key_version_id")

    @property
    @pulumi.getter
    def message(self) -> pulumi.Output[str]:
        """
        The base64-encoded binary data object denoting the message or message digest to sign. You can have a message up to 4096 bytes in size. To sign a larger message, provide the message digest.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter(name="messageType")
    def message_type(self) -> pulumi.Output[str]:
        """
        Denotes whether the value of the message parameter is a raw message or a message digest.  The default value, `RAW`, indicates a message. To indicate a message digest, use `DIGEST`.
        """
        return pulumi.get(self, "message_type")

    @property
    @pulumi.getter
    def signature(self) -> pulumi.Output[str]:
        """
        The base64-encoded binary data object denoting the cryptographic signature generated for the message or message digest.
        """
        return pulumi.get(self, "signature")

    @property
    @pulumi.getter(name="signingAlgorithm")
    def signing_algorithm(self) -> pulumi.Output[str]:
        """
        The algorithm to use to sign the message or message digest. For RSA keys, supported signature schemes include PKCS #1 and RSASSA-PSS, along with  different hashing algorithms.  For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms. When you pass a message digest for signing, ensure that you specify the same hashing algorithm  as used when creating the message digest.
        """
        return pulumi.get(self, "signing_algorithm")
