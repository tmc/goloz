// Package apidocs provides a means of serving up SwaggerUI and docs.
//
// Example Usage:
//
// http.Handle("/apidocs", http.StripPrefix("/apidocs", http.FileServer(http.FileSystem(Content())))
package apidocs

import (
	"embed"
	"io/fs"
	"net/http"
)

// content holds our static web server content.
//go:embed swagger-ui/build/*
var content embed.FS

// swagger-ui/build/*

// Content returns the content suitable to serve up with
// http.FileServer(http.FileSystem(Content()))
func Content() fs.FS {
	f, err := fs.Sub(content, "swagger-ui/build")
	if err != nil {
		panic(err)
	}
	return f
}

// Handler is the http handler to serve apidocs content.
func Handler() http.Handler {
	return http.FileServer(http.FS(Content()))
}
