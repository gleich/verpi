package lights

import (
	"strings"

	blinkt "github.com/alexellis/blinkt_go"
	"github.com/gleich/lumber"
	"github.com/gleich/verpi/pkg/api"
)

// Update the lights based off the deployment statuses
func Update(deployments []api.Deployment, display *blinkt.Blinkt) {
	lumber.Info("Updating lights")
	for i, deployment := range deployments {
		switch strings.ToUpper(deployment.State) {
		case "READY":
			display.SetPixel(i, 0, 255, 0)
		case "QUEUED":
			display.SetPixel(i, 255, 128, 0)
		case "BUILDING":
			display.SetPixel(i, 255, 128, 0)
		default:
			display.SetPixel(i, 255, 0, 0)
		}
	}
	display.Show()
	lumber.Success("Updated lights")
}
