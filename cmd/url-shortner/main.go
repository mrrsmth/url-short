package main

import (
	"fmt"
	"url-shortner/internal/config"
)

func main() {
	//todo clean env
	//logger slog
	//storage sqlite
	// router chi chi render
	// run server

	cfg := config.MustLoad()

	fmt.Println(cfg)
}
