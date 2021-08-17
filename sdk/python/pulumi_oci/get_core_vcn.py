# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetCoreVcnResult',
    'AwaitableGetCoreVcnResult',
    'get_core_vcn',
]

@pulumi.output_type
class GetCoreVcnResult:
    """
    A collection of values returned by GetCoreVcn.
    """
    def __init__(__self__, cidr_block=None, cidr_blocks=None, compartment_id=None, default_dhcp_options_id=None, default_route_table_id=None, default_security_list_id=None, defined_tags=None, display_name=None, dns_label=None, freeform_tags=None, id=None, ipv6cidr_blocks=None, is_ipv6enabled=None, state=None, time_created=None, vcn_domain_name=None, vcn_id=None):
        if cidr_block and not isinstance(cidr_block, str):
            raise TypeError("Expected argument 'cidr_block' to be a str")
        pulumi.set(__self__, "cidr_block", cidr_block)
        if cidr_blocks and not isinstance(cidr_blocks, list):
            raise TypeError("Expected argument 'cidr_blocks' to be a list")
        pulumi.set(__self__, "cidr_blocks", cidr_blocks)
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if default_dhcp_options_id and not isinstance(default_dhcp_options_id, str):
            raise TypeError("Expected argument 'default_dhcp_options_id' to be a str")
        pulumi.set(__self__, "default_dhcp_options_id", default_dhcp_options_id)
        if default_route_table_id and not isinstance(default_route_table_id, str):
            raise TypeError("Expected argument 'default_route_table_id' to be a str")
        pulumi.set(__self__, "default_route_table_id", default_route_table_id)
        if default_security_list_id and not isinstance(default_security_list_id, str):
            raise TypeError("Expected argument 'default_security_list_id' to be a str")
        pulumi.set(__self__, "default_security_list_id", default_security_list_id)
        if defined_tags and not isinstance(defined_tags, dict):
            raise TypeError("Expected argument 'defined_tags' to be a dict")
        pulumi.set(__self__, "defined_tags", defined_tags)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if dns_label and not isinstance(dns_label, str):
            raise TypeError("Expected argument 'dns_label' to be a str")
        pulumi.set(__self__, "dns_label", dns_label)
        if freeform_tags and not isinstance(freeform_tags, dict):
            raise TypeError("Expected argument 'freeform_tags' to be a dict")
        pulumi.set(__self__, "freeform_tags", freeform_tags)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ipv6cidr_blocks and not isinstance(ipv6cidr_blocks, list):
            raise TypeError("Expected argument 'ipv6cidr_blocks' to be a list")
        pulumi.set(__self__, "ipv6cidr_blocks", ipv6cidr_blocks)
        if is_ipv6enabled and not isinstance(is_ipv6enabled, bool):
            raise TypeError("Expected argument 'is_ipv6enabled' to be a bool")
        pulumi.set(__self__, "is_ipv6enabled", is_ipv6enabled)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if time_created and not isinstance(time_created, str):
            raise TypeError("Expected argument 'time_created' to be a str")
        pulumi.set(__self__, "time_created", time_created)
        if vcn_domain_name and not isinstance(vcn_domain_name, str):
            raise TypeError("Expected argument 'vcn_domain_name' to be a str")
        pulumi.set(__self__, "vcn_domain_name", vcn_domain_name)
        if vcn_id and not isinstance(vcn_id, str):
            raise TypeError("Expected argument 'vcn_id' to be a str")
        pulumi.set(__self__, "vcn_id", vcn_id)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> str:
        """
        Deprecated. The first CIDR IP address from cidrBlocks.  Example: `172.16.0.0/16`
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter(name="cidrBlocks")
    def cidr_blocks(self) -> Sequence[str]:
        """
        The list of IPv4 CIDR blocks the VCN will use.
        """
        return pulumi.get(self, "cidr_blocks")

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> str:
        """
        The OCID of the compartment containing the VCN.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="defaultDhcpOptionsId")
    def default_dhcp_options_id(self) -> str:
        """
        The OCID for the VCN's default set of DHCP options.
        """
        return pulumi.get(self, "default_dhcp_options_id")

    @property
    @pulumi.getter(name="defaultRouteTableId")
    def default_route_table_id(self) -> str:
        """
        The OCID for the VCN's default route table.
        """
        return pulumi.get(self, "default_route_table_id")

    @property
    @pulumi.getter(name="defaultSecurityListId")
    def default_security_list_id(self) -> str:
        """
        The OCID for the VCN's default security list.
        """
        return pulumi.get(self, "default_security_list_id")

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
        A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="dnsLabel")
    def dns_label(self) -> str:
        """
        A DNS label for the VCN, used in conjunction with the VNIC's hostname and subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be an alphanumeric string that begins with a letter. The value cannot be changed.
        """
        return pulumi.get(self, "dns_label")

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
        The VCN's Oracle ID (OCID).
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipv6cidrBlocks")
    def ipv6cidr_blocks(self) -> Sequence[str]:
        """
        For an IPv6-enabled VCN, this is the list of IPv6 CIDR blocks for the VCN's IP address space. The CIDRs are provided by Oracle and the sizes are always /56.
        """
        return pulumi.get(self, "ipv6cidr_blocks")

    @property
    @pulumi.getter(name="isIpv6enabled")
    def is_ipv6enabled(self) -> bool:
        return pulumi.get(self, "is_ipv6enabled")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The VCN's current state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        The date and time the VCN was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter(name="vcnDomainName")
    def vcn_domain_name(self) -> str:
        """
        The VCN's domain name, which consists of the VCN's DNS label, and the `oraclevcn.com` domain.
        """
        return pulumi.get(self, "vcn_domain_name")

    @property
    @pulumi.getter(name="vcnId")
    def vcn_id(self) -> str:
        return pulumi.get(self, "vcn_id")


class AwaitableGetCoreVcnResult(GetCoreVcnResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCoreVcnResult(
            cidr_block=self.cidr_block,
            cidr_blocks=self.cidr_blocks,
            compartment_id=self.compartment_id,
            default_dhcp_options_id=self.default_dhcp_options_id,
            default_route_table_id=self.default_route_table_id,
            default_security_list_id=self.default_security_list_id,
            defined_tags=self.defined_tags,
            display_name=self.display_name,
            dns_label=self.dns_label,
            freeform_tags=self.freeform_tags,
            id=self.id,
            ipv6cidr_blocks=self.ipv6cidr_blocks,
            is_ipv6enabled=self.is_ipv6enabled,
            state=self.state,
            time_created=self.time_created,
            vcn_domain_name=self.vcn_domain_name,
            vcn_id=self.vcn_id)


def get_core_vcn(vcn_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCoreVcnResult:
    """
    This data source provides details about a specific Vcn resource in Oracle Cloud Infrastructure Core service.

    Gets the specified VCN's information.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_vcn = oci.get_core_vcn(vcn_id=oci_core_vcn["test_vcn"]["id"])
    ```


    :param str vcn_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.
    """
    __args__ = dict()
    __args__['vcnId'] = vcn_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:index/getCoreVcn:GetCoreVcn', __args__, opts=opts, typ=GetCoreVcnResult).value

    return AwaitableGetCoreVcnResult(
        cidr_block=__ret__.cidr_block,
        cidr_blocks=__ret__.cidr_blocks,
        compartment_id=__ret__.compartment_id,
        default_dhcp_options_id=__ret__.default_dhcp_options_id,
        default_route_table_id=__ret__.default_route_table_id,
        default_security_list_id=__ret__.default_security_list_id,
        defined_tags=__ret__.defined_tags,
        display_name=__ret__.display_name,
        dns_label=__ret__.dns_label,
        freeform_tags=__ret__.freeform_tags,
        id=__ret__.id,
        ipv6cidr_blocks=__ret__.ipv6cidr_blocks,
        is_ipv6enabled=__ret__.is_ipv6enabled,
        state=__ret__.state,
        time_created=__ret__.time_created,
        vcn_domain_name=__ret__.vcn_domain_name,
        vcn_id=__ret__.vcn_id)