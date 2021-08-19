package gql

import "github.com/graphql-go/graphql"

var Cart = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Cart",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"totprice": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)