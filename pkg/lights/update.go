package lights

import (
	"strings"

	blinkt "github.com/alexellis/blinkt_go"
	"github.com/gleich/lumber"
)

// Update the lights based off the deployment statuses
func Update(deployments []string, display *blinkt.Blinkt) {
	lumber.Info("Updating lights")
	for i, deployment := range deployments {
		switch strings.ToUpper(deployment) {
		case "READY":
			display.SetPixel(i, 0, 255, 0) // Green
		case "QUEUED":
			display.SetPixel(i, 255, 128, 0) // Yellow
		case "BUILDING":
			display.SetPixel(i, 255, 128, 0) // Yellow
		default:
			display.SetPixel(i, 255, 0, 0) // Red
		}
	}
	display.Show()
	lumber.Success("Updated lights")
}
