package cmd

import (
	"log"
	"sync"

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

	const maxWorkers = 5 // tune based on CPU/network
	repoCh := make(chan mirror.Repo)
	var wg sync.WaitGroup

	// workers
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for repo := range repoCh {
				if err := mirror.CloneAndPush(repo, config, config.Push); err != nil {
					log.Printf("Error processing repo %s: %v", repo.Name, err)
				}
			}
		}()
	}

	// producer
	for _, repo := range repos {
		if repo.Fork || repo.Archived {
			continue
		}
		repoCh <- repo
	}

	close(repoCh)
	wg.Wait()
}

// func Mirror() {
// 	config, err := utils.LoadEnv()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	repos, err := github.FetchRepos(config.GitHubUsername, config.GitHubToken)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, repo := range repos {
// 		if repo.Fork || repo.Archived {
// 			continue
// 		}

// 		if err := mirror.CloneAndPush(repo, config); err != nil {
// 			log.Printf("Error processing repo %s: %v", repo.Name, err)
// 		}
// 	}
// }
