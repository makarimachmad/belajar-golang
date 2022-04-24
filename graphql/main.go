package main

import (
	"encoding/json"
	"fmt"
	"hello/graphql/schema/mutation"
	"hello/graphql/schema/query"
	"net/http"

	"github.com/graphql-go/graphql"
)

func executeeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: query,
	})
	if result.Errors != nil {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	var Shcema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: query.RootQuery,
		Mutation: mutation.Mutation,
	})
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeeQuery(r.URL.Query().Get("Query"), Shcema)
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8080", nil)
}