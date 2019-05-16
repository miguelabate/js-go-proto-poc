-install protoc and go plugin
go get -u github.com/golang/protobuf/protoc-gen-go
-generate code
protoc --go_out=paths=source_relative:./goServerProto message.proto