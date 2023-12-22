package main

import (
	"log/slog"
	"sdl/helper"
	"sdl/models"

	"github.com/joho/godotenv"
)

var logger *slog.Logger

func Init() {
	helper.Connect()

	helper.InitLogger()
	logger = helper.GetLogger()
}

func main() {
	err := godotenv.Load(".env.local")

	if err != nil {
		logger.Info(err.Error())
	}

	Init()

	// projects := models.InitProjects(db)
	models.InitProjects()
	// fmt.Println(projects.Value[0].Name)

	// for _, value := range projects.Value {
	// 	repositories := models.InitRepos(value.Name)

	// 	fmt.Printf("In project %s found %d repositories\n", value.Name, len(repositories.Value))
	// }

	// fmt.Println(repositories.Value[0].Name)

	// default_branch, err := repositories.GetDefaultBranch("ProjectIcarus")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Project icarus repo default branch: %s\n", default_branch)

	// policies := models.InitPolicies("ATR", "", "")
	// fmt.Println(policies.Value[0].ID)

	// // fmt.Println(policies)

	// pipelines := models.InitPipelines()
	// fmt.Println(pipelines.Value[0].Name)

	// pipeline := pipelines.Value[0].GetPipeline()
	// fmt.Println(pipeline.Configuration.Path)
}
