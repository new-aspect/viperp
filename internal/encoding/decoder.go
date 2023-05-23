package encoding

import "sync"

type Decoder interface {
	Decode(b []byte, v map[string]interface{}) error
}

const (
	// ErrDecoderNotFound is returned when there is no decoder registered for a format.
	ErrDecoderNotFound = encodingError("decoder not found for this format")

	// ErrDecoderFormatAlreadyRegistered is returned when an decoder is already registered for a format.
	ErrDecoderFormatAlreadyRegistered = encodingError("decoder already registered for this format")
)

type DecoderRegistry struct {
	decoders map[string]Decoder

	mu sync.RWMutex
}

func (e *DecoderRegistry) Decode(format string, b []byte, v map[string]interface{}) error {
	e.mu.RLock()
	decoder, ok := e.decoders[format]
	e.mu.RUnlock()

	if !ok {
		return ErrDecoderFormatAlreadyRegistered
	}

	return decoder.Decode(b, v)
}
