package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/RafaLezcano-telco/GoInicial/Testing/clase5/api/model"
)

//Psql

type Psql struct{
	db *sql.DB
}

func NewPsql() Psql{
	p:=Psql{}
	p.db= getPSQLConn()
	return p
}
func getPSQLConn() *sql.DB{
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "rafaTelco"
		dbname   = "clase6_golang"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err:= sql.Open("postgres",psqlInfo)
	if err!= nil{
		log.Fatalf("no se pudo conectar a la bd: %v",err)
	}
	if err= db.Ping(); err!=nil{
		log.Fatalf("no se puso hacer ping a la bd: %v",err)
	}
	return db
}

func (psql * Psql) Create(person *model.Person) error{
	if person==nil{
		return model.ErrPersonCanNotBeNil
	}
	stmt,err := psql.db.Prepare(`INSERT INTO persons VALUES ($1,$2)`)
	if err!= nil{
		log.Fatalf("no se pudo preparar la consulta: %v",err)
	}
	defer stmt.Close()
	_,err= stmt.Exec(person.Name,person.Age)
	if err!= nil{
		log.Fatalf("no se pudo insertar el registro %v",err)
	}
	return nil
}
func (psql * Psql) Update(ID int, person *model.Person) error{
	return errors.New("not implemented yet")
}

func (psql * Psql) Delete(ID int) error{
	return errors.New("not implemented yet")
}
func (psql * Psql) GetByID(ID int) (model.Person,error) {
	return model.Person{},errors.New("not implemented yet")
}

func (psql *Psql) GetAll() (model.Persons,error){
	stmt,err:= psql.db.Prepare(`SELECT name, age FROM persons`)
	if err != nil {
		log.Fatalf("no se pudo preparar la consulta: %v", err)
	}
	defer stmt.Close()

	rows, err:= stmt.Query()
	if err!= nil{
		log.Fatalf("no se pudo consultar los registros %v",err)
	}
	defer rows.Close()
	var persons model.Persons
	for rows.Next(){
		p:= model.Person{}
		err= rows.Scan(&p.Name, &p.Age)
		if err != nil{
			log.Fatalf("no se pudo mapear la persona %v",err)
		}
		persons=append(persons,p)
	}
	return persons,nil
}
func (psql * Psql) DeleteAll() error{
	stmt,err:= psql.db.Prepare(`TRUNCATE TABLE persons`)
	if err !=nil{
		log.Fatalf("no se pudo preparar la consulta: %v",err)
	}
	defer stmt.Close()
	_,err=stmt.Exec()
	if err!= nil{
		log.Fatalf("no se pudo borrar la tabla:%v",err)
	}
	return nil
}

func (psql *Psql) CloseDB(){
	psql.db.Close()
}
