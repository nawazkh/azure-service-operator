// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type NamespacesEventhubsAuthorizationRule_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties supplied to create or update AuthorizationRule
	Properties *Namespaces_Eventhubs_AuthorizationRule_Properties_STATUS `json:"properties,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

type Namespaces_Eventhubs_AuthorizationRule_Properties_STATUS struct {
	// Rights: The rights associated with the rule.
	Rights []Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS `json:"rights,omitempty"`
}

type Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS string

const (
	Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Listen = Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS("Listen")
	Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Manage = Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS("Manage")
	Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Send   = Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS("Send")
)

// Mapping from string to Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS
var namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Values = map[string]Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS{
	"listen": Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Listen,
	"manage": Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Manage,
	"send":   Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Send,
}
