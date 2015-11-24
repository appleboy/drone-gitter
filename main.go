package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

func main() {
	var repo = drone.Repo{}
	var build = drone.Build{}
	var sys = drone.System{}
	var vargs = struct {
		Webhook drone.StringSlice `json:"webhook"`
	}{}

	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("system", &sys)
	plugin.Param("vargs", &vargs)
	plugin.Parse()

	// data structure
	// old formats https://github.com/gitterHQ/services/blob/master/lib/drone/examples/enterprise_github_commit_success.json
	data := struct {
		Repo  drone.Repo    `json:"repo"`
		Build drone.Build   `json:"build"`
		System drone.System `json:"system"`
	}{repo, build, sys}

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
