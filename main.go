package main

import (
	"net/http"

	"github.com/gleich/lumber"
	"github.com/gleich/verpi/pkg/api"
	"github.com/gleich/verpi/pkg/conf"
)

func main() {
	config, err := conf.Read()
	if err != nil {
		lumber.Fatal(err, "Failed to read from configuration file")
	}

	deployments, err := api.Deployments(config, http.DefaultClient)
	if err != nil {
		lumber.Fatal(err, "Failed to get deployments")
	}
	lumber.Debug(deployments)
}
