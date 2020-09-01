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
	Quay   Quay
}

// Exec is the entrypoint
func (p Plugin) Exec() error {
	version, err := p.fetchTag()
	if err != nil {
		return err
	}

	dir := p.Config.Src
	file := p.Config.formatFile(version)

	if p.Quay.API != "" {
		sha, err := p.fetchSHA(version)
		if err != nil {
			return err
		}
		if err := findAndReplace(dir, "${SHA}", sha); err != nil {
			return err
		}
	}

	if err := findAndReplace(dir, "${VERSION}", version); err != nil {
		return err
	}

	if err := archive(dir, file); err != nil {
		return err
	}

	createPkg := p.Config.createPkgCMD(version, file)
	output, err := p.updateservicectl(createPkg)
	if err != nil {
		return err
	}
	log.Println(string(output))

	uploadPkg := p.Config.uploadPkgCMD(file)
	output, err = p.updateservicectl(uploadPkg)
	if err != nil {
		return err
	}
	log.Println(string(output))

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
