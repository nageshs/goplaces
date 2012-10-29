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
     // starts  proxy ol  starts proxy server on localhost:8090
     proxy.Serve()
}
```

```sh

curl -x localhost:8090 http://techcrunch.com/
```
