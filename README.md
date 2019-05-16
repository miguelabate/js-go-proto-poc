- include the protobuf js libs
https://github.com/protocolbuffers/protobuf/tree/master/js

- get the closure library from google to be able to import the proto generated
https://github.com/google/closure-library
https://developers.google.com/closure/library/docs/gettingstarted

- install protoc
https://github.com/protocolbuffers/protobuf/releases/tag/v3.8.0-rc1

- compile proto
protoc --proto_path=./ --js_out=library=MyMessages,binary:build/gen ./message.proto
