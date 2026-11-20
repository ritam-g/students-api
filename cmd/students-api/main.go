package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/ritam-g/students-api/internal"
)

func main() {
	// load config 

	cfg:= config.MustLoad()
	//db setup 
	//router setup
	router := http.NewServeMux() 

	router.HandleFunc("GET /",func(w http.ResponseWriter,req *http.Request){
		w.Write([]byte("hellow this is great api "))
	})
	// server setup 

	server:= http.Server{
		Addr: cfg.HttpServer.Address,
		Handler: router,
	}
	fmt.Printf("server  serted on port no %s",cfg.HttpServer.Address)
	err := server.ListenAndServe()

	if 	err !=nil {
		log.Fatal("fails to log server",err.Error())
	}



}