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
    'GetJmsFleetsResult',
    'AwaitableGetJmsFleetsResult',
    'get_jms_fleets',
]

@pulumi.output_type
class GetJmsFleetsResult:
    """
    A collection of values returned by GetJmsFleets.
    """
    def __init__(__self__, compartment_id=None, display_name=None, filters=None, fleet_collections=None, id=None, state=None):
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if fleet_collections and not isinstance(fleet_collections, list):
            raise TypeError("Expected argument 'fleet_collections' to be a list")
        pulumi.set(__self__, "fleet_collections", fleet_collections)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> Optional[str]:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of the Fleet.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The name of the Fleet.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetJmsFleetsFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="fleetCollections")
    def fleet_collections(self) -> Sequence['outputs.GetJmsFleetsFleetCollectionResult']:
        """
        The list of fleet_collection.
        """
        return pulumi.get(self, "fleet_collections")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        The lifecycle state of the Fleet.
        """
        return pulumi.get(self, "state")


class AwaitableGetJmsFleetsResult(GetJmsFleetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetJmsFleetsResult(
            compartment_id=self.compartment_id,
            display_name=self.display_name,
            filters=self.filters,
            fleet_collections=self.fleet_collections,
            id=self.id,
            state=self.state)


def get_jms_fleets(compartment_id: Optional[str] = None,
                   display_name: Optional[str] = None,
                   filters: Optional[Sequence[pulumi.InputType['GetJmsFleetsFilterArgs']]] = None,
                   id: Optional[str] = None,
                   state: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetJmsFleetsResult:
    """
    This data source provides the list of Fleets in Oracle Cloud Infrastructure Jms service.

    Returns a list of all the Fleets contained by a compartment.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_fleets = oci.get_jms_fleets(compartment_id=var["compartment_id"],
        display_name=var["fleet_display_name"],
        id=var["fleet_id"],
        state=var["fleet_state"])
    ```


    :param str compartment_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
    :param str display_name: The display name.
    :param str id: The ID of the Fleet.
    :param str state: The state of the lifecycle.
    """
    __args__ = dict()
    __args__['compartmentId'] = compartment_id
    __args__['displayName'] = display_name
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['state'] = state
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:index/getJmsFleets:GetJmsFleets', __args__, opts=opts, typ=GetJmsFleetsResult).value

    return AwaitableGetJmsFleetsResult(
        compartment_id=__ret__.compartment_id,
        display_name=__ret__.display_name,
        filters=__ret__.filters,
        fleet_collections=__ret__.fleet_collections,
        id=__ret__.id,
        state=__ret__.state)