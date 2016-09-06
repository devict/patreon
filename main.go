package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
	}
	p = ":" + p

	http.HandleFunc("/campaign", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get("https://api.patreon.com/campaigns/400650")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		if _, err := io.Copy(w, res.Body); err != nil {
			log.Println(err)
		}
	})

	log.Println("Listening to", p)
	log.Fatal(http.ListenAndServe(p, nil))
}
