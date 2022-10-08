package test

import (
	"awesomeProject/utils"
	"io"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestHomeEndpoint(t *testing.T) {
	homeEndpoint := "http://localhost:8080/home"
	testName := "Home endpoint response matching test"
	want := "Home sweet home"

	data, err := http.Get(homeEndpoint)
	if err != nil {
		t.Fatalf("%s. Error: %s", testName, err)
	}

	got, err := io.ReadAll(data.Body)
	if err != nil {
		t.Fatalf("%s. Error: %s", testName, err)
	}

	if want != string(got) {
		t.Fatalf("%s. wanted: %s, got: %s", testName, want, got)
	}
}

func TestApiEndpoint(t *testing.T) {
	apiEndpoint := "http://localhost:8080/api"
	testName := "Api endpoint response matching test"
	want := regexp.MustCompile("{\"message\":\"[A-Za-z0-9,:;\\-'\"\\s]*}")
	endpoints := []string{
		"/commits",
		"/repos",
		"/branches",
	}

	for _, endpoint := range endpoints {
		data, err := http.Get(apiEndpoint + endpoint)
		if err != nil {
			t.Fatalf("%s. Error: %s", testName, err)
		}

		readData, err := io.ReadAll(data.Body)
		if err != nil {
			t.Fatalf("%s. Error: %s", testName, err)
		}

		got := string(readData)

		if !want.MatchString(got) || err != nil {
			t.Fatalf("%s. wanted: %s, got: %s", testName, want, got)
		}
	}
}

func TestInternalEndpoint(t *testing.T) {
	apiEndpoint := "http://localhost:8080/api/internal/"
	testName := "Internal Api endpoint response matching test"
	want := "{\"message\":\"Internal Route\"}"

	data, err := http.Get(apiEndpoint)
	if err != nil {
		t.Fatalf("%s. Error: %s", testName, err)
	}

	got, err := io.ReadAll(data.Body)
	if err != nil {
		t.Fatalf("%s. Error: %s", testName, err)
	}

	if strings.Compare(want, string(got)) == 0 {
		t.Fatalf("%s. wanted: %s, got: %s", testName, want, got)
	}
}

func TestCalculateResponseTime(t *testing.T) {
	apiEndpoint := "http://localhost:8080/api/internal/calculateResponseTime"
	testName := "Internal calculate response time Api endpoint test"
	want := regexp.MustCompile("{\"message\":\"Time for processing request to \\w* endpoint: [0-9.]*s\"}")
	endpoints := []string{
		"/commits",
		"/repos",
		"/branches",
	}

	for _, endpoint := range endpoints {
		data, err := http.Get(apiEndpoint + endpoint)
		if err != nil {
			t.Fatalf("%s. Error: %s", testName, err)
		}

		readData, err := io.ReadAll(data.Body)
		if err != nil {
			t.Fatalf("%s. Error: %s", testName, err)
		}

		got := string(readData)

		if !want.MatchString(got) || err != nil {
			t.Fatalf("%s. wanted: %s, got: %s", testName, want, got)
		}
	}
}

func TestMockReposApi(t *testing.T) {
	apiEndpoint := "https://60a21d3f745cd70017576092.mockapi.io/api/v1/repos"
	testName := "Mock repos Api endpoint response JSON matching test"
	absPath, _ := filepath.Abs("./schemas/repos-schema.json")

	data, err := http.Get(apiEndpoint)
	if err != nil {
		t.Fatalf("%s. Error: %s", testName, err)
	}

	readData, err := io.ReadAll(data.Body)
	if err != nil {
		t.Fatalf("%s. Error: %s", testName, err)
	}

	validationErrors := utils.ValidateJson(absPath, readData)

	if validationErrors != nil {
		t.Fatalf("%s. Error: %s", testName, validationErrors)
	}
}

func TestMockBranchesApi(t *testing.T) {
	apiEndpoint := "https://60a21d3f745cd70017576092.mockapi.io/api/v1/repos/1/branches"
	testName := "Mock branches Api endpoint response JSON matching test"
	absPath, _ := filepath.Abs("./schemas/branches-schema.json")

	data, err := http.Get(apiEndpoint)
	if err != nil {
		t.Fatalf("%s. Error: %s", testName, err)
	}

	readData, err := io.ReadAll(data.Body)
	if err != nil {
		t.Fatalf("%s. Error: %s", testName, err)
	}

	validationErrors := utils.ValidateJson(absPath, readData)

	if validationErrors != nil {
		t.Fatalf("%s. Error: %s", testName, validationErrors)
	}
}

func TestMockCommitsApi(t *testing.T) {
	apiEndpoint := "https://60a21d3f745cd70017576092.mockapi.io/api/v1/repos/1/branches/1/commits"
	testName := "Mock commits Api endpoint response JSON matching test"
	absPath, _ := filepath.Abs("./schemas/commits-schema.json")

	data, err := http.Get(apiEndpoint)
	if err != nil {
		t.Fatalf("%s. Error: %s", testName, err)
	}

	readData, err := io.ReadAll(data.Body)
	if err != nil {
		t.Fatalf("%s. Error: %s", testName, err)
	}

	validationErrors := utils.ValidateJson(absPath, readData)

	if validationErrors != nil {
		t.Fatalf("%s. Error: %s", testName, validationErrors)
	}
}
