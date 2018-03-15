package main

import (
	"log"

	"github.com/amartelr/portal/dblayer"
	"github.com/amartelr/portal/webportal/portalapi"
)

func main() {
	db, err := dblayer.GetDatabaseHandler(dblayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(db.GetAvailableAnimals())
	log.Println("http://localhost:8080/api/portal/nickname/rex  http://localhost:8080/api/portal/type/Velociraptor")
	portalapi.RunApi("localhost:8080", db)

}
