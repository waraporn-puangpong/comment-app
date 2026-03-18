package router

import (
	"backend/internal/comment"
	"backend/internal/db"
	"backend/internal/server"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(apiServer *server.Server, dbconn db.IMongo) {
	handlers := comment.NewHandler(apiServer, dbconn)

	api := apiServer.Group("/api")
	api.POST("/comments/update-comment", handlers.SaveComment)
	api.GET("/comments", handlers.GetComments)

	swagger := apiServer.Group("/swagger")
	swagger.GET("/*any", func(ctx *server.Context) error {
		ginSwagger.WrapHandler(swaggerFiles.Handler)(ctx.Context)
		return nil
	})
}
