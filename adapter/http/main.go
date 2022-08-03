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

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/fanticheli/rest-api-clean-go/adapter/http/docs"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Clean GO API Docs
// @version 1.0.0
// @contact.name Igor Fanticheli
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	productService := di.ConfigProductDI(conn)

	router := mux.NewRouter()

	jsonApiRouter := router.PathPrefix("/").Subrouter()
	jsonApiRouter.Use(middleware.Cors)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	jsonApiRouter.Handle("/product", http.HandlerFunc(productService.Create)).Methods("POST", "OPTIONS")
	jsonApiRouter.Handle("/product", http.HandlerFunc(productService.Fetch)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET", "OPTIONS")

	port := viper.GetString("server.port")

	if port == "" {
		port = os.Getenv("PORT")
	}
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
