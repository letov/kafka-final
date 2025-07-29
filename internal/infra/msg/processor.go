package msg

import (
	"context"
	"kafka-final/domain"
	"kafka-final/internal/infra/config"
	"log"
	"regexp"

	"github.com/lovoo/goka"
	"go.uber.org/zap"
)

type Processor struct {
	gp     *goka.Processor
	config config.Config
	log    *zap.SugaredLogger
}

func (p *Processor) getProcess(_ctx context.Context, filterCh chan string) func(ctx goka.Context, msg any) {
	filter := ""
	var re *regexp.Regexp

	go func() {
		for {
			select {
			case <-_ctx.Done():
				return
			case filter = <-filterCh:
				re = regexp.MustCompile(filter)
			}
		}
	}()

	return func(ctx goka.Context, msg any) {
		var (
			m  *domain.Product
			ok bool
		)

		if m, ok = msg.(*domain.Product); !ok || m == nil {
			return
		}

		if len(filter) == 0 || re.MatchString(m.Name) {
			ctx.SetValue(m)
		}
	}
}

func (p *Processor) Run(ctx context.Context, filterCh chan string) {
	g := goka.DefineGroup(goka.Group(p.config.ProductFiltered),
		goka.Input(goka.Stream(p.config.ProductTopic), new(ProductCodec), p.getProcess(ctx, filterCh)),
		goka.Persist(new(ProductCodec)),
	)

	gp, err := goka.NewProcessor(p.config.Brokers, g)
	if err != nil {
		log.Fatal(err)
	}
	p.gp = gp

	err = p.gp.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}

	p.gp = gp
}

func (p *Processor) Stop() {
	if p.gp != nil {
		p.gp.Stop()
	}
}

func NewProcessor(conf config.Config, log *zap.SugaredLogger) *Processor {
	return &Processor{nil, conf, log}
}
