package mocking

import (
	"fmt"
	"io"
)

const countdownStart = 3

func Countdown(out io.Writer) {
	//fmt.Fprint(out, "3")
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")
}

/*
func main() {
	Countdown(os.Stdout)
}
 */