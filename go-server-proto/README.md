#Echo GO server

Starts up a server listening on 8081 that receives a message of type MessageMike, modifies it and sends it back.

#Prerequirements to run

- Install protoc and go plugin
go get -u github.com/golang/protobuf/protoc-gen-go

-Generate go code from proto
protoc --go_out=paths=source_relative:./go-server-roto proto/message.proto

-Run
go run main.go
