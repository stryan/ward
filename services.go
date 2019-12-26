package main

import (
	"encoding/json"

	"github.com/coreos/go-systemd/v22/dbus"
)

type Services struct {
	watched []string
	states  map[string]string
}

func (s *Services) Init() bool {
	s.watched = []string{"sshd.service", "brtfs.service", "mariadb.service"}
	s.states = make(map[string]string)
	for _, service := range s.watched {
		s.states[service] = "Unknown"
	}
	return true
}

func (s *Services) Path() string {
	return "services"
}

func (s *Services) Update() bool {
	c, err := dbus.NewSystemConnection()
	if err != nil {
		return false
	}
	un, err := c.ListUnits()
	if err != nil {
		return false
	}
	for _, service := range un {
		for _, w := range s.watched {
			//r, _ := regexp.Compile(w + "\\.(service|timer)")
			//if r.MatchString(service.Name) {
			if w == service.Name {
				s.states[service.Name] = service.ActiveState
			}
		}
	}
	c.Close()
	return true
}

func (s *Services) Output(args ...string) []byte {
	if len(args) == 0 {
		out, _ := json.Marshal(s.states)
		return out
	} else {
		out, _ := json.Marshal(s.states[args[0]])
		return out
	}
}
