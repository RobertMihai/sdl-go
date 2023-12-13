package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Repositories struct {
	Value []Repository `json:"value"`
	Count int32        `json:"count"`
}

type Repository struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Project struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		Url            string `json:"url"`
		State          string `json:"state"`
		Revision       int32  `json:"revision"`
		Visibility     string `json:"visibility"`
		LastUpdateTime string `json:"lastUpdateTime"`
	}
	DefaultBranch   string `json:"defaultBranch"`
	Size            int64  `json:"size"`
	RemoteUrl       string `json:"remoteUrl"`
	SshUrl          string `json:"sshUrl"`
	WebUrl          string `json:"webUrl"`
	IsDisabled      bool   `json:"isDisabled"`
	IsInMaintenance bool   `json:"isInMaintenance"`
}

func InitRepos() *Repositories {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories?api-version=%s", os.Getenv("ORG"), os.Getenv("PROJECT"), os.Getenv("APIVERSION"))
	repos := Repositories{}

	err := GetJson(url, &repos)
	if err != nil {
		fmt.Println(err)
	}

	return &repos
}

func GetJson(url string, target interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(":"+os.Getenv("PAT"))))
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&target)
}
