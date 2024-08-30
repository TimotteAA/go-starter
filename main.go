package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/TimotteAA/go-starter/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var mode = "development"
	flag.StringVar(&mode, "mode", "development", "Set the mode of app")

	flag.Parse()

	config, err := config.InitConfig(mode)
	if err != nil {
		log.Fatal("Fail to load config")
	}


	app := fiber.New(fiber.Config{
		AppName: config.AppName,
	})

	app.Listen(fmt.Sprintf(":%s", config.AppPort))
}