# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetTargetResult',
    'AwaitableGetTargetResult',
    'get_target',
]

@pulumi.output_type
class GetTargetResult:
    """
    A collection of values returned by getTarget.
    """
    def __init__(__self__, compartment_id=None, defined_tags=None, description=None, display_name=None, freeform_tags=None, id=None, inherited_by_compartments=None, lifecyle_details=None, recipe_count=None, state=None, system_tags=None, target_detector_recipes=None, target_id=None, target_resource_id=None, target_resource_type=None, target_responder_recipes=None, time_created=None, time_updated=None):
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if defined_tags and not isinstance(defined_tags, dict):
            raise TypeError("Expected argument 'defined_tags' to be a dict")
        pulumi.set(__self__, "defined_tags", defined_tags)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if freeform_tags and not isinstance(freeform_tags, dict):
            raise TypeError("Expected argument 'freeform_tags' to be a dict")
        pulumi.set(__self__, "freeform_tags", freeform_tags)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inherited_by_compartments and not isinstance(inherited_by_compartments, list):
            raise TypeError("Expected argument 'inherited_by_compartments' to be a list")
        pulumi.set(__self__, "inherited_by_compartments", inherited_by_compartments)
        if lifecyle_details and not isinstance(lifecyle_details, str):
            raise TypeError("Expected argument 'lifecyle_details' to be a str")
        pulumi.set(__self__, "lifecyle_details", lifecyle_details)
        if recipe_count and not isinstance(recipe_count, int):
            raise TypeError("Expected argument 'recipe_count' to be a int")
        pulumi.set(__self__, "recipe_count", recipe_count)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if system_tags and not isinstance(system_tags, dict):
            raise TypeError("Expected argument 'system_tags' to be a dict")
        pulumi.set(__self__, "system_tags", system_tags)
        if target_detector_recipes and not isinstance(target_detector_recipes, list):
            raise TypeError("Expected argument 'target_detector_recipes' to be a list")
        pulumi.set(__self__, "target_detector_recipes", target_detector_recipes)
        if target_id and not isinstance(target_id, str):
            raise TypeError("Expected argument 'target_id' to be a str")
        pulumi.set(__self__, "target_id", target_id)
        if target_resource_id and not isinstance(target_resource_id, str):
            raise TypeError("Expected argument 'target_resource_id' to be a str")
        pulumi.set(__self__, "target_resource_id", target_resource_id)
        if target_resource_type and not isinstance(target_resource_type, str):
            raise TypeError("Expected argument 'target_resource_type' to be a str")
        pulumi.set(__self__, "target_resource_type", target_resource_type)
        if target_responder_recipes and not isinstance(target_responder_recipes, list):
            raise TypeError("Expected argument 'target_responder_recipes' to be a list")
        pulumi.set(__self__, "target_responder_recipes", target_responder_recipes)
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
        Compartment Identifier
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="definedTags")
    def defined_tags(self) -> Mapping[str, Any]:
        """
        Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
        """
        return pulumi.get(self, "defined_tags")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        ResponderRule Description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        ResponderRule Display Name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="freeformTags")
    def freeform_tags(self) -> Mapping[str, Any]:
        """
        Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
        """
        return pulumi.get(self, "freeform_tags")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Unique identifier of TargetResponderRecipe that is immutable on creation
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inheritedByCompartments")
    def inherited_by_compartments(self) -> Sequence[str]:
        """
        List of inherited compartments
        """
        return pulumi.get(self, "inherited_by_compartments")

    @property
    @pulumi.getter(name="lifecyleDetails")
    def lifecyle_details(self) -> str:
        """
        A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
        """
        return pulumi.get(self, "lifecyle_details")

    @property
    @pulumi.getter(name="recipeCount")
    def recipe_count(self) -> int:
        """
        Total number of recipes attached to target
        """
        return pulumi.get(self, "recipe_count")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the ResponderRule.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="systemTags")
    def system_tags(self) -> Mapping[str, Any]:
        """
        System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
        """
        return pulumi.get(self, "system_tags")

    @property
    @pulumi.getter(name="targetDetectorRecipes")
    def target_detector_recipes(self) -> Sequence['outputs.GetTargetTargetDetectorRecipeResult']:
        """
        List of detector recipes associated with target
        """
        return pulumi.get(self, "target_detector_recipes")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> str:
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter(name="targetResourceId")
    def target_resource_id(self) -> str:
        """
        Resource ID which the target uses to monitor
        """
        return pulumi.get(self, "target_resource_id")

    @property
    @pulumi.getter(name="targetResourceType")
    def target_resource_type(self) -> str:
        """
        possible type of targets
        """
        return pulumi.get(self, "target_resource_type")

    @property
    @pulumi.getter(name="targetResponderRecipes")
    def target_responder_recipes(self) -> Sequence['outputs.GetTargetTargetResponderRecipeResult']:
        """
        List of responder recipes associated with target
        """
        return pulumi.get(self, "target_responder_recipes")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        The date and time the target was created. Format defined by RFC3339.
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter(name="timeUpdated")
    def time_updated(self) -> str:
        """
        The date and time the target was updated. Format defined by RFC3339.
        """
        return pulumi.get(self, "time_updated")


class AwaitableGetTargetResult(GetTargetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTargetResult(
            compartment_id=self.compartment_id,
            defined_tags=self.defined_tags,
            description=self.description,
            display_name=self.display_name,
            freeform_tags=self.freeform_tags,
            id=self.id,
            inherited_by_compartments=self.inherited_by_compartments,
            lifecyle_details=self.lifecyle_details,
            recipe_count=self.recipe_count,
            state=self.state,
            system_tags=self.system_tags,
            target_detector_recipes=self.target_detector_recipes,
            target_id=self.target_id,
            target_resource_id=self.target_resource_id,
            target_resource_type=self.target_resource_type,
            target_responder_recipes=self.target_responder_recipes,
            time_created=self.time_created,
            time_updated=self.time_updated)


def get_target(target_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTargetResult:
    """
    This data source provides details about a specific Target resource in Oracle Cloud Infrastructure Cloud Guard service.

    Returns a Target identified by targetId

    ## Example Usage

    ```python
    import pulumi
    import pulumi_oci as oci

    test_target = oci.cloudguard.get_target(target_id=oci_cloud_guard_target["test_target"]["id"])
    ```


    :param str target_id: OCID of target
    """
    __args__ = dict()
    __args__['targetId'] = target_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('oci:cloudguard/getTarget:getTarget', __args__, opts=opts, typ=GetTargetResult).value

    return AwaitableGetTargetResult(
        compartment_id=__ret__.compartment_id,
        defined_tags=__ret__.defined_tags,
        description=__ret__.description,
        display_name=__ret__.display_name,
        freeform_tags=__ret__.freeform_tags,
        id=__ret__.id,
        inherited_by_compartments=__ret__.inherited_by_compartments,
        lifecyle_details=__ret__.lifecyle_details,
        recipe_count=__ret__.recipe_count,
        state=__ret__.state,
        system_tags=__ret__.system_tags,
        target_detector_recipes=__ret__.target_detector_recipes,
        target_id=__ret__.target_id,
        target_resource_id=__ret__.target_resource_id,
        target_resource_type=__ret__.target_resource_type,
        target_responder_recipes=__ret__.target_responder_recipes,
        time_created=__ret__.time_created,
        time_updated=__ret__.time_updated)