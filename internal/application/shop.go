package application

import (
	"context"
	"kafka-final/internal/domain"
	"kafka-final/internal/infra/msg"
	"time"
)

func StartShop(
	e *msg.Emitter,
) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	productCh := make(chan *domain.Product)
	defer close(productCh)

	go e.EmitProduct(ctx, productCh)
	prodPerTickCnt := 30
	ticker := time.NewTicker(1000 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			domain.GenerateNewProducts(ctx, prodPerTickCnt, productCh)
		}
	}
}
