package main

import "log/slog"

var password, apiKey, token string

func main() {
	// 4. checking for sensitive data

	slog.Info("user password" + password) // want "contains sensitive data"
	slog.Debug("api key" + apiKey)        // want "contains sensitive data"
	slog.Info("token" + token)            // want "contains sensitive data"

	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")
	slog.Info("token validated")
}
