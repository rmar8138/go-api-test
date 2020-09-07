package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rmar8138/go-api-test/pkg/config"
)

// Index / http entrypoint
func Index(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, render.M{
		"serviceName": config.Configuration.Name,
	})
}
