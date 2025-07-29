package main

import (
	"kafka-final/internal/application/shop"
	"kafka-final/internal/infra/di"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.InjectApp(),
		fx.Invoke(shop.Start),
	).Run()
}
