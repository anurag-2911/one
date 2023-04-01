package mock01

import (
	"fmt"
	"testing"
)

func TestWPool(t *testing.T) {
	workerPool()
}
func TestBBuff(t *testing.T) {
	basicBuff()
}
func TestBufchn(t *testing.T) {
	buffchannel()
}
func TestSafeDev(t *testing.T) {
	res, err := safeDivide(10, 2)
	fmt.Println(res, err)

	res, err = safeDivide(10, 0)
	fmt.Println(res, err)

	fmt.Println("Normal execution")

}
func TestPanic(t *testing.T) {
	picnic()
}
func TestGor2(t *testing.T) {
	Gor02()
}
func TestWg(t *testing.T) {
	WaitG()
}
func TestGor(t *testing.T) {
	Gor01()
}
func TestFileOps(t *testing.T) {
	FileOps()
}

func TestLoops(t *testing.T) {
	Loops()
}
func TestCMaps(t *testing.T) {
	CMaps()
}
func TestSmap(t *testing.T) {
	CordsMap()
}
func TestMaps(t *testing.T) {
	aboutmaps()
}
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
