package file

import (
	"bufio"
	"errors"
	"os"
	"strings"

	gq "github.com/graphql-go/graphql"
)

type form struct {
	Path string
	Num  int
}

func (f *form) parse(args map[string]interface{}) error {
	path, ok := args["path"].(string)
	if !ok {
		return errors.New("invalid format 'path'")
	}
	num, ok := args["num"].(int)
	if !ok {
		num = 20
	}

	f.Path = path
	f.Num = num
	return nil
}

type Prof struct {
	Name        string `json:"name"`
	UpdatedTime string `json:"updated_time"`
	Content     string `json:"content"`
}

func Resolve(params gq.ResolveParams) (interface{}, error) {
	f := form{}
	if err := f.parse(params.Args); err != nil {
		return Prof{}, err
	}

	file, err := os.Open(f.Path)
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
	if len(lines) > f.Num {
		lines = lines[len(lines)-f.Num:]
	}
	content := strings.Join(lines, "\n")

	return Prof{
		Name:        f.Path,
		UpdatedTime: stat.ModTime().String(),
		Content:     content,
	}, nil
}
