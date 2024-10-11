package api

import (
	"ascii-art-web/ascii"
	"ascii-art-web/utils"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Ascii(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/api/v1/ascii/banners/", "/api/v1/ascii/banners":
		if req.Method != "GET" {
			http.Error(res, "405 - method not allowed", http.StatusMethodNotAllowed)
			return
		}
		banners, err := utils.GetBanners()
		if err != nil {
			http.Error(res, "500 - internal server error", http.StatusInternalServerError)
			return
		}
		res.WriteHeader(200)
		res.Write([]byte(banners))
	case "/api/v1/ascii/":
		if req.Method != "POST" {
			http.Error(res, "405 - method not allowed", http.StatusMethodNotAllowed)
			return
		}
		rawBody := ""
		for {
			buffer := make([]byte, 1)
			_, err := req.Body.Read(buffer)
			rawBody += string(buffer)
			if err == io.EOF {
				break
			}
		}
		body := strings.Split(rawBody, "&")

		if len(body[0]) < 6 || body[0][:6] != "input=" || len(body[1]) < 7 || body[1][:7] != "banner=" {
			http.Error(res, "400 - bad request", http.StatusBadRequest)
			return
		}
		var err error
		body[0], err = url.QueryUnescape(body[0][6:])
		if err != nil {
			http.Error(res, "400 - bad request", http.StatusBadRequest)
			return
		}
		body[1], err = url.QueryUnescape(body[1][7:])
		if err != nil {
			http.Error(res, "400 - bad request", http.StatusBadRequest)
			return
		}
		if len([]rune(body[0])) > 500 {
			http.Error(res, "413 - payload too large", http.StatusRequestEntityTooLarge)
			return
		}
		args := ascii.Args{Text: body[0], BannerName: body[1]}
		if !utils.IsValidBanner(args.BannerName) {
			http.Error(res, "400 - bad request", http.StatusBadRequest)
			return
		}
		asciiData, err := ascii.Generate(args)
		if err != nil {
			http.Error(res, "500 - internal server error", http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte("{" + "\"ascii\": " + "\"" + utils.Escape(asciiData.Value) + "\"" + ", \"message\": " + "\"" + asciiData.Message + "\"" + "}"))
	default:
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

}
