package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/graphql-services/idp"
	"github.com/graphql-services/idp/database"
)

const defaultPort = "80"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	urlString := os.Getenv("DATABASE_URL")
	if urlString == "" {
		panic(fmt.Errorf("database url must be provided"))
	}

	db := database.NewDBWithString(urlString)
	defer db.Close()
	db.AutoMigrate(&idp.User{})

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(idp.NewExecutableSchema(idp.Config{Resolvers: &idp.Resolver{DB: db}})))

	http.HandleFunc("/healthcheck", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("OK"))
		res.WriteHeader(200)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
