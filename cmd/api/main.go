package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crisnet-dev/fastfood/internal/config"
	"github.com/crisnet-dev/fastfood/internal/database"
	"github.com/crisnet-dev/fastfood/internal/routes"
)

func main() {

	config.LoadConfig()
	env := config.GetEnv()

	if err := database.SetUpDB(); err != nil {
		log.Println("\033[31mERROR TO CONNECT TO THE DATABASE!\033[0m")
		log.Fatal(err)
	}
	log.Println("\033[32mCONNECTED TO THE DATABASE!\033[0m")

	r := routes.SetupRoutes()

	log.Printf("Server listening in http://%s:%s\n", env.Host, env.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", env.Host, env.Port), r); err != nil {
		log.Fatal(err)
	}

}
