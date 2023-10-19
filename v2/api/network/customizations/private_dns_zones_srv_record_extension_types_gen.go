// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20200601 "github.com/Azure/azure-service-operator/v2/api/network/v1api20200601"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/network/v1api20200601/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type PrivateDnsZonesSRVRecordExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *PrivateDnsZonesSRVRecordExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20200601.PrivateDnsZonesSRVRecord{},
		&v20200601s.PrivateDnsZonesSRVRecord{}}
}
