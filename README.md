#General

Simple POC project that shows how to comunicate via websockets using proto messages between JS (in browser) and GO server 

#JS echo  client

Connects to a websocket server on port 8081, sends a MessageMike and receives it back

#Prerequirements to run

- include the protobuf js libs
git@github.com:protocolbuffers/protobuf.git
https://github.com/protocolbuffers/protobuf/tree/master/js

- get the closure library from google to be able to import the proto generated
https://github.com/google/closure-library
https://developers.google.com/closure/library/docs/gettingstarted

- install protoc
https://github.com/protocolbuffers/protobuf/releases/tag/v3.8.0-rc1

- compile proto for JS
protoc --proto_path=./ --js_out=library=MyMessages,binary:build/gen ./proto/message.proto

