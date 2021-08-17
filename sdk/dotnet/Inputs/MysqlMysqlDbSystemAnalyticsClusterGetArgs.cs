// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Inputs
{

    public sealed class MysqlMysqlDbSystemAnalyticsClusterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of analytics-processing compute instances, of the specified shape, in the HeatWave cluster.
        /// </summary>
        [Input("clusterSize")]
        public Input<int>? ClusterSize { get; set; }

        /// <summary>
        /// The name of the shape. The shape determines the resources allocated
        /// * CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20190415/ShapeSummary/ListShapes) operation.
        /// </summary>
        [Input("shapeName")]
        public Input<string>? ShapeName { get; set; }

        /// <summary>
        /// (Updatable) The target state for the DB System. Could be set to `ACTIVE` or `INACTIVE`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The date and time the DB System was created.
        /// </summary>
        [Input("timeCreated")]
        public Input<string>? TimeCreated { get; set; }

        /// <summary>
        /// The time the DB System was last updated.
        /// </summary>
        [Input("timeUpdated")]
        public Input<string>? TimeUpdated { get; set; }

        public MysqlMysqlDbSystemAnalyticsClusterGetArgs()
        {
        }
    }
}