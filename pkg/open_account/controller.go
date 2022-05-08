package open_account

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/EventBus"
	"github.com/hellokvn/go-bank-account-svc/pkg/open_account/commands"
	"gorm.io/gorm"
)

type handler struct {
	DB  *gorm.DB
	Bus EventBus.Bus
}

func OpenAccountCommandHandler(command *commands.OpenAccountCommand) {
	fmt.Println("OpenAccountCommandHandler")
	fmt.Println(command)
}

func RegisterRoutes(app *fiber.App) {
	h := &handler{
		Bus: EventBus.New(),
	}

	h.Bus.Subscribe("OpenAccountCommand", OpenAccountCommandHandler)

	routes := app.Group("/bank-account")
	routes.Post("/open", h.OpenAccount)
}
