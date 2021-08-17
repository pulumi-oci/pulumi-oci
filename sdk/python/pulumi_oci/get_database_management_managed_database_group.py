# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetDatabaseManagementManagedDatabaseGroupResult',
    'AwaitableGetDatabaseManagementManagedDatabaseGroupResult',
    'get_database_management_managed_database_group',
]

@pulumi.output_type
class GetDatabaseManagementManagedDatabaseGroupResult:
    """
    A collection of values returned by GetDatabaseManagementManagedDatabaseGroup.
    """
    def __init__(__self__, compartment_id=None, description=None, id=None, managed_database_group_id=None, managed_databases=None, name=None, state=None, time_created=None, time_updated=None):
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if managed_database_group_id and not isinstance(managed_database_group_id, str):
            raise TypeError("Expected argument 'managed_database_group_id' to be a str")
        pulumi.set(__self__, "managed_database_group_id", managed_database_group_id)
        if managed_databases and not isinstance(managed_databases, list):
            raise TypeError("Expected argument 'managed_databases' to be a list")
        pulumi.set(__self__, "managed_databases", managed_databases)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if time_created and not isinstance(time_created, str):
            raise TypeError("Expected argument 'time_created' to be a str")
        pulumi.set(__self__, "time_created", time_created)
        if time_updated and not isinstance(time_updated, str):
            raise TypeError("Expected argument 'time_updated' to be a str")
        pulumi.set(__self__, "time_updated", time_updated)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> str:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database resides.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The information specified by the user about the Managed Database Group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="managedDatabaseGroupId")
    def managed_database_group_id(self) -> str:
        return pulumi.get(self, "managed_database_group_id")

    @property
    @pulumi.getter(name="managedDatabases")
    def managed_databases(self) -> Sequence['outputs.GetDatabaseManagementManagedDatabaseGroupManagedDatabaseResult']:
        """
        A list of Managed Databases in the Managed Database Group.
        """
        return pulumi.get(self, "managed_databases")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Managed Database Group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current lifecycle state of the Managed Database Group.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        The date and time the Managed Database Group was created.
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter(name="timeUpdated")
    def time_updated(self) -> str:
        """
        The date and time the Managed Database Group was last updated.
        """
        return pulumi.get(self, "time_updated")


class AwaitableGetDatabaseManagementManagedDatabaseGroupResult(GetDatabaseManagementManagedDatabaseGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseManagementManagedDatabaseGroupResult(
            compartment_id=self.compartment_id,
            description=self.description,
            id=self.id,
            managed_database_group_id=self.managed_database_group_id,
            managed_databases=self.managed_databases,
            name=self.name,
            state=self.state,
            time_created=self.time_created,
            time_updated=self.time_updated)


def get_database_management_managed_database_group(managed_database_group_id: Optional[str] = None,
                                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseManagementManagedDatabaseGroupResult:
    """
    This data source provides details about a specific Managed Database Group resource in Oracle Cloud Infrastructure Database Management service.

    Gets the details for the Managed Database Group specified by managedDatabaseGroupId.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_managed_database_group = oci.get_database_management_managed_database_group(managed_database_group_id=oci_database_management_managed_database_group["test_managed_database_group"]["id"])
    ```


    :param str managed_database_group_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
    """
    __args__ = dict()
    __args__['managedDatabaseGroupId'] = managed_database_group_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:index/getDatabaseManagementManagedDatabaseGroup:GetDatabaseManagementManagedDatabaseGroup', __args__, opts=opts, typ=GetDatabaseManagementManagedDatabaseGroupResult).value

    return AwaitableGetDatabaseManagementManagedDatabaseGroupResult(
        compartment_id=__ret__.compartment_id,
        description=__ret__.description,
        id=__ret__.id,
        managed_database_group_id=__ret__.managed_database_group_id,
        managed_databases=__ret__.managed_databases,
        name=__ret__.name,
        state=__ret__.state,
        time_created=__ret__.time_created,
        time_updated=__ret__.time_updated)