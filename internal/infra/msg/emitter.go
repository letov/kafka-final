package msg

import (
	"context"
	"kafka-final/domain"
	"kafka-final/internal/infra/config"

	"github.com/lovoo/goka"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Emitter struct {
	ge  *goka.Emitter
	log *zap.SugaredLogger
}

func (pe Emitter) EmitProduct(ctx context.Context, productCh chan *domain.Product) {
	for {
		select {
		case <-ctx.Done():
			return
		case product, ok := <-productCh:
			if ok {
				err := pe.ge.EmitSync(product.ProductId, product)
				if err != nil {
					pe.log.Fatal(err)
				}
			}
		}
	}
}

func (pe Emitter) EmitFind(ctx context.Context, findCh chan *domain.Find) {
	for {
		select {
		case <-ctx.Done():
			return
		case f, ok := <-findCh:
			if ok {
				err := pe.ge.EmitSync(f.Id, f)
				if err != nil {
					pe.log.Fatal(err)
				}
			}
		}
	}
}

func NewEmitterProduct(lc fx.Lifecycle, log *zap.SugaredLogger, config config.Config, sch *Schema) *Emitter {
	codec := NewProductCodec(config.ProductTopic, sch)
	ge, err := goka.NewEmitter(config.Brokers, goka.Stream(config.ProductTopic), codec)
	if err != nil {
		log.Fatal(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return ge.Finish()
		},
	})

	return &Emitter{
		ge,
		log,
	}
}

func NewEmitterFind(lc fx.Lifecycle, log *zap.SugaredLogger, config config.Config, sch *Schema) *Emitter {
	codec := NewFindCodec(config.ProductFind, sch)
	ge, err := goka.NewEmitter(config.Brokers, goka.Stream(config.ProductFind), codec)
	if err != nil {
		log.Fatal(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return ge.Finish()
		},
	})

	return &Emitter{
		ge,
		log,
	}
}
