package schemas

import "github.com/graphql-go/graphql"

var TradeSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Trade",
		Fields: graphql.Fields{
			"amount": &graphql.Field{
				Type: graphql.Float,
			},
			"date": &graphql.Field{
				Type: graphql.Int,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"tid": &graphql.Field{
				Type: graphql.Int,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
