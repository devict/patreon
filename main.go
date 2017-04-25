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

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		u := "https://api.meetup.com/2/events?offset=0&format=json&limited_events=False&group_urlname=devICT&photo-host=public&page=20&fields=&order=time&status=upcoming&desc=false&sig_id=73273692&sig=8f90c3aff1c3055274bc6dffca9225f51754a928"
		res, err := http.Get(u)
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
