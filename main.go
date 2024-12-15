package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mairuu/nekopost-api/services/chapter"
	"github.com/mairuu/nekopost-api/services/comment"
	"github.com/mairuu/nekopost-api/services/project"
)

func main()  {
    port := flag.String("port", "8080", "port to listen")
    passpharse := flag.String("passphrase", "", "api passpharse")
    flag.Parse()

    chapterApi := chapter.NewApi(*passpharse)
    projectApi := project.NewApi(*passpharse)
    commentApi := comment.NewApi()

    chapterHandler := chapter.NewHandler(chapterApi)
    chapterHandler.RegisterRoutes(http.DefaultServeMux)

    projectHandler := project.NewHandler(projectApi)
    projectHandler.RegisterRoutes(http.DefaultServeMux)

    commentHandler := comment.NewHandler(commentApi)
    commentHandler.RegisterRoutes(http.DefaultServeMux)

    log.Printf("listening on port: %s\n", *port)
    log.Fatal(http.ListenAndServe(":" + *port, nil))
}

