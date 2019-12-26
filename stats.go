package main

import (
	"encoding/json"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

type Stats struct {
}

func (s *Stats) Init() bool {
	return true
}

func (s *Stats) Path() string {
	return "stats"
}

func (s *Stats) Update() bool {
	return true
}

func (s *Stats) Output(args ...string) []byte {
	h, _ := host.Info()
	m, _ := mem.VirtualMemory()
	l, _ := load.Avg()
	outh, _ := json.Marshal(h)
	outm, _ := json.Marshal(m)
	outl, _ := json.Marshal(l)
	m1 := append(outh, outm...)
	out := append(m1, outl...)
	return out
}
