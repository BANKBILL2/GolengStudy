package main

import (
	"fmt"
	"log"
	"net/http"
	"prospersof/auth"
	"prospersof/fizzbuzz"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Hello world, the web server
	r := mux.NewRouter()

	r.HandleFunc("/token", auth.TokenHandler)
	r.HandleFunc("/fizzbuzz/{number}", fizzbuzzHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numberString := vars["number"]
	n, err := strconv.Atoi(numberString)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	s := fizzbuzz.Say(n)
	fmt.Fprint(w, s)
}
