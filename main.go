package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/nicolasverle/slugifyer/docs"
	"github.com/nicolasverle/slugifyer/pkg/config"
	"github.com/nicolasverle/slugifyer/pkg/handlers"
	"github.com/nicolasverle/slugifyer/pkg/router"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Slugifyer API
//	@version		0.0.1
//	@description	Slugifyer API to alias URL into a short string ID

// @host		localhost:8080
// @BasePath	/
func main() {
	r := router.SetupRouter()

	// parse API configuration
	if err := config.Parse(); err != nil {
		log.Fatal().Msgf("unable to parse configuration, %s", err.Error())
	}

	// health route
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "up"})
	})

	// expose swagger API documentation
	r.GET("/doc/*any", func(c *gin.Context) {
		if c.Request.RequestURI == "/doc/" {
			c.Redirect(302, "/doc/index.html")
			return
		}
		ginSwagger.WrapHandler(swaggerfiles.Handler)(c)
	})

	handlers.NewShortener().AddRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal().Msgf("unable to run the API on :8080, %s", err.Error())
	}
}
