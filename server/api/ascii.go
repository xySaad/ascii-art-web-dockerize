package api

import (
	"ascii-art-web/ascii"
	"ascii-art-web/utils"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Ascii(res http.ResponseWriter, req *http.Request) {
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
	asciiText, err := ascii.Generate(args)
	if err != nil {
		http.Error(res, "500 - internal server error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte(asciiText))
}
