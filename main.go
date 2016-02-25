package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	buildCommit string
)

func main() {
	fmt.Printf("Drone Gitter Plugin built from %s\n", buildCommit)

	system := drone.System{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("system", &system)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	payload, err := json.Marshal(&Payload{
		message(&system, &repo, &build),
		icon(&build),
		level(&build),
	})

	if err != nil {
		fmt.Println("Failed to generate payload")

		os.Exit(1)
		return
	}

	for _, url := range vargs.Webhook.Slice() {
		resp, err := http.Post(
			url,
			"application/json",
			bytes.NewBuffer(payload))

		if err != nil {
			fmt.Println("Failed to submit payload")

			os.Exit(1)
			return
		}

		resp.Body.Close()
	}
}

func icon(build *drone.Build) string {
	switch build.Status {
	case drone.StatusSuccess:
		return "smile"
	default:
		return "frown"
	}
}

func level(build *drone.Build) string {
	switch build.Status {
	case drone.StatusSuccess:
		return "normal"
	default:
		return "error"
	}
}

func message(system *drone.System, repo *drone.Repo, build *drone.Build) string {
	switch build.Event {
	case drone.EventPull:
		return pr(system, repo, build)
	case drone.EventDeploy:
		return deploy(system, repo, build)
	case drone.EventTag:
		return tag(system, repo, build)
	default:
		return push(system, repo, build)
	}
}

func push(system *drone.System, repo *drone.Repo, build *drone.Build) string {
	return fmt.Sprintf(
		"Drone [%s](%s) (%s) [%s](%s/%s/%d) (%s)",
		repo.FullName,
		build.Link,
		build.Branch,
		build.Status,
		system.Link,
		repo.FullName,
		build.Number,
		build.Commit[:7],
	)
}

func pr(system *drone.System, repo *drone.Repo, build *drone.Build) string {
	return fmt.Sprintf(
		"Drone [%s](%s) (pull request \\#%s) [%s](%s/%s/%d)",
		repo.FullName,
		build.Link,
		extractNumber(build.Ref),
		build.Status,
		system.Link,
		repo.FullName,
		build.Number,
	)
}

func tag(system *drone.System, repo *drone.Repo, build *drone.Build) string {
	return fmt.Sprintf(
		"Drone [%s](%s) (tag %s) [%s](%s/%s/%d)",
		repo.FullName,
		build.Link,
		filepath.Base(build.Ref),
		build.Status,
		system.Link,
		repo.FullName,
		build.Number,
	)
}

func deploy(system *drone.System, repo *drone.Repo, build *drone.Build) string {
	return fmt.Sprintf(
		"Drone [%s](%s) (deploy to %s) [%s](%s/%s/%d)",
		repo.FullName,
		build.Link,
		build.Deploy,
		build.Status,
		system.Link,
		repo.FullName,
		build.Number,
	)
}

func extractNumber(ref string) string {
	return regexp.MustCompile("([0-9]+)").FindString(ref)
}
