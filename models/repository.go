package models

import (
	"fmt"
	"sdl/helper"

	"gorm.io/gorm"
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
	Type            string `json:"type"`
}

var db *gorm.DB

func init() {
	helper.Connect()
	db = helper.GetDB()
	db.AutoMigrate(&Repository{})
}

func InitRepos(project string) *Repositories {
	repos := Repositories{}
	// project := os.Getenv("PROJECT")
	path := "git/repositories"

	err := helper.GetJson(project, path, "", &repos, false)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range repos.Value {
		value.StoreRepo()
	}

	return &repos
}

func (r *Repositories) GetDefaultBranch(repo_name string) (string, error) {
	for _, repo := range r.Value {
		if repo_name == repo.Name {
			return repo.DefaultBranch, nil
		}
	}

	return "", fmt.Errorf("repo %s not found", repo_name)
}

func (r *Repository) StoreRepo() *Repository {
	db.Create(&r)

	return r
}

func GetAllRepositories() []Repository {
	var repositories []Repository
	db.Find(&repositories)

	return repositories
}
