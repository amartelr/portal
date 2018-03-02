# Steps

## First Run

go run portal.go

## Second Run

go build portal.go && ./portal
go build portal.go -o portal.run

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
	`zone` INT NULL,
	age INT NULL,
	PRIMARY KEY (id),
    UNIQUE INDEX animal_nick_UQ (nickname ASC)
);

INSERT INTO go_portal.animal (animal_type, nickname, zone, age) VALUES ('Tyrannosaurus Rex', 'rex', 1, 10);
INSERT INTO go_portal.animal (animal_type, nickname, zone, age) VALUES ('Velociraptor', 'rapto', 2, 15);
INSERT INTO go_portal.animal (animal_type, nickname, zone, age) VALUES ('Velociraptor', 'velo', 2, 12);

### Runing samplemysql

go get -u -v github.com/go-sql-driver/mysql

go run dblayer/samplemysql/samplemysql.go