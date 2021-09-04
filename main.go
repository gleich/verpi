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
	display := lights.Setup(config)

	for {
		deployments, err := api.Deployments(config, http.DefaultClient)
		if err != nil {
			lumber.Fatal(err, "Failed to get deployments")
		}

		lights.Update(deployments, display)

		time.Sleep(2 * time.Second)
	}
}
