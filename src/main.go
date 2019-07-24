package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	var flags []cli.Flag
	for _, a := range args {
		flags = append(flags, a...)
	}

	app := cli.NewApp()
	app.Name = "coreupdate plugin"
	app.Usage = "coreupdate plugin"
	app.Action = run
	app.Version = version + "+" + build
	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			API:    c.String("repo.api"),
			Header: c.String("repo.header"),
			Name:   c.String("repo.name"),
			Owner:  c.String("repo.owner"),
			Token:  c.String("repo.token"),
		},
		Build: Build{
			Tag:      c.String("build.tag"),
			Number:   c.Int("build.number"),
			Event:    c.String("build.event"),
			Status:   c.String("build.status"),
			Commit:   c.String("commit.sha"),
			Ref:      c.String("commit.ref"),
			Branch:   c.String("commit.branch"),
			Author:   c.String("commit.author"),
			Pull:     c.String("commit.pull"),
			Message:  c.String("commit.message"),
			DeployTo: c.String("build.deployTo"),
			Link:     c.String("build.link"),
			Started:  c.Int64("build.started"),
			Created:  c.Int64("build.created"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Config: Config{
			AppID:   c.String("app.id"),
			Key:     c.String("key"),
			User:    c.String("user"),
			Server:  c.String("server"),
			File:    c.String("pkg.file"),
			Src:     c.String("pkg.src"),
			Channel: c.String("channel"),
			Publish: c.String("publish"),
		},
	}
	return plugin.Exec()
}
