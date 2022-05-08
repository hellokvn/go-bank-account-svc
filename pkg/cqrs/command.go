package cqrs

type BaseCommand struct {
	Id string
}

func (c *BaseCommand) GetCommandId() string {
	return c.Id
}

func (c *BaseCommand) SetCommandId(id string) {
	c.Id = id
}
