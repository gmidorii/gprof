// Define GraphQL Type
package cpu

import gq "github.com/graphql-go/graphql"

var coreType = gq.NewObject(
	gq.ObjectConfig{
		Name: "Core",
		Fields: gq.Fields{
			"percent": &gq.Field{
				Type: gq.Float,
			},
		},
	})

var Type = gq.NewObject(
	gq.ObjectConfig{
		Name: "CPU",
		Fields: gq.Fields{
			"cores": &gq.Field{
				Type: gq.NewList(coreType),
			},
			"model": &gq.Field{
				Type: gq.String,
			},
			"model_name": &gq.Field{
				Type: gq.String,
			},
			"cache_size": &gq.Field{
				Type: gq.Int,
			},
		},
	})
