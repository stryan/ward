package main

import (
	"encoding/json"
)

type Dummy struct {
	text string
}

func (d *Dummy) Init() bool {
	d.text = "Dummy Module"
	return true
}

func (d *Dummy) Path() string {
	return "dummy"
}

func (d *Dummy) Update() bool {
	return true
}

func (d *Dummy) Output(args ...string) []byte {
	out, _ := json.Marshal(d.text)
	return out
}
