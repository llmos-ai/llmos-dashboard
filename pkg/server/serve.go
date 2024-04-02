package server

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/llmos/llmos-dashboard/pkg/database"
	"github.com/llmos/llmos-dashboard/pkg/router"
)

type ApiServer struct {
	Context context.Context
	Engine  *gin.Engine
	SqlDB   *sql.DB
}

func NewApiServer(ctx context.Context) ApiServer {
	sql, err := database.RegisterSQLiteDB()
	if err != nil {
		slog.Error("Failed to init auth", err)
		panic(0)
	}

	return ApiServer{
		Context: ctx,
		Engine:  gin.Default(),
		SqlDB:   sql,
	}
}

func (a ApiServer) ListenAndServe() error {
	// register routers
	router.RegisterRouters(a.Engine, a.SqlDB)

	a.Engine.Run()

	<-a.Context.Done()
	return nil
}
