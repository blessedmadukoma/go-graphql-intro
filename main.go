package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	fmt.Println("GRAPHQL Tutorial")

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "World", nil
			},
		},
	}

	// define the object config
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	// define the schema config
	schemaQuery := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	// create our schema
	schema, err := graphql.NewSchema(schemaQuery)
	if err != nil {
		log.Fatal("Failed to create GraphQL schema:", err)
	}

	query := `
			{
				hello
			}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len((r.Errors)) > 0 {
		log.Fatal("Failed to run graphql operation:", r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

}
