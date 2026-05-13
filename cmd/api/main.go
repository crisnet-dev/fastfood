package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crisnet-dev/fastfood/internal/config"
	"github.com/crisnet-dev/fastfood/internal/routes"
)

func main() {

	config.LoadConfig()
	env := config.GetEnv()
	if env.Host == "" || env.Port == "" {
		log.Fatal("Verefied your .env file")
	}

	r := routes.SetupRoutes()

	log.Printf("Server listening in http://%s:%s\n", env.Host, env.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", env.Host, env.Port), r); err != nil {
		log.Fatal(err)
	}

}
