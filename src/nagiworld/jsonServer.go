package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"fmt"
	"strconv"
)

// type JsonServer interface {
// 	func responseJson(r *http.Request)
// }

// Utility function used by the implementors
func marshallBytes(v interface{}) ([]byte, error) {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func RegisterHandler(path string, fn func(r *http.Request) (interface{}, error)) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		val, err := fn(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := marshallBytes(val)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
	http.HandleFunc(path, handler)
	return handler
}

// Example struct that is sent down by the fooHandler
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

// Example struct for bazHandler
type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary int
}

func main() {
	fooHander := func(r *http.Request) (interface{}, error) {
		group := ColorGroup{
			ID:     1,
			Name:   "Reds_" + time.Now().String(),
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		}
		return group, nil
	}

	// Path handler /baz/<employee-number>
	bazHandler := func(r *http.Request) (interface{}, error) {
		// check the request to see the id 
		p := r.URL.Path[5:]
		fmt.Printf("%s got value\n", p)
		v, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		return &Employee{ID: v, Name: "Nagesh", Age: 36, Salary: 10}, nil
	}

	RegisterHandler("/foo", fooHander)
	RegisterHandler("/baz/", bazHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
