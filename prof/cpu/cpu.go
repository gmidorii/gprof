package cpu

import (
	"time"

	gq "github.com/graphql-go/graphql"
	"github.com/shirou/gopsutil/cpu"
)

type CPU struct {
	Cores     []Core `json:"cores"`
	Model     string `json:"model"`
	ModelName string `json:"model_name"`
	CacheSize int32  `json:"cache_size"`
}

type Core struct {
	Percent float64 `json:"percent"`
}

func Resolve(params gq.ResolveParams) (interface{}, error) {
	p, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		return CPU{}, err
	}

	cores := make([]Core, len(p))
	for i, v := range p {
		cores[i] = Core{
			Percent: v,
		}
	}

	infos, err := cpu.Info()
	if err != nil {
		return CPU{}, err
	}
	var cpu CPU
	if len(infos) > 0 {
		cpu.Model = infos[0].Model
		cpu.ModelName = infos[0].ModelName
		cpu.CacheSize = infos[0].CacheSize
	}
	cpu.Cores = cores
	return cpu, nil
}
