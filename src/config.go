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
