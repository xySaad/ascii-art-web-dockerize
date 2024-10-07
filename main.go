package main

import (
	"ascii-art-web/server"
	"fmt"
	"os"
)

var (
	Port   = "8080"
	Adress = "0.0.0.0"
)

func main() {
	err := server.Run(Adress, Port)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
