package main

import (
	"net/http"
	"time"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/verpi/pkg/api"
	"github.com/gleich/verpi/pkg/conf"
	"github.com/gleich/verpi/pkg/lights"
)

func main() {
	log := lumber.NewCustomLogger()
	log.Timezone = time.Local
	config, err := conf.Read(log)
	if err != nil {
		lumber.Fatal(err, "Failed to read from configuration file")
	}
	client := http.DefaultClient
	username, err := api.Username(log, config, client)
	if err != nil {
		lumber.Fatal(err, "Failed to get vercel username")
	}
	display := lights.Setup(log, config)

	for {
		// Rereading from configuration file to load any new changes
		config, err := conf.Read(log)
		if err != nil {
			log.Fatal(err, "Failed to read from configuration file")
		}

		if *config.Brightness == 0.0 {
			log.Info("Not updating lights because brightness set to 0")
			display.Clear()
			display.Show()
			time.Sleep(20 * time.Millisecond)
			continue
		}

		deployments, err := api.ProjectDeployments(log, username, config, client)
		if err != nil {
			log.Fatal(err, "Failed to get deployments")
		}

		lights.Update(log, config, deployments, display)

		time.Sleep(4 * time.Second)
	}
}
