package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mairuu/nekopost-api/services/chapter"
	"github.com/mairuu/nekopost-api/services/project"
)

func main()  {
    port := flag.String("port", "8080", "port to listen")
    flag.Parse()

    chapterHandler := chapter.NewHandler()
    chapterHandler.RegisterRoutes(http.DefaultServeMux)

    projectHandler := project.NewHandler()
    projectHandler.RegisterRoutes(http.DefaultServeMux)

    log.Printf("listening on port: %s\n", *port)
    log.Fatal(http.ListenAndServe(":" + *port, nil))
}

