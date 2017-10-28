package disk

import (
	"fmt"
	"math"

	gq "github.com/graphql-go/graphql"
	"github.com/shirou/gopsutil/disk"
)

type Prof struct {
	IO    IO    `json:"io"`
	Usage Usage `json:"usage"`
}

type IO struct {
	ReadCount uint64 `json:"read_count"`
}
type Usage struct {
	Path        string  `json:"path"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

func Resolve(params gq.ResolveParams) (interface{}, error) {
	stat, err := disk.Usage("/Users/midori")
	if err != nil {
		return Prof{}, err
	}

	fmt.Println(stat.Total)
	fmt.Println(uint64(math.MaxInt32))
	return Prof{
		IO: IO{
			ReadCount: 10,
		},
		Usage: Usage{
			Path:        stat.Path,
			Total:       convBToMB(stat.Total),
			Free:        convBToMB(stat.Free),
			Used:        convBToMB(stat.Used),
			UsedPercent: stat.UsedPercent,
		},
	}, nil
}

func convBToMB(b uint64) uint64 {
	return b / 1024 / 1024
}
