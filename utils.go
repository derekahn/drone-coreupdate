package main

import (
	"bytes"
	"io/ioutil"

	"github.com/mholt/archiver"
)

// creates a tarball of files in dir with name
func archive(dir, name string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	var payload []string
	for _, f := range files {
		payload = append(payload, dir+"/"+f.Name())
	}
	return archiver.Archive(payload, name)
}

// findAndReplace handles interpolation: 'find' -> 'replace'
func findAndReplace(dir, find, replace string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		file := f.Name()
		path := dir + "/" + file

		input, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		output := bytes.Replace(input, []byte(find), []byte(replace), -1)
		if err = ioutil.WriteFile(path, output, 0644); err != nil {
			return err
		}
	}
	return nil
}
