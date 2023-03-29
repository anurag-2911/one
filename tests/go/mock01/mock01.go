package mock01

import (
	"errors"
	"fmt"
	"os"
	"time"
	"unicode/utf8"
)

type CError struct {
	Kode    string
	Message string
}

func (cer CError) Error() string {
	return fmt.Sprintf("error is %s , %s", cer.Kode, cer.Message)
}

func DoSomething(val int) error {
	if val < 100 {
		return CError{Kode: "101", Message: "itne kam me kya hoga"}
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
func dividebyZero(n int, m int)(int,error){
	if(m==0){
		return 0,errors.New("division by zero")
	}
	return n/m, nil
}
func DivideZeroCustom(n int, m int)(int,error){
if(m==0){
	return 0,CError{Kode: "102",Message: "kahe re zero kahe"}
}else{
	return n/m,nil
}
}

func unitWork(){
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%d\t",i)
	}
}
func callUnitWork(){
	go unitWork()
	time.Sleep(time.Second*7)
	fmt.Printf("%s\n","done")
}

func produce(ch chan int ){
	for i := 0; i < 5; i++ {
		ch<-i
	}
	close(ch)
}
func consume(ch chan int){
	for i:=range ch{
		fmt.Println(i)
	}
}
func testchan(){
	ch:=make(chan int)
	go produce(ch)
	go consume(ch)
	fmt.Println("producer consumer done")
}