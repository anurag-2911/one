package mock01

import (
	"fmt"
	"testing"
)

func TestMockfn(t *testing.T) {
	unicodeTest()
	mock01Functions()
}
func TestDoSomething(t *testing.T) {
	err := DoSomething(12)
	if err != nil {
		fmt.Println("err is ",err.Error())
		t.Errorf("error is %s", err.Error())
	}
}
func TestDivideByZero(t *testing.T){
	_,err:=dividebyZero(10,0)
	if(err!=nil){
		t.Errorf("%s",err.Error())
		fmt.Printf("%s",err.Error())
	}
	_,err=DivideZeroCustom(101,0)
	if(err!=nil){
		fmt.Printf("%s",err.Error())
	}
}
