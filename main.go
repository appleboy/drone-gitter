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

type Vargs struct {
	Webhook drone.StringSlice `json:"webhook"`
}

func main() {
	var (
		repo  = new(drone.Repo)
		build = new(drone.Build)
		sys   = new(drone.System)
		vargs = new(Vargs)
	)

	plugin.Param("repo", repo)
	plugin.Param("build", build)
	plugin.Param("system", sys)
	plugin.Param("vargs", vargs)
	plugin.Parse()

	// gitter data structure
	// old formats https://github.com/gitterHQ/services/blob/master/lib/drone/examples/enterprise_github_commit_success.json
	data := struct {
		Message string `json:"message"`
		Icon    string `json:"icon"`
		Level   string `json:"errorLevel"`
	}{
		message(repo, build, sys),
		icon(build),
		errorLevel(build),
	}

	// json payload that will be posted
	payload, err := json.Marshal(&data)
	if err != nil {
		os.Exit(1)
	}

	// post payload to each url
	for _, url := range vargs.Webhook.Slice() {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			os.Exit(1)
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

func errorLevel(build *drone.Build) string {
	switch build.Status {
	case drone.StatusSuccess:
		return "normal"
	default:
		return "error"
	}
}

// helper function that extracts the pull request
// number from the merge ref.
func extractNumber(ref string) string {
	return regexp.MustCompile("([0-9]+)").FindString(ref)
}

// helper function that returns a message based on
// the build event type.
func message(repo *drone.Repo, build *drone.Build, sys *drone.System) string {
	switch build.Event {
	case drone.EventPull:
		return pullRequest(repo, build, sys)
	case drone.EventDeploy:
		return deploy(repo, build, sys)
	case drone.EventTag:
		return tag(repo, build, sys)
	default:
		return push(repo, build, sys)
	}
}

func push(repo *drone.Repo, build *drone.Build, sys *drone.System) string {
	return fmt.Sprintf(
		"Drone [%s](%s) (%s) [%s](%s/%s/%d) (%s)",
		repo.FullName,
		build.Link,
		build.Branch,
		build.Status,
		sys.Link,
		repo.FullName,
		build.Number,
		build.Commit[:7],
	)
}

func pullRequest(repo *drone.Repo, build *drone.Build, sys *drone.System) string {
	return fmt.Sprintf(
		"Drone [%s](%s) (pull request \\#%s) [%s](%s/%s/%d)",
		repo.FullName,
		build.Link,
		extractNumber(build.Ref),
		build.Status,
		sys.Link,
		repo.FullName,
		build.Number,
	)
}

func tag(repo *drone.Repo, build *drone.Build, sys *drone.System) string {
	return fmt.Sprintf(
		"Drone [%s](%s) (tag %s) [%s](%s/%s/%d)",
		repo.FullName,
		build.Link,
		filepath.Base(build.Ref),
		build.Status,
		sys.Link,
		repo.FullName,
		build.Number,
	)
}

func deploy(repo *drone.Repo, build *drone.Build, sys *drone.System) string {
	return fmt.Sprintf(
		"Drone [%s](%s) (deploy to %s) [%s](%s/%s/%d)",
		repo.FullName,
		build.Link,
		build.Deploy,
		build.Status,
		sys.Link,
		repo.FullName,
		build.Number,
	)
}
