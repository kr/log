package timer

import (
	"testing"
)

func TestPrint(t *testing.T) {
	defer Start().Print("arg")
}

func TestPrintMulti(t *testing.T) {
	tm := Start()
	tm.Print("1")
	tm.Print("2")
}

func TestPrintf(t *testing.T) {
	defer Start().Printf("%.2f", .9999)
}

func TestDummy(t *testing.T) {
	dummy(0).method()
}

type dummy int

func (dummy) method() {
	defer Start().Print()
}
