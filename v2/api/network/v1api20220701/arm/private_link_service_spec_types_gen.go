// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type PrivateLinkService_Spec struct {
	// ExtendedLocation: The extended location of the load balancer.
	ExtendedLocation *ExtendedLocation `json:"extendedLocation,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the private link service.
	Properties *PrivateLinkServiceProperties `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &PrivateLinkService_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (service PrivateLinkService_Spec) GetAPIVersion() string {
	return "2022-07-01"
}

// GetName returns the Name of the resource
func (service *PrivateLinkService_Spec) GetName() string {
	return service.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateLinkServices"
func (service *PrivateLinkService_Spec) GetType() string {
	return "Microsoft.Network/privateLinkServices"
}

// Properties of the private link service.
type PrivateLinkServiceProperties struct {
	// AutoApproval: The auto-approval list of the private link service.
	AutoApproval *ResourceSet `json:"autoApproval,omitempty"`

	// EnableProxyProtocol: Whether the private link service is enabled for proxy protocol or not.
	EnableProxyProtocol *bool `json:"enableProxyProtocol,omitempty"`

	// Fqdns: The list of Fqdn.
	Fqdns []string `json:"fqdns,omitempty"`

	// IpConfigurations: An array of private link service IP configurations.
	IpConfigurations []PrivateLinkServiceIpConfiguration `json:"ipConfigurations,omitempty"`

	// LoadBalancerFrontendIpConfigurations: An array of references to the load balancer IP configurations.
	LoadBalancerFrontendIpConfigurations []FrontendIPConfiguration_PrivateLinkService_SubResourceEmbedded `json:"loadBalancerFrontendIpConfigurations,omitempty"`

	// Visibility: The visibility list of the private link service.
	Visibility *ResourceSet `json:"visibility,omitempty"`
}

// Frontend IP address of the load balancer.
type FrontendIPConfiguration_PrivateLinkService_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// The private link service ip configuration.
type PrivateLinkServiceIpConfiguration struct {
	// Name: The name of private link service ip configuration.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the private link service ip configuration.
	Properties *PrivateLinkServiceIpConfigurationProperties `json:"properties,omitempty"`
}

// The base resource set for visibility and auto-approval.
type ResourceSet struct {
	// Subscriptions: The list of subscriptions.
	Subscriptions []string `json:"subscriptions,omitempty"`
}

// Properties of private link service IP configuration.
type PrivateLinkServiceIpConfigurationProperties struct {
	// Primary: Whether the ip configuration is primary or not.
	Primary *bool `json:"primary,omitempty"`

	// PrivateIPAddress: The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	// PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4.
	PrivateIPAddressVersion *IPVersion `json:"privateIPAddressVersion,omitempty"`

	// PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`

	// Subnet: The reference to the subnet resource.
	Subnet *Subnet_PrivateLinkService_SubResourceEmbedded `json:"subnet,omitempty"`
}

// Subnet in a virtual network resource.
type Subnet_PrivateLinkService_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}
