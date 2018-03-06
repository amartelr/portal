package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type animal struct {
	//ID         int    `bson:"id"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zona       int    `bson:"zona"`
	Age        int    `bson:"age"`
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//get a collection
	animalcollection := session.DB("go_portal").C("animal")

	animals := []interface{}{
		animal{
			AnimalType: "Tyrannosaurus rex",
			Nickname:   "rex",
			Zona:       1,
			Age:        11,
		}, animal{
			AnimalType: "Velociraptor",
			Nickname:   "rapto",
			Zona:       2,
			Age:        17,
		}, animal{
			AnimalType: "Velociraptor",
			Nickname:   "Velo",
			Zona:       2,
			Age:        9,
		},
	}
	err = animalcollection.Insert(animals...)

	if err != nil {
		log.Fatal(err)
	}

	/*
		err = animalcollection.Update(bson.M{"nickname": "rapto"}, bson.M{"$set": bson.M{"age": 18}})
		if err != nil {
			log.Fatal(err)
		}
	*/
	/*
		err = animalcollection.Remove(bson.M{"nickname": "rapto"})
		if err != nil {
			log.Fatal(err)
		}
	*/
	// age > 10 and Zona in (1,2)
	query := bson.M{
		"age": bson.M{
			"$gt": 5,
		},
		"zona": bson.M{
			"$in": []int{1, 2},
		},
	}
	results := []animal{}
	animalcollection.Find(query).All(&results) //.one
	fmt.Println(results)

	//one result
	result := animal{}
	animalcollection.Find(query).One(&result)
	fmt.Println(result)
}
