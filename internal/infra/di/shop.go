package di

import (
	"kafka-final/internal/infra/config"
	"kafka-final/internal/infra/msg"
	"kafka-final/internal/logger"

	"go.uber.org/fx"
)

func GetShopAppConstructors() []interface{} {
	return []interface{}{
		logger.NewLogger,
		config.NewConfig,

		msg.NewSchema,
		msg.NewEmitterProduct,
	}
}

func InjectShopApp() fx.Option {
	return fx.Provide(
		GetShopAppConstructors()...,
	)
}
