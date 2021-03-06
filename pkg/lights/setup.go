package lights

import (
	blinkt "github.com/alexellis/blinkt_go"
	"github.com/gleich/lumber/v2"
	"github.com/gleich/verpi/pkg/conf"
)

// Setup the display
func Setup(log lumber.Logger, config conf.Conf) *blinkt.Blinkt {
	log.Info("Setting up display")
	display := blinkt.NewBlinkt(*config.Brightness)
	display.SetClearOnExit(true)
	display.Setup()
	display.Clear()
	log.Success("Setup display")
	return &display
}
