package msg

import (
	"encoding/json"
	"errors"
	"kafka-final/domain"
)

type ProductCodec struct {
	topic string
	sch   *Schema
}

func (mc ProductCodec) Encode(value any) ([]byte, error) {
	if _, isMsg := value.(*domain.Product); !isMsg {
		return nil, errors.New("value is not Product")
	}
	return json.Marshal(value)
}

func (mc ProductCodec) Decode(data []byte) (any, error) {
	var (
		p   domain.Product
		err error
	)
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, errors.New("unmarshal Product failed")
	}
	return &p, nil
}

func NewProductCodec(topic string, sch *Schema) ProductCodec {
	return ProductCodec{topic, sch}
}

type FindCodec struct {
	topic string
	sch   *Schema
}

func (mc FindCodec) Encode(value any) ([]byte, error) {
	if _, isMsg := value.(*domain.Find); !isMsg {
		return nil, errors.New("value is not Find")
	}
	return json.Marshal(value)
}

func (mc FindCodec) Decode(data []byte) (any, error) {
	var (
		p   domain.Find
		err error
	)
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, errors.New("unmarshal Find failed")
	}
	return &p, nil
}

func NewFindCodec(topic string, sch *Schema) FindCodec {
	return FindCodec{topic, sch}
}
