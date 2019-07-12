package main

import (
	"log"
	"os/exec"

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
	err := p.archive()
	if err != nil {
		return err
	}

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

func (p Plugin) archive() error {
	dir := p.Config.Src

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	var payload []string
	for _, f := range files {
		payload = append(payload, dir+"/"+f.Name())
	}

	return archiver.Archive(payload, p.fileName())
}

func (p Plugin) fileName() string {
	return p.Config.File + "." + p.Config.Version + ".tar"
}

func (p Plugin) createPackage() ([]byte, error) {
	action := []string{
		"package",
		"create",
		"--app-id=" + p.Config.AppID,
		"--version=" + p.Config.Version,
		"--file=" + p.fileName(),
		"--url=" + p.Config.URL + "." + p.fileName(),
	}

	cmd, args := p.baseCMD()
	args = append(args, action...)

	return exec.Command(cmd, args...).Output()
}

func (p Plugin) uploadPackage() ([]byte, error) {
	action := []string{
		"package",
		"upload",
		"--file=" + p.fileName(),
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
