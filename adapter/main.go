package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gabriel-Macedogmc/clean-archicture-golang/adapter/postgres"
	"github.com/Gabriel-Macedogmc/clean-archicture-golang/di"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func main() {
	databaseUrl := viper.GetString("database.url")
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

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
