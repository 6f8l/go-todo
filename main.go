package main

import (
	"fmt"
	"log"
	"net/http"

	"./handler"
)

func main() {
	mux := handler.SetUpRouting()

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
