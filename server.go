package main

import (
	"github.com/gaillard/graphql-example/schema"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	http.Handle("/query", handler.GraphQL(schema.NewExecutableSchema(schema.Config{Resolvers: &schema.Resolver{}})))
	log.Printf("connect to http://localhost:%s/query", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
