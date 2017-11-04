package file

import (
	"bufio"
	"errors"
	"os"
	"strings"

	gq "github.com/graphql-go/graphql"
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
		num = 20
	}

	file, err := os.Open(filePath)
	if err != nil {
		return Prof{}, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return Prof{}, err
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) > num {
		lines = lines[len(lines)-num:]
	}
	content := strings.Join(lines, "\n")

	return Prof{
		Name:        filePath,
		UpdatedTime: stat.ModTime().String(),
		Content:     content,
	}, nil
}
