package commands

type OpenAccountCommand struct {
	Id             string
	Holder         string
	Email          string
	Type           string
	OpeningBalance int32
}
