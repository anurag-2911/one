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
