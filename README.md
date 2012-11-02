goplaces
========

A random collection of go utilities/servers etc

build
======
go install nagiworld/proxy

Proxy Server Usage
==================
```go 

import  "nagiworld/proxy"

func main() {
     // starts proxy server on localhost:<port>
     proxy.Serve(8090)
}
```

```sh

curl -x localhost:8090 http://techcrunch.com/
```

JsonServer Usage:
=================
The json server is an abstraction for letting users write handlers which take in the request and return any json object which is then send down the wire to the client. 

All thats needed is to provide a __func_(r *http.Request)__ and returns the data as a json struct

```go
	fooHandler := func(r *http.Re(interface{}, error) {
       ///  Look into the request and return a struct 
       return &Employee{ID: 2, Name:"foo", Age:20, Salary:500}
	}

 // Once created we then register it with the handler 

 RegisterHandler("/foo", fooHandler)
```

Running
==========
To run the json server you can issue __go run jsonServer.go__