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
    'GetLoadBalancerRuleSetsResult',
    'AwaitableGetLoadBalancerRuleSetsResult',
    'get_load_balancer_rule_sets',
]

@pulumi.output_type
class GetLoadBalancerRuleSetsResult:
    """
    A collection of values returned by GetLoadBalancerRuleSets.
    """
    def __init__(__self__, filters=None, id=None, load_balancer_id=None, rule_sets=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if load_balancer_id and not isinstance(load_balancer_id, str):
            raise TypeError("Expected argument 'load_balancer_id' to be a str")
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        if rule_sets and not isinstance(rule_sets, list):
            raise TypeError("Expected argument 'rule_sets' to be a list")
        pulumi.set(__self__, "rule_sets", rule_sets)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetLoadBalancerRuleSetsFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> str:
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter(name="ruleSets")
    def rule_sets(self) -> Sequence['outputs.GetLoadBalancerRuleSetsRuleSetResult']:
        """
        The list of rule_sets.
        """
        return pulumi.get(self, "rule_sets")


class AwaitableGetLoadBalancerRuleSetsResult(GetLoadBalancerRuleSetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadBalancerRuleSetsResult(
            filters=self.filters,
            id=self.id,
            load_balancer_id=self.load_balancer_id,
            rule_sets=self.rule_sets)


def get_load_balancer_rule_sets(filters: Optional[Sequence[pulumi.InputType['GetLoadBalancerRuleSetsFilterArgs']]] = None,
                                load_balancer_id: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadBalancerRuleSetsResult:
    """
    This data source provides the list of Rule Sets in Oracle Cloud Infrastructure Load Balancer service.

    Lists all rule sets associated with the specified load balancer.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_rule_sets = oci.get_load_balancer_rule_sets(load_balancer_id=oci_load_balancer_load_balancer["test_load_balancer"]["id"])
    ```


    :param str load_balancer_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the specified load balancer.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['loadBalancerId'] = load_balancer_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:index/getLoadBalancerRuleSets:GetLoadBalancerRuleSets', __args__, opts=opts, typ=GetLoadBalancerRuleSetsResult).value

    return AwaitableGetLoadBalancerRuleSetsResult(
        filters=__ret__.filters,
        id=__ret__.id,
        load_balancer_id=__ret__.load_balancer_id,
        rule_sets=__ret__.rule_sets)