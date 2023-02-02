package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Jeffail/tunny"
)

func main() {

	rand.Seed(time.Now().Unix())

	pool := tunny.NewFunc(10, func(payload interface{}) interface{} {
		var result []byte

		// TODO: Something CPU heavy with payload

		return result
	})

	defer pool.Close()

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		input, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		// Funnel this work into our pool. This call is synchronous and will
		// block until the job is completed.
		result := pool.Process(input)

		time.Sleep(time.Duration(rand.Intn(8)+1) * time.Second)

		w.Write(result.([]byte))

		log.Print(result.([]byte))
	})

	http.ListenAndServe(":8082", nil)
}
