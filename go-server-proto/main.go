package main

import (
	"encoding/json"
	"flag"
	mymessages "go-server-proto/proto"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8081", "http service address")
var allPersons []*mymessages.Person

//var upgrader = websocket.Upgrader{} // default options dont work cross origin
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { //doing this to allow cross origin for now
		return true
	},
}

func ws(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("received: %s", message)
		clientEvent := &mymessages.ClientEvent{}
		err = proto.Unmarshal(message, clientEvent)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		log.Printf("received: %s", clientEvent)

		if clientEvent.GetUpsertPerson() != nil {
			handleUpsert(clientEvent.GetUpsertPerson())
		}
		if clientEvent.GetQueryPerson() != nil {
			handleQuery(clientEvent.GetQueryPerson(), c)
		}

	}
	log.Print("Ended a client session. Refresh persisted DB.")
	persistToDisk()
}

func persistToDisk() {
	personsdb_obj := mymessages.PersonDB{}
	personsdb_obj.Persons = allPersons
	persondb_bytes, _ := json.Marshal(personsdb_obj)
	err := ioutil.WriteFile("myPersonsDB", persondb_bytes, 0644)
	if err != nil {
		log.Fatal("persist persons to disk error: ", err)
	}
}

func loadFromDisk() {
	dat, err := ioutil.ReadFile("myPersonsDB")
	if err != nil {
		log.Print("read persons from disk error: ", err)
		return
	}
	personsdb_obj := &mymessages.PersonDB{}
	err = json.Unmarshal(dat, personsdb_obj)
	if err != nil {
		log.Fatal("unmarshalling from disk error: ", err)
	}

	allPersons = personsdb_obj.Persons
}

func handleUpsert(cliEvent *mymessages.UpsertPerson) {
	allPersons = append(allPersons, cliEvent.GetPerson())
}

func handleQuery(query *mymessages.QueryPerson, c *websocket.Conn) {
	foundPersons := make([]*mymessages.Person, 0)

	for _, person := range allPersons {
		if strings.Contains(person.GetName(), query.GetNameRegexp()) {
			foundPersons = append(foundPersons, person)
		}
	}

	result := new(mymessages.QueryPersonResult)
	result.Persons = foundPersons

	dataToSend, err := proto.Marshal(result)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	err = c.WriteMessage(websocket.BinaryMessage, dataToSend)
	if err != nil {
		log.Println("write error:", err)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	//init slice. Common for all clients.
	allPersons = make([]*mymessages.Person, 0)
	loadFromDisk()

	http.HandleFunc("/ws", ws)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
