package server

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/llmos-ai/llmos-dashboard/pkg/database"
	"github.com/llmos-ai/llmos-dashboard/pkg/generated/ent"
	"github.com/llmos-ai/llmos-dashboard/pkg/router"
)

type ApiServer struct {
	Context  context.Context
	Engine   *gin.Engine
	DBClient *ent.Client
}

func NewApiServer(ctx context.Context) ApiServer {
	client, err := database.RegisterDBClient(ctx)
	if err != nil {
		slog.Error("Failed to init auth", err)
		panic(0)
	}

	return ApiServer{
		Context:  ctx,
		Engine:   gin.Default(),
		DBClient: client,
	}
}

func (a ApiServer) ListenAndServe() error {
	// register routers
	err := router.RegisterRouters(a.Engine, a.DBClient, a.Context)
	if err != nil {
		return err
	}

	a.Engine.Run()

	<-a.Context.Done()

	return a.DBClient.Close()
}
