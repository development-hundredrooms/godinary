package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/trilopin/godinary/image"
	"github.com/trilopin/godinary/parser"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to Homepage!\n")
}

func Fetch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	job, err := parser.Parse(ps.ByName("info")[1:])
	if err != nil {
		log.Fatal("Invalid request")
	}
	resp, err := http.Get(job.SourceURL)
	if err != nil {
		log.Fatal("Cannot get url " + job.SourceURL)
	}
	w.Header().Set("Content-Type", "image/jpeg")
	err = image.Process(resp.Body, job, w)
}

func StartServer() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/v0.1/:account/fetch/*info", Fetch)

	log.Fatal(http.ListenAndServe(":3001", router))
}