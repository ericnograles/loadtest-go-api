# loadtest-go-api

## Overview

Source code for my [contrived load test](https://twitter.com/grales/status/904681614885756928) from Twitter.

## Prerequisites

1. [GoLang 1.9+](https://golang.org/dl/)
1. [Delve](https://github.com/derekparker/delve)
1. [Govendor](https://github.com/kardianos/govendor)
1. (Optional, Recommended) [Docker](https://www.docker.com/)
1. PostgreSQL
  * (macOS) [Postgres.app](https://postgresapp.com/)
  * (Docker, Recommended) [Postgres Image](https://hub.docker.com/_/postgres/)

## Installation

1. `go get github.com/ericnograles/loadtest-go-api`
1. `cd $GOPATH/src/github.com/ericnograles/loadtest-go-api`
1.  `govendor sync`
1. `go build && DATABASE_URL=postgres://postgres@127.0.0.1:5432/hello_go PORT=8080 ./loadtest-go-api`
    * **Note**: Swap out the value for DATABASE_URL for the proper local database on your setup
    * **Note**: Swap out the value for PORT with a proper port you'd like to use locally for the Web API
