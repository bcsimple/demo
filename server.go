package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			writer.WriteHeader(http.StatusOK)
			version := os.Getenv("APP_VERSION")
			if version == "" {
				version = "appv10"
			}
			if _, err := io.WriteString(writer, fmt.Sprintf("client_addr: %s,version: %s\n", request.RemoteAddr, version)); err != nil {
				log.Fatalln(err)
			}
		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}
	})

	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
