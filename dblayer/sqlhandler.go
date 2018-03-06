package dblayer

import (
	"database/sql"
	"fmt"
	"log"
)

type SQLHandler struct {
	*sql.DB
}

func (handler *SQLHandler) GetAvailableAnimals() ([]Animal, error) {
	return handler.sendQuery("select * from animal")
}

func (handler *SQLHandler) GetAnimalByNickname(nickname string) (Animal, error) {

	row := handler.QueryRow(fmt.Sprintf("select * from animal where nickname = '%s'", nickname)) //? for mysql or sqlite and it used to be $1 for pq
	a := Animal{}
	err := row.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zona, &a.Age)
	return a, err
}

func (handler *SQLHandler) GetAnymalByType(animalType string) ([]Animal, error) {
	return handler.sendQuery(fmt.Sprintf("select * from animal where Animal_type = '%s'", animalType))
}

func (handler *SQLHandler) AddAnimal(a Animal) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into animal (Animal_type,nickname,zona,age) values ('%s','%s',%d,%d)", a.AnimalType, a.Nickname, a.Zona, a.Age))
	return err
}
func (handler *SQLHandler) UpdateAnimal(a Animal, nname string) error {
	_, err := handler.Exec(fmt.Sprintf("Update animal set Animal_type = '%s' ,nickname = '%s',zona = %d,age = %d where nickname = '%s'", a.AnimalType, a.Nickname, a.Zona, a.Age, nname))
	return err
}

func (handler *SQLHandler) sendQuery(q string) ([]Animal, error) {
	Animals := []Animal{}
	rows, err := handler.Query(q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		a := Animal{}
		err := rows.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zona, &a.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		Animals = append(Animals, a)
	}
	return Animals, rows.Err()
}
