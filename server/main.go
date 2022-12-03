package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, os.Getenv("POD_NAME"))
	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
