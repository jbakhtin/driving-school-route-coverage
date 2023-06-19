package logger

import (
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

type Logger struct {
	zap.Logger
}

func New(cfg config.Config) *Logger {
	var tops []teeOption

	if cfg.AppEnv == "dev" {
		tops = append(tops, teeOption{
			W: os.Stdout,
			Lef: func(lvl zapcore.Level) bool {
				return true
			},
		})
	} else {
		time := time.Now()

		dir1 := fmt.Sprintf("%vinfo/", cfg.Log.Directory)
		err := os.MkdirAll(dir1,0777)
		if err != nil {
			panic(err)
		}

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%v/%v-info.log", dir1, time.Format("2006-01-02")),
			MaxSize:    cfg.Log.MaxSize,
			MaxBackups: cfg.Log.MaxBackups,
			MaxAge:     cfg.Log.MaxAge,
			Compress:   cfg.Log.Compress,
		})


		tops = append(tops, teeOption{
			W: w,
			Lef: func(lvl zapcore.Level) bool {
				return lvl <= zapcore.InfoLevel
			},
		})

		dir2 := fmt.Sprintf("%verror/", cfg.Log.Directory)
		err = os.MkdirAll(dir2,0777)
		if err != nil {
			panic(err)
		}

		w = zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%v/%v-error.log", dir2, time.Format("2006-01-02")),
			MaxSize:    cfg.Log.MaxSize,
			MaxBackups: cfg.Log.MaxBackups,
			MaxAge:     cfg.Log.MaxAge,
			Compress:   cfg.Log.Compress,
		})

		tops = append(tops, teeOption{
			W: w,
			Lef: func(lvl zapcore.Level) bool {
				return lvl > zapcore.InfoLevel
			},
		})
	}

	cores := newTee(tops)
	logger := &Logger{
		Logger: *zap.New(zapcore.NewTee(cores...)),
	}
	defer logger.Sync()

	return logger
}

type LevelEnablerFunc func(lvl zapcore.Level) bool

type RotateOptions struct {
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type teeOption struct {
	W   io.Writer
	Lef LevelEnablerFunc
}

func newTee(tops []teeOption) []zapcore.Core {
	var cores []zapcore.Core
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
	}

	for _, top := range tops {
		top := top
		if top.W == nil {
			panic("the writer is nil")
		}

		lv := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return top.Lef(zapcore.Level(lvl))
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(top.W),
			lv,
		)
		cores = append(cores, core)
	}

	return cores
}
