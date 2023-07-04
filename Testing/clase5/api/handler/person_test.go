package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestPerson_Create_wrong_structure(t *testing.T){
	data:= []byte(`{"name":123,"age":18}`)
	w:= httptest.NewRecorder()
	r:= httptest.NewRequest(http.MethodPost,"/",bytes.NewBuffer(data))
	r.Header.Set("Content-Type","application/json")
	e:=echo.New()
	ctx:=e.NewContext(r,w)

	p:=newPerson(&StorageMockOk{})
	err:= p.create(ctx)
	if err != nil{
		t.Errorf("no se esperaba error, pero se obtuvo %v",err)
	}
	

}