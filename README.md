# go-starlink
[![license](https://img.shields.io/github/license/b0ch3nski/go-starlink)](LICENSE)
[![commit](https://img.shields.io/github/last-commit/b0ch3nski/go-starlink/master)](https://github.com/b0ch3nski/go-starlink/commits/master)
[![go.dev](https://pkg.go.dev/badge/github.com/b0ch3nski/go-starlink)](https://pkg.go.dev/github.com/b0ch3nski/go-starlink)
[![goreportcard](https://goreportcard.com/badge/github.com/b0ch3nski/go-starlink)](https://goreportcard.com/report/github.com/b0ch3nski/go-starlink)
[![issues](https://img.shields.io/github/issues/b0ch3nski/go-starlink)](https://github.com/b0ch3nski/go-starlink/issues)
[![sourcegraph](https://sourcegraph.com/github.com/b0ch3nski/go-starlink/-/badge.svg)](https://sourcegraph.com/github.com/b0ch3nski/go-starlink)

Simple Golang client to interact with SpaceX Starlink dish, based on Proto files acquired from Dishy itself using
[GRPC Server Reflection][reflection] and [grpcurl][grpcurl] tool.

For a code generation procedure, see included [Makefile](Makefile).

[reflection]: https://github.com/grpc/grpc/blob/master/doc/server-reflection.md
[grpcurl]: https://github.com/fullstorydev/grpcurl

## install

```
go get github.com/b0ch3nski/go-starlink
```

## example

```go
ctx := context.Background()

starlinkClient, _ := client.NewClient(ctx, client.DefaultDishyAddr)
starlinkStatus, _ := starlinkClient.Status(ctx)

fmt.Println(starlinkStatus)
```
