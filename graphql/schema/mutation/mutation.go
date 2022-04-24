package mutation

import "github.com/graphql-go/graphql"
import "BELAJAR-GOLANG/graphql/schema/model"

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name:"Mutation",
	Fields:graphql.Fields{
		"CreateProducts":&graphql.Field{
			Type:graphql.NewList(model.ProductTypes),
			Args:graphql.FieldConfigArgument{
				"id" : &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.String),
				},
				"product_name": &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.String),
				},
				"price": &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.Float),
				},
				"qte_stock": &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.Int),
				},
				"product_gambar": &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.String),
				},
			},
			Resolve:  CreateProdcutMutation,
		},
	},
})