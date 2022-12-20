package base

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestPostScriptRequest(t *testing.T) {

	// Set the expected script
	script := "test_script"

	// Create the expected headers
	expectedHeaders := make(map[string]string)
	expectedHeaders["Content-Type"] = "application/x-www-form-urlencoded"
	expectedHeaders["User-Agent"] = "Flunky Jenkins Toolkit Version: " + FlunkyVersion + "-" + FlunkyCommit
	expectedHeaders["Authorization"] = "Basic dGVzdF9hcGlfdXNlcjp0ZXN0X2FwaV9rZXk="

	// Setup the mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Print the headers and check they are correct as in the expectedHeaders map
		fmt.Println("Headers:")
		for headerName, headerValue := range r.Header {
			if _, ok := expectedHeaders[headerName]; ok {
				assert.Equal(t, expectedHeaders[headerName], headerValue[0])
				fmt.Println(headerName, headerValue)
			} else {
				continue
			}
		}

		// Print the body
		fmt.Println("Body:")
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		// assert the body is correct
		assert.Equal(t, "script="+script, bodyString)

		// Return a 200 OK response
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// Set the API URL to the mock server
	viper.Set("jenkins_api_url", ts.URL)

	// Set the API key and user to something
	viper.Set("jenkins_api_key", "test_api_key")
	viper.Set("jenkins_api_user", "test_api_user")

	// Call the function we are testing
	PostScriptRequest("/scriptText", script)
}
