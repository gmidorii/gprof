package cpu

import "github.com/shirou/gopsutil/cpu"

type CPU struct {
	Count int
}

func Run() (CPU, error) {
	c, err := cpu.Counts(false)
	if err != nil {
		return CPU{}, err
	}
	cpu := CPU{
		Count: c,
	}
	return cpu, nil
}
