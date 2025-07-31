package application

import (
	"context"
	"kafka-final/internal/domain"
	"kafka-final/internal/infra/msg"

	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
)

func StartClient(
	e *msg.Emitter,
) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	findCh := make(chan *domain.Find)
	defer close(findCh)

	go e.EmitFind(ctx, findCh)

	prompt := promptui.Prompt{
		Label: "Find product",
	}

	for {
		p, _ := prompt.Run()
		if len(p) > 0 {
			findCh <- &domain.Find{
				Id:     uuid.New().String(),
				UserId: "uid",
				Find:   p,
			}
		}
	}
}
