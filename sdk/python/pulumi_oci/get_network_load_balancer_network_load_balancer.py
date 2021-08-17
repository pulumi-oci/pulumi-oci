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
    'GetNetworkLoadBalancerNetworkLoadBalancerResult',
    'AwaitableGetNetworkLoadBalancerNetworkLoadBalancerResult',
    'get_network_load_balancer_network_load_balancer',
]

@pulumi.output_type
class GetNetworkLoadBalancerNetworkLoadBalancerResult:
    """
    A collection of values returned by GetNetworkLoadBalancerNetworkLoadBalancer.
    """
    def __init__(__self__, compartment_id=None, defined_tags=None, display_name=None, freeform_tags=None, id=None, ip_addresses=None, is_preserve_source_destination=None, is_private=None, lifecycle_details=None, network_load_balancer_id=None, network_security_group_ids=None, reserved_ips=None, state=None, subnet_id=None, system_tags=None, time_created=None, time_updated=None):
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if defined_tags and not isinstance(defined_tags, dict):
            raise TypeError("Expected argument 'defined_tags' to be a dict")
        pulumi.set(__self__, "defined_tags", defined_tags)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if freeform_tags and not isinstance(freeform_tags, dict):
            raise TypeError("Expected argument 'freeform_tags' to be a dict")
        pulumi.set(__self__, "freeform_tags", freeform_tags)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_addresses and not isinstance(ip_addresses, list):
            raise TypeError("Expected argument 'ip_addresses' to be a list")
        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if is_preserve_source_destination and not isinstance(is_preserve_source_destination, bool):
            raise TypeError("Expected argument 'is_preserve_source_destination' to be a bool")
        pulumi.set(__self__, "is_preserve_source_destination", is_preserve_source_destination)
        if is_private and not isinstance(is_private, bool):
            raise TypeError("Expected argument 'is_private' to be a bool")
        pulumi.set(__self__, "is_private", is_private)
        if lifecycle_details and not isinstance(lifecycle_details, str):
            raise TypeError("Expected argument 'lifecycle_details' to be a str")
        pulumi.set(__self__, "lifecycle_details", lifecycle_details)
        if network_load_balancer_id and not isinstance(network_load_balancer_id, str):
            raise TypeError("Expected argument 'network_load_balancer_id' to be a str")
        pulumi.set(__self__, "network_load_balancer_id", network_load_balancer_id)
        if network_security_group_ids and not isinstance(network_security_group_ids, list):
            raise TypeError("Expected argument 'network_security_group_ids' to be a list")
        pulumi.set(__self__, "network_security_group_ids", network_security_group_ids)
        if reserved_ips and not isinstance(reserved_ips, list):
            raise TypeError("Expected argument 'reserved_ips' to be a list")
        pulumi.set(__self__, "reserved_ips", reserved_ips)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if system_tags and not isinstance(system_tags, dict):
            raise TypeError("Expected argument 'system_tags' to be a dict")
        pulumi.set(__self__, "system_tags", system_tags)
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
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancer.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="definedTags")
    def defined_tags(self) -> Mapping[str, Any]:
        """
        Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        """
        return pulumi.get(self, "defined_tags")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        A user-friendly name, which does not have to be unique, and can be changed.  Example: `example_load_balancer`
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="freeformTags")
    def freeform_tags(self) -> Mapping[str, Any]:
        """
        Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        """
        return pulumi.get(self, "freeform_tags")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        OCID of the reserved public IP address created with the virtual cloud network.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Sequence['outputs.GetNetworkLoadBalancerNetworkLoadBalancerIpAddressResult']:
        """
        An array of IP addresses.
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="isPreserveSourceDestination")
    def is_preserve_source_destination(self) -> bool:
        """
        When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC. Packets are sent to the backend set without any changes to the source and destination IP.
        """
        return pulumi.get(self, "is_preserve_source_destination")

    @property
    @pulumi.getter(name="isPrivate")
    def is_private(self) -> bool:
        """
        Whether the network load balancer has a virtual cloud network-local (private) IP address.
        """
        return pulumi.get(self, "is_private")

    @property
    @pulumi.getter(name="lifecycleDetails")
    def lifecycle_details(self) -> str:
        """
        A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
        """
        return pulumi.get(self, "lifecycle_details")

    @property
    @pulumi.getter(name="networkLoadBalancerId")
    def network_load_balancer_id(self) -> str:
        return pulumi.get(self, "network_load_balancer_id")

    @property
    @pulumi.getter(name="networkSecurityGroupIds")
    def network_security_group_ids(self) -> Sequence[str]:
        """
        An array of network security groups [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with the network load balancer.
        """
        return pulumi.get(self, "network_security_group_ids")

    @property
    @pulumi.getter(name="reservedIps")
    def reserved_ips(self) -> Sequence['outputs.GetNetworkLoadBalancerNetworkLoadBalancerReservedIpResult']:
        return pulumi.get(self, "reserved_ips")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the network load balancer.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The subnet in which the network load balancer is spawned [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)."
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="systemTags")
    def system_tags(self) -> Mapping[str, Any]:
        """
        Key-value pair representing system tags' keys and values scoped to a namespace. Example: `{"bar-key": "value"}`
        """
        return pulumi.get(self, "system_tags")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        The date and time the network load balancer was created, in the format defined by RFC3339.  Example: `2020-05-01T21:10:29.600Z`
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter(name="timeUpdated")
    def time_updated(self) -> str:
        """
        The time the network load balancer was updated. An RFC3339 formatted date-time string.  Example: `2020-05-01T22:10:29.600Z`
        """
        return pulumi.get(self, "time_updated")


class AwaitableGetNetworkLoadBalancerNetworkLoadBalancerResult(GetNetworkLoadBalancerNetworkLoadBalancerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkLoadBalancerNetworkLoadBalancerResult(
            compartment_id=self.compartment_id,
            defined_tags=self.defined_tags,
            display_name=self.display_name,
            freeform_tags=self.freeform_tags,
            id=self.id,
            ip_addresses=self.ip_addresses,
            is_preserve_source_destination=self.is_preserve_source_destination,
            is_private=self.is_private,
            lifecycle_details=self.lifecycle_details,
            network_load_balancer_id=self.network_load_balancer_id,
            network_security_group_ids=self.network_security_group_ids,
            reserved_ips=self.reserved_ips,
            state=self.state,
            subnet_id=self.subnet_id,
            system_tags=self.system_tags,
            time_created=self.time_created,
            time_updated=self.time_updated)


def get_network_load_balancer_network_load_balancer(network_load_balancer_id: Optional[str] = None,
                                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkLoadBalancerNetworkLoadBalancerResult:
    """
    This data source provides details about a specific Network Load Balancer resource in Oracle Cloud Infrastructure Network Load Balancer service.

    Retrieves network load balancer configuration information by identifier.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_network_load_balancer = oci.get_network_load_balancer_network_load_balancer(network_load_balancer_id=oci_network_load_balancer_network_load_balancer["test_network_load_balancer"]["id"])
    ```


    :param str network_load_balancer_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
    """
    __args__ = dict()
    __args__['networkLoadBalancerId'] = network_load_balancer_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:index/getNetworkLoadBalancerNetworkLoadBalancer:GetNetworkLoadBalancerNetworkLoadBalancer', __args__, opts=opts, typ=GetNetworkLoadBalancerNetworkLoadBalancerResult).value

    return AwaitableGetNetworkLoadBalancerNetworkLoadBalancerResult(
        compartment_id=__ret__.compartment_id,
        defined_tags=__ret__.defined_tags,
        display_name=__ret__.display_name,
        freeform_tags=__ret__.freeform_tags,
        id=__ret__.id,
        ip_addresses=__ret__.ip_addresses,
        is_preserve_source_destination=__ret__.is_preserve_source_destination,
        is_private=__ret__.is_private,
        lifecycle_details=__ret__.lifecycle_details,
        network_load_balancer_id=__ret__.network_load_balancer_id,
        network_security_group_ids=__ret__.network_security_group_ids,
        reserved_ips=__ret__.reserved_ips,
        state=__ret__.state,
        subnet_id=__ret__.subnet_id,
        system_tags=__ret__.system_tags,
        time_created=__ret__.time_created,
        time_updated=__ret__.time_updated)