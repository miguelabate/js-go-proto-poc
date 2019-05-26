package main

import (
	"go-server-proto/mymessages"

	"flag"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/gorilla/websocket"
)

//func main() {
//	msg := &mymessages.MessageMike{Text: "dfds"}
//	data, err := proto.Marshal(msg)
//	if err != nil {
//		log.Fatal("marshaling error: ", err)
//	}
//	newMsg := &mymessages.MessageMike{}
//	err = proto.Unmarshal(data, newMsg)
//	if err != nil {
//		log.Fatal("unmarshaling error: ", err)
//	}
//	// Now test and newTest contain the same data.
//	if msg.GetText() != newMsg.GetText() {
//		log.Fatalf("data mismatch %q != %q", msg.GetText(), newMsg.GetText())
//	}

var addr = flag.String("addr", ":8081", "http service address")

//var upgrader = websocket.Upgrader{} // default options dont work cross origin
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { //doing this to allow cross origin for now
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		//log.Printf("recv: %s", message)
		receivedMsg := &mymessages.MessageMike{}
		err = proto.Unmarshal(message, receivedMsg)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		log.Printf("recv: %s", receivedMsg)

		//now modify and send
		receivedMsg.Text = receivedMsg.GetText() + " Altered content"
		receivedMsg.Lang = receivedMsg.GetLang() + " Altered lang"
		receivedMsg.Length = receivedMsg.GetLength() * 1.5
		receivedMsg.Years = receivedMsg.GetYears() + 10

		dataToSend, err := proto.Marshal(receivedMsg)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}

		err = c.WriteMessage(websocket.BinaryMessage, dataToSend)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
