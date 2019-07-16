package main

import "github.com/urfave/cli"

type (
	// Build represents drone events
	Build struct {
		Tag      string
		Event    string
		Number   int
		Commit   string
		Ref      string
		Branch   string
		Author   string
		Pull     string
		Message  string
		DeployTo string
		Status   string
		Link     string
		Started  int64
		Created  int64
	}

	// Config is settings by user
	Config struct {
		AppID  string
		Key    string
		User   string
		Server string
		File   string
		Src    string
	}

	// Job captures drone runtime events
	Job struct {
		Started int64
	}

	// Repo is git details
	Repo struct {
		Owner string
		Name  string
	}

	flags []cli.Flag
)

var (
	args = []flags{
		buildArgs,
		configArgs,
		jobArgs,
		repoArgs,
	}

	buildArgs = []cli.Flag{
		cli.StringFlag{
			Name:   "build.event",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
			Value:  "push",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			EnvVar: "DRONE_BUILD_STATUS",
			Value:  "success",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "build.deployTo",
			Usage:  "environment deployed to",
			EnvVar: "DRONE_DEPLOY_TO",
		},
	}

	commitArgs = []cli.Flag{
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
			Value:  "00000000",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
			Value:  "refs/heads/master",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
			Value:  "master",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.pull",
			Usage:  "git pull request",
			EnvVar: "DRONE_PULL_REQUEST",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
	}

	configArgs = []cli.Flag{
		cli.StringFlag{
			Name:   "app.id",
			Usage:  "updateservicectl --app-id",
			EnvVar: "APP_ID,PLUGIN_APP_ID",
		},
		cli.StringFlag{
			Name:   "key",
			Usage:  "updateservicectl --key",
			EnvVar: "KEY,PLUGIN_KEY",
		},
		cli.StringFlag{
			Name:   "user",
			Usage:  "updateservicectl --user",
			EnvVar: "CTL_USER,PLUGIN_USER",
		},
		cli.StringFlag{
			Name:   "server",
			Usage:  "updateservicectl --server",
			EnvVar: "SERVER,PLUGIN_SERVER",
		},
		cli.StringFlag{
			Name:   "pkg.file",
			Usage:  "updateservicectl package [create || upload] --file",
			EnvVar: "PKG_FILE,PLUGIN_PKG_FILE",
		},
		cli.StringFlag{
			Name:   "pkg.src",
			Usage:  "target directory to tarball",
			EnvVar: "PKG_SRC,PLUGIN_PKG_SRC",
		},
	}

	jobArgs = []cli.Flag{
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
	}

	repoArgs = []cli.Flag{
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
	}
)
