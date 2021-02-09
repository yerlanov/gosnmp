package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"test/config"
	"test/controller"
)

type Conf struct {
	Port      string `envconfig:"PORT" required:"true"`
	Community string `envconfig:"COMMUNITY"`
	MibDir    string `envconfig:"MIBDIR"`
	IpAddress string `envconfig:"IP"`
}

func main() {
	conf := config.New()

	r := mux.NewRouter()
	r.HandleFunc("/operStatus", controller.GetOperStatus).Methods("POST")

	fmt.Println("Service running on :" + conf.Port)
	if err := http.ListenAndServe(":"+conf.Port, r); err != nil {
		panic(err)
	}
}
