package main

import (
	"config"
	"direct"
	"filter"
	"log"
	"net/http"
)

func main() {
	// Some initialization first
	// Loading the config for the ListeningPort and TargetHost URL
	config.LoadConfig()
	// LoadFilters will pull in all the Regex's from the filters.json
	// file and pre-compile them for use
	filter.LoadFilters("./")

	// Create a handler for all incoming traffic: send it to DirectIt
	http.HandleFunc("/", direct.DirectIt)
	log.Println("Starting server on port:" + config.MyConfig.ListeningPort)
	http.ListenAndServe(config.MyConfig.ListeningPort, nil)
}
