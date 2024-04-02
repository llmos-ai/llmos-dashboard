package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/llmos/llmos-dashboard/pkg/server"
)

func main() {
	ctx := context.Background()
	apiServer := server.NewApiServer(ctx)
	err := apiServer.ListenAndServe()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
