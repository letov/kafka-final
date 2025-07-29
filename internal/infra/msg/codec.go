package msg

import (
	"encoding/json"
	"errors"
	"kafka-final/domain"
)

type Codec struct {
	topic string
	sch   *Schema
}

func (mc Codec) Encode(value any) ([]byte, error) {
	if _, isMsg := value.(*domain.Product); !isMsg {
		return nil, errors.New("value is not Product")
	}
	return json.Marshal(value)
}

func (mc Codec) Decode(data []byte) (any, error) {
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

func NewMsgCodec(topic string, sch *Schema) Codec {
	return Codec{topic, sch}
}
