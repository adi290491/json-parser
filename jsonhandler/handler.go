package jsonhandler

import (
	"encoding/json"
	"fmt"
	"io"
	"json-parser/lexer"
	"json-parser/parser"
	"json-parser/token"
	"log"
	"net/http"
)

type RequestPayload struct {
	Json string `json:"json"`
}

type ResponsePayload struct {
	Status    string `json:"status"`
	Formatted string `json:"formatter,omitempty"`
	Error     string `json:"error,omitempty"`
}

func ParseJSONHandler(w http.ResponseWriter, r *http.Request) {

	var jsonData string
	var err error
	token.Tokens = nil
	r.ParseMultipartForm(10 << 20)

	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")

	if err == nil {
		defer file.Close()
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Unable to read uploaded file", http.StatusBadRequest)
			return
		}
		jsonData = string(fileBytes)
	} else {
		jsonData = r.FormValue("json")
		if jsonData == "" {
			http.Error(w, "No JSON input provided", http.StatusBadRequest)
			return
		}
	}

	lexer := lexer.NewLexer(jsonData)
	err = lexer.Run()

	if err != nil {
		// Return response if lexer fails
		response := ResponsePayload{
			Status: "failure",
			Error:  fmt.Sprintf("Lexing error: %v", err),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	parser := parser.NewParser(token.Tokens)
	result, err := parser.Parse()
	log.Println("Parsing completed...")
	var response ResponsePayload

	if err != nil {
		response = ResponsePayload{
			Status: "failure",
			Error:  fmt.Sprintf("Parsing error: %v", err),
		}
	} else {

		formattedoutput, err := json.MarshalIndent(result, "", "	")
		if err != nil {
			response = ResponsePayload{
				Status: "failure",
				Error:  "Error formatting JSON",
			}
		} else {
			response = ResponsePayload{
				Status:    "success",
				Formatted: string(formattedoutput),
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
