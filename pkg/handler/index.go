package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

// Index / http entrypoint
func Index(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, render.M{
		"hello": "world",
	})
}
