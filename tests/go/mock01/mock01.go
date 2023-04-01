package mock01

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
	"unicode/utf8"
)

func basicBuff() {
	bchn := make(chan string, 3)

	bchn <- "hallo"
	bchn <- "hola"
	bchn <- "hello"

	for v := range bchn {
		fmt.Println(v)
	}

}
func workerPool() {
	const numberOfJobs = 5
	const numberOfWorkers = 3
	jobs := make(chan int, numberOfJobs)
	results := make(chan int, numberOfWorkers)

	var wg sync.WaitGroup

	wg.Add(numberOfWorkers)

	for w := 0; w < numberOfWorkers; w++ {
		go func (id int)  {
			worker(id,jobs,results)
			wg.Done()
		}(w)	
	}

	for j := 0; j < numberOfJobs; j++ {
		jobs<-j
	}
	close(jobs)

	wg.Wait()

	close(results)

	for r := range results {
		fmt.Println(r)
	}

}

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d started job %d\n ", id, job)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finished with job %d\n ", id, job)
		result<-job*2
	}
}
func buffchannel() {
	ch := make(chan string, 5)
	go buffProd(ch)
	go buffConsume(ch)
	time.Sleep(time.Second * 10)
	fmt.Println(" all done")
}
func buffProd(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("%d", i)
		fmt.Printf("produced %v\t", i)
	}
	close(ch)
}
func buffConsume(ch chan string) {
	for v := range ch {
		fmt.Println("received", v)
	}
}
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error: %v", r)
		}
	}()

	if b == 0 {
		panic("Division by zero")
	}
	return a / b, nil
}
func picnic() {
	val := true
	defer func() {
		fmt.Printf("%s\n", "bohot bura hua")
	}()
	if val {
		panic("bachao bachao")
	}
}
func Gor02() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func(n int) {
			ch <- n
		}(i)
	}
	for v := range ch {
		fmt.Printf("%d\t", v)
	}
	fmt.Println("done")
}
func WaitG() {
	var wg sync.WaitGroup
	wg.Add(2)
	go PrintSmg(&wg, "hallo")
	go PrintSmg(&wg, "world")

	wg.Wait()
	fmt.Println("all go routines done")

}
func PrintSmg(wg *sync.WaitGroup, s string) {
	defer wg.Done()
	time.Sleep(time.Millisecond)
	fmt.Println("Ok! ", s)
}

func Gor01() {
	go Mes("hallo")
	go Mes("World")
	time.Sleep(time.Nanosecond * 5)
	fmt.Printf("%s\n", "from main go routine")
}
func Mes(msg string) {
	fmt.Printf("/n%s/n", msg)
}
func FileOps() {
	var filepath string = "C:/work/learn/one/tests/go/mock01/mock01_test.go"
	data, err := os.ReadFile(filepath)
	if err == nil {
		fmt.Println(len(data))
	} else {
		fmt.Println(err)
	}

	fs, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer fs.Close()
		scanner := bufio.NewScanner(fs)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
			time.Sleep(time.Second)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading file ", err)
		}
	}
}

type Coords struct {
	X int
	Y int
}

func Loops() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\t", i)
	}
	x := 0
	for {

		fmt.Printf("%d\t", x)
		x++
		if x > 10 {
			break
		}
	}
}

func CordsMap() {
	smap := make(map[Coords]int, 0)
	smap[Coords{X: 1, Y: 2}] = 0
	for k, v := range smap {
		fmt.Println(k, v)
	}
}
func CMaps() {
	concurrentMaps := sync.Map{}
	concurrentMaps.Store("one", 1)
	val, exists := concurrentMaps.Load("one")
	if exists {
		fmt.Println(val)
	}
	concurrentMaps.Delete("one")
	concurrentMaps.Store("two", 2)
	concurrentMaps.Range(func(key, value any) bool {
		fmt.Printf("%v,%v", key, value)
		return true

	})

}
func aboutmaps() {
	table1 := make(map[int]string)
	loop := 1000000
	starTime := time.Now()
	for i := 0; i < loop; i++ {
		table1[i] = fmt.Sprintf("%d", i)
	}
	endtime := time.Since(starTime).Seconds()
	fmt.Println("time taken", endtime)
	val, exists := table1[5]
	if exists {
		fmt.Println(val)
	}
	fmt.Println(len(table1))
	delete(table1, 5)
	fmt.Println(len(table1))
	for k, v := range table1 {
		fmt.Print(k)
		fmt.Print(" ")
		fmt.Print(v)
	}
	fmt.Println()
}

func ArraynSlices() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("length of array", len(numbers), cap(numbers))

	letters := [3]rune{'a', 'b', 'c'}

	for _, v := range letters {
		fmt.Printf("%d\t", v)
	}

	fmt.Printf("%c\t", letters[2])

	var num []int = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(len(num), cap(num))
	var num2 []int
	num2 = append(num2, 1)
	fmt.Println(len(num2), cap(num2))

	slc1 := make([]int, 3, 4)
	slc1[0] = 1
	slc1 = append(slc1, 1)
	slc1 = append(slc1, 2)
	// slc1 = append(slc1, 3)
	fmt.Println(len(slc1), cap(slc1))
	for index, v := range slc1 {
		fmt.Println(index, v)
	}
	var slc2 []int = make([]int, len(slc1))
	copy(slc2, slc1[0:2])
	fmt.Println(len(slc2), cap(slc2))
	for index, v := range slc2 {
		fmt.Println(index, v)
	}

	combined := append(slc1, slc2...)
	fmt.Println(combined)

}

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
func unicodeTest() {
	greeting := "こんにちは, 世界!"
	for _, v := range greeting {
		fmt.Printf("%c\t", v)
	}
	r, sz := utf8.DecodeRuneInString(greeting)
	fmt.Println(r, sz)

	fmt.Println(greeting)
}
func dividebyZero(n int, m int) (int, error) {
	if m == 0 {
		return 0, errors.New("division by zero")
	}
	return n / m, nil
}
func DivideZeroCustom(n int, m int) (int, error) {
	if m == 0 {
		return 0, CError{Kode: "102", Message: "kahe re zero kahe"}
	} else {
		return n / m, nil
	}
}

func unitWork() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%d\t", i)
	}
}
func callUnitWork() {
	go unitWork()
	time.Sleep(time.Second * 7)
	fmt.Printf("%s\n", "done")
}

func produce(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
func consume(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
func testchan() {
	ch := make(chan int)
	go produce(ch)
	go consume(ch)
	fmt.Println("producer consumer done")
	// fmt.Scanln()
}

func simpleChannel() {
	ch := make(chan string)

	go func(s string) {
		time.Sleep(time.Second * 5)
		ch <- "hallo " + s
	}("world")

	val := <-ch
	fmt.Printf("%s", val)

}
func multipleInputSingleOutput() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go func(n int) {

			ch <- n
			time.Sleep(10 * time.Second)
		}(i)
	}

	for {
		fmt.Printf("%d\t", <-ch)
	}

}
func mazdoor(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("starting ", n)
	time.Sleep(2 * time.Second)
	fmt.Println("finished ", n)
}
func callmazdoor() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go mazdoor(i, &wg)
	}
	wg.Wait()
	fmt.Println("all done")
}

type Writer interface {
	Write(data []byte) (int, error)
}

type File struct {
	name string
}
type Network struct {
	IP string
}

func (n *Network) Write(data []byte) (int, error) {
	fmt.Println("writing to network ", n.IP, string(data))
	return len(data), nil
}
func (f *File) Write(data []byte) (int, error) {
	fmt.Printf("writing to file %s data %s", f.name, string(data))
	return len(data), nil
}

func saveData(w Writer, data []byte) {
	w.Write(data)
}
func callSavedata() {
	f := &File{name: "file1.ttx"}
	data := []byte("hallo world")
	n := &Network{IP: "localhost"}
	saveData(f, data)
	saveData(n, data)

}

type Current struct {
	Counter int
}

func (c Current) Increment() {
	c.Counter++
}
func (c *Current) IncrementPointer() {
	c.Counter++
}

func checkvaluepointer() {
	c := &Current{Counter: 0}
	fmt.Printf("%d", c.Counter)
	c.Increment()
	fmt.Printf("%d", c.Counter)
	c.IncrementPointer()
	fmt.Printf("%d", c.Counter)

}
