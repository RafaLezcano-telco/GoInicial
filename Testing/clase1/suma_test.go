package clase1

import "testing"

func TestAdd(t *testing.T){
	want:=5
	got := Add(0,5)
	if got!=want{
		t.Errorf("Error:se esperaba %d, se obtuvo %d",want,got)
	}
}

func TestAddAcum(t *testing.T){
	want:=6
	got:=AddAcum(1,2,3)
	if(got != want){
		t.Errorf("Error:se esperaba %d, se obtuvo %d",want,got)
		t.Fail()
	}
}