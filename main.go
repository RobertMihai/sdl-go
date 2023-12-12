package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/RobertMihai/sdl-go/models"
	"github.com/joho/godotenv"
)

var repository = *models.NewRepository()

func main() {

	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal(err)
	}

	url := "https://dev.azure.com/SecurityRaccoon/ProjectIcarus/_apis/git/repositories?api-version=7.2-preview.1"

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	pat := os.Getenv("PAT")

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(":"+pat)))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(res.Body).Decode(&repository)

	fmt.Println(string(body[:]))

	fmt.Printf("status code: %d\n", res.StatusCode)

	fmt.Println(repository.ToString())
}
