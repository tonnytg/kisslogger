package main

import (
	"errors"
	"log/slog"
	"os"
)

var globalError error

func main() {
	slog.Info("Start kisslogger", "errors:", globalError)

	if len(os.Args) < 2 {
		globalError = errors.New("No arguments")
		slog.Error("main.go", "errors:", globalError)
		return
	}

	slog.Info("End kisslogger", "errors:", globalError)
}
