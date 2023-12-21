package main

import (
	"fmt"
	"log"
	"sdl/models"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal(err)
	}

	projects := models.InitProjects()
	fmt.Println(projects.Value[0].Name)

	for _, value := range projects.Value {
		repositories := models.InitRepos(value.Name)

		fmt.Printf("In project %s found %d repositories\n", value.Name, len(repositories.Value))
	}

	// fmt.Println(repositories.Value[0].Name)

	// default_branch, err := repositories.GetDefaultBranch("ProjectIcarus")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Project icarus repo default branch: %s\n", default_branch)

	models.InitPolicies()
	// fmt.Println(policies)

	pipelines := models.InitPipelines()
	fmt.Println(pipelines.Value[0].Name)

	pipeline := pipelines.Value[0].GetPipeline()
	fmt.Println(pipeline.Configuration.Path)
}
