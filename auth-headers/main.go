package main

import (
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"
)

var (
	listen = os.Getenv("LISTEN")
	logger *zap.Logger
)

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
}

func main() {
	http.HandleFunc("/", GatekeeperHeadersToEnvoyExtAuthHandler)
	logger.Info("Starting http server at", zap.String("listen", listen))
	log.Fatal(http.ListenAndServe(listen, nil))
}
