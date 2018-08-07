package cluster

import (
	"github.com/aergoio/aergo-actor/log"
)

var (
	plog = log.New(log.DebugLevel, "[CLUSTER]")
)

// SetLogLevel sets the log level for the logger.
//
// SetLogLevel is safe to call concurrently
func SetLogLevel(level log.Level) {
	plog.SetLevel(level)
}
