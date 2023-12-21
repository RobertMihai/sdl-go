package models

import (
	"fmt"
	"os"
	"sdl/helper"
)

type Policies struct {
	Value []Policy `json:"value"`
	Count int32    `json:"count"`
}

type Policy struct {
}

func InitPolicies() *Policies {
	policies := Policies{}

	project := os.Getenv("PROJECT")
	path := "git/policy/configurations"

	err := helper.GetJson(project, path, "", &policies, true)
	if err != nil {
		fmt.Println(err)
	}

	return &policies
}
