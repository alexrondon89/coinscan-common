package server

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

const (
	ReadTimeout  = 20 * time.Second
	WriteTimeout = 30 * time.Second
	IdleTimeout  = 3 * time.Second
)

func New(config ...fiber.Config) *fiber.App {
	var servConfig fiber.Config

	if config != nil {
		servConfig = config[0]
	} else {
		servConfig = defaultConfig()
	}

	return fiber.New(servConfig)

}

func defaultConfig() fiber.Config {
	return fiber.Config{
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		IdleTimeout:  IdleTimeout,
	}
}
