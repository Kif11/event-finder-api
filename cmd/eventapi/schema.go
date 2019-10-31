package main

import (
	"errors"

	"github.com/graphql-go/graphql"
)

// BuildSchema for GraphQL
func BuildSchema(s *Server) (graphql.Schema, error) {
	eventObject := graphql.NewObject(graphql.ObjectConfig{
		Name: "Event",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"time": &graphql.Field{
				Type: graphql.String,
			},
			"lat": &graphql.Field{
				Type: graphql.String,
			},
			"lon": &graphql.Field{
				Type: graphql.String,
			},
			"link": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	// Schema
	fields := graphql.Fields{
		"events": &graphql.Field{
			Type: graphql.NewList(eventObject),
			Args: graphql.FieldConfigArgument{
				"lat": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
				"lon": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
				"totalTime": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"distance": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: s.EventsResolver,
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return schema, errors.New("schema: " + err.Error())
	}

	return schema, nil
}
