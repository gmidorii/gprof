package file

import (
	"fmt"

	gq "github.com/graphql-go/graphql"
)

type Prof struct {
	Name        string `json:"name"`
	UpdatedTime string `json:"updated_time"`
	Content     string `json:"content"`
}

func Resolve(params gq.ResolveParams) (interface{}, error) {
	filePath := params.Args["path"]
	fmt.Println(filePath)

	return Prof{}, nil
}
