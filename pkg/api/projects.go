package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/verpi/pkg/conf"
)

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
