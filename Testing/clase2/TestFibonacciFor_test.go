package clase2

import "testing"

func TestFibonacciRecMem(t*testing.T){
	want:=610
	got := FibonacciRecMem(15)
	if got!=want{
		t.Errorf("Se esperaba %d, se obtuvo %d", want,got)
	}
}

func BenchmarkFibonacciRecMem(b *testing.B){
	for n:=0;n<b.N; n++{
		FibonacciRecMem(30)
	}
}