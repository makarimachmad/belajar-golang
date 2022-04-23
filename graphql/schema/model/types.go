package model

import "github.com/graphql-go/graphql"

var ProductTypes = graphql.NewObject(graphql.ObjectConfig{
	Name:"Products",
	Fields:graphql.Fields{
		"id":&graphql.Field{
			Type:graphql.Int,
		},
		"product_name":&graphql.Field{
			Type:graphql.String,
		},
		"qte_stock":&graphql.Field{
			Type:graphql.Int,
		},
		"price":&graphql.Field{
			Type:graphql.Float,
		},
		"product_gambar":&graphql.Field{
			Type:graphql.String,
		},
	},
})