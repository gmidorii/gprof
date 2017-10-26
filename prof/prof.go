package prof

import (
	gq "github.com/graphql-go/graphql"
	"github.com/midorigreen/gprof/prof/cpu"
)

var rootQuery = gq.NewObject(
	gq.ObjectConfig{
		Name: "RoorQuery",
		Fields: gq.Fields{
			"cpu": &gq.Field{
				Type:        cpu.Type,
				Description: "List of cpus",
				Resolve:     cpu.Resolve,
			},
		},
		Description: "Root",
	})

var Schema, _ = gq.NewSchema(
	gq.SchemaConfig{
		Query: rootQuery,
	},
)
