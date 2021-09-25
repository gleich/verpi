package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/verpi/pkg/conf"
)

// Get a list of the last 10 deployments
func ProjectDeployments(
	log lumber.Logger,
	username string,
	config conf.Conf,
	client *http.Client,
) ([]string, error) {
	log.Info("Getting deployments")
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
			if deployment.Creator.Username == username && len(deployments) <= 8 {
				deployments = append(deployments, deployment.ReadyState)
				break
			}
		}
	}

	log.Success("Got data for", len(deployments), "deployments")

	return deployments, nil
}
