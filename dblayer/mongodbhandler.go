package dblayer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongodbHandler struct {
	*mgo.Session
}

func NewMongodbHandler(connection string) (*MongodbHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongodbHandler{
		Session: s,
	}, err
}

func (handler *MongodbHandler) GetAvailableAnimals() ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("go_portal").C("animal").Find(nil).All(&animals)
	return animals, err
}

func (handler *MongodbHandler) GetAnimalByNickname(nickname string) (Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	a := Animal{}
	err := s.DB("go_portal").C("animal").Find(bson.M{"nickname": nickname}).One(&a)
	return a, err
}

func (handler *MongodbHandler) GetAnymalByType(animalType string) ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("go_portal").C("animal").Find(bson.M{"animal_type": animalType}).All(&animals)
	return animals, err
}

func (handler *MongodbHandler) AddAnimal(a Animal) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("go_portal").C("animal").Insert(a)
}

func (handler *MongodbHandler) UpdateAnimal(a Animal, nname string) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("go_portal").C("animal").Update(bson.M{"nickname": nname}, a)
}

func (handler *MongodbHandler) getFreshSession() *mgo.Session {
	return handler.Session.Copy()
}
