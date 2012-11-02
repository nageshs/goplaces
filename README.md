goplaces
========

A random collection of go utilities/servers etc

build
======
go install nagiworld/proxy

Usage
=====
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
