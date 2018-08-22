package main

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	schemaLoader := gojsonschema.NewStringLoader(schemaString)

	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		panic(err)
	}

	for idx, response := range responses {
		loader := gojsonschema.NewStringLoader(response)
		if valid, err := schema.Validate(loader); err != nil {
			fmt.Printf("#%d json vaildation failed with %s on %s\n", idx, err, response)
		} else if !valid.Valid() {
			fmt.Printf("#%d json vaildation failed with %s on %s\n", idx, valid.Errors(), response)
		} else {
			fmt.Printf("#%d json vaildation succeed\n", idx)
		}
	}
}

/*
# example of a full response
{
	"value": {[
	{"key_name", "key1", "value": true},
	{"key_name", "key2", "value": false},
	]},
}
*/

const schemaString = `
{
	"$schema": "http://json-schema.org/draft-04/schema#",
	"type": "object",
	"required": [
		"value"
	],
	"properties": {
		"value": {
			"type": "array",
			"items": {
				"type": "object",
				"required": [
					"name",
					"value"
				],
				"properties": {
					"name": {
						"type": "string"
					},
					"value": {
						"type": "boolean"
					}
				}
			}
		}
	}
}
  `

var responses [2]string = [2]string{
	`{"value":[{"name":"key1","value":true}, {"name":"key2","value":true}]}`,
	`{"value":[{"name":"key1","value":true}, {"name":"key2","value":[true]}]}`,
}
