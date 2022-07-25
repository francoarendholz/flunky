package base

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

func PostScriptRequest(apiPath string, data string) {

	apiKey := viper.GetString("jenkins_api_key")
	apiUser := viper.GetString("jenkins_api_user")
	apiURL := viper.GetString("jenkins_api_url")

	params := url.Values{}
	params.Add("script", data)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", apiURL+apiPath, body)
	if err != nil {

	}
	req.Header.Add("content-type", "text/plain")
	req.Header.Add("user-agent", "Flunky Jenkins Toolkit Version: "+FlunkyVersion+"-"+FlunkyCommit)
	req.SetBasicAuth(apiUser, apiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		println(bodyString)
	} else {
		println("Response was not 200 OK - printing debug information...\n")
		respDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("RESPONSE:\n%s", string(respDump))
	}
}
