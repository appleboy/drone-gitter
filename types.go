package main

import (
	"github.com/drone/drone-go/drone"
)

type Params struct {
	Webhook drone.StringSlice `json:"webhook"`
}

type Payload struct {
	Message string `json:"message"`
	Icon    string `json:"icon"`
	Level   string `json:"errorLevel"`
}
