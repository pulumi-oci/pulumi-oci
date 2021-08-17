# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetFileStorageExportsResult',
    'AwaitableGetFileStorageExportsResult',
    'get_file_storage_exports',
]

@pulumi.output_type
class GetFileStorageExportsResult:
    """
    A collection of values returned by GetFileStorageExports.
    """
    def __init__(__self__, compartment_id=None, export_set_id=None, exports=None, file_system_id=None, filters=None, id=None, state=None):
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if export_set_id and not isinstance(export_set_id, str):
            raise TypeError("Expected argument 'export_set_id' to be a str")
        pulumi.set(__self__, "export_set_id", export_set_id)
        if exports and not isinstance(exports, list):
            raise TypeError("Expected argument 'exports' to be a list")
        pulumi.set(__self__, "exports", exports)
        if file_system_id and not isinstance(file_system_id, str):
            raise TypeError("Expected argument 'file_system_id' to be a str")
        pulumi.set(__self__, "file_system_id", file_system_id)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> Optional[str]:
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="exportSetId")
    def export_set_id(self) -> Optional[str]:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this export's export set.
        """
        return pulumi.get(self, "export_set_id")

    @property
    @pulumi.getter
    def exports(self) -> Sequence['outputs.GetFileStorageExportsExportResult']:
        """
        The list of exports.
        """
        return pulumi.get(self, "exports")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> Optional[str]:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this export's file system.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetFileStorageExportsFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this export.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        The current state of this export.
        """
        return pulumi.get(self, "state")


class AwaitableGetFileStorageExportsResult(GetFileStorageExportsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileStorageExportsResult(
            compartment_id=self.compartment_id,
            export_set_id=self.export_set_id,
            exports=self.exports,
            file_system_id=self.file_system_id,
            filters=self.filters,
            id=self.id,
            state=self.state)


def get_file_storage_exports(compartment_id: Optional[str] = None,
                             export_set_id: Optional[str] = None,
                             file_system_id: Optional[str] = None,
                             filters: Optional[Sequence[pulumi.InputType['GetFileStorageExportsFilterArgs']]] = None,
                             id: Optional[str] = None,
                             state: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileStorageExportsResult:
    """
    This data source provides the list of Exports in Oracle Cloud Infrastructure File Storage service.

    Lists export resources by compartment, file system, or export
    set. You must specify an export set ID, a file system ID, and
    / or a compartment ID.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_exports = oci.get_file_storage_exports(compartment_id=var["compartment_id"],
        export_set_id=oci_file_storage_export_set["test_export_set"]["id"],
        file_system_id=oci_file_storage_file_system["test_file_system"]["id"],
        id=var["export_id"],
        state=var["export_state"])
    ```


    :param str compartment_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
    :param str export_set_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the export set.
    :param str file_system_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system.
    :param str id: Filter results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resouce type.
    :param str state: Filter results by the specified lifecycle state. Must be a valid state for the resource type.
    """
    __args__ = dict()
    __args__['compartmentId'] = compartment_id
    __args__['exportSetId'] = export_set_id
    __args__['fileSystemId'] = file_system_id
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['state'] = state
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:index/getFileStorageExports:GetFileStorageExports', __args__, opts=opts, typ=GetFileStorageExportsResult).value

    return AwaitableGetFileStorageExportsResult(
        compartment_id=__ret__.compartment_id,
        export_set_id=__ret__.export_set_id,
        exports=__ret__.exports,
        file_system_id=__ret__.file_system_id,
        filters=__ret__.filters,
        id=__ret__.id,
        state=__ret__.state)