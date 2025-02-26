# Slugifyer API

This simple API aims to shorten URL into little alphanum string and then process to redirection when a valid slug is called.

## Routes

`[POST] /shortenize?url=<url>` will convert the url into a slug

`[GET] /:slug` will redirect to the url aliased by the slug specified

## Run

Ensure `docker-compose` is installed locally and simply perfor `make run`