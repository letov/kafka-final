package shop

import (
	"context"
	"kafka-final/domain"
	"kafka-final/internal/infra/msg"
)

func Start(
	e *msg.ProductEmitter,
) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	productOutCh := make(chan *domain.Product)
	defer close(productOutCh)

	go e.Emit(ctx, productOutCh)

	prodCnt := 10
	domain.GenerateNewProducts(ctx, prodCnt, productOutCh)
}
