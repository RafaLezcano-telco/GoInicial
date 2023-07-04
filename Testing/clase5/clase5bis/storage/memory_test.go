package storage

import (
	"testing"

	"github.com/RafaLezcano-telco/GoInicial/Testing/clase5/api/model"
)

func TestCreate(t*testing.T){
	table:=[]struct{
		name string
		person *model.Person
		wantError error
	}{
		{"Empty person",nil,model.ErrPersonCanNotBeNil},
		{"Alexys",&model.Person{Name:"Alexys"},nil},
		{"Matthew",&model.Person{Name:"Matthew"},nil},
	}
	m:= NewMemory()
	for _,v:=range table{
		t.Run(v.name,func(t*testing.T){
			gotErr:=m.Create(v.person)
			if gotErr != v.wantError{
				t.Errorf("Se esperaba %v, se obtuvo %v",v.wantError,gotErr)
			}
			if m.currentID!= len(table){
				t.Errorf("Se esperaba %d registros, se obtuvo %d", len(table),m.currentID)
			}
		})
	}
}


func TestCreate_empty_person(t *testing.T) {
	m := NewMemory()
	err:= m.Create(nil)
	if err== nil{
		t.Errorf("Se esperaba error, se obtuvo nil")

	}
	if err != model.ErrPersonCanNotBeNil{
		t.Errorf("Se esperaba el error %v, se obtuvo el error %v",model.ErrPersonCanNotBeNil,err)
	}
}


func TestCreate_count_elements(t *testing.T){
	m := NewMemory()
	p:= model.Person{Name:"Alexys"}
	err:= m.Create(&p)
	if err !=nil {
		t.Errorf("No se esperaba error, se obtuvo %v",err)
	}
	if m.currentID != 1{
		t.Errorf("Se esperaba 1 elemento, se obvtuvo: %d", m.currentID)
	}

}