package server

import (
	"ascii-art-web/server/api"
	"net/http"
)

func Run(adress, port string) error {
	http.HandleFunc("/", StaticHandler)
	http.HandleFunc("/api/ascii", api.Ascii)

	err := http.ListenAndServe(adress+":"+port, nil)
	if err != nil {
		return err
	}
	return nil
}
