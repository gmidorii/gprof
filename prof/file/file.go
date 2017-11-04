package file

import (
	"errors"
	"fmt"
	"strings"

	gq "github.com/graphql-go/graphql"
	"github.com/hpcloud/tail"
)

type Prof struct {
	Name        string `json:"name"`
	UpdatedTime string `json:"updated_time"`
	Content     string `json:"content"`
}

func Resolve(params gq.ResolveParams) (interface{}, error) {
	filePath, ok := params.Args["path"].(string)
	if !ok {
		return Prof{}, errors.New("invalid format 'path'")
	}
	num, ok := params.Args["num"].(int)
	if !ok {
		return Prof{}, errors.New("invalid format 'num'")
	}
	fmt.Println(filePath)

	tail, err := tail.TailFile(filePath, tail.Config{})
	if err != nil {
		return Prof{}, err
	}

	lines := make([]string, num)
	i := 0
	for line := range tail.Lines {
		if i >= num {
			break
		}
		lines[i] = line.Text
		i++
	}
	content := strings.Join(lines, "\n")

	return Prof{
		Name:    filePath,
		Content: content,
	}, nil
}
