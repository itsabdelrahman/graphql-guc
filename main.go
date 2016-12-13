package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemaConfig)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()["query"][0]
		result := graphql.Do(graphql.Params{
			Schema: schema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("Now server is running on port 8080...")
	fmt.Println("Test with GET: curl -g 'http://localhost:8080/graphql?query={hero{name}}'")
	http.ListenAndServe(":8080", nil)
}
