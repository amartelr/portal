package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type animal struct {
	//gorm.Model
	ID         int    `gorm:"primary_key;not null;unique;AUTO_INCREMENT"`
	Animaltype string `gorm:"type:TEXT"`
	Nickname   string `gorm:"type:TEXT"`
	Zona       int    `gorm:"type:INTEGER"`
	Age        int    `gorm:"type:INTEGER"`
}

func main() {
	//db, err := gorm.Open("mysql", "gouser:!password@/go_portal?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("sqlite3", "../samplesqlite/portal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// // Disable table name's pluralization
	db.SingularTable(true)

	db.DropTableIfExists(&animal{})
	//	db.Table("animal").DropTableIfExists(&animal{})

	// will add any missing fields, will add 's' to the struct name

	db.AutoMigrate(&animal{})
	//db.Table("animal").AutoMigrate(&animal{})

	//inserts:
	a := animal{
		Animaltype: "Tyrannosaurus rex",
		Nickname:   "rex",
		Zona:       1,
		Age:        11,
	}
	db.Create(&a) //vs create()
	//db.Table("animal").Create(&a)

	a = animal{
		Animaltype: "Velociraptor",
		Nickname:   "rapto",
		Zona:       2,
		Age:        15,
	}
	db.Save(&a) //vs create()

	//updates
	//db.Table("animals").Where("nickname = ? and zone= ?", "rapto", 2).Update("age", 16)

	//queries
	animals := []animal{}
	db.Table("animal").Find(&animals, "age > ?", 10) //first
	fmt.Println(animals)

}
