package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sojebsikder/go-mirror/mirror"
)

func LoadEnv() (mirror.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return mirror.Config{}, fmt.Errorf("failed to load .env: %w", err)
	}

	config := mirror.Config{
		GitHubUsername: os.Getenv("GITHUB_USERNAME"),
		GitHubToken:    os.Getenv("GITHUB_TOKEN"),
		RemoteURL:      os.Getenv("GIT_REMOTE_URL"),
		RemoteUser:     os.Getenv("GIT_REMOTE_USER"),
		RemoteToken:    os.Getenv("GIT_REMOTE_TOKEN"),
	}

	if config.GitHubUsername == "" || config.GitHubToken == "" || config.RemoteURL == "" {
		return config, fmt.Errorf("missing required environment variables")
	}

	return config, nil
}
