package main

import (
	"net/http"
	"time"

	"github.com/gleich/lumber"
	"github.com/gleich/verpi/pkg/api"
	"github.com/gleich/verpi/pkg/conf"
	"github.com/gleich/verpi/pkg/lights"
)

func main() {
	config, err := conf.Read()
	if err != nil {
		lumber.Fatal(err, "Failed to read from configuration file")
	}
	client := http.DefaultClient
	username, err := api.Username(config, client)
	if err != nil {
		lumber.Fatal(err, "Failed to get vercel username")
	}
	display := lights.Setup(config)

	for {
		// Rereading from configuration file to load any new changes
		config, err := conf.Read()
		if err != nil {
			lumber.Fatal(err, "Failed to read from configuration file")
		}

		if *config.Brightness == 0.0 {
			lumber.Info("Not updating lights because brightness set to 0")
			display.Clear()
			time.Sleep(20 * time.Millisecond)
			continue
		}

		deployments, err := api.ProjectDeployments(username, config, client)
		if err != nil {
			lumber.Fatal(err, "Failed to get deployments")
		}

		lights.Update(config, deployments, display)

		time.Sleep(4 * time.Second)
	}
}
