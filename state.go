package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type State struct {
	modules map[string]Module
	Delta   time.Duration
	Port    string
}

func newState() *State {
	modMap := make(map[string]Module)
	m := &State{
		modules: modMap,
	}
	return m
}
func (s *State) registerModule(m Module) {
	s.modules[m.Path()] = m
}

func (s *State) handleModule(w http.ResponseWriter, req *http.Request) {
	modName := mux.Vars(req)["name"]
	w.Header().Set("Content-Type", "application/json")
	if val, ok := s.modules[modName]; ok {
		w.Write(val.Output())
		return
	}
	http.NotFound(w, req)
}

func (s *State) handleModuleOutput(w http.ResponseWriter, req *http.Request) {
	modName := mux.Vars(req)["name"]
	outName := mux.Vars(req)["output"]
	w.Header().Set("Content-Type", "application/json")
	if val, ok := s.modules[modName]; ok {
		w.Write(val.Output(outName))
		return
	}
	http.NotFound(w, req)
}

func (s *State) PrintRaw(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, i := range s.modules {
		w.Write(i.Output())
		fmt.Fprintf(w, "\n")
	}
}

func (s *State) UpdateChecks() {
	for true {
		time.Sleep(s.Delta * time.Second)
		for _, i := range s.modules {
			i.Update()
		}
	}
}
