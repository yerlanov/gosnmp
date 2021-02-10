package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"test/config"
	"test/controller"
)

func main() {
	conf := config.New()

	r := mux.NewRouter()
	r.HandleFunc("/operstatus", controller.GetOperStatus).Methods("POST")

	fmt.Println("Service running on :" + conf.Port)
	if err := http.ListenAndServe(":"+conf.Port, r); err != nil {
		panic(err)
	}
}
