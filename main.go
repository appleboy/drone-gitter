package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	_ "github.com/joho/godotenv/autoload"
)

var version string // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "gitter"
	app.Usage = "gitter plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "webhook",
			Usage:  "gitter webhook url",
			EnvVar: "GITTER_WEBHOOK",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.link",
			Usage:  "git commit link",
			EnvVar: "DRONE_COMMIT_LINK",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.deploy",
			Usage:  "build deployment target",
			EnvVar: "DRONE_DEPLOY_TO",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Commit: Commit{
			Sha:    c.String("commit.sha"),
			Ref:    c.String("commit.ref"),
			Branch: c.String("commit.branch"),
			Author: c.String("commit.author"),
			Link:   c.String("commit.link"),
		},
		Build: Build{
			Number: c.Int("build.number"),
			Event:  c.String("build.event"),
			Status: c.String("build.status"),
			Deploy: c.String("build.deploy"),
			Link:   c.String("build.link"),
		},
		Config: Config{
			Webhook: c.String("webhook"),
		},
	}

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
