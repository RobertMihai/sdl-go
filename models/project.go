package models

import (
	"log/slog"
	"sdl/helper"

	"gorm.io/gorm"
)

type Projects struct {
	Value []Project `json:"value"`
	Count int32     `json:"count"`
}

type Project struct {
	gorm.Model
	ID             string `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	State          string `json:"state"`
	Revision       int    `json:"revision"`
	Visibility     string `json:"visibility"`
	LastUpdateTime string `json:"lastUpdateTime"`
}

var db *gorm.DB
var logger *slog.Logger

func InitProjects() {
	logger = helper.GetLogger()
	db = helper.GetDB()
	db.AutoMigrate(&Project{})

	projects := Projects{}
	path := "projects"

	err := helper.GetJson("", path, "", &projects, false)
	if err != nil {
		logger.Error(err.Error())
	}

	for _, project := range projects.Value {
		project.StoreProject()
		// InitRepos(project.Name)
		InitPolicies(project.Name, "", "")
	}

}

func (p *Project) StoreProject() *Project {
	db.Create(&p)

	// logger.Error(stuff.Error.Error())

	return p
}
