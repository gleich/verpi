package main

import (
	"github.com/gleich/lumber"
	"github.com/gleich/verpi/pkg/conf"
)

func main() {
	_, err := conf.Read()
	if err != nil {
		lumber.FatalMsg(err, "Failed to read from configuration file")
	}
}
