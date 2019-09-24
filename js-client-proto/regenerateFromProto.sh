#!/bin/sh
protoc --proto_path=../ --js_out=library=MyMessages,binary:build/gen ../proto/message.proto
