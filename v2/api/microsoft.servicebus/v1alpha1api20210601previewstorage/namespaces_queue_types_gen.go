// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601previewstorage

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/microsoft.servicebus/v1alpha1api20210101previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210601preview.NamespacesQueue
//Generated from: https://schema.management.azure.com/schemas/2021-06-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_queues
type NamespacesQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesQueues_Spec `json:"spec,omitempty"`
	Status            SBQueue_Status        `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesQueue{}

// GetConditions returns the conditions of the resource
func (namespacesQueue *NamespacesQueue) GetConditions() conditions.Conditions {
	return namespacesQueue.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (namespacesQueue *NamespacesQueue) SetConditions(conditions conditions.Conditions) {
	namespacesQueue.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesQueue{}

// ConvertFrom populates our NamespacesQueue from the provided hub NamespacesQueue
func (namespacesQueue *NamespacesQueue) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20210101previewstorage.NamespacesQueue)
	if !ok {
		return fmt.Errorf("expected storage:microsoft.servicebus/v1alpha1api20210101previewstorage/NamespacesQueue but received %T instead", hub)
	}

	return namespacesQueue.AssignPropertiesFromNamespacesQueue(source)
}

// ConvertTo populates the provided hub NamespacesQueue from our NamespacesQueue
func (namespacesQueue *NamespacesQueue) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20210101previewstorage.NamespacesQueue)
	if !ok {
		return fmt.Errorf("expected storage:microsoft.servicebus/v1alpha1api20210101previewstorage/NamespacesQueue but received %T instead", hub)
	}

	return namespacesQueue.AssignPropertiesToNamespacesQueue(destination)
}

var _ genruntime.KubernetesResource = &NamespacesQueue{}

// AzureName returns the Azure name of the resource
func (namespacesQueue *NamespacesQueue) AzureName() string {
	return namespacesQueue.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (namespacesQueue *NamespacesQueue) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (namespacesQueue *NamespacesQueue) GetSpec() genruntime.ConvertibleSpec {
	return &namespacesQueue.Spec
}

// GetStatus returns the status of this resource
func (namespacesQueue *NamespacesQueue) GetStatus() genruntime.ConvertibleStatus {
	return &namespacesQueue.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/queues"
func (namespacesQueue *NamespacesQueue) GetType() string {
	return "Microsoft.ServiceBus/namespaces/queues"
}

// NewEmptyStatus returns a new empty (blank) status
func (namespacesQueue *NamespacesQueue) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SBQueue_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (namespacesQueue *NamespacesQueue) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(namespacesQueue.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: namespacesQueue.Namespace,
		Name:      namespacesQueue.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (namespacesQueue *NamespacesQueue) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SBQueue_Status); ok {
		namespacesQueue.Status = *st
		return nil
	}

	// Convert status to required version
	var st SBQueue_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	namespacesQueue.Status = st
	return nil
}

// AssignPropertiesFromNamespacesQueue populates our NamespacesQueue from the provided source NamespacesQueue
func (namespacesQueue *NamespacesQueue) AssignPropertiesFromNamespacesQueue(source *v1alpha1api20210101previewstorage.NamespacesQueue) error {

	// Spec
	var spec NamespacesQueues_Spec
	err := spec.AssignPropertiesFromNamespacesQueuesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesFromNamespacesQueuesSpec()")
	}
	namespacesQueue.Spec = spec

	// Status
	var status SBQueue_Status
	err = status.AssignPropertiesFromSBQueueStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesFromSBQueueStatus()")
	}
	namespacesQueue.Status = status

	// No error
	return nil
}

// AssignPropertiesToNamespacesQueue populates the provided destination NamespacesQueue from our NamespacesQueue
func (namespacesQueue *NamespacesQueue) AssignPropertiesToNamespacesQueue(destination *v1alpha1api20210101previewstorage.NamespacesQueue) error {

	// Spec
	var spec v1alpha1api20210101previewstorage.NamespacesQueues_Spec
	err := namespacesQueue.Spec.AssignPropertiesToNamespacesQueuesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesToNamespacesQueuesSpec()")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210101previewstorage.SBQueue_Status
	err = namespacesQueue.Status.AssignPropertiesToSBQueueStatus(&status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesToSBQueueStatus()")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (namespacesQueue *NamespacesQueue) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: namespacesQueue.Spec.OriginalVersion,
		Kind:    "NamespacesQueue",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210601preview.NamespacesQueue
//Generated from: https://schema.management.azure.com/schemas/2021-06-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_queues
type NamespacesQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesQueue `json:"items"`
}

//Storage version of v1alpha1api20210601preview.NamespacesQueues_Spec
type NamespacesQueues_Spec struct {
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// +kubebuilder:validation:MinLength=1
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                           string  `json:"azureName"`
	DeadLetteringOnMessageExpiration    *bool   `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive            *string `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool   `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool   `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool   `json:"enablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo       *string `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                           *string `json:"forwardTo,omitempty"`
	Location                            *string `json:"location,omitempty"`
	LockDuration                        *string `json:"lockDuration,omitempty"`
	MaxDeliveryCount                    *int    `json:"maxDeliveryCount,omitempty"`
	MaxMessageSizeInKilobytes           *int    `json:"maxMessageSizeInKilobytes,omitempty"`
	MaxSizeInMegabytes                  *int    `json:"maxSizeInMegabytes,omitempty"`
	OriginalVersion                     string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner                      genruntime.KnownResourceReference `group:"microsoft.servicebus.azure.com" json:"owner" kind:"Namespace"`
	PropertyBag                genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection *bool                             `json:"requiresDuplicateDetection,omitempty"`
	RequiresSession            *bool                             `json:"requiresSession,omitempty"`
	Tags                       map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesQueues_Spec{}

// ConvertSpecFrom populates our NamespacesQueues_Spec from the provided source
func (namespacesQueuesSpec *NamespacesQueues_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210101previewstorage.NamespacesQueues_Spec)
	if ok {
		// Populate our instance from source
		return namespacesQueuesSpec.AssignPropertiesFromNamespacesQueuesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210101previewstorage.NamespacesQueues_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = namespacesQueuesSpec.AssignPropertiesFromNamespacesQueuesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesQueues_Spec
func (namespacesQueuesSpec *NamespacesQueues_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210101previewstorage.NamespacesQueues_Spec)
	if ok {
		// Populate destination from our instance
		return namespacesQueuesSpec.AssignPropertiesToNamespacesQueuesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210101previewstorage.NamespacesQueues_Spec{}
	err := namespacesQueuesSpec.AssignPropertiesToNamespacesQueuesSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromNamespacesQueuesSpec populates our NamespacesQueues_Spec from the provided source NamespacesQueues_Spec
func (namespacesQueuesSpec *NamespacesQueues_Spec) AssignPropertiesFromNamespacesQueuesSpec(source *v1alpha1api20210101previewstorage.NamespacesQueues_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AutoDeleteOnIdle
	namespacesQueuesSpec.AutoDeleteOnIdle = genruntime.ClonePointerToString(source.AutoDeleteOnIdle)

	// AzureName
	namespacesQueuesSpec.AzureName = source.AzureName

	// DeadLetteringOnMessageExpiration
	if source.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *source.DeadLetteringOnMessageExpiration
		namespacesQueuesSpec.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		namespacesQueuesSpec.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	namespacesQueuesSpec.DefaultMessageTimeToLive = genruntime.ClonePointerToString(source.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	namespacesQueuesSpec.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(source.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if source.EnableBatchedOperations != nil {
		enableBatchedOperation := *source.EnableBatchedOperations
		namespacesQueuesSpec.EnableBatchedOperations = &enableBatchedOperation
	} else {
		namespacesQueuesSpec.EnableBatchedOperations = nil
	}

	// EnableExpress
	if source.EnableExpress != nil {
		enableExpress := *source.EnableExpress
		namespacesQueuesSpec.EnableExpress = &enableExpress
	} else {
		namespacesQueuesSpec.EnableExpress = nil
	}

	// EnablePartitioning
	if source.EnablePartitioning != nil {
		enablePartitioning := *source.EnablePartitioning
		namespacesQueuesSpec.EnablePartitioning = &enablePartitioning
	} else {
		namespacesQueuesSpec.EnablePartitioning = nil
	}

	// ForwardDeadLetteredMessagesTo
	namespacesQueuesSpec.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(source.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	namespacesQueuesSpec.ForwardTo = genruntime.ClonePointerToString(source.ForwardTo)

	// Location
	namespacesQueuesSpec.Location = genruntime.ClonePointerToString(source.Location)

	// LockDuration
	namespacesQueuesSpec.LockDuration = genruntime.ClonePointerToString(source.LockDuration)

	// MaxDeliveryCount
	namespacesQueuesSpec.MaxDeliveryCount = genruntime.ClonePointerToInt(source.MaxDeliveryCount)

	// MaxMessageSizeInKilobytes
	if propertyBag.Contains("MaxMessageSizeInKilobytes") {
		var maxMessageSizeInKilobyte int
		err := propertyBag.Pull("MaxMessageSizeInKilobytes", &maxMessageSizeInKilobyte)
		if err != nil {
			return errors.Wrap(err, "pulling 'MaxMessageSizeInKilobytes' from propertyBag")
		}

		namespacesQueuesSpec.MaxMessageSizeInKilobytes = &maxMessageSizeInKilobyte
	} else {
		namespacesQueuesSpec.MaxMessageSizeInKilobytes = nil
	}

	// MaxSizeInMegabytes
	namespacesQueuesSpec.MaxSizeInMegabytes = genruntime.ClonePointerToInt(source.MaxSizeInMegabytes)

	// OriginalVersion
	namespacesQueuesSpec.OriginalVersion = source.OriginalVersion

	// Owner
	namespacesQueuesSpec.Owner = source.Owner.Copy()

	// RequiresDuplicateDetection
	if source.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *source.RequiresDuplicateDetection
		namespacesQueuesSpec.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		namespacesQueuesSpec.RequiresDuplicateDetection = nil
	}

	// RequiresSession
	if source.RequiresSession != nil {
		requiresSession := *source.RequiresSession
		namespacesQueuesSpec.RequiresSession = &requiresSession
	} else {
		namespacesQueuesSpec.RequiresSession = nil
	}

	// Tags
	namespacesQueuesSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		namespacesQueuesSpec.PropertyBag = propertyBag
	} else {
		namespacesQueuesSpec.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNamespacesQueuesSpec populates the provided destination NamespacesQueues_Spec from our NamespacesQueues_Spec
func (namespacesQueuesSpec *NamespacesQueues_Spec) AssignPropertiesToNamespacesQueuesSpec(destination *v1alpha1api20210101previewstorage.NamespacesQueues_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(namespacesQueuesSpec.PropertyBag)

	// AutoDeleteOnIdle
	destination.AutoDeleteOnIdle = genruntime.ClonePointerToString(namespacesQueuesSpec.AutoDeleteOnIdle)

	// AzureName
	destination.AzureName = namespacesQueuesSpec.AzureName

	// DeadLetteringOnMessageExpiration
	if namespacesQueuesSpec.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *namespacesQueuesSpec.DeadLetteringOnMessageExpiration
		destination.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		destination.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	destination.DefaultMessageTimeToLive = genruntime.ClonePointerToString(namespacesQueuesSpec.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	destination.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(namespacesQueuesSpec.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if namespacesQueuesSpec.EnableBatchedOperations != nil {
		enableBatchedOperation := *namespacesQueuesSpec.EnableBatchedOperations
		destination.EnableBatchedOperations = &enableBatchedOperation
	} else {
		destination.EnableBatchedOperations = nil
	}

	// EnableExpress
	if namespacesQueuesSpec.EnableExpress != nil {
		enableExpress := *namespacesQueuesSpec.EnableExpress
		destination.EnableExpress = &enableExpress
	} else {
		destination.EnableExpress = nil
	}

	// EnablePartitioning
	if namespacesQueuesSpec.EnablePartitioning != nil {
		enablePartitioning := *namespacesQueuesSpec.EnablePartitioning
		destination.EnablePartitioning = &enablePartitioning
	} else {
		destination.EnablePartitioning = nil
	}

	// ForwardDeadLetteredMessagesTo
	destination.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(namespacesQueuesSpec.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	destination.ForwardTo = genruntime.ClonePointerToString(namespacesQueuesSpec.ForwardTo)

	// Location
	destination.Location = genruntime.ClonePointerToString(namespacesQueuesSpec.Location)

	// LockDuration
	destination.LockDuration = genruntime.ClonePointerToString(namespacesQueuesSpec.LockDuration)

	// MaxDeliveryCount
	destination.MaxDeliveryCount = genruntime.ClonePointerToInt(namespacesQueuesSpec.MaxDeliveryCount)

	// MaxMessageSizeInKilobytes
	if namespacesQueuesSpec.MaxMessageSizeInKilobytes != nil {
		propertyBag.Add("MaxMessageSizeInKilobytes", *namespacesQueuesSpec.MaxMessageSizeInKilobytes)
	}

	// MaxSizeInMegabytes
	destination.MaxSizeInMegabytes = genruntime.ClonePointerToInt(namespacesQueuesSpec.MaxSizeInMegabytes)

	// OriginalVersion
	destination.OriginalVersion = namespacesQueuesSpec.OriginalVersion

	// Owner
	destination.Owner = namespacesQueuesSpec.Owner.Copy()

	// RequiresDuplicateDetection
	if namespacesQueuesSpec.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *namespacesQueuesSpec.RequiresDuplicateDetection
		destination.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		destination.RequiresDuplicateDetection = nil
	}

	// RequiresSession
	if namespacesQueuesSpec.RequiresSession != nil {
		requiresSession := *namespacesQueuesSpec.RequiresSession
		destination.RequiresSession = &requiresSession
	} else {
		destination.RequiresSession = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(namespacesQueuesSpec.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20210601preview.SBQueue_Status
//Generated from:
type SBQueue_Status struct {
	AccessedAt                          *string                     `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                     `json:"autoDeleteOnIdle,omitempty"`
	Conditions                          []conditions.Condition      `json:"conditions,omitempty"`
	CountDetails                        *MessageCountDetails_Status `json:"countDetails,omitempty"`
	CreatedAt                           *string                     `json:"createdAt,omitempty"`
	DeadLetteringOnMessageExpiration    *bool                       `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive            *string                     `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                     `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                       `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                       `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                       `json:"enablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo       *string                     `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                           *string                     `json:"forwardTo,omitempty"`
	Id                                  *string                     `json:"id,omitempty"`
	LockDuration                        *string                     `json:"lockDuration,omitempty"`
	MaxDeliveryCount                    *int                        `json:"maxDeliveryCount,omitempty"`
	MaxMessageSizeInKilobytes           *int                        `json:"maxMessageSizeInKilobytes,omitempty"`
	MaxSizeInMegabytes                  *int                        `json:"maxSizeInMegabytes,omitempty"`
	MessageCount                        *int                        `json:"messageCount,omitempty"`
	Name                                *string                     `json:"name,omitempty"`
	PropertyBag                         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection          *bool                       `json:"requiresDuplicateDetection,omitempty"`
	RequiresSession                     *bool                       `json:"requiresSession,omitempty"`
	SizeInBytes                         *int                        `json:"sizeInBytes,omitempty"`
	Status                              *string                     `json:"status,omitempty"`
	SystemData                          *SystemData_Status          `json:"systemData,omitempty"`
	Type                                *string                     `json:"type,omitempty"`
	UpdatedAt                           *string                     `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SBQueue_Status{}

// ConvertStatusFrom populates our SBQueue_Status from the provided source
func (sbQueueStatus *SBQueue_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20210101previewstorage.SBQueue_Status)
	if ok {
		// Populate our instance from source
		return sbQueueStatus.AssignPropertiesFromSBQueueStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210101previewstorage.SBQueue_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = sbQueueStatus.AssignPropertiesFromSBQueueStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SBQueue_Status
func (sbQueueStatus *SBQueue_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20210101previewstorage.SBQueue_Status)
	if ok {
		// Populate destination from our instance
		return sbQueueStatus.AssignPropertiesToSBQueueStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210101previewstorage.SBQueue_Status{}
	err := sbQueueStatus.AssignPropertiesToSBQueueStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignPropertiesFromSBQueueStatus populates our SBQueue_Status from the provided source SBQueue_Status
func (sbQueueStatus *SBQueue_Status) AssignPropertiesFromSBQueueStatus(source *v1alpha1api20210101previewstorage.SBQueue_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AccessedAt
	sbQueueStatus.AccessedAt = genruntime.ClonePointerToString(source.AccessedAt)

	// AutoDeleteOnIdle
	sbQueueStatus.AutoDeleteOnIdle = genruntime.ClonePointerToString(source.AutoDeleteOnIdle)

	// Conditions
	sbQueueStatus.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CountDetails
	if source.CountDetails != nil {
		var countDetail MessageCountDetails_Status
		err := countDetail.AssignPropertiesFromMessageCountDetailsStatus(source.CountDetails)
		if err != nil {
			return errors.Wrap(err, "populating CountDetails from CountDetails, calling AssignPropertiesFromMessageCountDetailsStatus()")
		}
		sbQueueStatus.CountDetails = &countDetail
	} else {
		sbQueueStatus.CountDetails = nil
	}

	// CreatedAt
	sbQueueStatus.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// DeadLetteringOnMessageExpiration
	if source.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *source.DeadLetteringOnMessageExpiration
		sbQueueStatus.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		sbQueueStatus.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	sbQueueStatus.DefaultMessageTimeToLive = genruntime.ClonePointerToString(source.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	sbQueueStatus.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(source.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if source.EnableBatchedOperations != nil {
		enableBatchedOperation := *source.EnableBatchedOperations
		sbQueueStatus.EnableBatchedOperations = &enableBatchedOperation
	} else {
		sbQueueStatus.EnableBatchedOperations = nil
	}

	// EnableExpress
	if source.EnableExpress != nil {
		enableExpress := *source.EnableExpress
		sbQueueStatus.EnableExpress = &enableExpress
	} else {
		sbQueueStatus.EnableExpress = nil
	}

	// EnablePartitioning
	if source.EnablePartitioning != nil {
		enablePartitioning := *source.EnablePartitioning
		sbQueueStatus.EnablePartitioning = &enablePartitioning
	} else {
		sbQueueStatus.EnablePartitioning = nil
	}

	// ForwardDeadLetteredMessagesTo
	sbQueueStatus.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(source.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	sbQueueStatus.ForwardTo = genruntime.ClonePointerToString(source.ForwardTo)

	// Id
	sbQueueStatus.Id = genruntime.ClonePointerToString(source.Id)

	// LockDuration
	sbQueueStatus.LockDuration = genruntime.ClonePointerToString(source.LockDuration)

	// MaxDeliveryCount
	sbQueueStatus.MaxDeliveryCount = genruntime.ClonePointerToInt(source.MaxDeliveryCount)

	// MaxMessageSizeInKilobytes
	if propertyBag.Contains("MaxMessageSizeInKilobytes") {
		var maxMessageSizeInKilobyte int
		err := propertyBag.Pull("MaxMessageSizeInKilobytes", &maxMessageSizeInKilobyte)
		if err != nil {
			return errors.Wrap(err, "pulling 'MaxMessageSizeInKilobytes' from propertyBag")
		}

		sbQueueStatus.MaxMessageSizeInKilobytes = &maxMessageSizeInKilobyte
	} else {
		sbQueueStatus.MaxMessageSizeInKilobytes = nil
	}

	// MaxSizeInMegabytes
	sbQueueStatus.MaxSizeInMegabytes = genruntime.ClonePointerToInt(source.MaxSizeInMegabytes)

	// MessageCount
	sbQueueStatus.MessageCount = genruntime.ClonePointerToInt(source.MessageCount)

	// Name
	sbQueueStatus.Name = genruntime.ClonePointerToString(source.Name)

	// RequiresDuplicateDetection
	if source.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *source.RequiresDuplicateDetection
		sbQueueStatus.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		sbQueueStatus.RequiresDuplicateDetection = nil
	}

	// RequiresSession
	if source.RequiresSession != nil {
		requiresSession := *source.RequiresSession
		sbQueueStatus.RequiresSession = &requiresSession
	} else {
		sbQueueStatus.RequiresSession = nil
	}

	// SizeInBytes
	sbQueueStatus.SizeInBytes = genruntime.ClonePointerToInt(source.SizeInBytes)

	// Status
	sbQueueStatus.Status = genruntime.ClonePointerToString(source.Status)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "populating SystemData from SystemData, calling AssignPropertiesFromSystemDataStatus()")
		}
		sbQueueStatus.SystemData = &systemDatum
	} else {
		sbQueueStatus.SystemData = nil
	}

	// Type
	sbQueueStatus.Type = genruntime.ClonePointerToString(source.Type)

	// UpdatedAt
	sbQueueStatus.UpdatedAt = genruntime.ClonePointerToString(source.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		sbQueueStatus.PropertyBag = propertyBag
	} else {
		sbQueueStatus.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSBQueueStatus populates the provided destination SBQueue_Status from our SBQueue_Status
func (sbQueueStatus *SBQueue_Status) AssignPropertiesToSBQueueStatus(destination *v1alpha1api20210101previewstorage.SBQueue_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(sbQueueStatus.PropertyBag)

	// AccessedAt
	destination.AccessedAt = genruntime.ClonePointerToString(sbQueueStatus.AccessedAt)

	// AutoDeleteOnIdle
	destination.AutoDeleteOnIdle = genruntime.ClonePointerToString(sbQueueStatus.AutoDeleteOnIdle)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(sbQueueStatus.Conditions)

	// CountDetails
	if sbQueueStatus.CountDetails != nil {
		var countDetail v1alpha1api20210101previewstorage.MessageCountDetails_Status
		err := (*sbQueueStatus.CountDetails).AssignPropertiesToMessageCountDetailsStatus(&countDetail)
		if err != nil {
			return errors.Wrap(err, "populating CountDetails from CountDetails, calling AssignPropertiesToMessageCountDetailsStatus()")
		}
		destination.CountDetails = &countDetail
	} else {
		destination.CountDetails = nil
	}

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(sbQueueStatus.CreatedAt)

	// DeadLetteringOnMessageExpiration
	if sbQueueStatus.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *sbQueueStatus.DeadLetteringOnMessageExpiration
		destination.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		destination.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	destination.DefaultMessageTimeToLive = genruntime.ClonePointerToString(sbQueueStatus.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	destination.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(sbQueueStatus.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if sbQueueStatus.EnableBatchedOperations != nil {
		enableBatchedOperation := *sbQueueStatus.EnableBatchedOperations
		destination.EnableBatchedOperations = &enableBatchedOperation
	} else {
		destination.EnableBatchedOperations = nil
	}

	// EnableExpress
	if sbQueueStatus.EnableExpress != nil {
		enableExpress := *sbQueueStatus.EnableExpress
		destination.EnableExpress = &enableExpress
	} else {
		destination.EnableExpress = nil
	}

	// EnablePartitioning
	if sbQueueStatus.EnablePartitioning != nil {
		enablePartitioning := *sbQueueStatus.EnablePartitioning
		destination.EnablePartitioning = &enablePartitioning
	} else {
		destination.EnablePartitioning = nil
	}

	// ForwardDeadLetteredMessagesTo
	destination.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(sbQueueStatus.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	destination.ForwardTo = genruntime.ClonePointerToString(sbQueueStatus.ForwardTo)

	// Id
	destination.Id = genruntime.ClonePointerToString(sbQueueStatus.Id)

	// LockDuration
	destination.LockDuration = genruntime.ClonePointerToString(sbQueueStatus.LockDuration)

	// MaxDeliveryCount
	destination.MaxDeliveryCount = genruntime.ClonePointerToInt(sbQueueStatus.MaxDeliveryCount)

	// MaxMessageSizeInKilobytes
	if sbQueueStatus.MaxMessageSizeInKilobytes != nil {
		propertyBag.Add("MaxMessageSizeInKilobytes", *sbQueueStatus.MaxMessageSizeInKilobytes)
	}

	// MaxSizeInMegabytes
	destination.MaxSizeInMegabytes = genruntime.ClonePointerToInt(sbQueueStatus.MaxSizeInMegabytes)

	// MessageCount
	destination.MessageCount = genruntime.ClonePointerToInt(sbQueueStatus.MessageCount)

	// Name
	destination.Name = genruntime.ClonePointerToString(sbQueueStatus.Name)

	// RequiresDuplicateDetection
	if sbQueueStatus.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *sbQueueStatus.RequiresDuplicateDetection
		destination.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		destination.RequiresDuplicateDetection = nil
	}

	// RequiresSession
	if sbQueueStatus.RequiresSession != nil {
		requiresSession := *sbQueueStatus.RequiresSession
		destination.RequiresSession = &requiresSession
	} else {
		destination.RequiresSession = nil
	}

	// SizeInBytes
	destination.SizeInBytes = genruntime.ClonePointerToInt(sbQueueStatus.SizeInBytes)

	// Status
	destination.Status = genruntime.ClonePointerToString(sbQueueStatus.Status)

	// SystemData
	if sbQueueStatus.SystemData != nil {
		var systemDatum v1alpha1api20210101previewstorage.SystemData_Status
		err := (*sbQueueStatus.SystemData).AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "populating SystemData from SystemData, calling AssignPropertiesToSystemDataStatus()")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(sbQueueStatus.Type)

	// UpdatedAt
	destination.UpdatedAt = genruntime.ClonePointerToString(sbQueueStatus.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20210601preview.MessageCountDetails_Status
//Generated from:
type MessageCountDetails_Status struct {
	ActiveMessageCount             *int                   `json:"activeMessageCount,omitempty"`
	DeadLetterMessageCount         *int                   `json:"deadLetterMessageCount,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ScheduledMessageCount          *int                   `json:"scheduledMessageCount,omitempty"`
	TransferDeadLetterMessageCount *int                   `json:"transferDeadLetterMessageCount,omitempty"`
	TransferMessageCount           *int                   `json:"transferMessageCount,omitempty"`
}

// AssignPropertiesFromMessageCountDetailsStatus populates our MessageCountDetails_Status from the provided source MessageCountDetails_Status
func (messageCountDetailsStatus *MessageCountDetails_Status) AssignPropertiesFromMessageCountDetailsStatus(source *v1alpha1api20210101previewstorage.MessageCountDetails_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ActiveMessageCount
	messageCountDetailsStatus.ActiveMessageCount = genruntime.ClonePointerToInt(source.ActiveMessageCount)

	// DeadLetterMessageCount
	messageCountDetailsStatus.DeadLetterMessageCount = genruntime.ClonePointerToInt(source.DeadLetterMessageCount)

	// ScheduledMessageCount
	messageCountDetailsStatus.ScheduledMessageCount = genruntime.ClonePointerToInt(source.ScheduledMessageCount)

	// TransferDeadLetterMessageCount
	messageCountDetailsStatus.TransferDeadLetterMessageCount = genruntime.ClonePointerToInt(source.TransferDeadLetterMessageCount)

	// TransferMessageCount
	messageCountDetailsStatus.TransferMessageCount = genruntime.ClonePointerToInt(source.TransferMessageCount)

	// Update the property bag
	if len(propertyBag) > 0 {
		messageCountDetailsStatus.PropertyBag = propertyBag
	} else {
		messageCountDetailsStatus.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToMessageCountDetailsStatus populates the provided destination MessageCountDetails_Status from our MessageCountDetails_Status
func (messageCountDetailsStatus *MessageCountDetails_Status) AssignPropertiesToMessageCountDetailsStatus(destination *v1alpha1api20210101previewstorage.MessageCountDetails_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(messageCountDetailsStatus.PropertyBag)

	// ActiveMessageCount
	destination.ActiveMessageCount = genruntime.ClonePointerToInt(messageCountDetailsStatus.ActiveMessageCount)

	// DeadLetterMessageCount
	destination.DeadLetterMessageCount = genruntime.ClonePointerToInt(messageCountDetailsStatus.DeadLetterMessageCount)

	// ScheduledMessageCount
	destination.ScheduledMessageCount = genruntime.ClonePointerToInt(messageCountDetailsStatus.ScheduledMessageCount)

	// TransferDeadLetterMessageCount
	destination.TransferDeadLetterMessageCount = genruntime.ClonePointerToInt(messageCountDetailsStatus.TransferDeadLetterMessageCount)

	// TransferMessageCount
	destination.TransferMessageCount = genruntime.ClonePointerToInt(messageCountDetailsStatus.TransferMessageCount)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&NamespacesQueue{}, &NamespacesQueueList{})
}
