package models

import (
	"fmt"
	"sdl/helper"
)

type Projects struct {
	Value []Project `json:"value"`
	Count int32     `json:"count"`
}

type Project struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	State       string `json:"state"`
}

func InitProjects() *Projects {
	projects := Projects{}
	path := "projects"

	err := helper.GetJson("", path, "", &projects, false)
	if err != nil {
		fmt.Println(err)
	}

	return &projects

}
