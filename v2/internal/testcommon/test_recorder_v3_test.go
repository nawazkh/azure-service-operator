/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"io"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/config"
)

func TestRecorderV3_WhenRecordingAndRecordingDoesNotExist_MakesRecording(t *testing.T) {
	// NB: Can't run tests using Setenv() in parallel
	t.Setenv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	t.Setenv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")
	
	g := NewGomegaWithT(t)

	cfg := config.Values{}
	cassetteName := "recordings/" + t.Name()

	// Test prerequisite: The recording must not already exist
	// We delete the recording at the end of the test, so it shouldn't ever be committed
	// But if something goes awry, and we don't clean up, we need to flag the presence of
	// this file as a failure.
	// We're noisy about it (instead of just deleting the file proactively) in order to
	// ensure the dev knows the file shouldn't be committed.
	exists, err := cassetteFileExists(cassetteName)
	g.Expect(err).To(BeNil())
	g.Expect(exists).To(BeFalse())

	// Ensure we clean up the cassette file at the end of the test
	cassetteFile := cassetteFileName(cassetteName)
	defer func() {
		err := os.Remove(cassetteFile)
		g.Expect(err).To(BeNil())
	}()

	// Create our TestRecorder and ensure it's recording
	recorder, err := newTestRecorderV3(cassetteName, cfg, true)
	g.Expect(err).To(BeNil())
	g.Expect(recorder.IsReplaying()).To(BeFalse())

	url := "https://www.bing.com"
	client := recorder.CreateClient(t)

	// Make sure we can get a response from the internet
	resp, err := client.Get(url)
	g.Expect(err).To(BeNil())

	// Ensure the body is not empty
	body, err := io.ReadAll(resp.Body)
	g.Expect(err).To(BeNil())
	g.Expect(body).NotTo(HaveLen(0))

	// Stop the recorder
	err = recorder.Stop()
	g.Expect(err).To(BeNil())

	// Verify we created a recording
	exists, err = cassetteFileExists(cassetteName)
	g.Expect(err).To(BeNil())
	g.Expect(exists).To(BeTrue())
}

func TestRecorderV3_WhenRecordingAndRecordingExists_DoesPlayback(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := config.Values{}
	cassetteName := "recordings/" + t.Name()

	// Test prerequisite: The recording must already exist
	exists, err := cassetteFileExists(cassetteName)
	g.Expect(err).To(BeNil())
	g.Expect(exists).To(BeTrue())

	// Create our TestRecorder and ensure it's recording
	recorder, err := newTestRecorderV3(cassetteName, cfg, false)
	g.Expect(err).To(BeNil())
	g.Expect(recorder.IsReplaying()).To(BeTrue())

	url := "https://www.bing.com"
	client := recorder.CreateClient(t)

	// Make sure we can get a response from the internet
	resp, err := client.Get(url)
	g.Expect(err).To(BeNil())

	// Ensure the body is not empty
	body, err := io.ReadAll(resp.Body)
	g.Expect(err).To(BeNil())
	g.Expect(body).NotTo(HaveLen(0))

	// Stop the recorder
	err = recorder.Stop()
	g.Expect(err).To(BeNil())
}

func TestRecorderV3_WhenPlayingBackAndRecordingDoesNotExist_ReturnsErrorOnCreation(t *testing.T) {
	t.Parallel()
	
	g := NewGomegaWithT(t)

	cfg := config.Values{}
	cassetteName := "recordings/" + t.Name()

	// Create our TestRecorder and ensure it fails to create because we have no recording
	_, err := newTestRecorderV3(cassetteName, cfg, false)
	g.Expect(err).NotTo(BeNil())
}
