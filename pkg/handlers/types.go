package handlers

import "github.com/gin-gonic/gin"

type (
	Handler interface {
		AddRoutes(e *gin.Engine)
	}
)
