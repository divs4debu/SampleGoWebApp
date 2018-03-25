package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gowiki/handler/article"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s" , r.URL.Path[1:])
}


func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/edit/", article.EditHandler)
	http.HandleFunc("/view/", article.ViewHandler)
	http.HandleFunc("/save/", article.SaveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}