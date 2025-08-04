package cmd

import (
	"log"

	"github.com/sojebsikder/go-mirror/internal/github"
	"github.com/sojebsikder/go-mirror/internal/mirror"
	"github.com/sojebsikder/go-mirror/pkg/utils"
)

func Mirror() {
	config, err := utils.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	repos, err := github.FetchRepos(config.GitHubUsername, config.GitHubToken)
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		if repo.Fork || repo.Archived {
			continue
		}

		if err := mirror.CloneAndPush(repo, config); err != nil {
			log.Printf("Error processing repo %s: %v", repo.Name, err)
		}
	}
}
