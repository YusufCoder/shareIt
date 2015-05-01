package main

import (
	"log"
	"os"
	"net/http"
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	dataDir   = kingpin.Arg("data", "Path to data direcory").Required().String()
	adminPW   = kingpin.Arg("password", "Admin password").Required().String()
)

func main() {

	//Parse command line params
	kingpin.Version("0.1")
  	kingpin.Parse()

  	//Check data direcory
  	if _, err := os.Stat(*dataDir); os.IsNotExist(err) {
  		log.Fatalf("Data direcory doesn't exist")
  	}

  	//Start server
	log.Printf("Starting server...")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":9000", router))
}
