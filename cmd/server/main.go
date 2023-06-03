package main

import (
	"github.com/arvians-id/go-portfolio/cmd/config"
	"github.com/arvians-id/go-portfolio/internal/http/routes"
	"log"
	"os"
)

func main() {
	// Init Log File
	file, err := os.OpenFile("./logs/main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("There is something wrong with the log file", err)
	}
	defer file.Close()

	configuration := config.New()
	router, err := routes.NewInitializedRoutes(configuration, file)
	if err != nil {
		log.Fatalln(err)
	}

	err = router.Listen(configuration.Get("APP_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
}
