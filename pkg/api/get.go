package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gleich/lumber"
	"github.com/gleich/verpi/pkg/conf"
)

// Base URL for the vercel api
const baseURL = "https://api.vercel.com"

// Get a list of the last 10 deployments
func ProjectDeployments(username string, config conf.Conf, client *http.Client) ([]string, error) {
	lumber.Info("Getting deployments")
	// Making request
	req, err := http.NewRequest("GET", baseURL+"/v8/projects", nil)
	if err != nil {
		return []string{}, err
	}
	req.Header.Add("Authorization", "Bearer "+config.Token)

	resp, err := client.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	// Reading response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}

	// Parsing response
	var data struct {
		Projects []struct {
			LatestDeployments []struct {
				Creator struct {
					Username string
				}
				ReadyState string
			}
		}
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []string{}, err
	}

	// Filtering and cleaning response to only include deployments from the user
	deployments := []string{}
	for _, project := range data.Projects {
		for _, deployment := range project.LatestDeployments {
			if deployment.Creator.Username == username {
				deployments = append(deployments, deployment.ReadyState)
				break
			}
		}
	}

	lumber.Success("Got data for", len(deployments), "deployments")
	return deployments, nil
}

// Get the user's slack ID so we can filter
func Username(config conf.Conf, client *http.Client) (string, error) {
	lumber.Info("Getting the user's vercel username")
	// Making request
	req, err := http.NewRequest("GET", baseURL+"/www/user", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+config.Token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Reading response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parsing response
	var data struct {
		User struct {
			Username string
		}
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	if data.User.Username == "" {
		return "", errors.New("No username found with access token")
	}

	lumber.Success("Got", data.User.Username+"'s", "vercel username")
	return data.User.Username, nil
}
