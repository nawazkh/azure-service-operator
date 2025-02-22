// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_MongoDBDatabaseGetProperties_Resource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseGetProperties_Resource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUS, MongoDBDatabaseGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUS runs a test to see if a specific instance of MongoDBDatabaseGetProperties_Resource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUS(subject MongoDBDatabaseGetProperties_Resource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseGetProperties_Resource_STATUS
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

// Generator of MongoDBDatabaseGetProperties_Resource_STATUS instances for property testing - lazily instantiated by
// MongoDBDatabaseGetProperties_Resource_STATUSGenerator()
var mongoDBDatabaseGetProperties_Resource_STATUSGenerator gopter.Gen

// MongoDBDatabaseGetProperties_Resource_STATUSGenerator returns a generator of MongoDBDatabaseGetProperties_Resource_STATUS instances for property testing.
// We first initialize mongoDBDatabaseGetProperties_Resource_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBDatabaseGetProperties_Resource_STATUSGenerator() gopter.Gen {
	if mongoDBDatabaseGetProperties_Resource_STATUSGenerator != nil {
		return mongoDBDatabaseGetProperties_Resource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS(generators)
	mongoDBDatabaseGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetProperties_Resource_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS(generators)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS(generators)
	mongoDBDatabaseGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetProperties_Resource_STATUS{}), generators)

	return mongoDBDatabaseGetProperties_Resource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(CreateMode_STATUS_Default, CreateMode_STATUS_Restore))
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["RestoreParameters"] = gen.PtrOf(RestoreParametersBase_STATUSGenerator())
}

func Test_MongoDBDatabaseGetProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseGetProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseGetProperties_STATUS, MongoDBDatabaseGetProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseGetProperties_STATUS runs a test to see if a specific instance of MongoDBDatabaseGetProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseGetProperties_STATUS(subject MongoDBDatabaseGetProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseGetProperties_STATUS
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

// Generator of MongoDBDatabaseGetProperties_STATUS instances for property testing - lazily instantiated by
// MongoDBDatabaseGetProperties_STATUSGenerator()
var mongoDBDatabaseGetProperties_STATUSGenerator gopter.Gen

// MongoDBDatabaseGetProperties_STATUSGenerator returns a generator of MongoDBDatabaseGetProperties_STATUS instances for property testing.
func MongoDBDatabaseGetProperties_STATUSGenerator() gopter.Gen {
	if mongoDBDatabaseGetProperties_STATUSGenerator != nil {
		return mongoDBDatabaseGetProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_STATUS(generators)
	mongoDBDatabaseGetProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetProperties_STATUS{}), generators)

	return mongoDBDatabaseGetProperties_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResource_STATUSGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseGetProperties_Resource_STATUSGenerator())
}

func Test_MongodbDatabase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabase_STATUS, MongodbDatabase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabase_STATUS runs a test to see if a specific instance of MongodbDatabase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabase_STATUS(subject MongodbDatabase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabase_STATUS
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

// Generator of MongodbDatabase_STATUS instances for property testing - lazily instantiated by
// MongodbDatabase_STATUSGenerator()
var mongodbDatabase_STATUSGenerator gopter.Gen

// MongodbDatabase_STATUSGenerator returns a generator of MongodbDatabase_STATUS instances for property testing.
// We first initialize mongodbDatabase_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongodbDatabase_STATUSGenerator() gopter.Gen {
	if mongodbDatabase_STATUSGenerator != nil {
		return mongodbDatabase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabase_STATUS(generators)
	mongodbDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(MongodbDatabase_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabase_STATUS(generators)
	AddRelatedPropertyGeneratorsForMongodbDatabase_STATUS(generators)
	mongodbDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(MongodbDatabase_STATUS{}), generators)

	return mongodbDatabase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMongodbDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongodbDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongodbDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(MongoDBDatabaseGetProperties_STATUSGenerator())
}
