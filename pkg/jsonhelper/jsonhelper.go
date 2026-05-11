package jsonhelper

import (
	"encoding/json"
	"log/slog"
	"os"
)

func Encode[T any](t T) []byte {
	b, err := json.Marshal(t)
	if err != nil {
		slog.Error("failed to encode entity:", "err", err)
		os.Exit(1)
	}
	return b
}

func Decode[T any](b []byte) T {
	var t T

	err := json.Unmarshal(b, &t)
	if err != nil {
		slog.Error("failed to decode entity:", "err", err, "value", string(b))
		os.Exit(1)
	}
	return t
}
