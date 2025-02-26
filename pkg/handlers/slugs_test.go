package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/nicolasverle/slugifyer/pkg/config"
	"github.com/nicolasverle/slugifyer/pkg/router"
	"github.com/stretchr/testify/assert"
)

func TestShouldFailWithInvalidSlug(t *testing.T) {
	router := router.SetupRouter()

	NewShortener().AddRoutes(router)

	w := httptest.NewRecorder()

	slug2Test := "wZrtdffgfgsfzeferfgefdgfsgfdg"

	req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", slug2Test), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), fmt.Sprintf("%s is not a valid slug", slug2Test))
}

func TestShouldFailWithNoValidURL(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()

	url2Test := "foobar"

	req, _ := http.NewRequest("POST", fmt.Sprintf("/shortenize?url=%s", url2Test), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), fmt.Sprintf("%s is not a valid URL", url2Test))
}

func TestShouldFailWithURLCannotBeReached(t *testing.T) {
	router := setupServer()

	w := httptest.NewRecorder()

	url2Test := "https://foobar.com/bobTheSponge"

	req, _ := http.NewRequest("POST", fmt.Sprintf("/shortenize?url=%s", url2Test), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), fmt.Sprintf("%s cannot be reached", url2Test))
}

func setupServer() *gin.Engine {
	router := router.SetupRouter()

	NewShortener().AddRoutes(router)

	if err := config.Parse(); err != nil {
		log.Fatal().Msgf("unable to parse configuration, %s", err.Error())
	}

	return router
}
