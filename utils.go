package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

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

// findAndReplace handles interpolation "${VERSION}" -> "1.0.1"
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

// fetchTag 'GET' git[lab/hub] api for latest tag
func fetchTag(url, header, token string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set(header, token)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Both github and gitlab reponses have property "name"
	type tag struct {
		Version string `json:"name"`
	}

	var tags []tag
	err = json.NewDecoder(res.Body).Decode(&tags)
	if err != nil {
		return "", err
	}
	if len(tags) < 1 {
		return "", errors.New("No tags available")
	}

	latest, _ := tags[len(tags)-1], tags[:len(tags)-1]
	return latest.Version, nil
}
