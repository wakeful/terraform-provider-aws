// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logs

// Exports for use in tests only.
var (
	ResourceAccountPolicy        = resourceAccountPolicy
	ResourceDataProtectionPolicy = resourceDataProtectionPolicy
	ResourceDestination          = resourceDestination
	ResourceDestinationPolicy    = resourceDestinationPolicy
	ResourceGroup                = resourceGroup
	ResourceMetricFilter         = resourceMetricFilter
	ResourceQueryDefinition      = resourceQueryDefinition
	ResourceResourcePolicy       = resourceResourcePolicy
	ResourceStream               = resourceStream
	ResourceSubscriptionFilter   = resourceSubscriptionFilter
	ResourceAnomalyDetector      = newResourceAnomalyDetector

	FindAccountPolicyByTwoPartKey      = findAccountPolicyByTwoPartKey
	FindDestinationByName              = findDestinationByName
	FindLogGroupByName                 = findLogGroupByName
	FindLogStreamByTwoPartKey          = findLogStreamByTwoPartKey // nosemgrep:ci.logs-in-var-name
	FindMetricFilterByTwoPartKey       = findMetricFilterByTwoPartKey
	FindIndexPolicyByLogGroupName      = findIndexPolicyByLogGroupName
	FindQueryDefinitionByTwoPartKey    = findQueryDefinitionByTwoPartKey
	FindResourcePolicyByName           = findResourcePolicyByName
	FindSubscriptionFilterByTwoPartKey = findSubscriptionFilterByTwoPartKey
	FindLogAnomalyDetectorByARN        = findLogAnomalyDetectorByARN

	ValidLogGroupName                      = validLogGroupName
	ValidLogGroupNamePrefix                = validLogGroupNamePrefix
	ValidLogMetricFilterName               = validLogMetricFilterName
	ValidLogMetricFilterTransformationName = validLogMetricFilterTransformationName
	ValidStreamName                        = validStreamName
)
