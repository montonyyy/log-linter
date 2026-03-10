package main

import "log/slog"

func main() {
	// 1. checking for capital letter

	slog.Info("Starting server on port 8080")   // want "contains capital letter"
	slog.Error("Failed to connect to database") // want "contains capital letter"

	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")
}
