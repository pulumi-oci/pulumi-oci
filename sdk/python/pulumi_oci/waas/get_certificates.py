# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetCertificatesResult',
    'AwaitableGetCertificatesResult',
    'get_certificates',
]

@pulumi.output_type
class GetCertificatesResult:
    """
    A collection of values returned by getCertificates.
    """
    def __init__(__self__, certificates=None, compartment_id=None, display_names=None, filters=None, id=None, ids=None, states=None, time_created_greater_than_or_equal_to=None, time_created_less_than=None):
        if certificates and not isinstance(certificates, list):
            raise TypeError("Expected argument 'certificates' to be a list")
        pulumi.set(__self__, "certificates", certificates)
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if display_names and not isinstance(display_names, list):
            raise TypeError("Expected argument 'display_names' to be a list")
        pulumi.set(__self__, "display_names", display_names)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if states and not isinstance(states, list):
            raise TypeError("Expected argument 'states' to be a list")
        pulumi.set(__self__, "states", states)
        if time_created_greater_than_or_equal_to and not isinstance(time_created_greater_than_or_equal_to, str):
            raise TypeError("Expected argument 'time_created_greater_than_or_equal_to' to be a str")
        pulumi.set(__self__, "time_created_greater_than_or_equal_to", time_created_greater_than_or_equal_to)
        if time_created_less_than and not isinstance(time_created_less_than, str):
            raise TypeError("Expected argument 'time_created_less_than' to be a str")
        pulumi.set(__self__, "time_created_less_than", time_created_less_than)

    @property
    @pulumi.getter
    def certificates(self) -> Sequence['outputs.GetCertificatesCertificateResult']:
        """
        The list of certificates.
        """
        return pulumi.get(self, "certificates")

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> str:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SSL certificate's compartment.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="displayNames")
    def display_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "display_names")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetCertificatesFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def states(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "states")

    @property
    @pulumi.getter(name="timeCreatedGreaterThanOrEqualTo")
    def time_created_greater_than_or_equal_to(self) -> Optional[str]:
        return pulumi.get(self, "time_created_greater_than_or_equal_to")

    @property
    @pulumi.getter(name="timeCreatedLessThan")
    def time_created_less_than(self) -> Optional[str]:
        return pulumi.get(self, "time_created_less_than")


class AwaitableGetCertificatesResult(GetCertificatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificatesResult(
            certificates=self.certificates,
            compartment_id=self.compartment_id,
            display_names=self.display_names,
            filters=self.filters,
            id=self.id,
            ids=self.ids,
            states=self.states,
            time_created_greater_than_or_equal_to=self.time_created_greater_than_or_equal_to,
            time_created_less_than=self.time_created_less_than)


def get_certificates(compartment_id: Optional[str] = None,
                     display_names: Optional[Sequence[str]] = None,
                     filters: Optional[Sequence[pulumi.InputType['GetCertificatesFilterArgs']]] = None,
                     ids: Optional[Sequence[str]] = None,
                     states: Optional[Sequence[str]] = None,
                     time_created_greater_than_or_equal_to: Optional[str] = None,
                     time_created_less_than: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCertificatesResult:
    """
    This data source provides the list of Certificates in Oracle Cloud Infrastructure Web Application Acceleration and Security service.

    Gets a list of SSL certificates that can be used in a WAAS policy.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_certificates = oci.waas.get_certificates(compartment_id=var["compartment_id"],
        display_names=var["certificate_display_names"],
        ids=var["certificate_ids"],
        states=var["certificate_states"],
        time_created_greater_than_or_equal_to=var["certificate_time_created_greater_than_or_equal_to"],
        time_created_less_than=var["certificate_time_created_less_than"])
    ```


    :param str compartment_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
    :param Sequence[str] display_names: Filter certificates using a list of display names.
    :param Sequence[str] ids: Filter certificates using a list of certificates OCIDs.
    :param Sequence[str] states: Filter certificates using a list of lifecycle states.
    :param str time_created_greater_than_or_equal_to: A filter that matches certificates created on or after the specified date-time.
    :param str time_created_less_than: A filter that matches certificates created before the specified date-time.
    """
    __args__ = dict()
    __args__['compartmentId'] = compartment_id
    __args__['displayNames'] = display_names
    __args__['filters'] = filters
    __args__['ids'] = ids
    __args__['states'] = states
    __args__['timeCreatedGreaterThanOrEqualTo'] = time_created_greater_than_or_equal_to
    __args__['timeCreatedLessThan'] = time_created_less_than
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:waas/getCertificates:getCertificates', __args__, opts=opts, typ=GetCertificatesResult).value

    return AwaitableGetCertificatesResult(
        certificates=__ret__.certificates,
        compartment_id=__ret__.compartment_id,
        display_names=__ret__.display_names,
        filters=__ret__.filters,
        id=__ret__.id,
        ids=__ret__.ids,
        states=__ret__.states,
        time_created_greater_than_or_equal_to=__ret__.time_created_greater_than_or_equal_to,
        time_created_less_than=__ret__.time_created_less_than)