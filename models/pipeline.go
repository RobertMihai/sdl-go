package models

import (
	"fmt"
	"os"
	"sdl/helper"

	"github.com/RobertMihai/sdl-go/models"
)

type Pipelines struct {
	Value []Pipeline `json:"value"`
	Count int32      `json:"count"`
}

type Pipeline struct {
	Url           string                `json:"url"`
	Id            int32                 `json:"id"`
	Revision      int32                 `json:"revision"`
	Name          string                `json:"name"`
	Folder        string                `json:"folder"`
	Configuration PipelineConfiguration `json:"configuration"`
}

type PipelineConfiguration struct {
	Path       string            `json:"path"`
	Repository models.Repository `json:"repository"`
	Type       string            `json:"type"`
}

func InitPipelines() *Pipelines {
	pipelines := Pipelines{}

	project := os.Getenv("PROJECT")
	path := "/pipelines"

	err := helper.GetJson(project, path, "", &pipelines, false)
	if err != nil {
		logger.Info(err.Error())
	}

	return &pipelines
}

func (p *Pipeline) GetPipeline() *Pipeline {
	project := os.Getenv("PROJECT")
	path := fmt.Sprintf("/pipelines/%d", p.Id)

	err := helper.GetJson(project, path, "", &p, false)

	if err != nil {
		logger.Info(err.Error())
	}

	return p
}
