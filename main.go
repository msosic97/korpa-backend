package main

// Program starts and finish here

import (
	"log"
	"fmt"
	"net/http"
	"github.com/msosic97/korpa-backend/tools/db"
	"github.com/msosic97/korpa-backend/routes"
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

	r := routes.InitializeRouter()
	fmt.Print("Server is running on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))

}
