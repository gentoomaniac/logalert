package main

import (
	"log/slog"
	"os"

	safelog "github.com/gentoomaniac/logalert/pkg/safeLog"
)

type safeLog struct {
}

func safeLogCmd() {
	log := slog.New(safelog.NewSafeLogHandler(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug})).WithAttrs(
		[]slog.Attr{
			{Key: "password", Value: slog.StringValue("supersecretpassword")}}))

	log.Info("Hello World", "anotherpassword", "passwordissupersecret")

}
