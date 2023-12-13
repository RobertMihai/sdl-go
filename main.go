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

	repositories := models.InitRepos()
	fmt.Println(repositories.Value[0].Name)
}
