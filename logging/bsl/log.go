package bsl

import (
	"fmt"
	"os"
	"time"

	phlog "github.com/phuslu/log"
)

type Config struct {
	TimeFormat    string
	IncludeCaller bool
}

var log *phlog.Logger

func New(config ...Config) {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.TimeFormat == "" {
		cfg.TimeFormat = "060102 15:04:05.000"
	}

	log = &phlog.Logger{
		Level:      phlog.TraceLevel,
		TimeFormat: cfg.TimeFormat,
		Writer: &phlog.ConsoleWriter{
			ColorOutput: true,
			Writer:      os.Stdout,
		},
	}

	if cfg.IncludeCaller {
		log.Caller = 2
	}
}

func prefix(section string) string {
	return fmt.Sprintf("%-15s || ", section)
}

type Entry struct {
	entry   *phlog.Entry
	section string
}

func (e *Entry) Str(key, val string) *Entry               { e.entry.Str(key, val); return e }
func (e *Entry) Strs(key string, val []string) *Entry     { e.entry.Strs(key, val); return e }
func (e *Entry) Int(key string, val int) *Entry           { e.entry.Int(key, val); return e }
func (e *Entry) Dur(key string, val time.Duration) *Entry { e.entry.Dur(key, val); return e }
func (e *Entry) Err(err error) *Entry                     { e.entry.Err(err); return e }
func (e *Entry) Msg(msg string)                           { e.entry.Msg(prefix(e.section) + msg) }
func (e *Entry) Msgf(format string, v ...any)             { e.entry.Msgf(prefix(e.section)+format, v...) }

func Debug(section string) *Entry { return &Entry{entry: log.Debug(), section: section} }
func Info(section string) *Entry  { return &Entry{entry: log.Info(), section: section} }

//goland:noinspection GoUnusedExportedFunction
func Warn(section string) *Entry  { return &Entry{entry: log.Warn(), section: section} }
func Error(section string) *Entry { return &Entry{entry: log.Error(), section: section} }
func Fatal(section string) *Entry { return &Entry{entry: log.Fatal(), section: section} }
func Trace(section string) *Entry { return &Entry{entry: log.Trace(), section: section} }

//goland:noinspection GoUnusedExportedFunction
func Debugf(section, format string, v ...any) { log.Debug().Msgf(prefix(section)+format, v...) }

//goland:noinspection GoUnusedExportedFunction
func Infof(section, format string, v ...any) { log.Info().Msgf(prefix(section)+format, v...) }

//goland:noinspection GoUnusedExportedFunction
func Warnf(section, format string, v ...any) { log.Warn().Msgf(prefix(section)+format, v...) }

//goland:noinspection GoUnusedExportedFunction
func Errorf(section, format string, v ...any) { log.Error().Msgf(prefix(section)+format, v...) }

//goland:noinspection GoUnusedExportedFunction
func Fatalf(section, format string, v ...any) { log.Fatal().Msgf(prefix(section)+format, v...) }

//goland:noinspection GoUnusedExportedFunction
func Tracef(section, format string, v ...any) { log.Trace().Msgf(prefix(section)+format, v...) }
