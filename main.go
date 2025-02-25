package main

import (
	"net/http"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/nicolasverle/slugifyer/pkg/config"
	"github.com/nicolasverle/slugifyer/pkg/handlers"
	"github.com/rs/zerolog/log"
)

func main() {
	r := gin.New()
	r.Use(ginzerolog.Logger("gin"))

	if err := config.Parse(); err != nil {
		log.Fatal().Msgf("unable to parse configuration, %s", err.Error())
	}

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "up"})
	})

	handlers.NewShortener().AddRoutes(r)

	r.Run(":8080")
}
