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
    'GetMeteringComputationCustomTableResult',
    'AwaitableGetMeteringComputationCustomTableResult',
    'get_metering_computation_custom_table',
]

@pulumi.output_type
class GetMeteringComputationCustomTableResult:
    """
    A collection of values returned by GetMeteringComputationCustomTable.
    """
    def __init__(__self__, compartment_id=None, custom_table_id=None, id=None, saved_custom_table=None, saved_report_id=None):
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if custom_table_id and not isinstance(custom_table_id, str):
            raise TypeError("Expected argument 'custom_table_id' to be a str")
        pulumi.set(__self__, "custom_table_id", custom_table_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if saved_custom_table and not isinstance(saved_custom_table, dict):
            raise TypeError("Expected argument 'saved_custom_table' to be a dict")
        pulumi.set(__self__, "saved_custom_table", saved_custom_table)
        if saved_report_id and not isinstance(saved_report_id, str):
            raise TypeError("Expected argument 'saved_report_id' to be a str")
        pulumi.set(__self__, "saved_report_id", saved_report_id)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> str:
        """
        The custom table compartment OCID.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="customTableId")
    def custom_table_id(self) -> str:
        return pulumi.get(self, "custom_table_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The custom table OCID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="savedCustomTable")
    def saved_custom_table(self) -> 'outputs.GetMeteringComputationCustomTableSavedCustomTableResult':
        """
        The custom table for Cost Analysis UI rendering.
        """
        return pulumi.get(self, "saved_custom_table")

    @property
    @pulumi.getter(name="savedReportId")
    def saved_report_id(self) -> str:
        """
        The custom table associated saved report OCID.
        """
        return pulumi.get(self, "saved_report_id")


class AwaitableGetMeteringComputationCustomTableResult(GetMeteringComputationCustomTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMeteringComputationCustomTableResult(
            compartment_id=self.compartment_id,
            custom_table_id=self.custom_table_id,
            id=self.id,
            saved_custom_table=self.saved_custom_table,
            saved_report_id=self.saved_report_id)


def get_metering_computation_custom_table(custom_table_id: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMeteringComputationCustomTableResult:
    """
    This data source provides details about a specific Custom Table resource in Oracle Cloud Infrastructure Metering Computation service.

    Returns the saved custom table.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_custom_table = oci.get_metering_computation_custom_table(custom_table_id=oci_metering_computation_custom_table["test_custom_table"]["id"])
    ```


    :param str custom_table_id: The custom table unique OCID.
    """
    __args__ = dict()
    __args__['customTableId'] = custom_table_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:index/getMeteringComputationCustomTable:GetMeteringComputationCustomTable', __args__, opts=opts, typ=GetMeteringComputationCustomTableResult).value

    return AwaitableGetMeteringComputationCustomTableResult(
        compartment_id=__ret__.compartment_id,
        custom_table_id=__ret__.custom_table_id,
        id=__ret__.id,
        saved_custom_table=__ret__.saved_custom_table,
        saved_report_id=__ret__.saved_report_id)