package open_account

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/hellokvn/go-bank-account-svc/pkg/open_account/commands"
)

type OpenAccountRequestBody struct {
	Holder string `json:"holder"`
}

func (h handler) OpenAccount(c *fiber.Ctx) error {
	body := OpenAccountRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	fmt.Println(body)

	id, err := uuid.NewV4()

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	command := &commands.OpenAccountCommand{
		Id:     id.String(),
		Holder: body.Holder,
	}

	h.Bus.Execute(command)

	return c.Status(fiber.StatusCreated).JSON(&command)
}
