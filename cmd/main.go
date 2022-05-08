package main

import (
	"fmt"
	"log"

	"github.com/hellokvn/go-bank-account-svc/pkg/common/config"
	"github.com/hellokvn/go-bank-account-svc/pkg/open_account"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	// h := db.Init(&c)
	app := fiber.New()

	open_account.RegisterRoutes(app)

	app.Listen(fmt.Sprintf(":%s", c.Port))
}
