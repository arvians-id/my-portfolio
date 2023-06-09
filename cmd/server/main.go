package main

import (
	"github.com/arvians-id/go-portfolio/cmd/config"
	"github.com/arvians-id/go-portfolio/internal/http"
	"os"
)

func main() {
	// Init Log File
	file, err := os.OpenFile("./logs/main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	configuration := config.New()
	router, err := http.NewInitializedRoutes(configuration, file)
	if err != nil {
		panic(err)
	}

	err = router.Listen(configuration.Get("APP_PORT"))
	if err != nil {
		panic(err)
	}
}
