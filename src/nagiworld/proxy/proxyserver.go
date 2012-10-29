package proxy

import ("fmt"
	"net/http"
	"io/ioutil"
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
	// Add the headers 
	for k, v := range r.Header {
		if k != "ProxyConnection" {
			for _, d := range v {
				req.Header.Add(k, d)
			}
		}
	}
	// make the request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
		return
	}

	defer resp.Body.Close()
	// todo: use streaming read/write to avoid mem copy
	body, err := ioutil.ReadAll(resp.Body)

	for k,v := range resp.Header {
		for _, d := range v {
			w.Header().Add(k, d)
		}
	}
	
	w.WriteHeader(resp.StatusCode)

	o, err := w.Write(body)
	if err != nil {
		fmt.Println("Error writing to client")
		return
	}
	fmt.Printf("Served '%s', wrote %d bytes\n", r.URL.String(), o)
}

func Serve() {
	var h ProxyServer
	http.ListenAndServe("localhost:8090", h)
}