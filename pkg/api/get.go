package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gleich/lumber"
	"github.com/gleich/verpi/pkg/conf"
)

type Deployment struct {
	State string
}

// Get a list of the last 10 deployments
func Deployments(config conf.Conf, client *http.Client) ([]Deployment, error) {
	lumber.Info("Getting deployments")
	// Making request
	req, err := http.NewRequest("GET", "https://api.vercel.com/v5/now/deployments?limit=8", nil)
	if err != nil {
		return []Deployment{}, err
	}
	req.Header.Add("Authorization", "Bearer "+config.Token)

	resp, err := client.Do(req)
	if err != nil {
		return []Deployment{}, err
	}
	defer resp.Body.Close()

	// Reading response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Deployment{}, err
	}

	// Parsing response
	var data struct{ Deployments []Deployment }
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []Deployment{}, err
	}

	lumber.Success("Got data for", len(data.Deployments), "deployments")
	return data.Deployments, nil
}
