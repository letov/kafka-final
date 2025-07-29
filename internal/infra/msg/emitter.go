package msg

import (
	"context"
	"fmt"
	"kafka-final/domain"
	"kafka-final/internal/infra/config"

	"github.com/lovoo/goka"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ProductEmitter struct {
	ge  *goka.Emitter
	log *zap.SugaredLogger
}

func (pe ProductEmitter) Emit(ctx context.Context, productCh chan *domain.Product) {
	for {
		select {
		case <-ctx.Done():
			return
		case product, ok := <-productCh:
			if ok {
				err := pe.ge.EmitSync(fmt.Sprintf("emit product %s", product.Name), product)
				if err != nil {
					pe.log.Fatal(err)
				}
			}
		}
	}
}

func NewEmitter(lc fx.Lifecycle, log *zap.SugaredLogger, config config.Config, sch *Schema) *ProductEmitter {
	codec := NewMsgCodec(config.ProductTopic, sch)
	ge, err := goka.NewEmitter(config.Brokers, goka.Stream(config.ProductTopic), codec)
	if err != nil {
		log.Fatal(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return ge.Finish()
		},
	})

	return &ProductEmitter{
		ge,
		log,
	}
}
