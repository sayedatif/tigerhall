package main

import (
	"github.com/sayedatif/tigerhall/config"
	"github.com/sayedatif/tigerhall/server"
)

func main() {
	config.Init()

	server.Init()
}
