package main

import (
	"fmt"
	"log"
	"os/exec"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mholt/archiver"
)

// Plugin TODO
type Plugin struct {
	Repo   Repo
	Build  Build
	Config Config
	Job    Job
}

// Exec plugin (entry)
func (p Plugin) Exec() error {
	output, err := p.createPackage()
	if err != nil {
		return err
	}
	log.Println(string(output))

	output, err = p.uploadPackage()
	if err != nil {
		return err
	}
	log.Println(string(output))

	return nil
}

func (p Plugin) createPackage() ([]byte, error) {
	action := []string{
		"package",
		"create",
		"--app-id=" + p.Config.AppID,
		"--version=" + p.Config.Version,
		fmt.Sprintf("--file=%s.%s.tar", p.Config.File, p.Config.Version),
		fmt.Sprintf("--url=%s/%s.%s.tar",
			p.Config.URL,
			p.Config.File,
			p.Config.Version,
		),
	}

	cmd, args := p.baseCMD()
	args = append(args, action...)

	return exec.Command(cmd, args...).Output()
}

func (p Plugin) uploadPackage() ([]byte, error) {
	action := []string{
		"package",
		"upload",
		fmt.Sprintf("--file=%s.%s.tar", p.Config.File, p.Config.Version),
	}

	cmd, args := p.baseCMD()
	args = append(args, action...)

	return exec.Command(cmd, args...).Output()
}

func (p Plugin) baseCMD() (string, []string) {
	flags := []string{
		"--key=" + p.Config.Key,
		"--user=" + p.Config.User,
		"--server=" + p.Config.Server,
	}
	return "updateservicectl", flags
}
