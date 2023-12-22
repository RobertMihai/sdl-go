// Retrieve a list of policy configurations by a given set of scope/filtering criteria.
// Below is a short description of how all of the query parameters interact with each other:
// repositoryId set, refName set: returns all policy configurations that apply to a particular branch in a repository
// repositoryId set, refName unset: returns all policy configurations that apply to a particular repository
// repositoryId unset, refName unset: returns all policy configurations that are defined at the project level
// repositoryId unset, refName set: returns all project-level branch policies, plus the project level configurations For all of the examples above, when policyType is set, it'll restrict results to the given policy type

package models

import (
	"fmt"
	"sdl/helper"

	"github.com/RobertMihai/sdl-go/models"
	"gorm.io/gorm"
)

type Policies struct {
	Value []Policy `json:"value"`
	Count int32    `json:"count"`
}

type Policy struct {
	gorm.Model
	CreatedData         string         `json:"createdDate"`
	IsEnabled           bool           `json:"isEnabled"`
	IsBlocking          bool           `json:"isBlocking"`
	IsDeleted           bool           `json:"isDeleted"`
	Settings            PolicySettings `json:"settings" gorm:"embedded;embeddedPrefix:settings_"`
	IsEnterpriseManaged bool           `json:"isEnterpriseManaged"`
	Revision            int            `json:"revision"`
	ID                  int            `json:"id"`
	Url                 string         `json:"url"`
	Type                PolicyType     `json:"type" gorm:"embedded;embeddedPrefix:type_"`
}

type PolicySettings struct {
	Scopes                         []PolicyScope     `json:"scope" gorm:"-"`
	Scope                          models.Repository `gorm:"foreignKey"`
	EnforceConsistentCase          bool              `json:"enforceConsistentCase"`
	RejectDotGit                   bool              `json:"rejectDotGit"`
	OptimizedByDefault             bool              `json:"optimizedByDefault"`
	BreadcrumbDays                 bool              `json:"breadcrumbDays"`
	AllowedForkTargets             bool              `json:"allowedForkTargets"`
	StrictVoteMode                 bool              `json:"strictVoteMode"`
	InheritPullRequestCreationMode bool              `json:"inheritPullRequestCreationMode"`
	RepoPullRequestDraftByDefault  bool              `json:"repoPullRequestAsDraftByDefault"`
}

type PolicyScope struct {
	RepositoryID string `json:"repositoryId" gorm:"foreignKey"`
}

type PolicyType struct {
	ID          string `json:"id"`
	Url         string `json:"url"`
	DisplayName string `json:"displayName"`
}

func InitPolicies(project string, repositoryId string, refName string) {
	policies := Policies{}

	// project := os.Getenv("PROJECT")
	path := "git/policy/configurations"
	extra_args := ""

	if repositoryId != "" {
		extra_args += fmt.Sprintf("repositoryId=%s&", repositoryId)
	}

	if refName != "" {
		extra_args += fmt.Sprintf("refName=%s&", refName)
	}

	err := helper.GetJson(project, path, extra_args, &policies, false)
	if err != nil {
		fmt.Println(err)
	}

	for _, policy := range policies.Value {
		policy.StorePolicy()
	}
}

func (p *Policy) StorePolicy() {
	if len(p.Settings.Scopes) > 1 {
		logger.Info("Policy has multiple scopes",
			"count", len(p.Settings.Scopes),
			"policyID", p.ID)
	}

	for _, scope := range p.Settings.Scopes {
		p.Settings.Scope = Repository{ID: scope.RepositoryID}
		db.Create(&p)
	}
}
