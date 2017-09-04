# loadtest-go-api

## Overview

Source code for my [contrived load test](https://twitter.com/grales/status/904681614885756928) from Twitter.

## Prerequisites

1. [GoLang 1.8+](https://golang.org/dl/)
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
1. `go build && DATABASE_URL=postgres://postgres@127.0.0.1:5432/hello_go PORT=8080 LOADER_IO_TOKEN=some_loaderio_token ./loadtest-go-api`
    * **DATABASE_URL**: A valid Postgres URI of your local setup, including username and optionally password
    * **PORT**: Port on which to run this Web API
    * **LOADER_IO_TOKEN**: Your loader.io verification token
