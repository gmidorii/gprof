package prof

import (
	gq "github.com/graphql-go/graphql"
	"github.com/midorigreen/gprof/prof/cpu"
	"github.com/midorigreen/gprof/prof/disk"
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
			"disk": &gq.Field{
				Type:        disk.Type,
				Description: "prof disk",
				Resolve:     disk.Resolve,
			},
		},
		Description: "Root",
	})

var Schema, _ = gq.NewSchema(
	gq.SchemaConfig{
		Query: rootQuery,
	},
)
