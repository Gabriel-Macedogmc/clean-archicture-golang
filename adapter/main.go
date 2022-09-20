package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Gabriel-Macedogmc/clean-archicture-golang/adapter/postgres"
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/di"
	"github.com/gorilla/mux"
)

func main() {
	appMode := os.Getenv("APP_ENV")

	log.Println(appMode)

	databaseUrl := os.Getenv("DATABASE_URL")

	log.Println(databaseUrl)

	conn := postgres.ConnectToDatabase(databaseUrl)
	postgres.Migrate()

	productService := di.ConfigProductDI(conn)

	router := mux.NewRouter()
	router.Handle("/product", http.HandlerFunc(productService.Create)).Methods("POST")
	router.Handle("/products", http.HandlerFunc(productService.List)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET")

	port := os.Getenv("PORT")
	log.Printf("LISTEN ON PORT: %v", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
