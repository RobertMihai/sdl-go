package models

type Repository struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Project struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		Url            string `json:"url"`
		State          string `json:"state"`
		Revision       string `json:"revision"`
		Visibility     string `json:"visibility"`
		LastUpdateTime string `json:"lastUpdateTime"`
	}
	DefaultBranch   string `json:"defaultBranch"`
	Size            int64  `json:"size"`
	RemoteUrl       string `json:"remoteUrl"`
	SshUrl          string `json:"sshUrl"`
	WebUrl          string `json:"webUrl"`
	IsDisabled      string `json:"isDisabled"`
	IsInMaintenance string `json:"isInMaintenance"`
}
