package mock01

import (
	"fmt"
	"os"
	"unicode/utf8"
)

type ZError struct {
	Kode    string
	Message string
}

func (er ZError) Error() string {
	return fmt.Sprintf("error is %s , %s", er.Kode, er.Message)
}

func DoSomething(val int) error {
	if val < 100 {
		return ZError{Kode: "101", Message: "itne kam me kya hoga"}
	}
	return nil
}
func mock01Functions() {
	args := os.Args
	for _, val := range args[1:] {
		fmt.Println(val)
	}

}
func unicodeTest(){
	greeting := "こんにちは, 世界!"
	for _, v := range greeting {
		fmt.Printf("%c\t",v)
	}
	r,sz:=utf8.DecodeRuneInString(greeting)
	fmt.Println(r,sz)

	fmt.Println(greeting)
}
