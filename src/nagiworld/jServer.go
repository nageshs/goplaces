package main

import (
	"time"
	"fmt"
	"net/http"
	"strconv"
	"nagiworld/jsonserver"
)

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
	// Register a fooHandler to map to "/foo"
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
		fmt.Printf("Handling %s, parsed no: %s\n", r.URL.Path, p)
		v, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		return Employee{ID: v, Name: "Nagesh", Age: 36, Salary: 10}, nil
	}

	jsonserver.RegisterHandler("/foo", fooHander)
	jsonserver.RegisterHandler("/baz/", bazHandler)
	jsonserver.StartServer("localhost", 8080)
}

