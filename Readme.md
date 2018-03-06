# Steps

## First Run

go run portal.go

### Download database drivers

go get -u -v github.com/go-sql-driver/mysql
go get -u -v github.com/mattn/go-sqlite3

go get -u -v gopkg.in/mgo.v2
go get -u -v gopkg.in/mgo.v2/bson

go get -u -v github.com/lib/pq

go run dblayer/samplemysql/samplemysql.go

## ORM

go get -u -v github.com/jinzhu/gorm

## Mysql

mysql -uroot -p(yourpass)

CREATE DATABASE go_portal;
CREATE USER 'gouser'@'localhost' IDENTIFIED BY '!password';

GRANT ALL ON go_portal.* TO 'gouser'@'localhost';

exit

CREATE TABLE go_portal.animal (
	id INT NOT NULL AUTO_INCREMENT,
	animal_type varchar(45) NOT NULL,
	nickname varchar(45) NOT NULL,
	zona INT NULL,
	age INT NULL,
	PRIMARY KEY (id),
    UNIQUE INDEX animal_nick_UQ (nickname ASC)
);

INSERT INTO go_portal.animal (animal_type, nickname, zona, age) VALUES ('Tyrannosaurus Rex', 'rex', 1, 10);
INSERT INTO go_portal.animal (animal_type, nickname, zona, age) VALUES ('Velociraptor', 'rapto', 2, 15);
INSERT INTO go_portal.animal (animal_type, nickname, zona, age) VALUES ('Velociraptor', 'velo', 2, 12);

## sqlite3

sqlite3 ../samplesqlite/portal.db
.tables

## mongo

mongo
show dbs
use go_portal
db.animal.find({}).pretty()

## Postgresql 
sudo su - postgres
createuser --createdb --username postgres --no-createrole --no-superuser --pwprompt gouser
Enter password for new role: !password
Enter it again: !password

/*
sudo -u postgres psql postgres
\password postgres
*/

CREATE TABLE public.animal
(
    id serial,
    animal_type text,
    nickname text,
    zona integer,
    age integer,
    PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
);

ALTER TABLE public.animal
    OWNER to postgres;

insert INTO animal (animal_type,nickname,zona,age) VALUES ('Tyrannosaurus rex','rex', 1, 10),('Velociraptor', 'rapto', 2, 15);

## Protocol Buffers

go get -u -v github.com/golang/protobuf/protoc-gen-go
sudo apt install protobuf-compiler

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto

## proto3

cd comlayer/proto3
protoc -I=. --go_out=. ./protadata.proto

## runing 
./proto3test -op s
./proto3test -op c