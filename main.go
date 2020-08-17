package main

import (
	"time"

	tm "github.com/buger/goterm"
)

type Conway struct {
	Size   int
	State  [][]string
	Living int
	Dead   int
}

func main() {
	tm.Clear() // Clear current screen

	c := Conway{Size: 10}

	c.born()

	for {
		// By moving cursor to top-left position we ensure that console output
		// will be overwritten each time, instead of adding new.
		tm.MoveCursor(1, 1)

		c.show()

		tm.Flush() // Call it every time at the end of rendering

		time.Sleep(time.Second)
	}
}

func (c Conway) born() {
	for i := 0; i < c.Size; i++ {
		for j := 0; j < c.Size; j++ {
			c.State[i][j] = " "
		}
	}
}

func live() {

}

func (c Conway) show() {
	for i := 0; i < c.Size; i++ {
		row := ""
		for j := 0; j < c.Size; j++ {
			row += c.State[i][j]
		}
		tm.Println(row)
	}
}
