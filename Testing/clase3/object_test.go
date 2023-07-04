package clase3

import (
	"reflect"
	"testing"
)

func TestPerro(t*testing.T){
	want :=&Perro{
		Name:"Firulais",
		Age:1,
		Kind:Kind{
			Name:"criollo",
		},
	}
	got :=&Perro{
		Name:"Firulaiss",
		Age:1,
		Kind:Kind{
			Name:"criollo",
		},
	}
	//t.Logf("want %p, got %p",want,got)
	if !reflect.DeepEqual(want,got){
		t.Errorf("Se esperaba %v, se esperaba %v",want,got)
	}
}