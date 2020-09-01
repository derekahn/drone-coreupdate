package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

// fetchTag 'GET' gitlab or github api for latest tag
func (p Plugin) fetchTag() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", p.Repo.API, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set(p.Repo.Header, p.Repo.Token)

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

	// Using '/tags?order_by=updated' query param we
	// need to take the first element
	latest, _ := tags[0], tags[1:]
	return latest.Version, nil
}

// fetchSHA 'GET' quay.io repo's manifest_digest sha256
func (p Plugin) fetchSHA(version string) (string, error) {
	url := p.Quay.API + version

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+p.Quay.Token)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	type (
		Tag struct {
			SHA string `json:"manifest_digest"`
		}
		body struct {
			Tags []Tag `json:"tags"`
		}
	)

	var b body
	if err := json.NewDecoder(res.Body).Decode(&b); err != nil {
		return "", nil
	}
	if len(b.Tags) < 1 {
		return "", errors.New("No tags available")
	}
	return b.Tags[0].SHA, nil
}
