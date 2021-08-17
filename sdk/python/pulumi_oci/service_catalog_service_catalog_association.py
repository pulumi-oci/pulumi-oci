# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServiceCatalogServiceCatalogAssociationArgs', 'ServiceCatalogServiceCatalogAssociation']

@pulumi.input_type
class ServiceCatalogServiceCatalogAssociationArgs:
    def __init__(__self__, *,
                 entity_id: pulumi.Input[str],
                 service_catalog_id: pulumi.Input[str],
                 entity_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceCatalogServiceCatalogAssociation resource.
        :param pulumi.Input[str] entity_id: Identifier of the entity being associated with service catalog.
        :param pulumi.Input[str] service_catalog_id: Identifier of the service catalog.
        :param pulumi.Input[str] entity_type: The type of the entity that is associated with the service catalog.
        """
        pulumi.set(__self__, "entity_id", entity_id)
        pulumi.set(__self__, "service_catalog_id", service_catalog_id)
        if entity_type is not None:
            pulumi.set(__self__, "entity_type", entity_type)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> pulumi.Input[str]:
        """
        Identifier of the entity being associated with service catalog.
        """
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter(name="serviceCatalogId")
    def service_catalog_id(self) -> pulumi.Input[str]:
        """
        Identifier of the service catalog.
        """
        return pulumi.get(self, "service_catalog_id")

    @service_catalog_id.setter
    def service_catalog_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_catalog_id", value)

    @property
    @pulumi.getter(name="entityType")
    def entity_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the entity that is associated with the service catalog.
        """
        return pulumi.get(self, "entity_type")

    @entity_type.setter
    def entity_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_type", value)


@pulumi.input_type
class _ServiceCatalogServiceCatalogAssociationState:
    def __init__(__self__, *,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 entity_type: Optional[pulumi.Input[str]] = None,
                 service_catalog_id: Optional[pulumi.Input[str]] = None,
                 time_created: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceCatalogServiceCatalogAssociation resources.
        :param pulumi.Input[str] entity_id: Identifier of the entity being associated with service catalog.
        :param pulumi.Input[str] entity_type: The type of the entity that is associated with the service catalog.
        :param pulumi.Input[str] service_catalog_id: Identifier of the service catalog.
        :param pulumi.Input[str] time_created: Timestamp of when the resource was associated with service catalog.
        """
        if entity_id is not None:
            pulumi.set(__self__, "entity_id", entity_id)
        if entity_type is not None:
            pulumi.set(__self__, "entity_type", entity_type)
        if service_catalog_id is not None:
            pulumi.set(__self__, "service_catalog_id", service_catalog_id)
        if time_created is not None:
            pulumi.set(__self__, "time_created", time_created)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the entity being associated with service catalog.
        """
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter(name="entityType")
    def entity_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the entity that is associated with the service catalog.
        """
        return pulumi.get(self, "entity_type")

    @entity_type.setter
    def entity_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_type", value)

    @property
    @pulumi.getter(name="serviceCatalogId")
    def service_catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the service catalog.
        """
        return pulumi.get(self, "service_catalog_id")

    @service_catalog_id.setter
    def service_catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_catalog_id", value)

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> Optional[pulumi.Input[str]]:
        """
        Timestamp of when the resource was associated with service catalog.
        """
        return pulumi.get(self, "time_created")

    @time_created.setter
    def time_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_created", value)


class ServiceCatalogServiceCatalogAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 entity_type: Optional[pulumi.Input[str]] = None,
                 service_catalog_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource provides the Service Catalog Association resource in Oracle Cloud Infrastructure Service Catalog service.

        Creates an association between service catalog and a resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_service_catalog_association = oci.ServiceCatalogServiceCatalogAssociation("testServiceCatalogAssociation",
            entity_id=oci_service_catalog_entity["test_entity"]["id"],
            service_catalog_id=oci_service_catalog_service_catalog["test_service_catalog"]["id"],
            entity_type=var["service_catalog_association_entity_type"])
        ```

        ## Import

        ServiceCatalogAssociations can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:index/serviceCatalogServiceCatalogAssociation:ServiceCatalogServiceCatalogAssociation test_service_catalog_association "id"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] entity_id: Identifier of the entity being associated with service catalog.
        :param pulumi.Input[str] entity_type: The type of the entity that is associated with the service catalog.
        :param pulumi.Input[str] service_catalog_id: Identifier of the service catalog.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceCatalogServiceCatalogAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource provides the Service Catalog Association resource in Oracle Cloud Infrastructure Service Catalog service.

        Creates an association between service catalog and a resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_service_catalog_association = oci.ServiceCatalogServiceCatalogAssociation("testServiceCatalogAssociation",
            entity_id=oci_service_catalog_entity["test_entity"]["id"],
            service_catalog_id=oci_service_catalog_service_catalog["test_service_catalog"]["id"],
            entity_type=var["service_catalog_association_entity_type"])
        ```

        ## Import

        ServiceCatalogAssociations can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:index/serviceCatalogServiceCatalogAssociation:ServiceCatalogServiceCatalogAssociation test_service_catalog_association "id"
        ```

        :param str resource_name: The name of the resource.
        :param ServiceCatalogServiceCatalogAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceCatalogServiceCatalogAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 entity_type: Optional[pulumi.Input[str]] = None,
                 service_catalog_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ServiceCatalogServiceCatalogAssociationArgs.__new__(ServiceCatalogServiceCatalogAssociationArgs)

            if entity_id is None and not opts.urn:
                raise TypeError("Missing required property 'entity_id'")
            __props__.__dict__["entity_id"] = entity_id
            __props__.__dict__["entity_type"] = entity_type
            if service_catalog_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_catalog_id'")
            __props__.__dict__["service_catalog_id"] = service_catalog_id
            __props__.__dict__["time_created"] = None
        super(ServiceCatalogServiceCatalogAssociation, __self__).__init__(
            'oci:index/serviceCatalogServiceCatalogAssociation:ServiceCatalogServiceCatalogAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            entity_id: Optional[pulumi.Input[str]] = None,
            entity_type: Optional[pulumi.Input[str]] = None,
            service_catalog_id: Optional[pulumi.Input[str]] = None,
            time_created: Optional[pulumi.Input[str]] = None) -> 'ServiceCatalogServiceCatalogAssociation':
        """
        Get an existing ServiceCatalogServiceCatalogAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] entity_id: Identifier of the entity being associated with service catalog.
        :param pulumi.Input[str] entity_type: The type of the entity that is associated with the service catalog.
        :param pulumi.Input[str] service_catalog_id: Identifier of the service catalog.
        :param pulumi.Input[str] time_created: Timestamp of when the resource was associated with service catalog.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceCatalogServiceCatalogAssociationState.__new__(_ServiceCatalogServiceCatalogAssociationState)

        __props__.__dict__["entity_id"] = entity_id
        __props__.__dict__["entity_type"] = entity_type
        __props__.__dict__["service_catalog_id"] = service_catalog_id
        __props__.__dict__["time_created"] = time_created
        return ServiceCatalogServiceCatalogAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> pulumi.Output[str]:
        """
        Identifier of the entity being associated with service catalog.
        """
        return pulumi.get(self, "entity_id")

    @property
    @pulumi.getter(name="entityType")
    def entity_type(self) -> pulumi.Output[str]:
        """
        The type of the entity that is associated with the service catalog.
        """
        return pulumi.get(self, "entity_type")

    @property
    @pulumi.getter(name="serviceCatalogId")
    def service_catalog_id(self) -> pulumi.Output[str]:
        """
        Identifier of the service catalog.
        """
        return pulumi.get(self, "service_catalog_id")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> pulumi.Output[str]:
        """
        Timestamp of when the resource was associated with service catalog.
        """
        return pulumi.get(self, "time_created")
