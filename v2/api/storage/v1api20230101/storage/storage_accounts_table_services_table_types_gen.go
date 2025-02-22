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

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccountstableservicestables,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccountstableservicestables/status,storageaccountstableservicestables/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230101.StorageAccountsTableServicesTable
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2023-01-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default/tables/{tableName}
type StorageAccountsTableServicesTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsTableServicesTable_Spec   `json:"spec,omitempty"`
	Status            StorageAccountsTableServicesTable_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsTableServicesTable{}

// GetConditions returns the conditions of the resource
func (table *StorageAccountsTableServicesTable) GetConditions() conditions.Conditions {
	return table.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (table *StorageAccountsTableServicesTable) SetConditions(conditions conditions.Conditions) {
	table.Status.Conditions = conditions
}

var _ configmaps.Exporter = &StorageAccountsTableServicesTable{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (table *StorageAccountsTableServicesTable) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if table.Spec.OperatorSpec == nil {
		return nil
	}
	return table.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &StorageAccountsTableServicesTable{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (table *StorageAccountsTableServicesTable) SecretDestinationExpressions() []*core.DestinationExpression {
	if table.Spec.OperatorSpec == nil {
		return nil
	}
	return table.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &StorageAccountsTableServicesTable{}

// AzureName returns the Azure name of the resource
func (table *StorageAccountsTableServicesTable) AzureName() string {
	return table.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-01"
func (table StorageAccountsTableServicesTable) GetAPIVersion() string {
	return "2023-01-01"
}

// GetResourceScope returns the scope of the resource
func (table *StorageAccountsTableServicesTable) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (table *StorageAccountsTableServicesTable) GetSpec() genruntime.ConvertibleSpec {
	return &table.Spec
}

// GetStatus returns the status of this resource
func (table *StorageAccountsTableServicesTable) GetStatus() genruntime.ConvertibleStatus {
	return &table.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (table *StorageAccountsTableServicesTable) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/tableServices/tables"
func (table *StorageAccountsTableServicesTable) GetType() string {
	return "Microsoft.Storage/storageAccounts/tableServices/tables"
}

// NewEmptyStatus returns a new empty (blank) status
func (table *StorageAccountsTableServicesTable) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccountsTableServicesTable_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (table *StorageAccountsTableServicesTable) Owner() *genruntime.ResourceReference {
	if table.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(table.Spec)
	return table.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (table *StorageAccountsTableServicesTable) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccountsTableServicesTable_STATUS); ok {
		table.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccountsTableServicesTable_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	table.Status = st
	return nil
}

// Hub marks that this StorageAccountsTableServicesTable is the hub type for conversion
func (table *StorageAccountsTableServicesTable) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (table *StorageAccountsTableServicesTable) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: table.Spec.OriginalVersion,
		Kind:    "StorageAccountsTableServicesTable",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230101.StorageAccountsTableServicesTable
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2023-01-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default/tables/{tableName}
type StorageAccountsTableServicesTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsTableServicesTable `json:"items"`
}

// Storage version of v1api20230101.StorageAccountsTableServicesTable_Spec
type StorageAccountsTableServicesTable_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                         `json:"azureName,omitempty"`
	OperatorSpec    *StorageAccountsTableServicesTableOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                         `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccountsTableService resource
	Owner             *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccountsTableService"`
	PropertyBag       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	SignedIdentifiers []TableSignedIdentifier            `json:"signedIdentifiers,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccountsTableServicesTable_Spec{}

// ConvertSpecFrom populates our StorageAccountsTableServicesTable_Spec from the provided source
func (table *StorageAccountsTableServicesTable_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == table {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(table)
}

// ConvertSpecTo populates the provided destination from our StorageAccountsTableServicesTable_Spec
func (table *StorageAccountsTableServicesTable_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == table {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(table)
}

// Storage version of v1api20230101.StorageAccountsTableServicesTable_STATUS
type StorageAccountsTableServicesTable_STATUS struct {
	Conditions        []conditions.Condition         `json:"conditions,omitempty"`
	Id                *string                        `json:"id,omitempty"`
	Name              *string                        `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	SignedIdentifiers []TableSignedIdentifier_STATUS `json:"signedIdentifiers,omitempty"`
	TableName         *string                        `json:"tableName,omitempty"`
	Type              *string                        `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccountsTableServicesTable_STATUS{}

// ConvertStatusFrom populates our StorageAccountsTableServicesTable_STATUS from the provided source
func (table *StorageAccountsTableServicesTable_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == table {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(table)
}

// ConvertStatusTo populates the provided destination from our StorageAccountsTableServicesTable_STATUS
func (table *StorageAccountsTableServicesTable_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == table {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(table)
}

// Storage version of v1api20230101.StorageAccountsTableServicesTableOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type StorageAccountsTableServicesTableOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20230101.TableSignedIdentifier
// Object to set Table Access Policy.
type TableSignedIdentifier struct {
	AccessPolicy *TableAccessPolicy     `json:"accessPolicy,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// Reference: unique-64-character-value of the stored access policy.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20230101.TableSignedIdentifier_STATUS
// Object to set Table Access Policy.
type TableSignedIdentifier_STATUS struct {
	AccessPolicy *TableAccessPolicy_STATUS `json:"accessPolicy,omitempty"`
	Id           *string                   `json:"id,omitempty"`
	PropertyBag  genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.TableAccessPolicy
// Table Access Policy Properties Object.
type TableAccessPolicy struct {
	ExpiryTime  *string                `json:"expiryTime,omitempty"`
	Permission  *string                `json:"permission,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartTime   *string                `json:"startTime,omitempty"`
}

// Storage version of v1api20230101.TableAccessPolicy_STATUS
// Table Access Policy Properties Object.
type TableAccessPolicy_STATUS struct {
	ExpiryTime  *string                `json:"expiryTime,omitempty"`
	Permission  *string                `json:"permission,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartTime   *string                `json:"startTime,omitempty"`
}

func init() {
	SchemeBuilder.Register(&StorageAccountsTableServicesTable{}, &StorageAccountsTableServicesTableList{})
}
