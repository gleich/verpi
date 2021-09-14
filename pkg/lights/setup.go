package lights

import (
	blinkt "github.com/alexellis/blinkt_go"
	"github.com/gleich/lumber"
	"github.com/gleich/verpi/pkg/conf"
)

// Setup the display
func Setup(config conf.Conf) *blinkt.Blinkt {
	lumber.Info("Setting up display")
	display := blinkt.NewBlinkt(*config.Brightness)
	display.SetClearOnExit(true)
	display.Setup()
	lumber.Success("Setup display")
	return &display
}
