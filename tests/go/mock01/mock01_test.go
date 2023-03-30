package mock01

import (
	"fmt"
	"testing"
)

func TestArraynSlices(t *testing.T) {
	ArraynSlices()
}

func TestMockfn(t *testing.T) {
	// callUnitWork()
	// unicodeTest()
	// mock01Functions()
	testchan()
}
func TestPointernValue(t *testing.T) {
	checkvaluepointer()
}
func TestMore(t *testing.T) {
	multipleInputSingleOutput()
	// simpleChannel()
}
func TestIWrite(t *testing.T) {
	callSavedata()
}
func TestMazdoor(t *testing.T) {
	callmazdoor()
}
func TestDoSomething(t *testing.T) {
	err := DoSomething(12)
	if err != nil {
		fmt.Println("err is ", err.Error())
		t.Errorf("error is %s", err.Error())
	}
}
func TestDivideByZero(t *testing.T) {
	_, err := dividebyZero(10, 0)
	if err != nil {
		t.Errorf("%s", err.Error())
		fmt.Printf("%s", err.Error())
	}
	_, err = DivideZeroCustom(101, 0)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}
