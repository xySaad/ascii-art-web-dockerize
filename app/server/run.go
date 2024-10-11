package server

import (
	"ascii-art-web/server/api"
	"fmt"
	"net/http"
)

func Run(adress, port string) error {
	temp := adress

	if temp == "0.0.0.0" {
		temp = "http://localhost"
	}

	fmt.Println("server is running on:", temp+":"+port)

	http.HandleFunc("/", StaticHandler)
	http.HandleFunc("/api/v1/ascii/", api.Ascii)

	err := http.ListenAndServe(adress+":"+port, nil)
	if err != nil {
		return err
	}
	return nil
}
