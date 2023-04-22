package stdlog

import (
	"io"
	"log"
	"strings"

	logger "github.com/pubgo/funk/log"

	"github.com/pubgo/lava/core/logging"
)

func init() {
	logging.Register("stdLog", New)
}

// New 替换std默认log
func New(logger logger.Logger) {
	stdLog := log.Default()
	// 接管系统默认log
	*stdLog = *log.New(&std{l: logger.WithName("std-log").WithCallerSkip(3)}, "", 0)
}

var _ io.Writer = (*std)(nil)

type std struct {
	l logger.Logger
}

func (s *std) Write(p []byte) (n int, err error) {
	s.l.Info().Msg(strings.TrimSpace(string(p)))
	return len(p), err
}
