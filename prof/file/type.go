package file

import gq "github.com/graphql-go/graphql"

var Type = gq.NewObject(
	gq.ObjectConfig{
		Name: "Prof",
		Fields: gq.Fields{
			"name": &gq.Field{
				Type: gq.String,
			},
		},
	})
