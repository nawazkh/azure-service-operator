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

func Test_ServersDatabasesAuditingSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesAuditingSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesAuditingSetting, ServersDatabasesAuditingSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesAuditingSetting runs a test to see if a specific instance of ServersDatabasesAuditingSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesAuditingSetting(subject ServersDatabasesAuditingSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesAuditingSetting
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

// Generator of ServersDatabasesAuditingSetting instances for property testing - lazily instantiated by
// ServersDatabasesAuditingSettingGenerator()
var serversDatabasesAuditingSettingGenerator gopter.Gen

// ServersDatabasesAuditingSettingGenerator returns a generator of ServersDatabasesAuditingSetting instances for property testing.
func ServersDatabasesAuditingSettingGenerator() gopter.Gen {
	if serversDatabasesAuditingSettingGenerator != nil {
		return serversDatabasesAuditingSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersDatabasesAuditingSetting(generators)
	serversDatabasesAuditingSettingGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesAuditingSetting{}), generators)

	return serversDatabasesAuditingSettingGenerator
}

// AddRelatedPropertyGeneratorsForServersDatabasesAuditingSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesAuditingSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = ServersDatabasesAuditingSetting_SpecGenerator()
	gens["Status"] = ServersDatabasesAuditingSetting_STATUSGenerator()
}

func Test_ServersDatabasesAuditingSettingOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesAuditingSettingOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesAuditingSettingOperatorSpec, ServersDatabasesAuditingSettingOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesAuditingSettingOperatorSpec runs a test to see if a specific instance of ServersDatabasesAuditingSettingOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesAuditingSettingOperatorSpec(subject ServersDatabasesAuditingSettingOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesAuditingSettingOperatorSpec
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

// Generator of ServersDatabasesAuditingSettingOperatorSpec instances for property testing - lazily instantiated by
// ServersDatabasesAuditingSettingOperatorSpecGenerator()
var serversDatabasesAuditingSettingOperatorSpecGenerator gopter.Gen

// ServersDatabasesAuditingSettingOperatorSpecGenerator returns a generator of ServersDatabasesAuditingSettingOperatorSpec instances for property testing.
func ServersDatabasesAuditingSettingOperatorSpecGenerator() gopter.Gen {
	if serversDatabasesAuditingSettingOperatorSpecGenerator != nil {
		return serversDatabasesAuditingSettingOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	serversDatabasesAuditingSettingOperatorSpecGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesAuditingSettingOperatorSpec{}), generators)

	return serversDatabasesAuditingSettingOperatorSpecGenerator
}

func Test_ServersDatabasesAuditingSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesAuditingSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesAuditingSetting_STATUS, ServersDatabasesAuditingSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesAuditingSetting_STATUS runs a test to see if a specific instance of ServersDatabasesAuditingSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesAuditingSetting_STATUS(subject ServersDatabasesAuditingSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesAuditingSetting_STATUS
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

// Generator of ServersDatabasesAuditingSetting_STATUS instances for property testing - lazily instantiated by
// ServersDatabasesAuditingSetting_STATUSGenerator()
var serversDatabasesAuditingSetting_STATUSGenerator gopter.Gen

// ServersDatabasesAuditingSetting_STATUSGenerator returns a generator of ServersDatabasesAuditingSetting_STATUS instances for property testing.
func ServersDatabasesAuditingSetting_STATUSGenerator() gopter.Gen {
	if serversDatabasesAuditingSetting_STATUSGenerator != nil {
		return serversDatabasesAuditingSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesAuditingSetting_STATUS(generators)
	serversDatabasesAuditingSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesAuditingSetting_STATUS{}), generators)

	return serversDatabasesAuditingSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServersDatabasesAuditingSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersDatabasesAuditingSetting_STATUS(gens map[string]gopter.Gen) {
	gens["AuditActionsAndGroups"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsAzureMonitorTargetEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsManagedIdentityInUse"] = gen.PtrOf(gen.Bool())
	gens["IsStorageSecondaryKeyInUse"] = gen.PtrOf(gen.Bool())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["QueueDelayMs"] = gen.PtrOf(gen.Int())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersDatabasesAuditingSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesAuditingSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesAuditingSetting_Spec, ServersDatabasesAuditingSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesAuditingSetting_Spec runs a test to see if a specific instance of ServersDatabasesAuditingSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesAuditingSetting_Spec(subject ServersDatabasesAuditingSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesAuditingSetting_Spec
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

// Generator of ServersDatabasesAuditingSetting_Spec instances for property testing - lazily instantiated by
// ServersDatabasesAuditingSetting_SpecGenerator()
var serversDatabasesAuditingSetting_SpecGenerator gopter.Gen

// ServersDatabasesAuditingSetting_SpecGenerator returns a generator of ServersDatabasesAuditingSetting_Spec instances for property testing.
// We first initialize serversDatabasesAuditingSetting_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersDatabasesAuditingSetting_SpecGenerator() gopter.Gen {
	if serversDatabasesAuditingSetting_SpecGenerator != nil {
		return serversDatabasesAuditingSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesAuditingSetting_Spec(generators)
	serversDatabasesAuditingSetting_SpecGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesAuditingSetting_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesAuditingSetting_Spec(generators)
	AddRelatedPropertyGeneratorsForServersDatabasesAuditingSetting_Spec(generators)
	serversDatabasesAuditingSetting_SpecGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesAuditingSetting_Spec{}), generators)

	return serversDatabasesAuditingSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersDatabasesAuditingSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersDatabasesAuditingSetting_Spec(gens map[string]gopter.Gen) {
	gens["AuditActionsAndGroups"] = gen.SliceOf(gen.AlphaString())
	gens["IsAzureMonitorTargetEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsManagedIdentityInUse"] = gen.PtrOf(gen.Bool())
	gens["IsStorageSecondaryKeyInUse"] = gen.PtrOf(gen.Bool())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["QueueDelayMs"] = gen.PtrOf(gen.Int())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServersDatabasesAuditingSetting_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesAuditingSetting_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(ServersDatabasesAuditingSettingOperatorSpecGenerator())
}
