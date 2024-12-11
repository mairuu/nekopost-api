package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mairuu/nekopost-api/internal/handler"
)

func main()  {
    port := flag.String("port", "8080", "port to listen")
    flag.Parse()

    http.HandleFunc("GET /api/chapters/latest", handler.HandleLatestChapters)
    http.HandleFunc("GET /api/chapters/{chapter_id}/comments", handler.HandleChapterComments)

    http.HandleFunc("GET /api/projects/random", handler.HandleRandomProjects)
    http.HandleFunc("GET /api/projects/popular", handler.HandlePopularProjects)
    http.HandleFunc("GET /api/projects", handler.HandleProjects)
    http.HandleFunc("GET /api/projects/{project_id}", handler.HandleProject)
    http.HandleFunc("GET /api/projects/{project_id}/cover.jpg", handler.HandleProjectCover)
    http.HandleFunc("GET /api/projects/{project_id}/chapters/{chapter_id}", handler.HandleProjectChapter)

    log.Printf("listening on port: %s\n", *port)
    log.Fatal(http.ListenAndServe(":" + *port, nil))
}

