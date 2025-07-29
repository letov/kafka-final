package di

import (
	"kafka-final/internal/infra/config"
	"kafka-final/internal/infra/msg"
	"kafka-final/internal/logger"

	"go.uber.org/fx"
)

func GetAppConstructors() []interface{} {
	return []interface{}{
		logger.NewLogger,
		config.NewConfig,

		msg.NewSchema,
		msg.NewEmitter,
	}
}

func InjectApp() fx.Option {
	return fx.Provide(
		GetAppConstructors()...,
	)
}
