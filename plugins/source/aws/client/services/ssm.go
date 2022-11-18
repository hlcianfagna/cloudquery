// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

//go:generate mockgen -package=mocks -destination=../mocks/ssm.go -source=ssm.go SsmClient
type SsmClient interface {
	DescribeActivations(context.Context, *ssm.DescribeActivationsInput, ...func(*ssm.Options)) (*ssm.DescribeActivationsOutput, error)
	DescribeAssociation(context.Context, *ssm.DescribeAssociationInput, ...func(*ssm.Options)) (*ssm.DescribeAssociationOutput, error)
	DescribeAssociationExecutionTargets(context.Context, *ssm.DescribeAssociationExecutionTargetsInput, ...func(*ssm.Options)) (*ssm.DescribeAssociationExecutionTargetsOutput, error)
	DescribeAssociationExecutions(context.Context, *ssm.DescribeAssociationExecutionsInput, ...func(*ssm.Options)) (*ssm.DescribeAssociationExecutionsOutput, error)
	DescribeAutomationExecutions(context.Context, *ssm.DescribeAutomationExecutionsInput, ...func(*ssm.Options)) (*ssm.DescribeAutomationExecutionsOutput, error)
	DescribeAutomationStepExecutions(context.Context, *ssm.DescribeAutomationStepExecutionsInput, ...func(*ssm.Options)) (*ssm.DescribeAutomationStepExecutionsOutput, error)
	DescribeAvailablePatches(context.Context, *ssm.DescribeAvailablePatchesInput, ...func(*ssm.Options)) (*ssm.DescribeAvailablePatchesOutput, error)
	DescribeDocument(context.Context, *ssm.DescribeDocumentInput, ...func(*ssm.Options)) (*ssm.DescribeDocumentOutput, error)
	DescribeDocumentPermission(context.Context, *ssm.DescribeDocumentPermissionInput, ...func(*ssm.Options)) (*ssm.DescribeDocumentPermissionOutput, error)
	DescribeEffectiveInstanceAssociations(context.Context, *ssm.DescribeEffectiveInstanceAssociationsInput, ...func(*ssm.Options)) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error)
	DescribeEffectivePatchesForPatchBaseline(context.Context, *ssm.DescribeEffectivePatchesForPatchBaselineInput, ...func(*ssm.Options)) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error)
	DescribeInstanceAssociationsStatus(context.Context, *ssm.DescribeInstanceAssociationsStatusInput, ...func(*ssm.Options)) (*ssm.DescribeInstanceAssociationsStatusOutput, error)
	DescribeInstanceInformation(context.Context, *ssm.DescribeInstanceInformationInput, ...func(*ssm.Options)) (*ssm.DescribeInstanceInformationOutput, error)
	DescribeInstancePatchStates(context.Context, *ssm.DescribeInstancePatchStatesInput, ...func(*ssm.Options)) (*ssm.DescribeInstancePatchStatesOutput, error)
	DescribeInstancePatchStatesForPatchGroup(context.Context, *ssm.DescribeInstancePatchStatesForPatchGroupInput, ...func(*ssm.Options)) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error)
	DescribeInstancePatches(context.Context, *ssm.DescribeInstancePatchesInput, ...func(*ssm.Options)) (*ssm.DescribeInstancePatchesOutput, error)
	DescribeInventoryDeletions(context.Context, *ssm.DescribeInventoryDeletionsInput, ...func(*ssm.Options)) (*ssm.DescribeInventoryDeletionsOutput, error)
	DescribeMaintenanceWindowExecutionTaskInvocations(context.Context, *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error)
	DescribeMaintenanceWindowExecutionTasks(context.Context, *ssm.DescribeMaintenanceWindowExecutionTasksInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error)
	DescribeMaintenanceWindowExecutions(context.Context, *ssm.DescribeMaintenanceWindowExecutionsInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error)
	DescribeMaintenanceWindowSchedule(context.Context, *ssm.DescribeMaintenanceWindowScheduleInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowScheduleOutput, error)
	DescribeMaintenanceWindowTargets(context.Context, *ssm.DescribeMaintenanceWindowTargetsInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowTargetsOutput, error)
	DescribeMaintenanceWindowTasks(context.Context, *ssm.DescribeMaintenanceWindowTasksInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowTasksOutput, error)
	DescribeMaintenanceWindows(context.Context, *ssm.DescribeMaintenanceWindowsInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowsOutput, error)
	DescribeMaintenanceWindowsForTarget(context.Context, *ssm.DescribeMaintenanceWindowsForTargetInput, ...func(*ssm.Options)) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error)
	DescribeOpsItems(context.Context, *ssm.DescribeOpsItemsInput, ...func(*ssm.Options)) (*ssm.DescribeOpsItemsOutput, error)
	DescribeParameters(context.Context, *ssm.DescribeParametersInput, ...func(*ssm.Options)) (*ssm.DescribeParametersOutput, error)
	DescribePatchBaselines(context.Context, *ssm.DescribePatchBaselinesInput, ...func(*ssm.Options)) (*ssm.DescribePatchBaselinesOutput, error)
	DescribePatchGroupState(context.Context, *ssm.DescribePatchGroupStateInput, ...func(*ssm.Options)) (*ssm.DescribePatchGroupStateOutput, error)
	DescribePatchGroups(context.Context, *ssm.DescribePatchGroupsInput, ...func(*ssm.Options)) (*ssm.DescribePatchGroupsOutput, error)
	DescribePatchProperties(context.Context, *ssm.DescribePatchPropertiesInput, ...func(*ssm.Options)) (*ssm.DescribePatchPropertiesOutput, error)
	DescribeSessions(context.Context, *ssm.DescribeSessionsInput, ...func(*ssm.Options)) (*ssm.DescribeSessionsOutput, error)
	GetAutomationExecution(context.Context, *ssm.GetAutomationExecutionInput, ...func(*ssm.Options)) (*ssm.GetAutomationExecutionOutput, error)
	GetCalendarState(context.Context, *ssm.GetCalendarStateInput, ...func(*ssm.Options)) (*ssm.GetCalendarStateOutput, error)
	GetCommandInvocation(context.Context, *ssm.GetCommandInvocationInput, ...func(*ssm.Options)) (*ssm.GetCommandInvocationOutput, error)
	GetConnectionStatus(context.Context, *ssm.GetConnectionStatusInput, ...func(*ssm.Options)) (*ssm.GetConnectionStatusOutput, error)
	GetDefaultPatchBaseline(context.Context, *ssm.GetDefaultPatchBaselineInput, ...func(*ssm.Options)) (*ssm.GetDefaultPatchBaselineOutput, error)
	GetDeployablePatchSnapshotForInstance(context.Context, *ssm.GetDeployablePatchSnapshotForInstanceInput, ...func(*ssm.Options)) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error)
	GetDocument(context.Context, *ssm.GetDocumentInput, ...func(*ssm.Options)) (*ssm.GetDocumentOutput, error)
	GetInventory(context.Context, *ssm.GetInventoryInput, ...func(*ssm.Options)) (*ssm.GetInventoryOutput, error)
	GetInventorySchema(context.Context, *ssm.GetInventorySchemaInput, ...func(*ssm.Options)) (*ssm.GetInventorySchemaOutput, error)
	GetMaintenanceWindow(context.Context, *ssm.GetMaintenanceWindowInput, ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowOutput, error)
	GetMaintenanceWindowExecution(context.Context, *ssm.GetMaintenanceWindowExecutionInput, ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowExecutionOutput, error)
	GetMaintenanceWindowExecutionTask(context.Context, *ssm.GetMaintenanceWindowExecutionTaskInput, ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error)
	GetMaintenanceWindowExecutionTaskInvocation(context.Context, *ssm.GetMaintenanceWindowExecutionTaskInvocationInput, ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error)
	GetMaintenanceWindowTask(context.Context, *ssm.GetMaintenanceWindowTaskInput, ...func(*ssm.Options)) (*ssm.GetMaintenanceWindowTaskOutput, error)
	GetOpsItem(context.Context, *ssm.GetOpsItemInput, ...func(*ssm.Options)) (*ssm.GetOpsItemOutput, error)
	GetOpsMetadata(context.Context, *ssm.GetOpsMetadataInput, ...func(*ssm.Options)) (*ssm.GetOpsMetadataOutput, error)
	GetOpsSummary(context.Context, *ssm.GetOpsSummaryInput, ...func(*ssm.Options)) (*ssm.GetOpsSummaryOutput, error)
	GetParameter(context.Context, *ssm.GetParameterInput, ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
	GetParameterHistory(context.Context, *ssm.GetParameterHistoryInput, ...func(*ssm.Options)) (*ssm.GetParameterHistoryOutput, error)
	GetParameters(context.Context, *ssm.GetParametersInput, ...func(*ssm.Options)) (*ssm.GetParametersOutput, error)
	GetParametersByPath(context.Context, *ssm.GetParametersByPathInput, ...func(*ssm.Options)) (*ssm.GetParametersByPathOutput, error)
	GetPatchBaseline(context.Context, *ssm.GetPatchBaselineInput, ...func(*ssm.Options)) (*ssm.GetPatchBaselineOutput, error)
	GetPatchBaselineForPatchGroup(context.Context, *ssm.GetPatchBaselineForPatchGroupInput, ...func(*ssm.Options)) (*ssm.GetPatchBaselineForPatchGroupOutput, error)
	GetResourcePolicies(context.Context, *ssm.GetResourcePoliciesInput, ...func(*ssm.Options)) (*ssm.GetResourcePoliciesOutput, error)
	GetServiceSetting(context.Context, *ssm.GetServiceSettingInput, ...func(*ssm.Options)) (*ssm.GetServiceSettingOutput, error)
	ListAssociationVersions(context.Context, *ssm.ListAssociationVersionsInput, ...func(*ssm.Options)) (*ssm.ListAssociationVersionsOutput, error)
	ListAssociations(context.Context, *ssm.ListAssociationsInput, ...func(*ssm.Options)) (*ssm.ListAssociationsOutput, error)
	ListCommandInvocations(context.Context, *ssm.ListCommandInvocationsInput, ...func(*ssm.Options)) (*ssm.ListCommandInvocationsOutput, error)
	ListCommands(context.Context, *ssm.ListCommandsInput, ...func(*ssm.Options)) (*ssm.ListCommandsOutput, error)
	ListComplianceItems(context.Context, *ssm.ListComplianceItemsInput, ...func(*ssm.Options)) (*ssm.ListComplianceItemsOutput, error)
	ListComplianceSummaries(context.Context, *ssm.ListComplianceSummariesInput, ...func(*ssm.Options)) (*ssm.ListComplianceSummariesOutput, error)
	ListDocumentMetadataHistory(context.Context, *ssm.ListDocumentMetadataHistoryInput, ...func(*ssm.Options)) (*ssm.ListDocumentMetadataHistoryOutput, error)
	ListDocumentVersions(context.Context, *ssm.ListDocumentVersionsInput, ...func(*ssm.Options)) (*ssm.ListDocumentVersionsOutput, error)
	ListDocuments(context.Context, *ssm.ListDocumentsInput, ...func(*ssm.Options)) (*ssm.ListDocumentsOutput, error)
	ListInventoryEntries(context.Context, *ssm.ListInventoryEntriesInput, ...func(*ssm.Options)) (*ssm.ListInventoryEntriesOutput, error)
	ListOpsItemEvents(context.Context, *ssm.ListOpsItemEventsInput, ...func(*ssm.Options)) (*ssm.ListOpsItemEventsOutput, error)
	ListOpsItemRelatedItems(context.Context, *ssm.ListOpsItemRelatedItemsInput, ...func(*ssm.Options)) (*ssm.ListOpsItemRelatedItemsOutput, error)
	ListOpsMetadata(context.Context, *ssm.ListOpsMetadataInput, ...func(*ssm.Options)) (*ssm.ListOpsMetadataOutput, error)
	ListResourceComplianceSummaries(context.Context, *ssm.ListResourceComplianceSummariesInput, ...func(*ssm.Options)) (*ssm.ListResourceComplianceSummariesOutput, error)
	ListResourceDataSync(context.Context, *ssm.ListResourceDataSyncInput, ...func(*ssm.Options)) (*ssm.ListResourceDataSyncOutput, error)
	ListTagsForResource(context.Context, *ssm.ListTagsForResourceInput, ...func(*ssm.Options)) (*ssm.ListTagsForResourceOutput, error)
}
