package lights

import (
	"strings"

	blinkt "github.com/alexellis/blinkt_go"
	"github.com/gleich/lumber"
	"github.com/gleich/verpi/pkg/conf"
)

// Update the lights based off the deployment statuses
func Update(config conf.Conf, deployments []string, display *blinkt.Blinkt) {
	display.SetBrightness(*config.Brightness)
	lumber.Info("Updating lights")

	if *config.Brightness == 0.0 {
		display.SetAll(0, 0, 0)
		lumber.Success("Let lights to off because of brightness set to 0.0")
		return
	}

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
