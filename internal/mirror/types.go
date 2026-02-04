package mirror

type Repo struct {
	Name     string `json:"name"`
	CloneURL string `json:"clone_url"`
	Fork     bool   `json:"fork"`
	Archived bool   `json:"archived"`
}

type Config struct {
	GitHubUsername string
	GitHubToken    string
	RemoteURL      string
	RemoteUser     string
	RemoteToken    string
	Push           bool
}
