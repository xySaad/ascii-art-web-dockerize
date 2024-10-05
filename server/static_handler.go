package server

import (
	"net/http"
	"os"
)

func StaticHandler(res http.ResponseWriter, req *http.Request) {

	fs := http.FileServer(http.Dir("./static"))
	// Attempt to open the requested file
	_, err := os.Stat("./static" + req.URL.Path)
	if os.IsNotExist(err) {
		// File doesn't exist, serve 404.html
		http.ServeFile(res, req, "./static/pages/404.html")
		return
	}
	fs.ServeHTTP(res, req)
}
