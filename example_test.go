package logdur_test

import (
	"github.com/kr/logdur"
)

func ExampleTimer_Print() {
	defer logdur.Start().Print()
	// time-consuming code
}

func ExampleTimer_Printf() {
	defer logdur.Start().Printf("%.2f", .9999)
	// time-consuming code
}
