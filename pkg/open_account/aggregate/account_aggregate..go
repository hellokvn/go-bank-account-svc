package aggregates

import "github.com/hellokvn/go-bank-account-svc/pkg/cqrs"

type AccountAggregate struct {
	cqrs.BaseAggregate
	Active  bool
	Balance int32
}
