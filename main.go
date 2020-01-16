package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func setup() *State {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/ward/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %v \n", err)
	}
	viper.SetConfigType("yaml")
	state := newState()
	state.Delta = time.Duration(viper.GetInt("time_delta"))
	state.Port = viper.GetString("port")
	if viper.IsSet("services") && viper.GetBool("services.active") {
		s := &Services{}
		s.Init()
		state.registerModule(s)
	}
	if viper.IsSet("ip") && viper.GetBool("ip.active") {
		i := &Ip{}
		i.Init()
		state.registerModule(i)
	}
	if viper.IsSet("stats") && viper.GetBool("stats.active") {
		st := &Stats{}
		st.Init()
		state.registerModule(st)
	}
	if viper.IsSet("dnsmasq") && viper.GetBool("dnsmasq.active") {
		d := &Dnsmasq{}
		d.Init()
		state.registerModule(d)
	}
	for _, i := range state.modules {
		i.Update()
	}
	return state
}

func main() {
	log.Println("Placing Wards")
	state := setup()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", state.PrintRaw).Methods("GET")
	router.HandleFunc("/module/{name}", state.handleModule).Methods("GET")
	router.HandleFunc("/module/{name}/{output}", state.handleModuleOutput).Methods("GET")

	log.Println("Starting ward server")
	h := &http.Server{Addr: ":" + state.Port, Handler: router}

	go func() {
		if err := h.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stop

	log.Println("Removing Wards")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	h.Shutdown(ctx)

	log.Println("Wards removed")

}
