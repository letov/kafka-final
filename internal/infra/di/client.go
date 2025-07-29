package di

import (
	"kafka-final/internal/infra/config"
	"kafka-final/internal/infra/msg"
	"kafka-final/internal/logger"

	"go.uber.org/fx"
)

func GetClientAppConstructors() []interface{} {
	return []interface{}{
		logger.NewLogger,
		config.NewConfig,

		msg.NewSchema,
		msg.NewEmitterFind,
	}
}

func InjectClientApp() fx.Option {
	return fx.Provide(
		GetClientAppConstructors()...,
	)
}
