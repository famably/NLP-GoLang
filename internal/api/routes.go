package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, handler *Handler) {
    router.POST("/groups/:group_id/documents", handler.StoreDocument)
    router.GET("/groups/:group_id/documents", handler.GetDocuments)
    router.GET("/groups/:group_id/summary", handler.SummarizeDocuments)
}