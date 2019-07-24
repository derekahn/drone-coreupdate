package main

import "github.com/urfave/cli"

// These are the core settings and requirements for the plugin to run

// Config is configuration settings by user
type Config struct {
	User    string
	Key     string
	Server  string
	AppID   string
	File    string
	Src     string
	Channel string
	Publish string
}

// formatFile creates a file name like "awesome.2.0.0.tar"
func (c Config) formatFile(version string) string {
	return c.File + "." + version + ".tar"
}

func (c Config) createPkgCMD(version, file string) []string {
	return []string{
		"package",
		"create",
		"--app-id=" + c.AppID,
		"--version=" + version,
		"--file=" + file,
		"--url=" + c.Server + "/packages/" + file,
	}
}

func (c Config) uploadPkgCMD(file string) []string {
	return []string{
		"package",
		"upload",
		"--file=" + file,
	}
}

func (c Config) updateChanCMD(version string) []string {
	return []string{
		"channel",
		"update",
		"--app-id=" + c.AppID,
		"--channel=" + c.Channel,
		"--version=" + version,
		"--publish=" + c.Publish,
	}
}

func (c Config) credFlags() []string {
	return []string{
		"--key=" + c.Key,
		"--user=" + c.User,
		"--server=" + c.Server,
	}
}

var configArgs = []cli.Flag{
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
	cli.StringFlag{
		Name:   "channel",
		Usage:  "updateservicectl channel",
		EnvVar: "CHANNEL,PLUGIN_CHANNEL",
	},
	cli.StringFlag{
		Name:   "publish",
		Usage:  "updateservicectl channel --publish",
		EnvVar: "PUBLISH,PLUGIN_PUBLISH",
	},
}
