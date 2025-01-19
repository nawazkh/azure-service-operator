// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101/storage"
	v20240101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20240101/storage"
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

func Test_NamespacesEventhubsConsumerGroup_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsConsumerGroup to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesEventhubsConsumerGroup, NamespacesEventhubsConsumerGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesEventhubsConsumerGroup tests if a specific instance of NamespacesEventhubsConsumerGroup round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesEventhubsConsumerGroup(subject NamespacesEventhubsConsumerGroup) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20240101s.NamespacesEventhubsConsumerGroup
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesEventhubsConsumerGroup
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

func Test_NamespacesEventhubsConsumerGroup_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsConsumerGroup to NamespacesEventhubsConsumerGroup via AssignProperties_To_NamespacesEventhubsConsumerGroup & AssignProperties_From_NamespacesEventhubsConsumerGroup returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup, NamespacesEventhubsConsumerGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup tests if a specific instance of NamespacesEventhubsConsumerGroup can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup(subject NamespacesEventhubsConsumerGroup) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.NamespacesEventhubsConsumerGroup
	err := copied.AssignProperties_To_NamespacesEventhubsConsumerGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubsConsumerGroup
	err = actual.AssignProperties_From_NamespacesEventhubsConsumerGroup(&other)
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

func Test_NamespacesEventhubsConsumerGroup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsConsumerGroup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsConsumerGroup, NamespacesEventhubsConsumerGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsConsumerGroup runs a test to see if a specific instance of NamespacesEventhubsConsumerGroup round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsConsumerGroup(subject NamespacesEventhubsConsumerGroup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsConsumerGroup
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

// Generator of NamespacesEventhubsConsumerGroup instances for property testing - lazily instantiated by
// NamespacesEventhubsConsumerGroupGenerator()
var namespacesEventhubsConsumerGroupGenerator gopter.Gen

// NamespacesEventhubsConsumerGroupGenerator returns a generator of NamespacesEventhubsConsumerGroup instances for property testing.
func NamespacesEventhubsConsumerGroupGenerator() gopter.Gen {
	if namespacesEventhubsConsumerGroupGenerator != nil {
		return namespacesEventhubsConsumerGroupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup(generators)
	namespacesEventhubsConsumerGroupGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumerGroup{}), generators)

	return namespacesEventhubsConsumerGroupGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesEventhubsConsumerGroup_SpecGenerator()
	gens["Status"] = NamespacesEventhubsConsumerGroup_STATUSGenerator()
}

func Test_NamespacesEventhubsConsumerGroupOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsConsumerGroupOperatorSpec to NamespacesEventhubsConsumerGroupOperatorSpec via AssignProperties_To_NamespacesEventhubsConsumerGroupOperatorSpec & AssignProperties_From_NamespacesEventhubsConsumerGroupOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroupOperatorSpec, NamespacesEventhubsConsumerGroupOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroupOperatorSpec tests if a specific instance of NamespacesEventhubsConsumerGroupOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroupOperatorSpec(subject NamespacesEventhubsConsumerGroupOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.NamespacesEventhubsConsumerGroupOperatorSpec
	err := copied.AssignProperties_To_NamespacesEventhubsConsumerGroupOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubsConsumerGroupOperatorSpec
	err = actual.AssignProperties_From_NamespacesEventhubsConsumerGroupOperatorSpec(&other)
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

func Test_NamespacesEventhubsConsumerGroupOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsConsumerGroupOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsConsumerGroupOperatorSpec, NamespacesEventhubsConsumerGroupOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsConsumerGroupOperatorSpec runs a test to see if a specific instance of NamespacesEventhubsConsumerGroupOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsConsumerGroupOperatorSpec(subject NamespacesEventhubsConsumerGroupOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsConsumerGroupOperatorSpec
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

// Generator of NamespacesEventhubsConsumerGroupOperatorSpec instances for property testing - lazily instantiated by
// NamespacesEventhubsConsumerGroupOperatorSpecGenerator()
var namespacesEventhubsConsumerGroupOperatorSpecGenerator gopter.Gen

// NamespacesEventhubsConsumerGroupOperatorSpecGenerator returns a generator of NamespacesEventhubsConsumerGroupOperatorSpec instances for property testing.
func NamespacesEventhubsConsumerGroupOperatorSpecGenerator() gopter.Gen {
	if namespacesEventhubsConsumerGroupOperatorSpecGenerator != nil {
		return namespacesEventhubsConsumerGroupOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	namespacesEventhubsConsumerGroupOperatorSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumerGroupOperatorSpec{}), generators)

	return namespacesEventhubsConsumerGroupOperatorSpecGenerator
}

func Test_NamespacesEventhubsConsumerGroup_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsConsumerGroup_STATUS to NamespacesEventhubsConsumerGroup_STATUS via AssignProperties_To_NamespacesEventhubsConsumerGroup_STATUS & AssignProperties_From_NamespacesEventhubsConsumerGroup_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup_STATUS, NamespacesEventhubsConsumerGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup_STATUS tests if a specific instance of NamespacesEventhubsConsumerGroup_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup_STATUS(subject NamespacesEventhubsConsumerGroup_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.NamespacesEventhubsConsumerGroup_STATUS
	err := copied.AssignProperties_To_NamespacesEventhubsConsumerGroup_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubsConsumerGroup_STATUS
	err = actual.AssignProperties_From_NamespacesEventhubsConsumerGroup_STATUS(&other)
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

func Test_NamespacesEventhubsConsumerGroup_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsConsumerGroup_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsConsumerGroup_STATUS, NamespacesEventhubsConsumerGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsConsumerGroup_STATUS runs a test to see if a specific instance of NamespacesEventhubsConsumerGroup_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsConsumerGroup_STATUS(subject NamespacesEventhubsConsumerGroup_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsConsumerGroup_STATUS
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

// Generator of NamespacesEventhubsConsumerGroup_STATUS instances for property testing - lazily instantiated by
// NamespacesEventhubsConsumerGroup_STATUSGenerator()
var namespacesEventhubsConsumerGroup_STATUSGenerator gopter.Gen

// NamespacesEventhubsConsumerGroup_STATUSGenerator returns a generator of NamespacesEventhubsConsumerGroup_STATUS instances for property testing.
// We first initialize namespacesEventhubsConsumerGroup_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhubsConsumerGroup_STATUSGenerator() gopter.Gen {
	if namespacesEventhubsConsumerGroup_STATUSGenerator != nil {
		return namespacesEventhubsConsumerGroup_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumerGroup_STATUS(generators)
	namespacesEventhubsConsumerGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumerGroup_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumerGroup_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup_STATUS(generators)
	namespacesEventhubsConsumerGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumerGroup_STATUS{}), generators)

	return namespacesEventhubsConsumerGroup_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumerGroup_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumerGroup_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["UserMetadata"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_NamespacesEventhubsConsumerGroup_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsConsumerGroup_Spec to NamespacesEventhubsConsumerGroup_Spec via AssignProperties_To_NamespacesEventhubsConsumerGroup_Spec & AssignProperties_From_NamespacesEventhubsConsumerGroup_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup_Spec, NamespacesEventhubsConsumerGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup_Spec tests if a specific instance of NamespacesEventhubsConsumerGroup_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup_Spec(subject NamespacesEventhubsConsumerGroup_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.NamespacesEventhubsConsumerGroup_Spec
	err := copied.AssignProperties_To_NamespacesEventhubsConsumerGroup_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubsConsumerGroup_Spec
	err = actual.AssignProperties_From_NamespacesEventhubsConsumerGroup_Spec(&other)
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

func Test_NamespacesEventhubsConsumerGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsConsumerGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsConsumerGroup_Spec, NamespacesEventhubsConsumerGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsConsumerGroup_Spec runs a test to see if a specific instance of NamespacesEventhubsConsumerGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsConsumerGroup_Spec(subject NamespacesEventhubsConsumerGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsConsumerGroup_Spec
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

// Generator of NamespacesEventhubsConsumerGroup_Spec instances for property testing - lazily instantiated by
// NamespacesEventhubsConsumerGroup_SpecGenerator()
var namespacesEventhubsConsumerGroup_SpecGenerator gopter.Gen

// NamespacesEventhubsConsumerGroup_SpecGenerator returns a generator of NamespacesEventhubsConsumerGroup_Spec instances for property testing.
// We first initialize namespacesEventhubsConsumerGroup_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhubsConsumerGroup_SpecGenerator() gopter.Gen {
	if namespacesEventhubsConsumerGroup_SpecGenerator != nil {
		return namespacesEventhubsConsumerGroup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumerGroup_Spec(generators)
	namespacesEventhubsConsumerGroup_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumerGroup_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumerGroup_Spec(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup_Spec(generators)
	namespacesEventhubsConsumerGroup_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumerGroup_Spec{}), generators)

	return namespacesEventhubsConsumerGroup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumerGroup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumerGroup_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["UserMetadata"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(NamespacesEventhubsConsumerGroupOperatorSpecGenerator())
}
