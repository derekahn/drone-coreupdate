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

	updateChan := p.Config.updateChanCMD(version)
	output, err = p.updateservicectl(updateChan)
	if err != nil {
		return err
	}
	log.Println(string(output))

	return nil
}

// updateservicectl executes the cli with sub commands and flags
func (p Plugin) updateservicectl(action []string) ([]byte, error) {
	creds := p.Config.credFlags()
	args := append(creds, action...)
	return exec.Command("updateservicectl", args...).Output()
}
