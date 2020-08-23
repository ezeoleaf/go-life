package main

import (
	"fmt"
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

	c := born()
	for {
		// By moving cursor to top-left position we ensure that console output
		// will be overwritten each time, instead of adding new.
		tm.MoveCursor(1, 1)

		show(c)
		tm.Flush() // Call it every time at the end of rendering

		time.Sleep(time.Second)
	}
}

func born() Conway {
	c := Conway{Size: 10}

	c.State = make([][]string, c.Size)
	for i := 0; i < c.Size; i++ {
		c.State[i] = make([]string, c.Size)
	}

	for i, innerArray := range c.State {
		for j := range innerArray {
			c.State[i][j] = " "
		}
	}

	return c
}

func live() {

}

func show(c Conway) {
	// for i := 0; i < c.Size; i++ {
	// 	row := ""
	// 	for j := 0; j < c.Size; j++ {
	// 		row += c.State[i][j]
	// 	}
	// 	tm.Println(row)
	// }
	// fmt.Println(c.State)
	for i, innerArray := range c.State {
		row := ""
		for j := range innerArray {
			row += c.State[i][j] + " "
		}
		fmt.Println(row)
		// tm.Println(row)
	}
}
