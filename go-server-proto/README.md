# Persons DB - GO server

Starts up a server listening on 8081 that receives a message of type PersonUpsert and PersonQuery. 

# Prerequirements to run

- Install protoc and go plugin  
go get -u github.com/golang/protobuf/protoc-gen-go

- Generate go code from proto  
protoc --proto_path=../ --go_out=. ../proto/message.proto

- Run  
go run main.go

