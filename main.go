package main

// Program starts and finish here

import (
	"log"
	"github.com/msosic97/korpa-backend/Tools/DB"
	"github.com/msosic97/korpa-backend/config"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.StartDB(config)
	if err != nil {
		log.Fatal(err)
	}
	// defer keyword is used to delay the execution of a function until the surrounding function completes.
	defer db.Close()

}
