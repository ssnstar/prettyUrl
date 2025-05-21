package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestShortUrlGet(t *testing.T) {

	apitest.New(). // configuration
			HandlerFunc(shortUrl).
			Get("/shorturl").
			Expect(t).
			Status(http.StatusMethodNotAllowed).
			End()
}

func TestShortUrlFail(t *testing.T) {

	apitest.New(). // configuration
			HandlerFunc(shortUrl).
			Post("/shorturl").
			Expect(t).
			Status(http.StatusBadRequest).
			End()
}

func TestShortUrlOk(t *testing.T) {

	apitest.New(). // configuration
			HandlerFunc(shortUrl).
			Post("/shorturl").
			Body(`{"url": "http://www.youtube.com"}`).
			ContentType("application/json").
			Expect(t).
			Status(http.StatusOK).
			End()
}
func TestRedirectPost(t *testing.T) {

	apitest.New(). // configuration
			HandlerFunc(redirect).
			Post("/redirect").
			Body(`{"url": "e62e2446"}`).
			ContentType("application/json").
			Expect(t).
			Status(http.StatusSeeOther).
			End()
}

func TestRedirectGet(t *testing.T) {

	apitest.New(). // configuration
			HandlerFunc(redirect).
			Get("/redirect").
			ContentType("application/json").
			Expect(t).
			Status(http.StatusMethodNotAllowed).
			End()
}
