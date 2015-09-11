package main

import (
	"github.com/johnstall/xwp"
	"log"
	"net/http"
)

func main() {
	db, err := xwp.Open("db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	h := &xwp.Handler{
		DB: db,
	}
	log.Fatal(http.ListenAndServe(":8888", h))
}
