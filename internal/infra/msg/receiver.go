package msg

import (
	"context"
	"errors"
	"kafka-final/internal/domain"
	"kafka-final/internal/infra/config"

	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Receiver struct {
	cons  *kafka.Consumer
	conf  config.Config
	l     *zap.SugaredLogger
	sch   *Schema
	codec *ProductCodec
}

func (r Receiver) Receive(
	ctx context.Context,
	productCh chan *domain.Product,
) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ev := r.cons.Poll(1000)
				if ev == nil {
					continue
				}
				switch e := ev.(type) {
				case *kafka.Message:
					r.l.Info("Get message")
					raw, err := r.codec.Decode(e.Value)
					data, ok := raw.(*domain.Product)
					if !ok {
						r.l.Warn(errors.New("decode message error"))
					}
					if err != nil {
						r.l.Warn(err.Error())
					} else {
						_, _ = r.cons.Commit()
						productCh <- data
					}
				case kafka.Error:
					r.l.Warn("Error: ", e)
				default:
					r.l.Warn("Some event: ", e)
				}
			}
		}
	}()
}

func NewReceiver(
	lc fx.Lifecycle,
	conf config.Config,
	l *zap.SugaredLogger,
	sch *Schema,
) *Receiver {
	cons, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  strings.Join(conf.Brokers, ","),
		"group.id":           "analytic_consumer_group",
		"session.timeout.ms": "10000",
		"auto.offset.reset":  "earliest",
		"acks":               "all",
	})
	if err != nil {
		l.Fatal("Error creating consumer: ", err)
	}

	l.Info("Consumer created")

	if cons.SubscribeTopics([]string{conf.ProductTopic}, nil) != nil {
		l.Fatal("Subscribe error:", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return cons.Close()
		},
	})

	codec := new(ProductCodec)
	return &Receiver{cons, conf, l, sch, codec}
}
