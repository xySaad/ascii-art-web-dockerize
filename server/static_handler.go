package server

import (
	"fmt"
	"net/http"
	"os"
)

func StaticHandler(res http.ResponseWriter, req *http.Request) {

	fs := http.FileServer(http.Dir("./static"))

	// Attempt to open the requested file
	_, err := os.Stat("./static" + req.URL.Path)

	if os.IsNotExist(err) {
		// Set the 404 status code
		res.WriteHeader(http.StatusNotFound)

		// Read the 404 page content and write it to the response
		page404, err := os.ReadFile("./static/pages/404.html")
		if err != nil {
			http.Error(res, "404 page not found", http.StatusNotFound)
			return
		}
		// Write the 404 page content to the response
		_, err = res.Write(page404)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	if req.Method != "GET" {
		http.Error(res, "405 - method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Serve the file if found
	fs.ServeHTTP(res, req)
}
