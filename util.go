package viperp

import (
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// ConfigParseError denotes failing to parse configuration file
type ConfigParseError struct {
	err error
}

// Error returns the formatted configuration error
func (pe ConfigParseError) Error() string {
	return fmt.Sprintf("while parsing config: %s", pe.err.Error())
}

func insensitiviseMap(m map[string]interface{}) {
	for key, val := range m {
		val = insensitiviseVal(val)
		lower := strings.ToLower(key)
		if key != lower {
			// remove old key (not lower-cased)
			delete(m, key)
		}
		// update map
		m[lower] = val
	}
}

func insensitiviseVal(val interface{}) interface{} {
	switch val.(type) {
	case map[interface{}]interface{}:
		// nested map: cast and recursively insensitivise
		val = cast.ToStringMap(val)
		insensitiviseMap(val.(map[string]interface{}))
	case map[string]interface{}:
		insensitiviseMap(val.(map[string]interface{}))
	case []interface{}:
		insensitiviseArray(val.([]interface{}))
	}
	return val
}

func insensitiviseArray(a []interface{}) {
	for i, val := range a {
		a[i] = insensitiviseVal(val)
	}
}
