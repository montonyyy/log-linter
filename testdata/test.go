package main

import (
	"log/slog"
)

var password, apiKey, token string

func test() {
	slog.Error("sdа")
	// 1. capital letter
	slog.Info("Starting server on port 8080")   // wrong
	slog.Error("Failed to connect to database") // wrong

	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")

	// 2. not an english letters
	slog.Info("запуск сервера")                    // wrong
	slog.Error("ошибка подключения к базе данных") // wrong

	slog.Info("starting server")
	slog.Error("failed to connect to database")

	// 3. symbol letters
	slog.Info("server started! 🚀")               // wrong
	slog.Error("connection failed!!!")           // wrong
	slog.Warn("waring: something went wrong...") // wrong

	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")

	// 4. sensitive data
	slog.Info("user password:" + password) // wrong
	slog.Debug("api_key=" + apiKey)        // wrong
	slog.Info("token:" + token)            // wrong

	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")
	slog.Info("token validated")
}
