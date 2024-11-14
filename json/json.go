package json

type JsonValue interface{}

type JsonObject map[string]JsonValue

type JsonArray []JsonValue

type JsonString string

type JsonNumber float64

type JsonBoolean bool

type JsonNull struct{}
