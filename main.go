package main

import (
	"fmt"
	"json-parser/jsonhandler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/parse", jsonhandler.ParseJSONHandler)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("-----JSON PARSER-----")

}
