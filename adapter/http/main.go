package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fanticheli/rest-api-clean-go/adapter/http/rest/middleware"
	"github.com/fanticheli/rest-api-clean-go/adapter/postgres"
	"github.com/fanticheli/rest-api-clean-go/di"
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
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	productService := di.ConfigProductDI(conn)

	router := mux.NewRouter()

	jsonApiRouter := router.PathPrefix("/").Subrouter()
	jsonApiRouter.Use(middleware.Cors)

	jsonApiRouter.Handle("/product", http.HandlerFunc(productService.Create)).Methods("POST", "OPTIONS")

	port := viper.GetString("server.port")

	if port == "" {
		port = os.Getenv("PORT")
	}
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
