package main

import (
	"log"
	"log/slog"
)

func test() {
	password := "123"

	log.Println("!", password)
	slog.Info("asds!" + "sadsa;" + password)
}
