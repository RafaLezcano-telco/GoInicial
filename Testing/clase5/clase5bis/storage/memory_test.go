package storage

import "testing"

func TestCreate_empty_person(t *testing.T) {
	m := NewMemory()
	err:= m.Create(nil)
	if err== nil{
		t.Errorf("Se esperaba error, se obtuvo nil")

	}
	if err !=
}