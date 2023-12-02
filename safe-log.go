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
			{Key: "IBAN", Value: slog.StringValue("supersecretpassword")},
			{Key: "foo", Value: slog.StringValue("bar")}}))

	log.Info("Hello World", "anotherField", "fizzbuzz", "piiFieldBroken", "SE45 5000 0000 0583 9825 7466", "piiField", "DE12345678901234567800")

}
