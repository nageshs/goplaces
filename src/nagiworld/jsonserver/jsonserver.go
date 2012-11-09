package jsonserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Utility function used by the implementors
func marshallBytes(v interface{}) ([]byte, error) {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Public method to register a handler 
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

func StartServer(host string, port int) {
	if len(host) == 0 {
		host = "localhost"
	}

	err := http.ListenAndServe(host + ":" + strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
}


