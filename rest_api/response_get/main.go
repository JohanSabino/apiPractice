package main

import (
	"github.com/JohanSabino/apiPractice/rest_api/server"
	"log"
)

func main() {
	srv := server.New(":8080")

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Fatal error")
	}
}
