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

__all__ = ['QueryArgs', 'Query']

@pulumi.input_type
class QueryArgs:
    def __init__(__self__, *,
                 compartment_id: pulumi.Input[str],
                 query_definition: pulumi.Input['QueryQueryDefinitionArgs']):
        """
        The set of arguments for constructing a Query resource.
        :param pulumi.Input[str] compartment_id: The compartment OCID.
        :param pulumi.Input['QueryQueryDefinitionArgs'] query_definition: (Updatable) The common fields for queries.
        """
        pulumi.set(__self__, "compartment_id", compartment_id)
        pulumi.set(__self__, "query_definition", query_definition)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> pulumi.Input[str]:
        """
        The compartment OCID.
        """
        return pulumi.get(self, "compartment_id")

    @compartment_id.setter
    def compartment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "compartment_id", value)

    @property
    @pulumi.getter(name="queryDefinition")
    def query_definition(self) -> pulumi.Input['QueryQueryDefinitionArgs']:
        """
        (Updatable) The common fields for queries.
        """
        return pulumi.get(self, "query_definition")

    @query_definition.setter
    def query_definition(self, value: pulumi.Input['QueryQueryDefinitionArgs']):
        pulumi.set(self, "query_definition", value)


@pulumi.input_type
class _QueryState:
    def __init__(__self__, *,
                 compartment_id: Optional[pulumi.Input[str]] = None,
                 query_definition: Optional[pulumi.Input['QueryQueryDefinitionArgs']] = None):
        """
        Input properties used for looking up and filtering Query resources.
        :param pulumi.Input[str] compartment_id: The compartment OCID.
        :param pulumi.Input['QueryQueryDefinitionArgs'] query_definition: (Updatable) The common fields for queries.
        """
        if compartment_id is not None:
            pulumi.set(__self__, "compartment_id", compartment_id)
        if query_definition is not None:
            pulumi.set(__self__, "query_definition", query_definition)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The compartment OCID.
        """
        return pulumi.get(self, "compartment_id")

    @compartment_id.setter
    def compartment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compartment_id", value)

    @property
    @pulumi.getter(name="queryDefinition")
    def query_definition(self) -> Optional[pulumi.Input['QueryQueryDefinitionArgs']]:
        """
        (Updatable) The common fields for queries.
        """
        return pulumi.get(self, "query_definition")

    @query_definition.setter
    def query_definition(self, value: Optional[pulumi.Input['QueryQueryDefinitionArgs']]):
        pulumi.set(self, "query_definition", value)


class Query(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compartment_id: Optional[pulumi.Input[str]] = None,
                 query_definition: Optional[pulumi.Input[pulumi.InputType['QueryQueryDefinitionArgs']]] = None,
                 __props__=None):
        """
        This resource provides the Query resource in Oracle Cloud Infrastructure Metering Computation service.

        Returns the created query.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_query = oci.meteringcomputation.Query("testQuery",
            compartment_id=var["compartment_id"],
            query_definition=oci.meteringcomputation.QueryQueryDefinitionArgs(
                cost_analysis_ui=oci.meteringcomputation.QueryQueryDefinitionCostAnalysisUiArgs(
                    graph=var["query_query_definition_cost_analysis_ui_graph"],
                    is_cumulative_graph=var["query_query_definition_cost_analysis_ui_is_cumulative_graph"],
                ),
                display_name=var["query_query_definition_display_name"],
                report_query=oci.meteringcomputation.QueryQueryDefinitionReportQueryArgs(
                    granularity=var["query_query_definition_report_query_granularity"],
                    tenant_id=oci_metering_computation_tenant["test_tenant"]["id"],
                    compartment_depth=var["query_query_definition_report_query_compartment_depth"],
                    date_range_name=var["query_query_definition_report_query_date_range_name"],
                    filter=var["query_query_definition_report_query_filter"],
                    forecast=oci.meteringcomputation.QueryQueryDefinitionReportQueryForecastArgs(
                        time_forecast_ended=var["query_query_definition_report_query_forecast_time_forecast_ended"],
                        forecast_type=var["query_query_definition_report_query_forecast_forecast_type"],
                        time_forecast_started=var["query_query_definition_report_query_forecast_time_forecast_started"],
                    ),
                    group_bies=var["query_query_definition_report_query_group_by"],
                    group_by_tags=[oci.meteringcomputation.QueryQueryDefinitionReportQueryGroupByTagArgs(
                        key=var["query_query_definition_report_query_group_by_tag_key"],
                        namespace=var["query_query_definition_report_query_group_by_tag_namespace"],
                        value=var["query_query_definition_report_query_group_by_tag_value"],
                    )],
                    is_aggregate_by_time=var["query_query_definition_report_query_is_aggregate_by_time"],
                    query_type=var["query_query_definition_report_query_query_type"],
                    time_usage_ended=var["query_query_definition_report_query_time_usage_ended"],
                    time_usage_started=var["query_query_definition_report_query_time_usage_started"],
                ),
                version=var["query_query_definition_version"],
            ))
        ```

        ## Import

        Queries can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:meteringcomputation/query:Query test_query "id"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compartment_id: The compartment OCID.
        :param pulumi.Input[pulumi.InputType['QueryQueryDefinitionArgs']] query_definition: (Updatable) The common fields for queries.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QueryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource provides the Query resource in Oracle Cloud Infrastructure Metering Computation service.

        Returns the created query.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_oci as oci

        test_query = oci.meteringcomputation.Query("testQuery",
            compartment_id=var["compartment_id"],
            query_definition=oci.meteringcomputation.QueryQueryDefinitionArgs(
                cost_analysis_ui=oci.meteringcomputation.QueryQueryDefinitionCostAnalysisUiArgs(
                    graph=var["query_query_definition_cost_analysis_ui_graph"],
                    is_cumulative_graph=var["query_query_definition_cost_analysis_ui_is_cumulative_graph"],
                ),
                display_name=var["query_query_definition_display_name"],
                report_query=oci.meteringcomputation.QueryQueryDefinitionReportQueryArgs(
                    granularity=var["query_query_definition_report_query_granularity"],
                    tenant_id=oci_metering_computation_tenant["test_tenant"]["id"],
                    compartment_depth=var["query_query_definition_report_query_compartment_depth"],
                    date_range_name=var["query_query_definition_report_query_date_range_name"],
                    filter=var["query_query_definition_report_query_filter"],
                    forecast=oci.meteringcomputation.QueryQueryDefinitionReportQueryForecastArgs(
                        time_forecast_ended=var["query_query_definition_report_query_forecast_time_forecast_ended"],
                        forecast_type=var["query_query_definition_report_query_forecast_forecast_type"],
                        time_forecast_started=var["query_query_definition_report_query_forecast_time_forecast_started"],
                    ),
                    group_bies=var["query_query_definition_report_query_group_by"],
                    group_by_tags=[oci.meteringcomputation.QueryQueryDefinitionReportQueryGroupByTagArgs(
                        key=var["query_query_definition_report_query_group_by_tag_key"],
                        namespace=var["query_query_definition_report_query_group_by_tag_namespace"],
                        value=var["query_query_definition_report_query_group_by_tag_value"],
                    )],
                    is_aggregate_by_time=var["query_query_definition_report_query_is_aggregate_by_time"],
                    query_type=var["query_query_definition_report_query_query_type"],
                    time_usage_ended=var["query_query_definition_report_query_time_usage_ended"],
                    time_usage_started=var["query_query_definition_report_query_time_usage_started"],
                ),
                version=var["query_query_definition_version"],
            ))
        ```

        ## Import

        Queries can be imported using the `id`, e.g.

        ```sh
         $ pulumi import oci:meteringcomputation/query:Query test_query "id"
        ```

        :param str resource_name: The name of the resource.
        :param QueryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QueryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compartment_id: Optional[pulumi.Input[str]] = None,
                 query_definition: Optional[pulumi.Input[pulumi.InputType['QueryQueryDefinitionArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QueryArgs.__new__(QueryArgs)

            if compartment_id is None and not opts.urn:
                raise TypeError("Missing required property 'compartment_id'")
            __props__.__dict__["compartment_id"] = compartment_id
            if query_definition is None and not opts.urn:
                raise TypeError("Missing required property 'query_definition'")
            __props__.__dict__["query_definition"] = query_definition
        super(Query, __self__).__init__(
            'oci:meteringcomputation/query:Query',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            compartment_id: Optional[pulumi.Input[str]] = None,
            query_definition: Optional[pulumi.Input[pulumi.InputType['QueryQueryDefinitionArgs']]] = None) -> 'Query':
        """
        Get an existing Query resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compartment_id: The compartment OCID.
        :param pulumi.Input[pulumi.InputType['QueryQueryDefinitionArgs']] query_definition: (Updatable) The common fields for queries.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QueryState.__new__(_QueryState)

        __props__.__dict__["compartment_id"] = compartment_id
        __props__.__dict__["query_definition"] = query_definition
        return Query(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> pulumi.Output[str]:
        """
        The compartment OCID.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="queryDefinition")
    def query_definition(self) -> pulumi.Output['outputs.QueryQueryDefinition']:
        """
        (Updatable) The common fields for queries.
        """
        return pulumi.get(self, "query_definition")
