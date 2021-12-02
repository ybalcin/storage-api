# RESTful Storage API

The project is a storage api developed in Go. Currently, only supports memory caching.
It developed in DDD and Ports&Adapters architecture. Follows the SOLID principles and Idiomatic Go.

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer.

```shell
# clone repo
git clone https://github.com/ybalcin/storage-api

cd storage-api/cmd

# run api server
go run main.go
```

At this time, you have a RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

* `POST /`: gets mongo query
* `GET /in-memory?key=dummykey`: gets cache entry
* `POST /in-memory`: adds cache entry

Heroku Address: https://storage-api-study.herokuapp.com
