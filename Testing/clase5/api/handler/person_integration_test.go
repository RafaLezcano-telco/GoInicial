package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RafaLezcano-telco/GoInicial/Testing/clase5/api/model"
	"github.com/RafaLezcano-telco/GoInicial/Testing/clase5/api/storage"
	"github.com/labstack/echo"
)
func TestPerson_Create_integration(t *testing.T){
	
	t.Cleanup(func(){
		cleanData(t)
	})
	
	data:= []byte(`{"name":"Alexys","age":18}`)
	w:= httptest.NewRecorder()
	r:= httptest.NewRequest(http.MethodPost,"/",bytes.NewBuffer(data))
	r.Header.Set("Content-Type","application/json")
	e:=echo.New()
	ctx:=e.NewContext(r,w)

	store:=storage.NewPsql()

	p:=newPerson(&store)
	err:= p.create(ctx)
	if err != nil{
		t.Errorf("no se esperaba error, pero se obtuvo %v",err)
	}
	if w.Code != http.StatusCreated{
		t.Errorf("Codigo de estado, se esperaba %d, se obtuvo %d", http.StatusCreated,w.Code)
	}
	resp:= response{}
	err= json.NewDecoder(w.Body).Decode(&resp)
	if err !=nil{
		t.Errorf("no se pudo hacer unmarshal al body: %v",err)
	}
	wantMessage:="Persona creada correctamente"
	if resp.Message!= wantMessage{
		t.Errorf("La respuesta  no es la esperada, se obtuvo %q, se esperaba %q",resp.Message,wantMessage)
	}	
}

type responsePerson struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        []model.Person `json:"data"`
}
func TestPerson_GetAll_integration(t *testing.T){
	t.Cleanup(func(){
		cleanData(t)
	})
	insertData(t)
	w:= httptest.NewRecorder()
	r:= httptest.NewRequest(http.MethodPost,"/",nil)
	r.Header.Set("Content-Type","application/json")
	e:=echo.New()
	ctx:=e.NewContext(r,w)

	store:=storage.NewPsql()

	p:=newPerson(&store)
	err:= p.getAll(ctx)
	if err!= nil{
		t.Fatalf("no se esperaba error, se obtuvo %v",err)
	}
	if w.Code != http.StatusOK{
		t.Errorf("Codigo de estado, se esperaba %d, se obtuvo %d", http.StatusOK,w.Code)
	}

	resp:= responsePerson{}
	err=json.NewDecoder(w.Body).Decode(&resp)
	if err!= nil{
		t.Fatalf("no se pudo hacer unmarshal de la respuesta %v",err)
	}
	lenPersonsWant:=2
	lenPersonsGot:=len(resp.Data)
	if lenPersonsGot!=lenPersonsWant{
		t.Fatalf("Se esperaban %d personas, se obtuvieron %d",lenPersonsWant,lenPersonsGot)
	}
	firstNameWant:="Alexys"
	if resp.Data[0].Name!=firstNameWant{
		t.Fatalf("Se esperaban nombre %q , se obtuvo %q ",resp.Data[0].Name,firstNameWant)
	}
}
func cleanData(t *testing.T){
	store:=storage.NewPsql()
	defer store.CloseDB()

	err:= store.DeleteAll()
	if err!= nil{
		t.Fatalf("No se pudo limpiar la tabla %v",err)
	}
	
}

func insertData(t *testing.T){
	store:= storage.NewPsql()
	defer store.CloseDB()
	err:= store.Create(&model.Person{Name:"Alexys",Age:40})
	if err != nil {
		t.Fatalf("no se pudo registrar la persona %v",err)
	}
	err= store.Create(&model.Person{Name:"Mat",Age:41})
	if err != nil {
		t.Fatalf("no se pudo registrar la persona %v",err)
	}

}