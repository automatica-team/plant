package core

import (
	"log/slog"
	"os"

	"github.com/go-telebot/pkg/monitor"
	"github.com/mitchellh/mapstructure"
	tele "gopkg.in/telebot.v3"
)

func (mod *Core) Logger() tele.MiddlewareFunc {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	logger := slog.New(handler)
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if update, ok := monitor.NewUpdate(c); ok {
				var fields map[string]any

				dec, err := mapstructure.NewDecoder(
					&mapstructure.DecoderConfig{
						TagName: "json",
						Result:  &fields,
					},
				)
				if err != nil {
					return nil
				}

				if err := dec.Decode(update); err == nil {
					delete(fields, "text")
					delete(fields, "time")
					delete(fields, "date")

					var args []any
					for k, v := range fields {
						args = append(args, k, v)
					}

					logger.Info(update.Text, args...)
				}
			}
			return next(c)
		}
	}
}
