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

	// Init Configuration
	configuration := config.New()

	// Init DB
	db, err := config.NewPostgresSQLGorm(configuration)
	if err != nil {
		panic(err)
	}

	// Init Server
	router, err := http.NewInitializedRoutes(configuration, file, db)
	if err != nil {
		panic(err)
	}

	err = router.Listen(configuration.Get("APP_PORT"))
	if err != nil {
		panic(err)
	}
}
