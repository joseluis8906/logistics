package bus

import (
	"encoding/json"
	"errors"
)

var (
	ErrEmptyEmitter = errors.New("attempt to use an empty emitter")
)

type (
	Emiter interface {
		Emit(string, []byte) error
	}

	Event struct {
		Event   string `json:"event"`
		Payload []byte `json:"payload"`
	}
)

var (
	output Emiter
)

func SetOutput(aOutput Emiter) {
	output = aOutput
}

func Emit(event string, data interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if output == nil {
		return ErrEmptyEmitter
	}

	return output.Emit(event, payload)
}
