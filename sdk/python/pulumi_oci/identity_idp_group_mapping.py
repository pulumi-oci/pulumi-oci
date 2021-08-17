# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IdentityIdpGroupMappingArgs', 'IdentityIdpGroupMapping']

@pulumi.input_type
class IdentityIdpGroupMappingArgs:
    def __init__(__self__, *,
                 group_id: pulumi.Input[str],
                 identity_provider_id: pulumi.Input[str],
                 idp_group_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a IdentityIdpGroupMapping resource.
        :param pulumi.Input[str] group_id: (Updatable) The OCID of the IAM Service [group](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/Group/) you want to map to the IdP group.
        :param pulumi.Input[str] identity_provider_id: The OCID of the identity provider.
        :param pulumi.Input[str] idp_group_name: (Updatable) The name of the IdP group you want to map.
        """
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "identity_provider_id", identity_provider_id)
        pulumi.set(__self__, "idp_group_name", idp_group_name)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        (Updatable) The OCID of the IAM Service [group](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/Group/) you want to map to the IdP group.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="identityProviderId")
    def identity_provider_id(self) -> pulumi.Input[str]:
        """
        The OCID of the identity provider.
        """
        return pulumi.get(self, "identity_provider_id")

    @identity_provider_id.setter
    def identity_provider_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_provider_id", value)

    @property
    @pulumi.getter(name="idpGroupName")
    def idp_group_name(self) -> pulumi.Input[str]:
        """
        (Updatable) The name of the IdP group you want to map.
        """
        return pulumi.get(self, "idp_group_name")

    @idp_group_name.setter
    def idp_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "idp_group_name", value)


@pulumi.input_type
class _IdentityIdpGroupMappingState:
    def __init__(__self__, *,
                 compartment_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 identity_provider_id: Optional[pulumi.Input[str]] = None,
                 idp_group_name: Optional[pulumi.Input[str]] = None,
                 inactive_state: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 time_created: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IdentityIdpGroupMapping resources.
        :param pulumi.Input[str] compartment_id: The OCID of the tenancy containing the `IdentityProvider`.
        :param pulumi.Input[str] group_id: (Updatable) The OCID of the IAM Service [group](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/Group/) you want to map to the IdP group.
        :param pulumi.Input[str] identity_provider_id: The OCID of the identity provider.
        :param pulumi.Input[str] idp_group_name: (Updatable) The name of the IdP group you want to map.
        :param pulumi.Input[str] inactive_state: The detailed status of INACTIVE lifecycleState.
        :param pulumi.Input[str] state: The mapping's current state.
        :param pulumi.Input[str] time_created: Date and time the mapping was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
        """
        if compartment_id is not None:
            pulumi.set(__self__, "compartment_id", compartment_id)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if identity_provider_id is not None:
            pulumi.set(__self__, "identity_provider_id", identity_provider_id)
        if idp_group_name is not None:
            pulumi.set(__self__, "idp_group_name", idp_group_name)
        if inactive_state is not None:
            pulumi.set(__self__, "inactive_state", inactive_state)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if time_created is not None:
            pulumi.set(__self__, "time_created", time_created)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The OCID of the tenancy containing the `IdentityProvider`.
        """
        return pulumi.get(self, "compartment_id")

    @compartment_id.setter
    def compartment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compartment_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        (Updatable) The OCID of the IAM Service [group](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/Group/) you want to map to the IdP group.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="identityProviderId")
    def identity_provider_id(self) -> Optional[pulumi.Input[str]]:
        """
        The OCID of the identity provider.
        """
        return pulumi.get(self, "identity_provider_id")

    @identity_provider_id.setter
    def identity_provider_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_provider_id", value)

    @property
    @pulumi.getter(name="idpGroupName")
    def idp_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        (Updatable) The name of the IdP group you want to map.
        """
        return pulumi.get(self, "idp_group_name")

    @idp_group_name.setter
    def idp_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idp_group_name", value)

    @property
    @pulumi.getter(name="inactiveState")
    def inactive_state(self) -> Optional[pulumi.Input[str]]:
        """
        The detailed status of INACTIVE lifecycleState.
        """
        return pulumi.get(self, "inactive_state")

    @inactive_state.setter
    def inactive_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "inactive_state", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The mapping's current state.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time the mapping was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
        """
        return pulumi.get(self, "time_created")

    @time_created.setter
    def time_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_created", value)


class IdentityIdpGroupMapping(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 identity_provider_id: Optional[pulumi.Input[str]] = None,
                 idp_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource provides the Idp Group Mapping resource in Oracle Cloud Infrastructure Identity service.

        Creates a single mapping between an IdP group and an IAM Service
        [group](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/Group/).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_idp_group_mapping = oci.IdentityIdpGroupMapping("testIdpGroupMapping",
            group_id=oci_identity_group["test_group"]["id"],
            identity_provider_id=oci_identity_identity_provider["test_identity_provider"]["id"],
            idp_group_name=var["idp_group_mapping_idp_group_name"])
        ```

        ## Import

        IdpGroupMappings can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:index/identityIdpGroupMapping:IdentityIdpGroupMapping test_idp_group_mapping "identityProviders/{identityProviderId}/groupMappings/{mappingId}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_id: (Updatable) The OCID of the IAM Service [group](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/Group/) you want to map to the IdP group.
        :param pulumi.Input[str] identity_provider_id: The OCID of the identity provider.
        :param pulumi.Input[str] idp_group_name: (Updatable) The name of the IdP group you want to map.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdentityIdpGroupMappingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource provides the Idp Group Mapping resource in Oracle Cloud Infrastructure Identity service.

        Creates a single mapping between an IdP group and an IAM Service
        [group](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/Group/).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_idp_group_mapping = oci.IdentityIdpGroupMapping("testIdpGroupMapping",
            group_id=oci_identity_group["test_group"]["id"],
            identity_provider_id=oci_identity_identity_provider["test_identity_provider"]["id"],
            idp_group_name=var["idp_group_mapping_idp_group_name"])
        ```

        ## Import

        IdpGroupMappings can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:index/identityIdpGroupMapping:IdentityIdpGroupMapping test_idp_group_mapping "identityProviders/{identityProviderId}/groupMappings/{mappingId}"
        ```

        :param str resource_name: The name of the resource.
        :param IdentityIdpGroupMappingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdentityIdpGroupMappingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 identity_provider_id: Optional[pulumi.Input[str]] = None,
                 idp_group_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = IdentityIdpGroupMappingArgs.__new__(IdentityIdpGroupMappingArgs)

            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            if identity_provider_id is None and not opts.urn:
                raise TypeError("Missing required property 'identity_provider_id'")
            __props__.__dict__["identity_provider_id"] = identity_provider_id
            if idp_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'idp_group_name'")
            __props__.__dict__["idp_group_name"] = idp_group_name
            __props__.__dict__["compartment_id"] = None
            __props__.__dict__["inactive_state"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["time_created"] = None
        super(IdentityIdpGroupMapping, __self__).__init__(
            'oci:index/identityIdpGroupMapping:IdentityIdpGroupMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compartment_id: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            identity_provider_id: Optional[pulumi.Input[str]] = None,
            idp_group_name: Optional[pulumi.Input[str]] = None,
            inactive_state: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            time_created: Optional[pulumi.Input[str]] = None) -> 'IdentityIdpGroupMapping':
        """
        Get an existing IdentityIdpGroupMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compartment_id: The OCID of the tenancy containing the `IdentityProvider`.
        :param pulumi.Input[str] group_id: (Updatable) The OCID of the IAM Service [group](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/Group/) you want to map to the IdP group.
        :param pulumi.Input[str] identity_provider_id: The OCID of the identity provider.
        :param pulumi.Input[str] idp_group_name: (Updatable) The name of the IdP group you want to map.
        :param pulumi.Input[str] inactive_state: The detailed status of INACTIVE lifecycleState.
        :param pulumi.Input[str] state: The mapping's current state.
        :param pulumi.Input[str] time_created: Date and time the mapping was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdentityIdpGroupMappingState.__new__(_IdentityIdpGroupMappingState)

        __props__.__dict__["compartment_id"] = compartment_id
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["identity_provider_id"] = identity_provider_id
        __props__.__dict__["idp_group_name"] = idp_group_name
        __props__.__dict__["inactive_state"] = inactive_state
        __props__.__dict__["state"] = state
        __props__.__dict__["time_created"] = time_created
        return IdentityIdpGroupMapping(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> pulumi.Output[str]:
        """
        The OCID of the tenancy containing the `IdentityProvider`.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        (Updatable) The OCID of the IAM Service [group](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/Group/) you want to map to the IdP group.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="identityProviderId")
    def identity_provider_id(self) -> pulumi.Output[str]:
        """
        The OCID of the identity provider.
        """
        return pulumi.get(self, "identity_provider_id")

    @property
    @pulumi.getter(name="idpGroupName")
    def idp_group_name(self) -> pulumi.Output[str]:
        """
        (Updatable) The name of the IdP group you want to map.
        """
        return pulumi.get(self, "idp_group_name")

    @property
    @pulumi.getter(name="inactiveState")
    def inactive_state(self) -> pulumi.Output[str]:
        """
        The detailed status of INACTIVE lifecycleState.
        """
        return pulumi.get(self, "inactive_state")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The mapping's current state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> pulumi.Output[str]:
        """
        Date and time the mapping was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
        """
        return pulumi.get(self, "time_created")
