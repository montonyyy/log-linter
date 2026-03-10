package main

import "log/slog"

func main() {
	// 3. checking for symbol letters

	slog.Info("server started! 🚀")               // want "contains symbol letter"
	slog.Error("connection failed!!!")           // want "contains symbol letter"
	slog.Warn("waring: something went wrong...") // want "contains symbol letter"

	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")
}
