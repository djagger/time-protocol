# time-protocol

**Time protocol server** builds by `cd server/ && go build`,
then runs by `./server [-p port]` (by default it runs on _11037_ port).
```
Usage of ./server:
  -p int
    	listen port (default 11037)
```

**Time protocol client** builds by `cd client/ && go build` and it runs by
`./client`. First and second arguments is _host_ and _port_. For example:

- `./client` - connects to _localhost_ by _11037_ port
- `./client time.nist.gov 37` - connects to _time.nist.gov_ by _37_ port