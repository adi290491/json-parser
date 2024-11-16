# JSON-PARSER
This is a JSON parser implemented in Go and working behind a server via HTTP. It supports both raw JSON strings and file uploads. I have implemented a custom lexer to tokenize the input and a parser that parses the tokens. Finally the program writes the status as well as formatted output or error depending on the result, to the response writer.

## Features
- Handles JSON input via HTTP POST requests.
- Supports both JSON strings and file uploads (multipart/form-data).
- Custom-built lexer and parser for JSON with support for:
  Strings (including escaped characters).
  Numbers (including negative and scientific notation).
  Booleans (true / false) and null.
  Arrays and nested objects.
- Provides detailed error messages for invalid JSON input.

## Usage
### Running the Server
To start the server, run:
```bash
go run main.go
```
By default, the server runs on http://localhost:8080

### API Endpoint
#### 1. POST /parse (JSON string)

##### Example Request
```bash
curl -X POST http://localhost:8080/parse \
-H "Content-Type: multipart/form-data" \
-F "json={\"name\":\"Alice\",\"age\":30}"
```
##### Example Response
```json
{
  "status": "success",
  "formatted": "{\n  \"name\": \"Alice\",\n  \"age\": 30\n}"
}
```

#### 2. POST /parse (File upload)
##### Example Request
```bash
curl -X POST http://localhost:8080/parse \
-F "file=@input.json"
```
##### Example Response
```json
{
  "status": "success",
  "formatted": "{\n  \"city\": \"NYC\",\n  \"temperature\": 75\n}"
}
```

### Error Handling
The server will send a response with an error message depending on the invalidity of JSON string.

##### Unterminated string
```json
{
  "status": "failure",
  "error": "Lexing error: unterminated string at position 14"
}
```

##### Invalid number format
```json
{
  "status": "failure",
  "error": "Lexing error: invalid number format at position 12"
}
```
##### Unexpected character
```json
{
  "status": "failure",
  "error": "unexpected character: @ at position 7"
}
```
