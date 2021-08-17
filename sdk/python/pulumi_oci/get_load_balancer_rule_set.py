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
    'GetLoadBalancerRuleSetResult',
    'AwaitableGetLoadBalancerRuleSetResult',
    'get_load_balancer_rule_set',
]

@pulumi.output_type
class GetLoadBalancerRuleSetResult:
    """
    A collection of values returned by GetLoadBalancerRuleSet.
    """
    def __init__(__self__, id=None, items=None, load_balancer_id=None, name=None, state=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)
        if load_balancer_id and not isinstance(load_balancer_id, str):
            raise TypeError("Expected argument 'load_balancer_id' to be a str")
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.GetLoadBalancerRuleSetItemResult']:
        """
        An array of rules that compose the rule set.
        """
        return pulumi.get(self, "items")

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> str:
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name for this set of rules. It must be unique and it cannot be changed. Avoid entering confidential information.  Example: `example_rule_set`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")


class AwaitableGetLoadBalancerRuleSetResult(GetLoadBalancerRuleSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLoadBalancerRuleSetResult(
            id=self.id,
            items=self.items,
            load_balancer_id=self.load_balancer_id,
            name=self.name,
            state=self.state)


def get_load_balancer_rule_set(load_balancer_id: Optional[str] = None,
                               name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLoadBalancerRuleSetResult:
    """
    This data source provides details about a specific Rule Set resource in Oracle Cloud Infrastructure Load Balancer service.

    Gets the specified set of rules.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_rule_set = oci.get_load_balancer_rule_set(load_balancer_id=oci_load_balancer_load_balancer["test_load_balancer"]["id"],
        name=var["rule_set_name"])
    ```


    :param str load_balancer_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the specified load balancer.
    :param str name: The name of the rule set to retrieve.  Example: `example_rule_set`
    """
    __args__ = dict()
    __args__['loadBalancerId'] = load_balancer_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:index/getLoadBalancerRuleSet:GetLoadBalancerRuleSet', __args__, opts=opts, typ=GetLoadBalancerRuleSetResult).value

    return AwaitableGetLoadBalancerRuleSetResult(
        id=__ret__.id,
        items=__ret__.items,
        load_balancer_id=__ret__.load_balancer_id,
        name=__ret__.name,
        state=__ret__.state)