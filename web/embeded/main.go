package main

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
)

var (
	//go:embed frontend
	content embed.FS

	//go:embed frontend/index.html
	index string
)

func main() {
	static, _ := fs.Sub(fs.FS(content), "frontend")

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, index)
	})
	http.Handle(
		"/static/", http.FileServer(http.FS(static)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
