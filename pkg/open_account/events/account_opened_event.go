package events

import (
	"time"
)

type AccountOpenedEvent struct {
	Holder         string
	Email          string
	Type           string
	OpeningBalance int32
	CreatedDate    time.Time
}
