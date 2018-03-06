package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"github.com/amartelr/portal/comlayer/proto3"
	"github.com/amartelr/portal/dblayer"
	"github.com/golang/protobuf/proto"
)

func main() {
	//proto3test -op s => will run as a server, proto3test -op c => will run as a client, d=> will run mongo

	op := flag.String("op", "s", "s for server, c for client")
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		RunProto3Server()
	case "c":
		RunProto3Client()
	case "d":
		SendDBToServer()

	}
}

func RunProto3Client() {
	a := &proto3.Animal{
		Id:         1,
		AnimalType: "Raptor",
		Nickname:   "rapto",
		Zona:       3,
		Age:        21,
	}
	data, err := proto.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	SendData(data)
}

func SendData(data []byte) {
	c, err := net.Dial("tcp", "127.0.0.1:8181")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Write(data)
}

func RunProto3Server() {
	l, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()
		go func(c net.Conn) {
			defer c.Close()
			data, err := ioutil.ReadAll(c)
			if err != nil {
				return
			}
			a := &proto3.Animal{}
			err = proto.Unmarshal(data, a)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(a)
		}(c)
	}
}

func SendDBToServer() {
	handler, err := dblayer.GetDatabaseHandler(dblayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	animals, err := handler.GetAvailableAnimals()
	for _, animal := range animals {
		a := &proto3.Animal{
			Id:         int32(animal.ID),
			AnimalType: animal.AnimalType,
			Nickname:   animal.Nickname,
			Zona:       int32(animal.Zona),
			Age:        int32(animal.Age),
		}
		data, err := proto.Marshal(a)
		if err != nil {
			log.Fatal(err)
		}
		SendData(data)
	}
}
