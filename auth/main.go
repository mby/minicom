package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mby/minicom/auth/internal/auth"
	"github.com/mby/minicom/auth/internal/cfg"
	"github.com/mby/minicom/auth/internal/helpers"
)

func GetEnv() string {
	env := os.Getenv("ENV")

	if env == "" {
		panic("ENV is not set")
	}

	return env
}

func main() {
	env := GetEnv()
	cfg := cfg.GetConfig(env)

	authHandler := auth.NewHandler(cfg)
	defer authHandler.Cleanup()

	port := ":8080"
	helpers.PrintBanner("auth", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
