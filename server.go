package main

import (
	"log"
	"myapp/auth"
	"myapp/config"
	"myapp/dataloader"
	"myapp/directive"
	"myapp/graph"
	"myapp/graph/generated"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(auth.AuthMiddleware)
	router.Use(dataloader.LoaderMiddleware)

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.HasRole = directive.HasRole

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	handler := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
	}).Handler(srv)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main1() {
	config.MigrateSql()
}
