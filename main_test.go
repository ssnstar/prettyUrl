package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestShortUrlFail(t *testing.T) {

	apitest.New(). // configuration
			HandlerFunc(shortUrl).
			Post("http://localhost/shorturl").
			Expect(t).
			Status(http.StatusBadRequest).
			End()
}

func TestShortUrlOk(t *testing.T) {

	apitest.New(). // configuration
			HandlerFunc(shortUrl).
			Post("http://localhost/shorturl").
			Expect(t).
			Status(http.StatusOK).
			End()
}
func TestRedirect(t *testing.T) {

	apitest.New(). // configuration
			HandlerFunc(shortUrl).
			Get("http://localhost/redirect").
			ContentType("application/json").
			Expect(t).
			Status(http.StatusSeeOther).
			End()
}
