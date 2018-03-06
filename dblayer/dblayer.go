package dblayer

import (
	"errors"
)

const (
	MYSQL uint8 = iota
	SQLITE
	MONGODB
	COCKROACHDB
	POSTGRESQL
)

type PortalDBHandler interface {
	GetAvailableAnimals() ([]Animal, error)
	GetAnimalByNickname(string) (Animal, error)
	GetAnymalByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zona       int    `bson:"zona"`
	Age        int    `bson:"age"`
}

var DBTypeNotSupported = errors.New("Database type provided is not supported")

func GetDatabaseHandler(dbtype uint8, connection string) (PortalDBHandler, error) {
	switch dbtype {
	case MYSQL:
		return NewMySQLHandler(connection)
	case MONGODB:
		return NewMongodbHandler(connection)
	case SQLITE:
		return NewSQLiteHandler(connection)
	case POSTGRESQL:
		return NewPQHandler(connection)
	}
	return nil, DBTypeNotSupported
}
