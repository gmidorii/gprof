package cpu

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CPU struct {
	Cores []Core
}

type Core struct {
	Percent float64
}

func Run() (CPU, error) {
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
