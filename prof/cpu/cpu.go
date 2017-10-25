package cpu

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/shirou/gopsutil/cpu"
)

type CPU struct {
	Cores []Core `json:"cores"`
}

type Core struct {
	Percent float64 `json:"percent"`
}

func Resolve(params graphql.ResolveParams) (interface{}, error) {
	p, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		return CPU{}, err
	}
	cores := make([]Core, len(p))
	for i, v := range p {
		cores[i] = Core{v}
	}

	return CPU{
		Cores: cores,
	}, nil
}
