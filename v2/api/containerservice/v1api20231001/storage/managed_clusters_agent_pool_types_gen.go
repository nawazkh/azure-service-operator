// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=containerservice.azure.com,resources=managedclustersagentpools,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=containerservice.azure.com,resources={managedclustersagentpools/status,managedclustersagentpools/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20231001.ManagedClustersAgentPool
// Generator information:
// - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-10-01/managedClusters.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}
type ManagedClustersAgentPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedClusters_AgentPool_Spec   `json:"spec,omitempty"`
	Status            ManagedClusters_AgentPool_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ManagedClustersAgentPool{}

// GetConditions returns the conditions of the resource
func (pool *ManagedClustersAgentPool) GetConditions() conditions.Conditions {
	return pool.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (pool *ManagedClustersAgentPool) SetConditions(conditions conditions.Conditions) {
	pool.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &ManagedClustersAgentPool{}

// AzureName returns the Azure name of the resource
func (pool *ManagedClustersAgentPool) AzureName() string {
	return pool.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-10-01"
func (pool ManagedClustersAgentPool) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (pool *ManagedClustersAgentPool) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (pool *ManagedClustersAgentPool) GetSpec() genruntime.ConvertibleSpec {
	return &pool.Spec
}

// GetStatus returns the status of this resource
func (pool *ManagedClustersAgentPool) GetStatus() genruntime.ConvertibleStatus {
	return &pool.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (pool *ManagedClustersAgentPool) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerService/managedClusters/agentPools"
func (pool *ManagedClustersAgentPool) GetType() string {
	return "Microsoft.ContainerService/managedClusters/agentPools"
}

// NewEmptyStatus returns a new empty (blank) status
func (pool *ManagedClustersAgentPool) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ManagedClusters_AgentPool_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (pool *ManagedClustersAgentPool) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(pool.Spec)
	return pool.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (pool *ManagedClustersAgentPool) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ManagedClusters_AgentPool_STATUS); ok {
		pool.Status = *st
		return nil
	}

	// Convert status to required version
	var st ManagedClusters_AgentPool_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	pool.Status = st
	return nil
}

// Hub marks that this ManagedClustersAgentPool is the hub type for conversion
func (pool *ManagedClustersAgentPool) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (pool *ManagedClustersAgentPool) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: pool.Spec.OriginalVersion,
		Kind:    "ManagedClustersAgentPool",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20231001.ManagedClustersAgentPool
// Generator information:
// - Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-10-01/managedClusters.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}
type ManagedClustersAgentPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedClustersAgentPool `json:"items"`
}

// Storage version of v1api20231001.ManagedClusters_AgentPool_Spec
type ManagedClusters_AgentPool_Spec struct {
	AvailabilityZones []string `json:"availabilityZones,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// CapacityReservationGroupReference: AKS will associate the specified agent pool with the Capacity Reservation Group.
	CapacityReservationGroupReference *genruntime.ResourceReference `armReference:"CapacityReservationGroupID" json:"capacityReservationGroupReference,omitempty"`
	Count                             *int                          `json:"count,omitempty"`
	CreationData                      *CreationData                 `json:"creationData,omitempty"`
	EnableAutoScaling                 *bool                         `json:"enableAutoScaling,omitempty"`
	EnableEncryptionAtHost            *bool                         `json:"enableEncryptionAtHost,omitempty"`
	EnableFIPS                        *bool                         `json:"enableFIPS,omitempty"`
	EnableNodePublicIP                *bool                         `json:"enableNodePublicIP,omitempty"`
	EnableUltraSSD                    *bool                         `json:"enableUltraSSD,omitempty"`
	GpuInstanceProfile                *string                       `json:"gpuInstanceProfile,omitempty"`

	// HostGroupReference: This is of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}.
	// For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).
	HostGroupReference *genruntime.ResourceReference `armReference:"HostGroupID" json:"hostGroupReference,omitempty"`
	KubeletConfig      *KubeletConfig                `json:"kubeletConfig,omitempty"`
	KubeletDiskType    *string                       `json:"kubeletDiskType,omitempty"`
	LinuxOSConfig      *LinuxOSConfig                `json:"linuxOSConfig,omitempty"`
	MaxCount           *int                          `json:"maxCount,omitempty"`
	MaxPods            *int                          `json:"maxPods,omitempty"`
	MinCount           *int                          `json:"minCount,omitempty"`
	Mode               *string                       `json:"mode,omitempty"`
	NetworkProfile     *AgentPoolNetworkProfile      `json:"networkProfile,omitempty"`
	NodeLabels         map[string]string             `json:"nodeLabels,omitempty" serializationType:"explicitEmptyCollection"`

	// NodePublicIPPrefixReference: This is of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}
	NodePublicIPPrefixReference *genruntime.ResourceReference `armReference:"NodePublicIPPrefixID" json:"nodePublicIPPrefixReference,omitempty"`
	NodeTaints                  []string                      `json:"nodeTaints,omitempty" serializationType:"explicitEmptyCollection"`
	OrchestratorVersion         *string                       `json:"orchestratorVersion,omitempty"`
	OriginalVersion             string                        `json:"originalVersion,omitempty"`
	OsDiskSizeGB                *int                          `json:"osDiskSizeGB,omitempty"`
	OsDiskType                  *string                       `json:"osDiskType,omitempty"`
	OsSKU                       *string                       `json:"osSKU,omitempty"`
	OsType                      *string                       `json:"osType,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a containerservice.azure.com/ManagedCluster resource
	Owner *genruntime.KnownResourceReference `group:"containerservice.azure.com" json:"owner,omitempty" kind:"ManagedCluster"`

	// PodSubnetReference: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details).
	// This is of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	PodSubnetReference *genruntime.ResourceReference `armReference:"PodSubnetID" json:"podSubnetReference,omitempty"`
	PowerState         *PowerState                   `json:"powerState,omitempty"`
	PropertyBag        genruntime.PropertyBag        `json:"$propertyBag,omitempty"`

	// ProximityPlacementGroupReference: The ID for Proximity Placement Group.
	ProximityPlacementGroupReference *genruntime.ResourceReference `armReference:"ProximityPlacementGroupID" json:"proximityPlacementGroupReference,omitempty"`
	ScaleDownMode                    *string                       `json:"scaleDownMode,omitempty"`
	ScaleSetEvictionPolicy           *string                       `json:"scaleSetEvictionPolicy,omitempty"`
	ScaleSetPriority                 *string                       `json:"scaleSetPriority,omitempty"`
	SpotMaxPrice                     *float64                      `json:"spotMaxPrice,omitempty"`
	Tags                             map[string]string             `json:"tags,omitempty" serializationType:"explicitEmptyCollection"`
	Type                             *string                       `json:"type,omitempty"`
	UpgradeSettings                  *AgentPoolUpgradeSettings     `json:"upgradeSettings,omitempty"`
	VmSize                           *string                       `json:"vmSize,omitempty"`

	// VnetSubnetReference: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is
	// specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	VnetSubnetReference *genruntime.ResourceReference `armReference:"VnetSubnetID" json:"vnetSubnetReference,omitempty"`
	WorkloadRuntime     *string                       `json:"workloadRuntime,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ManagedClusters_AgentPool_Spec{}

// ConvertSpecFrom populates our ManagedClusters_AgentPool_Spec from the provided source
func (pool *ManagedClusters_AgentPool_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == pool {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(pool)
}

// ConvertSpecTo populates the provided destination from our ManagedClusters_AgentPool_Spec
func (pool *ManagedClusters_AgentPool_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == pool {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(pool)
}

// Storage version of v1api20231001.ManagedClusters_AgentPool_STATUS
type ManagedClusters_AgentPool_STATUS struct {
	AvailabilityZones          []string                         `json:"availabilityZones,omitempty"`
	CapacityReservationGroupID *string                          `json:"capacityReservationGroupID,omitempty"`
	Conditions                 []conditions.Condition           `json:"conditions,omitempty"`
	Count                      *int                             `json:"count,omitempty"`
	CreationData               *CreationData_STATUS             `json:"creationData,omitempty"`
	CurrentOrchestratorVersion *string                          `json:"currentOrchestratorVersion,omitempty"`
	EnableAutoScaling          *bool                            `json:"enableAutoScaling,omitempty"`
	EnableEncryptionAtHost     *bool                            `json:"enableEncryptionAtHost,omitempty"`
	EnableFIPS                 *bool                            `json:"enableFIPS,omitempty"`
	EnableNodePublicIP         *bool                            `json:"enableNodePublicIP,omitempty"`
	EnableUltraSSD             *bool                            `json:"enableUltraSSD,omitempty"`
	GpuInstanceProfile         *string                          `json:"gpuInstanceProfile,omitempty"`
	HostGroupID                *string                          `json:"hostGroupID,omitempty"`
	Id                         *string                          `json:"id,omitempty"`
	KubeletConfig              *KubeletConfig_STATUS            `json:"kubeletConfig,omitempty"`
	KubeletDiskType            *string                          `json:"kubeletDiskType,omitempty"`
	LinuxOSConfig              *LinuxOSConfig_STATUS            `json:"linuxOSConfig,omitempty"`
	MaxCount                   *int                             `json:"maxCount,omitempty"`
	MaxPods                    *int                             `json:"maxPods,omitempty"`
	MinCount                   *int                             `json:"minCount,omitempty"`
	Mode                       *string                          `json:"mode,omitempty"`
	Name                       *string                          `json:"name,omitempty"`
	NetworkProfile             *AgentPoolNetworkProfile_STATUS  `json:"networkProfile,omitempty"`
	NodeImageVersion           *string                          `json:"nodeImageVersion,omitempty"`
	NodeLabels                 map[string]string                `json:"nodeLabels,omitempty"`
	NodePublicIPPrefixID       *string                          `json:"nodePublicIPPrefixID,omitempty"`
	NodeTaints                 []string                         `json:"nodeTaints,omitempty"`
	OrchestratorVersion        *string                          `json:"orchestratorVersion,omitempty"`
	OsDiskSizeGB               *int                             `json:"osDiskSizeGB,omitempty"`
	OsDiskType                 *string                          `json:"osDiskType,omitempty"`
	OsSKU                      *string                          `json:"osSKU,omitempty"`
	OsType                     *string                          `json:"osType,omitempty"`
	PodSubnetID                *string                          `json:"podSubnetID,omitempty"`
	PowerState                 *PowerState_STATUS               `json:"powerState,omitempty"`
	PropertiesType             *string                          `json:"properties_type,omitempty"`
	PropertyBag                genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                          `json:"provisioningState,omitempty"`
	ProximityPlacementGroupID  *string                          `json:"proximityPlacementGroupID,omitempty"`
	ScaleDownMode              *string                          `json:"scaleDownMode,omitempty"`
	ScaleSetEvictionPolicy     *string                          `json:"scaleSetEvictionPolicy,omitempty"`
	ScaleSetPriority           *string                          `json:"scaleSetPriority,omitempty"`
	SpotMaxPrice               *float64                         `json:"spotMaxPrice,omitempty"`
	Tags                       map[string]string                `json:"tags,omitempty"`
	Type                       *string                          `json:"type,omitempty"`
	UpgradeSettings            *AgentPoolUpgradeSettings_STATUS `json:"upgradeSettings,omitempty"`
	VmSize                     *string                          `json:"vmSize,omitempty"`
	VnetSubnetID               *string                          `json:"vnetSubnetID,omitempty"`
	WorkloadRuntime            *string                          `json:"workloadRuntime,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ManagedClusters_AgentPool_STATUS{}

// ConvertStatusFrom populates our ManagedClusters_AgentPool_STATUS from the provided source
func (pool *ManagedClusters_AgentPool_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == pool {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(pool)
}

// ConvertStatusTo populates the provided destination from our ManagedClusters_AgentPool_STATUS
func (pool *ManagedClusters_AgentPool_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == pool {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(pool)
}

// Storage version of v1api20231001.AgentPoolNetworkProfile
// Network settings of an agent pool.
type AgentPoolNetworkProfile struct {
	AllowedHostPorts                    []PortRange                    `json:"allowedHostPorts,omitempty"`
	ApplicationSecurityGroupsReferences []genruntime.ResourceReference `armReference:"ApplicationSecurityGroups" json:"applicationSecurityGroupsReferences,omitempty"`
	NodePublicIPTags                    []IPTag                        `json:"nodePublicIPTags,omitempty"`
	PropertyBag                         genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231001.AgentPoolNetworkProfile_STATUS
// Network settings of an agent pool.
type AgentPoolNetworkProfile_STATUS struct {
	AllowedHostPorts          []PortRange_STATUS     `json:"allowedHostPorts,omitempty"`
	ApplicationSecurityGroups []string               `json:"applicationSecurityGroups,omitempty"`
	NodePublicIPTags          []IPTag_STATUS         `json:"nodePublicIPTags,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231001.AgentPoolUpgradeSettings
// Settings for upgrading an agentpool
type AgentPoolUpgradeSettings struct {
	DrainTimeoutInMinutes *int                   `json:"drainTimeoutInMinutes,omitempty"`
	MaxSurge              *string                `json:"maxSurge,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231001.AgentPoolUpgradeSettings_STATUS
// Settings for upgrading an agentpool
type AgentPoolUpgradeSettings_STATUS struct {
	DrainTimeoutInMinutes *int                   `json:"drainTimeoutInMinutes,omitempty"`
	MaxSurge              *string                `json:"maxSurge,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231001.CreationData
// Data used when creating a target resource from a source resource.
type CreationData struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// SourceResourceReference: This is the ARM ID of the source object to be used to create the target object.
	SourceResourceReference *genruntime.ResourceReference `armReference:"SourceResourceId" json:"sourceResourceReference,omitempty"`
}

// Storage version of v1api20231001.CreationData_STATUS
// Data used when creating a target resource from a source resource.
type CreationData_STATUS struct {
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceResourceId *string                `json:"sourceResourceId,omitempty"`
}

// Storage version of v1api20231001.KubeletConfig
// See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.
type KubeletConfig struct {
	AllowedUnsafeSysctls  []string               `json:"allowedUnsafeSysctls,omitempty"`
	ContainerLogMaxFiles  *int                   `json:"containerLogMaxFiles,omitempty"`
	ContainerLogMaxSizeMB *int                   `json:"containerLogMaxSizeMB,omitempty"`
	CpuCfsQuota           *bool                  `json:"cpuCfsQuota,omitempty"`
	CpuCfsQuotaPeriod     *string                `json:"cpuCfsQuotaPeriod,omitempty"`
	CpuManagerPolicy      *string                `json:"cpuManagerPolicy,omitempty"`
	FailSwapOn            *bool                  `json:"failSwapOn,omitempty"`
	ImageGcHighThreshold  *int                   `json:"imageGcHighThreshold,omitempty"`
	ImageGcLowThreshold   *int                   `json:"imageGcLowThreshold,omitempty"`
	PodMaxPids            *int                   `json:"podMaxPids,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TopologyManagerPolicy *string                `json:"topologyManagerPolicy,omitempty"`
}

// Storage version of v1api20231001.KubeletConfig_STATUS
// See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.
type KubeletConfig_STATUS struct {
	AllowedUnsafeSysctls  []string               `json:"allowedUnsafeSysctls,omitempty"`
	ContainerLogMaxFiles  *int                   `json:"containerLogMaxFiles,omitempty"`
	ContainerLogMaxSizeMB *int                   `json:"containerLogMaxSizeMB,omitempty"`
	CpuCfsQuota           *bool                  `json:"cpuCfsQuota,omitempty"`
	CpuCfsQuotaPeriod     *string                `json:"cpuCfsQuotaPeriod,omitempty"`
	CpuManagerPolicy      *string                `json:"cpuManagerPolicy,omitempty"`
	FailSwapOn            *bool                  `json:"failSwapOn,omitempty"`
	ImageGcHighThreshold  *int                   `json:"imageGcHighThreshold,omitempty"`
	ImageGcLowThreshold   *int                   `json:"imageGcLowThreshold,omitempty"`
	PodMaxPids            *int                   `json:"podMaxPids,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TopologyManagerPolicy *string                `json:"topologyManagerPolicy,omitempty"`
}

// Storage version of v1api20231001.LinuxOSConfig
// See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.
type LinuxOSConfig struct {
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SwapFileSizeMB             *int                   `json:"swapFileSizeMB,omitempty"`
	Sysctls                    *SysctlConfig          `json:"sysctls,omitempty"`
	TransparentHugePageDefrag  *string                `json:"transparentHugePageDefrag,omitempty"`
	TransparentHugePageEnabled *string                `json:"transparentHugePageEnabled,omitempty"`
}

// Storage version of v1api20231001.LinuxOSConfig_STATUS
// See [AKS custom node configuration](https://docs.microsoft.com/azure/aks/custom-node-configuration) for more details.
type LinuxOSConfig_STATUS struct {
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SwapFileSizeMB             *int                   `json:"swapFileSizeMB,omitempty"`
	Sysctls                    *SysctlConfig_STATUS   `json:"sysctls,omitempty"`
	TransparentHugePageDefrag  *string                `json:"transparentHugePageDefrag,omitempty"`
	TransparentHugePageEnabled *string                `json:"transparentHugePageEnabled,omitempty"`
}

// Storage version of v1api20231001.PowerState
// Describes the Power State of the cluster
type PowerState struct {
	Code        *string                `json:"code,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20231001.IPTag
// Contains the IPTag associated with the object.
type IPTag struct {
	IpTagType   *string                `json:"ipTagType,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tag         *string                `json:"tag,omitempty"`
}

// Storage version of v1api20231001.IPTag_STATUS
// Contains the IPTag associated with the object.
type IPTag_STATUS struct {
	IpTagType   *string                `json:"ipTagType,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tag         *string                `json:"tag,omitempty"`
}

// Storage version of v1api20231001.PortRange
// The port range.
type PortRange struct {
	PortEnd     *int                   `json:"portEnd,omitempty"`
	PortStart   *int                   `json:"portStart,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol    *string                `json:"protocol,omitempty"`
}

// Storage version of v1api20231001.PortRange_STATUS
// The port range.
type PortRange_STATUS struct {
	PortEnd     *int                   `json:"portEnd,omitempty"`
	PortStart   *int                   `json:"portStart,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol    *string                `json:"protocol,omitempty"`
}

// Storage version of v1api20231001.SysctlConfig
// Sysctl settings for Linux agent nodes.
type SysctlConfig struct {
	FsAioMaxNr                     *int                   `json:"fsAioMaxNr,omitempty"`
	FsFileMax                      *int                   `json:"fsFileMax,omitempty"`
	FsInotifyMaxUserWatches        *int                   `json:"fsInotifyMaxUserWatches,omitempty"`
	FsNrOpen                       *int                   `json:"fsNrOpen,omitempty"`
	KernelThreadsMax               *int                   `json:"kernelThreadsMax,omitempty"`
	NetCoreNetdevMaxBacklog        *int                   `json:"netCoreNetdevMaxBacklog,omitempty"`
	NetCoreOptmemMax               *int                   `json:"netCoreOptmemMax,omitempty"`
	NetCoreRmemDefault             *int                   `json:"netCoreRmemDefault,omitempty"`
	NetCoreRmemMax                 *int                   `json:"netCoreRmemMax,omitempty"`
	NetCoreSomaxconn               *int                   `json:"netCoreSomaxconn,omitempty"`
	NetCoreWmemDefault             *int                   `json:"netCoreWmemDefault,omitempty"`
	NetCoreWmemMax                 *int                   `json:"netCoreWmemMax,omitempty"`
	NetIpv4IpLocalPortRange        *string                `json:"netIpv4IpLocalPortRange,omitempty"`
	NetIpv4NeighDefaultGcThresh1   *int                   `json:"netIpv4NeighDefaultGcThresh1,omitempty"`
	NetIpv4NeighDefaultGcThresh2   *int                   `json:"netIpv4NeighDefaultGcThresh2,omitempty"`
	NetIpv4NeighDefaultGcThresh3   *int                   `json:"netIpv4NeighDefaultGcThresh3,omitempty"`
	NetIpv4TcpFinTimeout           *int                   `json:"netIpv4TcpFinTimeout,omitempty"`
	NetIpv4TcpKeepaliveProbes      *int                   `json:"netIpv4TcpKeepaliveProbes,omitempty"`
	NetIpv4TcpKeepaliveTime        *int                   `json:"netIpv4TcpKeepaliveTime,omitempty"`
	NetIpv4TcpMaxSynBacklog        *int                   `json:"netIpv4TcpMaxSynBacklog,omitempty"`
	NetIpv4TcpMaxTwBuckets         *int                   `json:"netIpv4TcpMaxTwBuckets,omitempty"`
	NetIpv4TcpTwReuse              *bool                  `json:"netIpv4TcpTwReuse,omitempty"`
	NetIpv4TcpkeepaliveIntvl       *int                   `json:"netIpv4TcpkeepaliveIntvl,omitempty"`
	NetNetfilterNfConntrackBuckets *int                   `json:"netNetfilterNfConntrackBuckets,omitempty"`
	NetNetfilterNfConntrackMax     *int                   `json:"netNetfilterNfConntrackMax,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VmMaxMapCount                  *int                   `json:"vmMaxMapCount,omitempty"`
	VmSwappiness                   *int                   `json:"vmSwappiness,omitempty"`
	VmVfsCachePressure             *int                   `json:"vmVfsCachePressure,omitempty"`
}

// Storage version of v1api20231001.SysctlConfig_STATUS
// Sysctl settings for Linux agent nodes.
type SysctlConfig_STATUS struct {
	FsAioMaxNr                     *int                   `json:"fsAioMaxNr,omitempty"`
	FsFileMax                      *int                   `json:"fsFileMax,omitempty"`
	FsInotifyMaxUserWatches        *int                   `json:"fsInotifyMaxUserWatches,omitempty"`
	FsNrOpen                       *int                   `json:"fsNrOpen,omitempty"`
	KernelThreadsMax               *int                   `json:"kernelThreadsMax,omitempty"`
	NetCoreNetdevMaxBacklog        *int                   `json:"netCoreNetdevMaxBacklog,omitempty"`
	NetCoreOptmemMax               *int                   `json:"netCoreOptmemMax,omitempty"`
	NetCoreRmemDefault             *int                   `json:"netCoreRmemDefault,omitempty"`
	NetCoreRmemMax                 *int                   `json:"netCoreRmemMax,omitempty"`
	NetCoreSomaxconn               *int                   `json:"netCoreSomaxconn,omitempty"`
	NetCoreWmemDefault             *int                   `json:"netCoreWmemDefault,omitempty"`
	NetCoreWmemMax                 *int                   `json:"netCoreWmemMax,omitempty"`
	NetIpv4IpLocalPortRange        *string                `json:"netIpv4IpLocalPortRange,omitempty"`
	NetIpv4NeighDefaultGcThresh1   *int                   `json:"netIpv4NeighDefaultGcThresh1,omitempty"`
	NetIpv4NeighDefaultGcThresh2   *int                   `json:"netIpv4NeighDefaultGcThresh2,omitempty"`
	NetIpv4NeighDefaultGcThresh3   *int                   `json:"netIpv4NeighDefaultGcThresh3,omitempty"`
	NetIpv4TcpFinTimeout           *int                   `json:"netIpv4TcpFinTimeout,omitempty"`
	NetIpv4TcpKeepaliveProbes      *int                   `json:"netIpv4TcpKeepaliveProbes,omitempty"`
	NetIpv4TcpKeepaliveTime        *int                   `json:"netIpv4TcpKeepaliveTime,omitempty"`
	NetIpv4TcpMaxSynBacklog        *int                   `json:"netIpv4TcpMaxSynBacklog,omitempty"`
	NetIpv4TcpMaxTwBuckets         *int                   `json:"netIpv4TcpMaxTwBuckets,omitempty"`
	NetIpv4TcpTwReuse              *bool                  `json:"netIpv4TcpTwReuse,omitempty"`
	NetIpv4TcpkeepaliveIntvl       *int                   `json:"netIpv4TcpkeepaliveIntvl,omitempty"`
	NetNetfilterNfConntrackBuckets *int                   `json:"netNetfilterNfConntrackBuckets,omitempty"`
	NetNetfilterNfConntrackMax     *int                   `json:"netNetfilterNfConntrackMax,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	VmMaxMapCount                  *int                   `json:"vmMaxMapCount,omitempty"`
	VmSwappiness                   *int                   `json:"vmSwappiness,omitempty"`
	VmVfsCachePressure             *int                   `json:"vmVfsCachePressure,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ManagedClustersAgentPool{}, &ManagedClustersAgentPoolList{})
}
