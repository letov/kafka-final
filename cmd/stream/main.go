package main

import (
	"kafka-final/internal/application"
	"kafka-final/internal/infra/di"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.InjectStreamApp(),
		fx.Invoke(application.StartStream),
	).Run()
}
