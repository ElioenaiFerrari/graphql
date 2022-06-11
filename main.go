package main

import (
	"log"
	"net/http"

	"github.com/ElioenaiFerrari/super-api/api"
	"github.com/ElioenaiFerrari/super-api/schemas"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	fields := graphql.Fields{
		"ticker": &graphql.Field{
			Type:        schemas.TickerSchema,
			Description: "Get the ticker for a given symbol",
			Args: graphql.FieldConfigArgument{
				"asset": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				mb := api.NewMercadoBitcoinAPI()

				asset := p.Args["asset"].(string)

				ticker, err := mb.GetTicker(asset)
				if err != nil {
					log.Fatal(err)
				}

				return ticker, nil
			},
		},
		"trades": &graphql.Field{
			Type:        graphql.NewList(schemas.TradeSchema),
			Description: "Get the trades for a given symbol",
			Args: graphql.FieldConfigArgument{
				"asset": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				mb := api.NewMercadoBitcoinAPI()

				asset := p.Args["asset"].(string)

				trades, err := mb.GetTrades(asset)
				if err != nil {
					log.Fatal(err)
				}

				return trades, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	root := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/graphql", root)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
