// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
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

func Test_Replica_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Replica via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReplica, ReplicaGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReplica runs a test to see if a specific instance of Replica round trips to JSON and back losslessly
func RunJSONSerializationTestForReplica(subject Replica) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Replica
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

// Generator of Replica instances for property testing - lazily instantiated by ReplicaGenerator()
var replicaGenerator gopter.Gen

// ReplicaGenerator returns a generator of Replica instances for property testing.
func ReplicaGenerator() gopter.Gen {
	if replicaGenerator != nil {
		return replicaGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForReplica(generators)
	replicaGenerator = gen.Struct(reflect.TypeOf(Replica{}), generators)

	return replicaGenerator
}

// AddRelatedPropertyGeneratorsForReplica is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReplica(gens map[string]gopter.Gen) {
	gens["Spec"] = Replica_SpecGenerator()
	gens["Status"] = Replica_STATUSGenerator()
}

func Test_ReplicaOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReplicaOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReplicaOperatorSpec, ReplicaOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReplicaOperatorSpec runs a test to see if a specific instance of ReplicaOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForReplicaOperatorSpec(subject ReplicaOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReplicaOperatorSpec
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

// Generator of ReplicaOperatorSpec instances for property testing - lazily instantiated by
// ReplicaOperatorSpecGenerator()
var replicaOperatorSpecGenerator gopter.Gen

// ReplicaOperatorSpecGenerator returns a generator of ReplicaOperatorSpec instances for property testing.
func ReplicaOperatorSpecGenerator() gopter.Gen {
	if replicaOperatorSpecGenerator != nil {
		return replicaOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	replicaOperatorSpecGenerator = gen.Struct(reflect.TypeOf(ReplicaOperatorSpec{}), generators)

	return replicaOperatorSpecGenerator
}

func Test_Replica_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Replica_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReplica_STATUS, Replica_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReplica_STATUS runs a test to see if a specific instance of Replica_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForReplica_STATUS(subject Replica_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Replica_STATUS
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

// Generator of Replica_STATUS instances for property testing - lazily instantiated by Replica_STATUSGenerator()
var replica_STATUSGenerator gopter.Gen

// Replica_STATUSGenerator returns a generator of Replica_STATUS instances for property testing.
// We first initialize replica_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Replica_STATUSGenerator() gopter.Gen {
	if replica_STATUSGenerator != nil {
		return replica_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReplica_STATUS(generators)
	replica_STATUSGenerator = gen.Struct(reflect.TypeOf(Replica_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReplica_STATUS(generators)
	AddRelatedPropertyGeneratorsForReplica_STATUS(generators)
	replica_STATUSGenerator = gen.Struct(reflect.TypeOf(Replica_STATUS{}), generators)

	return replica_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForReplica_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReplica_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["RegionEndpointEnabled"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceStopped"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReplica_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReplica_STATUS(gens map[string]gopter.Gen) {
	gens["Sku"] = gen.PtrOf(ResourceSku_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Replica_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Replica_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReplica_Spec, Replica_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReplica_Spec runs a test to see if a specific instance of Replica_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForReplica_Spec(subject Replica_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Replica_Spec
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

// Generator of Replica_Spec instances for property testing - lazily instantiated by Replica_SpecGenerator()
var replica_SpecGenerator gopter.Gen

// Replica_SpecGenerator returns a generator of Replica_Spec instances for property testing.
// We first initialize replica_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Replica_SpecGenerator() gopter.Gen {
	if replica_SpecGenerator != nil {
		return replica_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReplica_Spec(generators)
	replica_SpecGenerator = gen.Struct(reflect.TypeOf(Replica_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReplica_Spec(generators)
	AddRelatedPropertyGeneratorsForReplica_Spec(generators)
	replica_SpecGenerator = gen.Struct(reflect.TypeOf(Replica_Spec{}), generators)

	return replica_SpecGenerator
}

// AddIndependentPropertyGeneratorsForReplica_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReplica_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["RegionEndpointEnabled"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceStopped"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReplica_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReplica_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(ReplicaOperatorSpecGenerator())
	gens["Sku"] = gen.PtrOf(ResourceSkuGenerator())
}

func Test_ResourceSku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceSku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceSku, ResourceSkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceSku runs a test to see if a specific instance of ResourceSku round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceSku(subject ResourceSku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceSku
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

// Generator of ResourceSku instances for property testing - lazily instantiated by ResourceSkuGenerator()
var resourceSkuGenerator gopter.Gen

// ResourceSkuGenerator returns a generator of ResourceSku instances for property testing.
func ResourceSkuGenerator() gopter.Gen {
	if resourceSkuGenerator != nil {
		return resourceSkuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceSku(generators)
	resourceSkuGenerator = gen.Struct(reflect.TypeOf(ResourceSku{}), generators)

	return resourceSkuGenerator
}

// AddIndependentPropertyGeneratorsForResourceSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceSku(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_ResourceSku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceSku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceSku_STATUS, ResourceSku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceSku_STATUS runs a test to see if a specific instance of ResourceSku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceSku_STATUS(subject ResourceSku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceSku_STATUS
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

// Generator of ResourceSku_STATUS instances for property testing - lazily instantiated by ResourceSku_STATUSGenerator()
var resourceSku_STATUSGenerator gopter.Gen

// ResourceSku_STATUSGenerator returns a generator of ResourceSku_STATUS instances for property testing.
func ResourceSku_STATUSGenerator() gopter.Gen {
	if resourceSku_STATUSGenerator != nil {
		return resourceSku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceSku_STATUS(generators)
	resourceSku_STATUSGenerator = gen.Struct(reflect.TypeOf(ResourceSku_STATUS{}), generators)

	return resourceSku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForResourceSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceSku_STATUS(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Size"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}
