package handlers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nicolasverle/slugifyer/pkg/sql"
	"github.com/spf13/viper"
)

type (
	slugsHandler struct{}

	slugObject struct {
		URL  string `db:"url"`
		Slug string `db:"slug"`
	}
)

func NewShortener() Handler {
	return &slugsHandler{}
}

func (s *slugsHandler) AddRoutes(e *gin.Engine) {
	e.POST("/shortenize", shortenize)
	e.GET("/:slug", redirect)
}

// redirect godoc
//
//	@Summary		Will redirect to the URL specified by the slug
//	@Description	take an URL as a parameter, validate it and then create an alphanumeric alias
//	@Tags			slugs
//	@Param			slug	body		string	true	"slug to be processd for redirection"
//	@Success		303		{object}	nil
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Failure		404		{object}	error
//	@Router			/:slug [get]
func redirect(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse the slug"})
		return
	}

	if len(slug) != viper.GetInt("slugs.chars_length") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s is not a valid slug", slug)})
		return
	}

	db, err := sql.NewSQLHandler[slugObject]()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	retrieved, err := db.Read(slugObject{Slug: slug})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if retrieved != nil {
		http.Redirect(ctx.Writer, ctx.Request, retrieved.URL, http.StatusSeeOther)
		return
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("no url recorded for the slug %s", slug)})
		return
	}
}

// shortenize godoc
//
//	@Summary		Will shorten an URL
//	@Description	take an URL as a parameter, validate it and then create an alphanumeric alias
//	@Tags			slugs
//	@Accept			json
//	@Produce		json
//	@Param			url		query			string	true	"url to create slug from"
//	@Success		201		{object}	handlers.slugObject
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/shortenize [post]
func shortenize(ctx *gin.Context) {
	url2Process := ctx.Query("url")
	if url2Process == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no valid url found in your request"})
		return
	}

	checkTO, err := time.ParseDuration(viper.GetString("slugs.validation.url_timeout"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something went wrong while trying to parse the check timeout with value %q, %s", viper.GetString("slugs.validation.url_timeout"), err.Error())})
		return
	}

	if err := checkURL(url2Process, checkTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	obj := slugObject{URL: url2Process}

	db, err := sql.NewSQLHandler[slugObject]()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	retrieved, err := db.Read(obj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if retrieved != nil {
		ctx.JSON(http.StatusOK, gin.H{"url": retrieved.URL, "slug": retrieved.Slug})
		return
	} else {
		generatedSlug, err := generateSlug(obj.URL)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		obj.Slug = generatedSlug

		if err := db.Create(obj); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"url": obj.URL, "slug": obj.Slug})
		return
	}
}

func checkURL(url2Validate string, checkTimeout time.Duration) error {
	_, err := url.ParseRequestURI(url2Validate)
	if err != nil {
		return fmt.Errorf("%s is not a valid URL", url2Validate)
	}

	client := http.Client{Timeout: checkTimeout}

	if _, err = client.Get(url2Validate); err != nil {
		return fmt.Errorf("%s cannot be reached", url2Validate)
	}

	return nil
}

func generateSlug(url2Slug string) (string, error) {
	// use sha1 to prevent at maximum any collisions with the generations
	hasher := sha1.New()
	_, err := hasher.Write([]byte(url2Slug))
	if err != nil {
		return "", fmt.Errorf("unable to generate url slug for %s, %s", url2Slug, err.Error())
	}

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))[:viper.GetInt("slugs.chars_length")], nil
}

func (s slugObject) Get() string {
	return fmt.Sprintf("select * from %s where url='%s' or slug='%s'", viper.GetString("slugs.storage.table_name"), s.URL, s.Slug)
}
func (s slugObject) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (url, slug) VALUES(:url, :slug)", viper.GetString("slugs.storage.table_name"))
}
