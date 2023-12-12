package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sdl/models"
)

func main() {
	url := "https://dev.azure.com/SecurityRaccoon/ProjectIcarus/_apis/git/repositories?api-version=7.2-preview.1"

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(":PAT")))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	repository := models.Repository
	json.NewDecoder(res.Body).Decode(&repository)

	fmt.Println(string(body[:]))

	fmt.Printf("status code: %d\n", res.StatusCode)

	fmt.Println(repository)
}
