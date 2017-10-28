package cpu

import (
	"time"

	gq "github.com/graphql-go/graphql"
	"github.com/shirou/gopsutil/cpu"
)

type Prof struct {
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
		return Prof{}, err
	}

	cores := make([]Core, len(p))
	for i, v := range p {
		cores[i] = Core{
			Percent: v,
		}
	}

	infos, err := cpu.Info()
	if err != nil {
		return Prof{}, err
	}
	var prof Prof
	if len(infos) > 0 {
		prof.Model = infos[0].Model
		prof.ModelName = infos[0].ModelName
		prof.CacheSize = infos[0].CacheSize
	}
	prof.Cores = cores
	return prof, nil
}
