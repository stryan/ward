package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

func main() {
	state := newState()
	s := &Services{}
	i := &Ip{}
	s.Init()
	s.Update()
	i.Init()
	i.Update()
	state.registerModule(s)
	state.registerModule(i)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", state.PrintRaw).Methods("GET")
	router.HandleFunc("/module/{name}", state.handleModule).Methods("GET")
	router.HandleFunc("/module/{name}/{output}", state.handleModuleOutput).Methods("GET")
	http.ListenAndServe(":8000", router)
}
