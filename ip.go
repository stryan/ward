package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Ip struct {
	ips map[string]string
}

func (i *Ip) Init() bool {
	i.ips = make(map[string]string)
	i.ips["external"] = ""
	i.ips["internal"] = ""
	return true
}

func (i *Ip) Path() string {
	return "ip"
}

func (i *Ip) Update() bool {
	resp, err := http.Get("https://ifconfig.co")
	if err != nil {
		log.Printf("TODO: Handle error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("TODO: Handle body read error")
	}
	i.ips["internal"] = "192.168.1.50"
	i.ips["external"] = strings.TrimSpace(string(body))
	return true
}

func (i *Ip) Output(args ...string) []byte {
	if len(args) == 0 {
		out, _ := json.Marshal(i.ips)
		return out
	} else {
		out, _ := json.Marshal(i.ips[args[0]])
		return out
	}
}
