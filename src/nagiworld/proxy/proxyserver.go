package proxy

import ("fmt"
	"net/http"
	"io"
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

	for k,v := range resp.Header {
		for _, d := range v {
			w.Header().Add(k, d)
		}
	}
	
	w.WriteHeader(resp.StatusCode)
	count, err := io.Copy(w, resp.Body)

	// following can both be in a defer func
	if err != nil {
		// fmt.Printf("Error writing to client %s \n", err)
		fmt.Printf("%s\t%s\t%d\t%d\n", r.URL.String(), r.Method, resp.StatusCode, 0)
		return
	}

	fmt.Printf("%s\t%s\t%d\t%d\n", r.URL.String(), r.Method, resp.StatusCode, count)
}

func Serve() {
	var h ProxyServer
	http.ListenAndServe("localhost:8090", h)
}