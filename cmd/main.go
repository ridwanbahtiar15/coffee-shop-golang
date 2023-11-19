package main

import (
	"coffee-shop-golang/internal/routers"
	"coffee-shop-golang/pkg"
	"log"
)

func main() {
	database, err := pkg.PostgreSQLDB()
	if(err != nil) {
		log.Fatal(err)
	}

	routers := routers.New(database)
	server := pkg.Server(routers)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}