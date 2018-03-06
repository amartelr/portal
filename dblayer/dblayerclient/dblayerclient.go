package main

import "github.com/amartelr/portal/dblayer"
import "log"

func main() {
	//handler, err := dblayer.GetDatabaseHandler(dblayer.MYSQL, "gouser:!password@/go_portal")
	//handler, err := dblayer.GetDatabaseHandler(dblayer.SQLITE, "../samplesqlite/portal.db")
	//handler, err := dblayer.GetDatabaseHandler(dblayer.POSTGRESQL, "user=postgres dbname=dino sslmode=disable")
	handler, err := dblayer.GetDatabaseHandler(dblayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	/*
		err = handler.UpdateAnimal(dblayer.Animal{
			AnimalType: "Carnotaurus",
			Nickname:   "Carno",
			Zona:       3,
			Age:        23,
		}, "Carno")
		log.Println(err)
	*/
	log.Println(handler.GetAvailableAnimals())
	log.Println(handler.GetAnimalByNickname("rex"))
	log.Println(handler.GetAnymalByType("Velociraptor"))
}
