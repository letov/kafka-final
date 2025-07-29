package di

import (
	"kafka-final/internal/infra/config"
	"kafka-final/internal/infra/msg"
	"kafka-final/internal/logger"

	"go.uber.org/fx"
)

func GetStreamAppConstructors() []interface{} {
	return []interface{}{
		logger.NewLogger,
		config.NewConfig,
		msg.NewProcessor,
	}
}

func InjectStreamApp() fx.Option {
	return fx.Provide(
		GetStreamAppConstructors()...,
	)
}
