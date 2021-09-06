# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetEmailDomainResult',
    'AwaitableGetEmailDomainResult',
    'get_email_domain',
]

@pulumi.output_type
class GetEmailDomainResult:
    """
    A collection of values returned by getEmailDomain.
    """
    def __init__(__self__, active_dkim_id=None, compartment_id=None, defined_tags=None, description=None, email_domain_id=None, freeform_tags=None, id=None, is_spf=None, name=None, state=None, system_tags=None, time_created=None):
        if active_dkim_id and not isinstance(active_dkim_id, str):
            raise TypeError("Expected argument 'active_dkim_id' to be a str")
        pulumi.set(__self__, "active_dkim_id", active_dkim_id)
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if defined_tags and not isinstance(defined_tags, dict):
            raise TypeError("Expected argument 'defined_tags' to be a dict")
        pulumi.set(__self__, "defined_tags", defined_tags)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if email_domain_id and not isinstance(email_domain_id, str):
            raise TypeError("Expected argument 'email_domain_id' to be a str")
        pulumi.set(__self__, "email_domain_id", email_domain_id)
        if freeform_tags and not isinstance(freeform_tags, dict):
            raise TypeError("Expected argument 'freeform_tags' to be a dict")
        pulumi.set(__self__, "freeform_tags", freeform_tags)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_spf and not isinstance(is_spf, bool):
            raise TypeError("Expected argument 'is_spf' to be a bool")
        pulumi.set(__self__, "is_spf", is_spf)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if system_tags and not isinstance(system_tags, dict):
            raise TypeError("Expected argument 'system_tags' to be a dict")
        pulumi.set(__self__, "system_tags", system_tags)
        if time_created and not isinstance(time_created, str):
            raise TypeError("Expected argument 'time_created' to be a str")
        pulumi.set(__self__, "time_created", time_created)

    @property
    @pulumi.getter(name="activeDkimId")
    def active_dkim_id(self) -> str:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DKIM key that will be used to sign mail sent from this email domain.
        """
        return pulumi.get(self, "active_dkim_id")

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> str:
        """
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this email domain.
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
    @pulumi.getter
    def description(self) -> str:
        """
        The description of a email domain.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="emailDomainId")
    def email_domain_id(self) -> str:
        return pulumi.get(self, "email_domain_id")

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
        The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email domain.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isSpf")
    def is_spf(self) -> bool:
        """
        Value of the SPF field. For more information about SPF, please see [SPF Authentication](https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm#components).
        """
        return pulumi.get(self, "is_spf")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the email domain in the Internet Domain Name System (DNS).  Example: `example.net`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the email domain.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="systemTags")
    def system_tags(self) -> Mapping[str, Any]:
        """
        Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
        """
        return pulumi.get(self, "system_tags")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        The time the email domain was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".  Example: `2021-02-12T22:47:12.613Z`
        """
        return pulumi.get(self, "time_created")


class AwaitableGetEmailDomainResult(GetEmailDomainResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEmailDomainResult(
            active_dkim_id=self.active_dkim_id,
            compartment_id=self.compartment_id,
            defined_tags=self.defined_tags,
            description=self.description,
            email_domain_id=self.email_domain_id,
            freeform_tags=self.freeform_tags,
            id=self.id,
            is_spf=self.is_spf,
            name=self.name,
            state=self.state,
            system_tags=self.system_tags,
            time_created=self.time_created)


def get_email_domain(email_domain_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEmailDomainResult:
    """
    This data source provides details about a specific Email Domain resource in Oracle Cloud Infrastructure Email service.

    Retrieves the specified email domain.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_email_domain = oci.email.get_email_domain(email_domain_id=oci_email_email_domain["test_email_domain"]["id"])
    ```


    :param str email_domain_id: The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this email domain.
    """
    __args__ = dict()
    __args__['emailDomainId'] = email_domain_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:email/getEmailDomain:getEmailDomain', __args__, opts=opts, typ=GetEmailDomainResult).value

    return AwaitableGetEmailDomainResult(
        active_dkim_id=__ret__.active_dkim_id,
        compartment_id=__ret__.compartment_id,
        defined_tags=__ret__.defined_tags,
        description=__ret__.description,
        email_domain_id=__ret__.email_domain_id,
        freeform_tags=__ret__.freeform_tags,
        id=__ret__.id,
        is_spf=__ret__.is_spf,
        name=__ret__.name,
        state=__ret__.state,
        system_tags=__ret__.system_tags,
        time_created=__ret__.time_created)