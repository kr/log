package logdur

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"strings"
	"text/template"
	"time"
)

// The default template string used to format output lines.
// See SetTemplate for a description of the template params.
const DefaultTemplate = `{{.Func}} {{.Elapsed}}{{range .V}} {{.}}{{end}}`

var (
	tpl = template.Must(template.New("t").Parse(DefaultTemplate))
)

// SetTemplate sets the template string for formatting output lines.
// It is executed with a value of the following type:
//   struct {
//       File    string        // position of Print's caller
//       Line    int           // position of Print's caller
//       Func    string        // short name of Print's calling function
//       Name    string        // full name of Print's calling function
//       Elapsed time.Duration // T1 - T0
//       T0      time.Time     // time of call to Start
//       T1      time.Time     // time of call to Print
//       V       []interface{} // values passed to Print
//   }
//
// If SetTemplate has not been called, DefaultTemplate is used.
func SetTemplate(s string) error {
	t, err := template.New("t").Parse(s)
	if err != nil {
		return err
	}
	tpl = t
	return nil
}

// Timer records the time between its creation and any
// subsequent call to Print or Printf.
type Timer struct {
	t0 time.Time
}

// Start returns a new Timer, measuring from the current time.
func Start() Timer {
	return Timer{time.Now()}
}

// Print prints the calling function's name and the elapsed time
// since t was created, followed by values v, separated by spaces.
// Output can be controlled with SetTemplate.
func (t Timer) Print(v ...interface{}) {
	logt(t.t0, time.Now(), v...)
}

// Print prints the calling function's name and the elapsed time
// since t was created, followed by values v, formatted according
// to format as in package fmt. Output can be further controlled
// with SetTemplate.
func (t Timer) Printf(format string, v ...interface{}) {
	logt(t.t0, time.Now(), fmt.Sprintf(format, v...))
}

func logt(t0, t1 time.Time, v ...interface{}) {
	var x struct {
		File    string
		Line    int
		Func    string
		Name    string
		Elapsed time.Duration
		T0, T1  time.Time
		V       []interface{}
	}
	x.Name, x.File, x.Line = caller(2)
	x.Func = x.Name[strings.LastIndex(x.Name, ".")+1:]
	x.T0 = t0
	x.T1 = t1
	x.Elapsed = t1.Sub(t0)
	x.V = v
	var buf bytes.Buffer
	err := tpl.Execute(&buf, x)
	if err != nil {
		fmt.Fprint(&buf, err)
	}
	log.Println(buf.String())
}

func caller(skip int) (name, file string, line int) {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return "func", "file", 0
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return "func", "file", 0
	}
	return f.Name(), file, line
}
