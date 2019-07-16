package main

import (
	"log"
	"os/exec"
)

// Plugin contains all (master)
type Plugin struct {
	Repo   Repo
	Build  Build
	Config Config
	Job    Job
}

// Exec is the entrypoint
func (p Plugin) Exec() error {
	version, err := fetchTag(p.Repo.API, p.Repo.Header, p.Repo.Token)
	if err != nil {
		return err
	}

	dir := p.Config.Src
	file := p.Config.File + "." + version + ".tar"

	err = findAndReplace(dir, "${VERSION}", version)
	if err != nil {
		return err
	}

	err = archive(dir, file)
	if err != nil {
		return err
	}

	output, err := p.createPackage(file, version)
	if err != nil {
		return err
	}
	log.Println(string(output))

	output, err = p.uploadPackage(file)
	if err != nil {
		return err
	}
	log.Println(string(output))

	return nil
}

// runs '$ updateservicectl package create'
func (p Plugin) createPackage(file, version string) ([]byte, error) {
	action := []string{
		"package",
		"create",
		"--app-id=" + p.Config.AppID,
		"--version=" + version,
		"--file=" + file,
		"--url=" + p.Config.Server + "/packages/" + file,
	}

	cmd, args := p.baseCMD()
	args = append(args, action...)
	return exec.Command(cmd, args...).Output()
}

// runs '$ updateservicectl package upload'
func (p Plugin) uploadPackage(file string) ([]byte, error) {
	action := []string{
		"package",
		"upload",
		"--file=" + file,
	}

	cmd, args := p.baseCMD()
	args = append(args, action...)
	return exec.Command(cmd, args...).Output()
}

// runs 'updateservicectl' with the required flags
func (p Plugin) baseCMD() (string, []string) {
	flags := []string{
		"--key=" + p.Config.Key,
		"--user=" + p.Config.User,
		"--server=" + p.Config.Server,
	}
	return "updateservicectl", flags
}
