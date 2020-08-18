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
		// fmt.Println(c.State)
		tm.Flush() // Call it every time at the end of rendering

		time.Sleep(time.Second)
	}
}

func (c Conway) born() {
	initialState := make([][]string, c.Size)
	for i := 0; i < c.Size; i++ {
		initialState[i] = make([]string, c.Size)
	}

	for i, innerArray := range c.State {
		for j := range innerArray {
			initialState[i][j] = "x"
		}
	}

	c.State = initialState
}

func live() {

}

func (c Conway) show() {
	// for i := 0; i < c.Size; i++ {
	// 	row := ""
	// 	for j := 0; j < c.Size; j++ {
	// 		row += c.State[i][j]
	// 	}
	// 	tm.Println(row)
	// }
	for i, innerArray := range c.State {
		row := ""
		for j := range innerArray {
			row += c.State[i][j]
			tm.Println(c.State[i][j])
		}
		tm.Println(row)
		tm.Println("x")
	}
	// fmt.Println(c.State)
}
