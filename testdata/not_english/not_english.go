package main

import "log/slog"

func main() {
	// 2. checking for not an english letters

	slog.Info("запуск сервера")                    // want "contains not an english letter"
	slog.Error("ошибка подключения к базе данных") // want "contains not an english letter"

	slog.Info("starting server")
	slog.Error("failed to connect to database")
}
