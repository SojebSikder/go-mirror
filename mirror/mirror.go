package mirror

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CloneAndPush(repo Repo, cfg Config) error {
	name := repo.Name
	fmt.Println("Cloning:", name)

	cloneCmd := exec.Command("git", "clone", "--mirror", repo.CloneURL, name+".git")
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	if err := cloneCmd.Run(); err != nil {
		return fmt.Errorf("clone failed: %w", err)
	}

	pushURL := fmt.Sprintf("https://%s:%s@%s/%s/%s.git",
		cfg.RemoteUser,
		cfg.RemoteToken,
		strings.TrimPrefix(cfg.RemoteURL, "https://"),
		cfg.RemoteUser,
		name,
	)
	fmt.Println("Pushing to:", strings.Replace(pushURL, cfg.RemoteToken, "********", 1))

	pushCmd := exec.Command("git", "push", "--mirror", pushURL)
	pushCmd.Dir = name + ".git"
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	if err := pushCmd.Run(); err != nil {
		return fmt.Errorf("push failed: %w", err)
	}

	return nil
}
