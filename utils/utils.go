// utils package contains all the helper functions for the application that can be used from different packages(src code, test code).
package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/xeipuuv/gojsonschema"
)

func BuildURL(r *http.Request) string {
	apiURL := "https://60a21d3f745cd70017576092.mockapi.io/api/v1/repos"
	endpoint := mux.Vars(r)["endpoint"]

	switch endpoint {
	case "repos":
		return fmt.Sprintf(apiURL)
	case "branches":
		return fmt.Sprintf(apiURL + "/1/branches")
	case "commits":
		return fmt.Sprintf(apiURL + "/1/branches/1/commits")
	default:
		return apiURL
	}
}

func BuildInternalURL(r *http.Request) string {
	apiURL := "http://localhost:8080/api"
	endpoint := mux.Vars(r)["endpoint"]

	switch endpoint {
	case "repos":
		return fmt.Sprintf(apiURL + "/repos")
	case "branches":
		return fmt.Sprintf(apiURL + "/branches")
	case "commits":
		return fmt.Sprintf(apiURL + "/commits")
	default:
		return apiURL
	}
}

func ValidateJson(JSONSchemaFilePath string, JSONResponse []byte) []gojsonschema.ResultError {
	// note that absPath returns different results on different operating systems.
	// for example, in Windows it will return "C:\\Users\\user" format, which will throw error when loading json with
	// gojsonschema reference loader. But in Linux everything works okay. Threfore all '\\' replaced with '/'
	JSONSchemaFilePath = strings.ReplaceAll(JSONSchemaFilePath, "\\", "/")

	schemaLoader := gojsonschema.NewReferenceLoader("file://" + string(JSONSchemaFilePath))
	apiJSONLoader := gojsonschema.NewBytesLoader(JSONResponse)

	result, err := gojsonschema.Validate(schemaLoader, apiJSONLoader)
	if !result.Valid() || err != nil {
		return result.Errors()
	}

	return nil
}
