package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		randInt := 1 + rand.Intn(6)
		w.Header().Set("Content-Type", "text/json")
		data, _ := json.Marshal(strconv.Itoa(randInt))
		w.Write((data))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
