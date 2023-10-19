// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101

import (
	"encoding/json"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_VirtualNetwork_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetwork to hub returns original",
		prop.ForAll(RunResourceConversionTestForVirtualNetwork, VirtualNetworkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForVirtualNetwork tests if a specific instance of VirtualNetwork round trips to the hub storage version and back losslessly
func RunResourceConversionTestForVirtualNetwork(subject VirtualNetwork) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20201101s.VirtualNetwork
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual VirtualNetwork
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetwork_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetwork to VirtualNetwork via AssignProperties_To_VirtualNetwork & AssignProperties_From_VirtualNetwork returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetwork, VirtualNetworkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetwork tests if a specific instance of VirtualNetwork can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetwork(subject VirtualNetwork) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.VirtualNetwork
	err := copied.AssignProperties_To_VirtualNetwork(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetwork
	err = actual.AssignProperties_From_VirtualNetwork(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetwork_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork, VirtualNetworkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork runs a test to see if a specific instance of VirtualNetwork round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork(subject VirtualNetwork) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetwork instances for property testing - lazily instantiated by VirtualNetworkGenerator()
var virtualNetworkGenerator gopter.Gen

// VirtualNetworkGenerator returns a generator of VirtualNetwork instances for property testing.
func VirtualNetworkGenerator() gopter.Gen {
	if virtualNetworkGenerator != nil {
		return virtualNetworkGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForVirtualNetwork(generators)
	virtualNetworkGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork{}), generators)

	return virtualNetworkGenerator
}

// AddRelatedPropertyGeneratorsForVirtualNetwork is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork(gens map[string]gopter.Gen) {
	gens["Spec"] = VirtualNetwork_SpecGenerator()
	gens["Status"] = VirtualNetwork_STATUSGenerator()
}

func Test_VirtualNetwork_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetwork_Spec to VirtualNetwork_Spec via AssignProperties_To_VirtualNetwork_Spec & AssignProperties_From_VirtualNetwork_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetwork_Spec, VirtualNetwork_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetwork_Spec tests if a specific instance of VirtualNetwork_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetwork_Spec(subject VirtualNetwork_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.VirtualNetwork_Spec
	err := copied.AssignProperties_To_VirtualNetwork_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetwork_Spec
	err = actual.AssignProperties_From_VirtualNetwork_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetwork_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork_Spec, VirtualNetwork_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork_Spec runs a test to see if a specific instance of VirtualNetwork_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork_Spec(subject VirtualNetwork_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetwork_Spec instances for property testing - lazily instantiated by
// VirtualNetwork_SpecGenerator()
var virtualNetwork_SpecGenerator gopter.Gen

// VirtualNetwork_SpecGenerator returns a generator of VirtualNetwork_Spec instances for property testing.
// We first initialize virtualNetwork_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetwork_SpecGenerator() gopter.Gen {
	if virtualNetwork_SpecGenerator != nil {
		return virtualNetwork_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec(generators)
	virtualNetwork_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_Spec(generators)
	AddRelatedPropertyGeneratorsForVirtualNetwork_Spec(generators)
	virtualNetwork_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Spec{}), generators)

	return virtualNetwork_SpecGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetwork_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetwork_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EnableDdosProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableVmProtection"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetwork_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork_Spec(gens map[string]gopter.Gen) {
	gens["AddressSpace"] = gen.PtrOf(AddressSpaceGenerator())
	gens["BgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesGenerator())
	gens["DdosProtectionPlan"] = gen.PtrOf(SubResourceGenerator())
	gens["DhcpOptions"] = gen.PtrOf(DhcpOptionsGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResourceGenerator())
}

func Test_VirtualNetwork_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetwork_STATUS to VirtualNetwork_STATUS via AssignProperties_To_VirtualNetwork_STATUS & AssignProperties_From_VirtualNetwork_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetwork_STATUS, VirtualNetwork_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetwork_STATUS tests if a specific instance of VirtualNetwork_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetwork_STATUS(subject VirtualNetwork_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.VirtualNetwork_STATUS
	err := copied.AssignProperties_To_VirtualNetwork_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetwork_STATUS
	err = actual.AssignProperties_From_VirtualNetwork_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetwork_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork_STATUS, VirtualNetwork_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork_STATUS runs a test to see if a specific instance of VirtualNetwork_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork_STATUS(subject VirtualNetwork_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetwork_STATUS instances for property testing - lazily instantiated by
// VirtualNetwork_STATUSGenerator()
var virtualNetwork_STATUSGenerator gopter.Gen

// VirtualNetwork_STATUSGenerator returns a generator of VirtualNetwork_STATUS instances for property testing.
// We first initialize virtualNetwork_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetwork_STATUSGenerator() gopter.Gen {
	if virtualNetwork_STATUSGenerator != nil {
		return virtualNetwork_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_STATUS(generators)
	virtualNetwork_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualNetwork_STATUS(generators)
	virtualNetwork_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_STATUS{}), generators)

	return virtualNetwork_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetwork_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetwork_STATUS(gens map[string]gopter.Gen) {
	gens["EnableDdosProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableVmProtection"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetwork_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork_STATUS(gens map[string]gopter.Gen) {
	gens["AddressSpace"] = gen.PtrOf(AddressSpace_STATUSGenerator())
	gens["BgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_STATUSGenerator())
	gens["DdosProtectionPlan"] = gen.PtrOf(SubResource_STATUSGenerator())
	gens["DhcpOptions"] = gen.PtrOf(DhcpOptions_STATUSGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResource_STATUSGenerator())
}

func Test_AddressSpace_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AddressSpace to AddressSpace via AssignProperties_To_AddressSpace & AssignProperties_From_AddressSpace returns original",
		prop.ForAll(RunPropertyAssignmentTestForAddressSpace, AddressSpaceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAddressSpace tests if a specific instance of AddressSpace can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAddressSpace(subject AddressSpace) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.AddressSpace
	err := copied.AssignProperties_To_AddressSpace(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AddressSpace
	err = actual.AssignProperties_From_AddressSpace(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AddressSpace_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AddressSpace via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAddressSpace, AddressSpaceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAddressSpace runs a test to see if a specific instance of AddressSpace round trips to JSON and back losslessly
func RunJSONSerializationTestForAddressSpace(subject AddressSpace) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AddressSpace
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AddressSpace instances for property testing - lazily instantiated by AddressSpaceGenerator()
var addressSpaceGenerator gopter.Gen

// AddressSpaceGenerator returns a generator of AddressSpace instances for property testing.
func AddressSpaceGenerator() gopter.Gen {
	if addressSpaceGenerator != nil {
		return addressSpaceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAddressSpace(generators)
	addressSpaceGenerator = gen.Struct(reflect.TypeOf(AddressSpace{}), generators)

	return addressSpaceGenerator
}

// AddIndependentPropertyGeneratorsForAddressSpace is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAddressSpace(gens map[string]gopter.Gen) {
	gens["AddressPrefixes"] = gen.SliceOf(gen.AlphaString())
}

func Test_AddressSpace_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AddressSpace_STATUS to AddressSpace_STATUS via AssignProperties_To_AddressSpace_STATUS & AssignProperties_From_AddressSpace_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForAddressSpace_STATUS, AddressSpace_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAddressSpace_STATUS tests if a specific instance of AddressSpace_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAddressSpace_STATUS(subject AddressSpace_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.AddressSpace_STATUS
	err := copied.AssignProperties_To_AddressSpace_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AddressSpace_STATUS
	err = actual.AssignProperties_From_AddressSpace_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AddressSpace_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AddressSpace_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAddressSpace_STATUS, AddressSpace_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAddressSpace_STATUS runs a test to see if a specific instance of AddressSpace_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAddressSpace_STATUS(subject AddressSpace_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AddressSpace_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AddressSpace_STATUS instances for property testing - lazily instantiated by
// AddressSpace_STATUSGenerator()
var addressSpace_STATUSGenerator gopter.Gen

// AddressSpace_STATUSGenerator returns a generator of AddressSpace_STATUS instances for property testing.
func AddressSpace_STATUSGenerator() gopter.Gen {
	if addressSpace_STATUSGenerator != nil {
		return addressSpace_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAddressSpace_STATUS(generators)
	addressSpace_STATUSGenerator = gen.Struct(reflect.TypeOf(AddressSpace_STATUS{}), generators)

	return addressSpace_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAddressSpace_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAddressSpace_STATUS(gens map[string]gopter.Gen) {
	gens["AddressPrefixes"] = gen.SliceOf(gen.AlphaString())
}

func Test_DhcpOptions_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DhcpOptions to DhcpOptions via AssignProperties_To_DhcpOptions & AssignProperties_From_DhcpOptions returns original",
		prop.ForAll(RunPropertyAssignmentTestForDhcpOptions, DhcpOptionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDhcpOptions tests if a specific instance of DhcpOptions can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDhcpOptions(subject DhcpOptions) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.DhcpOptions
	err := copied.AssignProperties_To_DhcpOptions(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DhcpOptions
	err = actual.AssignProperties_From_DhcpOptions(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DhcpOptions_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DhcpOptions via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDhcpOptions, DhcpOptionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDhcpOptions runs a test to see if a specific instance of DhcpOptions round trips to JSON and back losslessly
func RunJSONSerializationTestForDhcpOptions(subject DhcpOptions) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DhcpOptions
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of DhcpOptions instances for property testing - lazily instantiated by DhcpOptionsGenerator()
var dhcpOptionsGenerator gopter.Gen

// DhcpOptionsGenerator returns a generator of DhcpOptions instances for property testing.
func DhcpOptionsGenerator() gopter.Gen {
	if dhcpOptionsGenerator != nil {
		return dhcpOptionsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDhcpOptions(generators)
	dhcpOptionsGenerator = gen.Struct(reflect.TypeOf(DhcpOptions{}), generators)

	return dhcpOptionsGenerator
}

// AddIndependentPropertyGeneratorsForDhcpOptions is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDhcpOptions(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
}

func Test_DhcpOptions_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DhcpOptions_STATUS to DhcpOptions_STATUS via AssignProperties_To_DhcpOptions_STATUS & AssignProperties_From_DhcpOptions_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDhcpOptions_STATUS, DhcpOptions_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDhcpOptions_STATUS tests if a specific instance of DhcpOptions_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDhcpOptions_STATUS(subject DhcpOptions_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.DhcpOptions_STATUS
	err := copied.AssignProperties_To_DhcpOptions_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DhcpOptions_STATUS
	err = actual.AssignProperties_From_DhcpOptions_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DhcpOptions_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DhcpOptions_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDhcpOptions_STATUS, DhcpOptions_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDhcpOptions_STATUS runs a test to see if a specific instance of DhcpOptions_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDhcpOptions_STATUS(subject DhcpOptions_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DhcpOptions_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of DhcpOptions_STATUS instances for property testing - lazily instantiated by DhcpOptions_STATUSGenerator()
var dhcpOptions_STATUSGenerator gopter.Gen

// DhcpOptions_STATUSGenerator returns a generator of DhcpOptions_STATUS instances for property testing.
func DhcpOptions_STATUSGenerator() gopter.Gen {
	if dhcpOptions_STATUSGenerator != nil {
		return dhcpOptions_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDhcpOptions_STATUS(generators)
	dhcpOptions_STATUSGenerator = gen.Struct(reflect.TypeOf(DhcpOptions_STATUS{}), generators)

	return dhcpOptions_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDhcpOptions_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDhcpOptions_STATUS(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
}

func Test_VirtualNetworkBgpCommunities_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworkBgpCommunities to VirtualNetworkBgpCommunities via AssignProperties_To_VirtualNetworkBgpCommunities & AssignProperties_From_VirtualNetworkBgpCommunities returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworkBgpCommunities, VirtualNetworkBgpCommunitiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworkBgpCommunities tests if a specific instance of VirtualNetworkBgpCommunities can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworkBgpCommunities(subject VirtualNetworkBgpCommunities) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.VirtualNetworkBgpCommunities
	err := copied.AssignProperties_To_VirtualNetworkBgpCommunities(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworkBgpCommunities
	err = actual.AssignProperties_From_VirtualNetworkBgpCommunities(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetworkBgpCommunities_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunities via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunities, VirtualNetworkBgpCommunitiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunities runs a test to see if a specific instance of VirtualNetworkBgpCommunities round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunities(subject VirtualNetworkBgpCommunities) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunities
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetworkBgpCommunities instances for property testing - lazily instantiated by
// VirtualNetworkBgpCommunitiesGenerator()
var virtualNetworkBgpCommunitiesGenerator gopter.Gen

// VirtualNetworkBgpCommunitiesGenerator returns a generator of VirtualNetworkBgpCommunities instances for property testing.
func VirtualNetworkBgpCommunitiesGenerator() gopter.Gen {
	if virtualNetworkBgpCommunitiesGenerator != nil {
		return virtualNetworkBgpCommunitiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities(generators)
	virtualNetworkBgpCommunitiesGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunities{}), generators)

	return virtualNetworkBgpCommunitiesGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities(gens map[string]gopter.Gen) {
	gens["VirtualNetworkCommunity"] = gen.PtrOf(gen.AlphaString())
}

func Test_VirtualNetworkBgpCommunities_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworkBgpCommunities_STATUS to VirtualNetworkBgpCommunities_STATUS via AssignProperties_To_VirtualNetworkBgpCommunities_STATUS & AssignProperties_From_VirtualNetworkBgpCommunities_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworkBgpCommunities_STATUS, VirtualNetworkBgpCommunities_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworkBgpCommunities_STATUS tests if a specific instance of VirtualNetworkBgpCommunities_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworkBgpCommunities_STATUS(subject VirtualNetworkBgpCommunities_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.VirtualNetworkBgpCommunities_STATUS
	err := copied.AssignProperties_To_VirtualNetworkBgpCommunities_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworkBgpCommunities_STATUS
	err = actual.AssignProperties_From_VirtualNetworkBgpCommunities_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetworkBgpCommunities_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunities_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUS, VirtualNetworkBgpCommunities_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUS runs a test to see if a specific instance of VirtualNetworkBgpCommunities_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUS(subject VirtualNetworkBgpCommunities_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunities_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetworkBgpCommunities_STATUS instances for property testing - lazily instantiated by
// VirtualNetworkBgpCommunities_STATUSGenerator()
var virtualNetworkBgpCommunities_STATUSGenerator gopter.Gen

// VirtualNetworkBgpCommunities_STATUSGenerator returns a generator of VirtualNetworkBgpCommunities_STATUS instances for property testing.
func VirtualNetworkBgpCommunities_STATUSGenerator() gopter.Gen {
	if virtualNetworkBgpCommunities_STATUSGenerator != nil {
		return virtualNetworkBgpCommunities_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUS(generators)
	virtualNetworkBgpCommunities_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunities_STATUS{}), generators)

	return virtualNetworkBgpCommunities_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUS(gens map[string]gopter.Gen) {
	gens["RegionalCommunity"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkCommunity"] = gen.PtrOf(gen.AlphaString())
}
