package helper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var url = "https://dev.azure.com/"
var args = "?api-version="
var logger = GetLogger()

func GetJson(project string, path string, extra_args string, target interface{}, debug bool) error {
	full_url := ""
	if extra_args == "" {
		full_url = url + os.Getenv("ORG") + "/" + project + "/_apis/" + path + args + os.Getenv("APIVERSION")
	} else {
		full_url = url + os.Getenv("ORG") + "/" + project + "/_apis/" + path + args + os.Getenv("APIVERSION") + "&" + extra_args
	}

	// fmt.Println("full url for request: " + full_url)
	// logger.Info("full url for request: " + full_url)

	req, err := http.NewRequest(http.MethodGet, full_url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(":"+os.Getenv("PAT"))))
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if debug {
		value, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Info(err.Error())
		}
		fmt.Println(string(value))
		// logger.Info(string(value))
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&target)
}
