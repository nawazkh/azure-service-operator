// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=servicebus.azure.com,resources=namespacestopicssubscriptions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=servicebus.azure.com,resources={namespacestopicssubscriptions/status,namespacestopicssubscriptions/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240101.NamespacesTopicsSubscription
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/stable/2024-01-01/subscriptions.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}
type NamespacesTopicsSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesTopicsSubscription_Spec   `json:"spec,omitempty"`
	Status            NamespacesTopicsSubscription_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesTopicsSubscription{}

// GetConditions returns the conditions of the resource
func (subscription *NamespacesTopicsSubscription) GetConditions() conditions.Conditions {
	return subscription.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (subscription *NamespacesTopicsSubscription) SetConditions(conditions conditions.Conditions) {
	subscription.Status.Conditions = conditions
}

var _ configmaps.Exporter = &NamespacesTopicsSubscription{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (subscription *NamespacesTopicsSubscription) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if subscription.Spec.OperatorSpec == nil {
		return nil
	}
	return subscription.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &NamespacesTopicsSubscription{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (subscription *NamespacesTopicsSubscription) SecretDestinationExpressions() []*core.DestinationExpression {
	if subscription.Spec.OperatorSpec == nil {
		return nil
	}
	return subscription.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &NamespacesTopicsSubscription{}

// AzureName returns the Azure name of the resource
func (subscription *NamespacesTopicsSubscription) AzureName() string {
	return subscription.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-01-01"
func (subscription NamespacesTopicsSubscription) GetAPIVersion() string {
	return "2024-01-01"
}

// GetResourceScope returns the scope of the resource
func (subscription *NamespacesTopicsSubscription) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (subscription *NamespacesTopicsSubscription) GetSpec() genruntime.ConvertibleSpec {
	return &subscription.Spec
}

// GetStatus returns the status of this resource
func (subscription *NamespacesTopicsSubscription) GetStatus() genruntime.ConvertibleStatus {
	return &subscription.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (subscription *NamespacesTopicsSubscription) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics/subscriptions"
func (subscription *NamespacesTopicsSubscription) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics/subscriptions"
}

// NewEmptyStatus returns a new empty (blank) status
func (subscription *NamespacesTopicsSubscription) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NamespacesTopicsSubscription_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (subscription *NamespacesTopicsSubscription) Owner() *genruntime.ResourceReference {
	if subscription.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(subscription.Spec)
	return subscription.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (subscription *NamespacesTopicsSubscription) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NamespacesTopicsSubscription_STATUS); ok {
		subscription.Status = *st
		return nil
	}

	// Convert status to required version
	var st NamespacesTopicsSubscription_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	subscription.Status = st
	return nil
}

// Hub marks that this NamespacesTopicsSubscription is the hub type for conversion
func (subscription *NamespacesTopicsSubscription) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (subscription *NamespacesTopicsSubscription) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: subscription.Spec.OriginalVersion,
		Kind:    "NamespacesTopicsSubscription",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240101.NamespacesTopicsSubscription
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/stable/2024-01-01/subscriptions.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}
type NamespacesTopicsSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesTopicsSubscription `json:"items"`
}

// Storage version of v1api20240101.NamespacesTopicsSubscription_Spec
type NamespacesTopicsSubscription_Spec struct {
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                                 string                                    `json:"azureName,omitempty"`
	ClientAffineProperties                    *SBClientAffineProperties                 `json:"clientAffineProperties,omitempty"`
	DeadLetteringOnFilterEvaluationExceptions *bool                                     `json:"deadLetteringOnFilterEvaluationExceptions,omitempty"`
	DeadLetteringOnMessageExpiration          *bool                                     `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive                  *string                                   `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow       *string                                   `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations                   *bool                                     `json:"enableBatchedOperations,omitempty"`
	ForwardDeadLetteredMessagesTo             *string                                   `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                                 *string                                   `json:"forwardTo,omitempty"`
	IsClientAffine                            *bool                                     `json:"isClientAffine,omitempty"`
	LockDuration                              *string                                   `json:"lockDuration,omitempty"`
	MaxDeliveryCount                          *int                                      `json:"maxDeliveryCount,omitempty"`
	OperatorSpec                              *NamespacesTopicsSubscriptionOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion                           string                                    `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a servicebus.azure.com/NamespacesTopic resource
	Owner           *genruntime.KnownResourceReference `group:"servicebus.azure.com" json:"owner,omitempty" kind:"NamespacesTopic"`
	PropertyBag     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RequiresSession *bool                              `json:"requiresSession,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesTopicsSubscription_Spec{}

// ConvertSpecFrom populates our NamespacesTopicsSubscription_Spec from the provided source
func (subscription *NamespacesTopicsSubscription_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == subscription {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(subscription)
}

// ConvertSpecTo populates the provided destination from our NamespacesTopicsSubscription_Spec
func (subscription *NamespacesTopicsSubscription_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == subscription {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(subscription)
}

// Storage version of v1api20240101.NamespacesTopicsSubscription_STATUS
type NamespacesTopicsSubscription_STATUS struct {
	AccessedAt                                *string                          `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                          *string                          `json:"autoDeleteOnIdle,omitempty"`
	ClientAffineProperties                    *SBClientAffineProperties_STATUS `json:"clientAffineProperties,omitempty"`
	Conditions                                []conditions.Condition           `json:"conditions,omitempty"`
	CountDetails                              *MessageCountDetails_STATUS      `json:"countDetails,omitempty"`
	CreatedAt                                 *string                          `json:"createdAt,omitempty"`
	DeadLetteringOnFilterEvaluationExceptions *bool                            `json:"deadLetteringOnFilterEvaluationExceptions,omitempty"`
	DeadLetteringOnMessageExpiration          *bool                            `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive                  *string                          `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow       *string                          `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations                   *bool                            `json:"enableBatchedOperations,omitempty"`
	ForwardDeadLetteredMessagesTo             *string                          `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                                 *string                          `json:"forwardTo,omitempty"`
	Id                                        *string                          `json:"id,omitempty"`
	IsClientAffine                            *bool                            `json:"isClientAffine,omitempty"`
	Location                                  *string                          `json:"location,omitempty"`
	LockDuration                              *string                          `json:"lockDuration,omitempty"`
	MaxDeliveryCount                          *int                             `json:"maxDeliveryCount,omitempty"`
	MessageCount                              *int                             `json:"messageCount,omitempty"`
	Name                                      *string                          `json:"name,omitempty"`
	PropertyBag                               genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	RequiresSession                           *bool                            `json:"requiresSession,omitempty"`
	Status                                    *string                          `json:"status,omitempty"`
	SystemData                                *SystemData_STATUS               `json:"systemData,omitempty"`
	Type                                      *string                          `json:"type,omitempty"`
	UpdatedAt                                 *string                          `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NamespacesTopicsSubscription_STATUS{}

// ConvertStatusFrom populates our NamespacesTopicsSubscription_STATUS from the provided source
func (subscription *NamespacesTopicsSubscription_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == subscription {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(subscription)
}

// ConvertStatusTo populates the provided destination from our NamespacesTopicsSubscription_STATUS
func (subscription *NamespacesTopicsSubscription_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == subscription {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(subscription)
}

// Storage version of v1api20240101.NamespacesTopicsSubscriptionOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type NamespacesTopicsSubscriptionOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20240101.SBClientAffineProperties
// Properties specific to client affine subscriptions.
type SBClientAffineProperties struct {
	ClientId    *string                `json:"clientId,omitempty"`
	IsDurable   *bool                  `json:"isDurable,omitempty"`
	IsShared    *bool                  `json:"isShared,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240101.SBClientAffineProperties_STATUS
// Properties specific to client affine subscriptions.
type SBClientAffineProperties_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	IsDurable   *bool                  `json:"isDurable,omitempty"`
	IsShared    *bool                  `json:"isShared,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NamespacesTopicsSubscription{}, &NamespacesTopicsSubscriptionList{})
}
