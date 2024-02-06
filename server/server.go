package server

import (
	"fmt"

	"github.com/sayedatif/tigerhall/config"
)

func Init() {
	config := config.GetConfig()
	router := NewRouter()
	port := config.GetString("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}
