package main

import (
	"bytes"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/damianino/gradient-graphql/database"
	"github.com/damianino/gradient-graphql/gradient"
	"github.com/damianino/gradient-graphql/graph"
	"github.com/damianino/gradient-graphql/graph/generated"
)

const defaultPort = "3000"
var db *database.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db = database.Connect()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	
	http.Handle("/view/", View())

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func View() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path, "/")
		log.Println(path)
		id:= path[len(path)-1]
		
		g, err := db.FindById(id)
		if err != nil{
			w.WriteHeader(http.StatusNotFound)
			return
		}
		m, err := gradient.CreateHex(g.Params.X, g.Params.Y, g.Params.Stops)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		b := new(bytes.Buffer)
		png.Encode(b, m)
		
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Length", strconv.Itoa(b.Len()))
		if _, err := w.Write(b.Bytes()); err != nil {
			w.WriteHeader(500)
			return
		}
	})
}