package main

import (
	"time"

	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear() // Clear current screen

	for {
		// By moving cursor to top-left position we ensure that console output
		// will be overwritten each time, instead of adding new.
		tm.MoveCursor(1, 1)

		rows := 10
		cols := 10

		for i := 1; i <= rows; i++ {
			row := ""
			for j := 1; j <= cols; j++ {
				row += "- "
			}
			tm.Println(row)
		}

		tm.Flush() // Call it every time at the end of rendering

		time.Sleep(time.Second)
	}
}
