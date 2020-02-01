# intro
(very) simple fileserver running on localhost:5000

# how to build
```shell script
$ go build
```

# how to run
```shell script
$ ./fileserver
``` 

# caution
- method only provides `POST`
- `Content-Type` header should be `multipart/form-data`
- form name should be `file`
- can build when `go` installed, but note it's os dependent
  > must build and run in the same os, platform
- when file name is the same, it will be overwritten  

# example  
```shell script
$ curl -XPOST http://localhost:5000/ -F "file=@buy.mp4" -vvv

*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 5000 (#0)
> POST / HTTP/1.1
> Host: localhost:5000
> User-Agent: curl/7.64.1
> Accept: */*
> Content-Length: 32343
> Content-Type: multipart/form-data; boundary=------------------------1ef85dcecbcbedcf
> Expect: 100-continue
>
< HTTP/1.1 100 Continue
* We are completely uploaded and fine
< HTTP/1.1 200 OK
< Date: Sat, 01 Feb 2020 11:22:48 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
* Closing connection 0
```

> then, you can find `buy.mp4` on current directory 

