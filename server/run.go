package server

import (
	"net/http"
)

func Run(adress, port string) error {
	http.HandleFunc("/", StaticHandler)

	err := http.ListenAndServe(adress+":"+port, nil)
	if err != nil {
		return err
	}
	return nil
}
