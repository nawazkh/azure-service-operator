// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20211101 "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101/storage"
	v20240101 "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20240101"
	v20240101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20240101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type NamespacesAuthorizationRuleExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *NamespacesAuthorizationRuleExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20211101.NamespacesAuthorizationRule{},
		&v20211101s.NamespacesAuthorizationRule{},
		&v20240101.NamespacesAuthorizationRule{},
		&v20240101s.NamespacesAuthorizationRule{}}
}
