# GO PPROF Example

A forcefylly simple example of Gorilla MUX running a simple web server and a Pprof Server.

PProf Server has all endpoints enabled:

* "/"
* "/cmdline"
* "/symbol"
* "/trace"
* "/profile"

https://stackoverflow.com/questions/19591065/profiling-go-web-application-built-with-gorillas-mux-with-net-http-pprof
	
* "/goroutine"
* "/heap"
* "/threadcreate"
* "/block"
* "/vars"

## Usage

Simply executing below should do the trick:

```
./docker_run.sh

```

Pprof runs on port 15121 and webserver on 15120

## Performance client

I used Apache Benchmark. Some examples in perf.sh

## Dockerfile

I provided two small utilities to build and run the docker containers

## Testing

If everything goes well open your browser and go to:

```
http://localhost:15121/

You shoould see: "Pprof!"

http://localhost:15121/

You shoould see: "Go Away!"
```



