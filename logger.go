package viperp

import (
	"fmt"
	jww "github.com/spf13/jwalterweatherman"
)

// Logger is a unified interface for various logging use cases and practices, including:
//   - leveled logging
//   - structured logging
type Logger interface {
	// Info logs an Info event
	// General information about what's happening inside the system
	Info(msg string, keyvals ...interface{})
}

type jwwLogger struct {
}

func (jwwLogger) Info(msg string, keyvals ...interface{}) {
	jww.INFO.Printf(jwwLogMessage(msg, keyvals...))
}

func jwwLogMessage(msg string, keyvals ...interface{}) string {
	out := msg

	if len(keyvals) > 0 && len(keyvals)%2 == 1 {
		keyvals = append(keyvals, nil)
	}

	for i := 0; i < len(keyvals)-2; i += 2 {
		out = fmt.Sprintf("%s %v=%v", out, keyvals[i], keyvals[i+1])
	}

	return out
}
