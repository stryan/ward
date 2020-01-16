package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

//1579224183 d4:25:8b:33:d1:78 192.168.1.160 xanadu *
type Dnshost struct {
	Name string
	Mac  string
	IP   string
	ID   string
}

type Dnsmasq struct {
	Hosts []Dnshost
}

func (d *Dnsmasq) Init() bool {
	d.Hosts = make([]Dnshost, 5)
	return true
}

func (d *Dnsmasq) Path() string {
	return "dnsmasq"
}

func (d *Dnsmasq) Update() bool {
	file, err := os.Open("test.txt")

	if err != nil {
		log.Printf("Dnsmasq:failed opening file: %s", err)
		return false
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	var tmphosts []Dnshost
	for _, line := range txtlines {
		line_s := strings.Split(line, " ")
		tmphosts = append(tmphosts, Dnshost{line_s[3], line_s[1], line_s[2], line_s[4]})
	}
	d.Hosts = tmphosts
	return true
}

func (d *Dnsmasq) Output(args ...string) []byte {
	var out []byte
	for _, m := range d.Hosts {
		tmp, _ := json.Marshal(m)
		out = append(out, tmp...)
	}
	return out
}
