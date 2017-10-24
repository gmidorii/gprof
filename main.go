package main

import (
	"fmt"
	"log"

	"github.com/midorigreen/gprof/prof/cpu"
)

func run() error {
	cpu, err := cpu.Run()
	if err != nil {
		return err
	}
	fmt.Println(cpu)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
