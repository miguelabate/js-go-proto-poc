package main

import (
	"goServerProto/mymessages"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	msg := &mymessages.MessageMike{Text: "dfds"}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newMsg := &mymessages.MessageMike{}
	err = proto.Unmarshal(data, newMsg)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if msg.GetText() != newMsg.GetText() {
		log.Fatalf("data mismatch %q != %q", msg.GetText(), newMsg.GetText())
	}
}
