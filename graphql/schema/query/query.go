package query

import(
	"github.com/graphql-go/graphql"
	"BELAJAR-GOLANG/graphql/schema/model"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:"Query",
	Fields:graphql.Fields{
		"Products":&graphql.Field{
			Type:graphql.NewList(model.ProductTypes),
			//config param argument
			Args:graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: ProductResolve,
		},
		//for create other object
	},
})