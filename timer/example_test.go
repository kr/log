package timer_test

import (
	"github.com/kr/log/timer"
)

func ExampleTimer_Print() {
	defer timer.Start().Print()
	// time-consuming code
}

func ExampleTimer_Printf() {
	defer timer.Start().Printf("%.2f", .9999)
	// time-consuming code
}
