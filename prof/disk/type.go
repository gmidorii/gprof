package disk

import gq "github.com/graphql-go/graphql"

var Type = gq.NewObject(
	gq.ObjectConfig{
		Name: "Prof",
		Fields: gq.Fields{
			"io": &gq.Field{
				Type: ioType,
			},
			"usage": &gq.Field{
				Type: usageType,
			},
		},
	})

var ioType = gq.NewObject(
	gq.ObjectConfig{
		Name: "IO",
		Fields: gq.Fields{
			"read_count": &gq.Field{
				Type: gq.Int,
			},
		},
	})

var usageType = gq.NewObject(
	gq.ObjectConfig{
		Name: "Usage",
		Fields: gq.Fields{
			"path": &gq.Field{
				Type: gq.String,
			},
			"total": &gq.Field{
				Type: gq.Int,
			},
			"free": &gq.Field{
				Type: gq.Int,
			},
			"used": &gq.Field{
				Type: gq.Int,
			},
			"used_percent": &gq.Field{
				Type: gq.Float,
			},
		},
	})
