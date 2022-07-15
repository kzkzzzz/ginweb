package logz

import (
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"testing"
)

func TestZapStd_Log(t *testing.T) {
	l := NewLogger(DefaultConfig, zap.AddCaller(), zap.AddCallerSkip(1))
	l.Log(log.LevelDebug, "msg", "qwe123")
	l.Log(log.LevelDebug, "msg", "qwe123", "qqq", "fff", "rrr", "rrr666", "rrr", "rrr666", "rrr", "rrr666")
}

func BenchmarkZapStd_Log(b *testing.B) {
	l := NewLogger(DefaultConfig, zap.AddCaller(), zap.AddCallerSkip(1))
	for i := 0; i < b.N; i++ {
		l.Log(log.LevelDebug, "msg", "qwe123", "qqq", "fff", "rrr1", "rrr666", "rrr2", "rrr666", "rrr3", "rrr666")
	}
}
