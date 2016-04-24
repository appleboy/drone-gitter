package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"
)

type (
	Repo struct {
		Owner string
		Name  string
	}

	Build struct {
		Event  string
		Number int
		Status string
		Deploy string
		Link   string
	}

	Commit struct {
		Sha    string
		Ref    string
		Branch string
		Author string
		Link   string
	}

	Config struct {
		Webhook string
	}

	Plugin struct {
		Repo   Repo
		Commit Commit
		Build  Build
		Config Config
	}

	Payload struct {
		Message string `json:"message"`
		Icon    string `json:"icon"`
		Level   string `json:"errorLevel"`
	}
)

func (p Plugin) Exec() error {

	in, err := json.Marshal(&Payload{
		message(p.Repo, p.Commit, p.Build),
		icon(p.Build),
		level(p.Build),
	})

	if err != nil {
		return err
	}

	resp, err := http.Post(p.Config.Webhook, "application/json", bytes.NewBuffer(in))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func icon(build Build) string {
	switch build.Status {
	case "success":
		return "smile"
	default:
		return "frown"
	}
}

func level(build Build) string {
	switch build.Status {
	case "success":
		return "normal"
	default:
		return "error"
	}
}

func message(repo Repo, commit Commit, build Build) string {
	switch build.Event {
	case "pull_request":
		return pr(repo, commit, build)
	case "deployment":
		return deploy(repo, commit, build)
	case "tag":
		return tag(repo, commit, build)
	default:
		return push(repo, commit, build)
	}
}

func push(repo Repo, commit Commit, build Build) string {
	return fmt.Sprintf(
		"Drone [%s/%s](%s) (%s) [%s](%s) (%s)",
		repo.Owner,
		repo.Name,
		commit.Link,
		commit.Branch,
		build.Status,
		build.Link,
		commit.Sha[:7],
	)
}

func pr(repo Repo, commit Commit, build Build) string {
	return fmt.Sprintf(
		"Drone [%s/%s](%s) (pull request \\#%s) [%s](%s)",
		repo.Owner,
		repo.Name,
		commit.Link,
		extractNumber(commit.Ref),
		build.Status,
		build.Link,
	)
}

func tag(repo Repo, commit Commit, build Build) string {
	return fmt.Sprintf(
		"Drone [%s/%s](%s) (tag %s) [%s](%s)",
		repo.Owner,
		repo.Name,
		commit.Link,
		filepath.Base(commit.Ref),
		build.Status,
		build.Link,
	)
}

func deploy(repo Repo, commit Commit, build Build) string {
	return fmt.Sprintf(
		"Drone [%s/%s](%s) (deploy to %s) [%s](%s)",
		repo.Owner,
		repo.Name,
		commit.Link,
		build.Deploy,
		build.Status,
		build.Link,
	)
}

func extractNumber(ref string) string {
	return regexp.MustCompile("([0-9]+)").FindString(ref)
}
