package proxy

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"log"
)

// server code 
type ProxyServer struct {

}

var client = &http.Client{}

func (h ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(r.Method, r.URL.String(), r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
		return
	}
	copyHeaders(r.Header, req.Header)

	// make the request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	copyHeaders(resp.Header, w.Header())

	w.WriteHeader(resp.StatusCode)
	count, err := io.Copy(w, resp.Body)

	if err != nil {
		// fmt.Printf("Error writing to client %s \n", err)
		fmt.Printf("%s\t%s\t%d\t%d\n", r.URL.String(), r.Method, resp.StatusCode, 0)
		return
	}

	fmt.Printf("%s\t%s\t%d\t%d\n", r.URL.String(), r.Method, resp.StatusCode, count)
}

// utility method to copy header from source to destination
func copyHeaders(from http.Header, to http.Header) {
	for k, v := range from {
		for _, d := range v {
			to.Add(k, d)
		}
	}
}

// Takes the port number the proxy server should start on. 
func Serve(port int) {
	var h ProxyServer
	fmt.Printf("Starting server on localhost:%s ...\n", strconv.Itoa(port))
	err := http.ListenAndServe("localhost:"+strconv.Itoa(port), h)
	if err != nil {
		fmt.Printf(" Error %s \n", err)
		log.Fatal(err)
	}
}
