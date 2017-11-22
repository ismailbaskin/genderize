package main

import (
	"net/http"
	"log"
	"strings"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Accuracy int `json:"accuracy"`
}
//go:generate go run generator/generator.go data/*
func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vals := strings.Split(strings.Trim(r.URL.Path[1:], " "), " ")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=31536000")

		results := make([]Person, 0)
		for i, val := range vals {
			if i > 4 {
				break
			}

			r := strings.NewReplacer(
				"ç", "c",
				"ğ", "g",
				"ı", "i",
				"İ", "i",
				"ö", "o",
				"ş", "s",
				"ü", "u",
			)
			name := r.Replace(strings.ToLower(val))
			var person Person;
			if malePercent, ok := getNames(name); ok {
				if malePercent > 50 {
					person = Person{val, "male", malePercent}
				} else {
					person = Person{val, "female", 100 - malePercent}
				}
				results = append(results, person)
			}
		}
		jData, _ := json.Marshal(results)

		// JSONP support
		callback := r.FormValue("callback")
		if callback != "" {
			jData = []byte(fmt.Sprintf("%s(%s)", callback, jData))
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		}

		w.Write(jData)
	})
	port := "8080"
	if value, ok := os.LookupEnv("PORT"); ok {
		port = value
	}
	err := http.ListenAndServe(":" + port, h)
	log.Fatal(err)
}
