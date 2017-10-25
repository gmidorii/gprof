package prof

import (
	gq "github.com/graphql-go/graphql"
	"github.com/midorigreen/gprof/prof/cpu"
)

var coreType = gq.NewObject(
	gq.ObjectConfig{
		Name: "Core",
		Fields: gq.Fields{
			"percent": &gq.Field{
				Type: gq.Float,
			},
		},
	})

var cpuType = gq.NewObject(
	gq.ObjectConfig{
		Name: "CPU",
		Fields: gq.Fields{
			"cores": &gq.Field{
				Type: gq.NewList(coreType),
			},
		},
	})

var rootQuery = gq.NewObject(
	gq.ObjectConfig{
		Name: "RoorQuery",
		Fields: gq.Fields{
			"cpu": &gq.Field{
				Type:        cpuType,
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
