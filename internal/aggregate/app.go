package aggregate

import (
	"context"
	"fmt"
	"github.com/wuyazi/gddd"
)

type UserRepositorySaveListener struct{}

func (l *UserRepositorySaveListener) Handle(ctx context.Context, domainEvent gddd.DomainEventMessage) {
	fmt.Println("\n------- userRepositorySaveListener --------")
}
