# go-starlink

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
