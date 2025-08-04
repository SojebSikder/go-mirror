package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sojebsikder/go-mirror/mirror"
)

func FetchRepos(username, token string) ([]mirror.Repo, error) {
	var allRepos []mirror.Repo
	page := 1

	for {
		url := fmt.Sprintf("https://api.github.com/user/repos?per_page=100&page=%d&affiliation=owner", page)
		req, _ := http.NewRequest("GET", url, nil)
		req.SetBasicAuth(username, token)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var repos []mirror.Repo
		if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
			return nil, err
		}
		if len(repos) == 0 {
			break
		}

		allRepos = append(allRepos, repos...)
		page++
	}

	return allRepos, nil
}
