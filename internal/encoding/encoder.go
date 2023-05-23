package encoding

// Encoder encodes the contents fo v into a byte representation.
// It's primarily used for encoding a map[string]interface{} into a file format.
type Encoder interface {
	Encode(v map[string]interface{}) ([]byte, error)
}
