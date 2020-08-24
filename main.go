package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	tm "github.com/buger/goterm"
)

const cell = "x"
const space = " "

type Conway struct {
	Size   int
	State  [][]string
	Living int
	Dead   int
}

func main() {
	tm.Clear() // Clear current screen
	iteration := 0
	c := born()
	var moved bool
	for {
		// By moving cursor to top-left position we ensure that console output
		// will be overwritten each time, instead of adding new.
		tm.MoveCursor(1, 1)

		show(c)
		c, moved = live(c)
		if moved {
			iteration++
		}
		tm.Println("Iteration: ", iteration)
		if !moved {
			tm.Println("End of simulation")
			break
		}
		tm.Flush() // Call it every time at the end of rendering
		time.Sleep(time.Second * 5)
	}
}

func getRandomForms() []string {
	return []string{"2-2", "3-3", "4-1", "4-2", "4-3"}
}

func born() Conway {
	c := Conway{Size: 10}

	rc := getRandomForms()

	c.State = make([][]string, c.Size)
	for i := 0; i < c.Size; i++ {
		c.State[i] = make([]string, c.Size)
	}

	for i, innerArray := range c.State {
		for j := range innerArray {
			s := strconv.Itoa(i) + "-" + strconv.Itoa(j)
			if isIn(rc, s) {
				c.State[i][j] = cell
			} else {
				c.State[i][j] = space
			}
		}
	}

	return c
}

// func born() Conway {
// 	c := Conway{Size: 10}

// 	rc := getRandomCells(10, c.Size)

// 	c.State = make([][]string, c.Size)
// 	for i := 0; i < c.Size; i++ {
// 		c.State[i] = make([]string, c.Size)
// 	}

// 	for i, innerArray := range c.State {
// 		for j := range innerArray {
// 			s := string(i) + "-" + string(j)
// 			if isIn(rc, s) {
// 				c.State[i][j] = cell
// 			} else {
// 				c.State[i][j] = space
// 			}
// 		}
// 	}

// 	return c
// }

func getRandomCells(q int, s int) []string {
	rand.Seed(time.Now().UTC().UnixNano())
	rc := make([]string, q)
	for i := 0; i < q; i++ {
		for {
			r1 := rand.Intn(s)
			r2 := rand.Intn(s)
			s := string(r1) + "-" + string(r2)
			if !isIn(rc, s) {
				rc = append(rc, s)
				break
			}
		}
	}
	return rc
}

func isIn(s []string, v string) bool {
	fmt.Println(s, v)
	for _, val := range s {
		if val == v {
			return true
		}
	}
	return false
}

func live(c Conway) (Conway, bool) {
	nS := make([][]string, c.Size)
	copy(nS, c.State)
	nS = append(c.State[:0:0], c.State...)
	moved := false
	for i, innerArray := range c.State {
		for j := range innerArray {
			cl := nS[i][j]
			lc := getLivingCells(nS, c.Size, i, j)
			if cl == cell {
				if !(lc == 2 || lc == 3) {
					moved = true
					c.State[i][j] = space
				}
			} else {
				if lc == 3 {
					moved = true
					c.State[i][j] = cell
				}
			}
		}
	}

	// copy(c.State, nS)

	return c, moved
}

func getLivingCells(c [][]string, s int, i int, j int) int {
	livingCells := 0

	for k := -1; k <= 1; k++ {
		for l := -1; l <= 1; l++ {
			ni := i + k
			nj := j + l

			if (ni < 0 || nj < 0) || (ni > s-1 || nj > s-1) {
				continue
			}

			if ni == i && nj == j {
				continue
			}

			if c[ni][nj] == cell {
				livingCells++
			}
		}
	}

	return livingCells
}

func show(c Conway) {
	for i, innerArray := range c.State {
		row := ""
		for j := range innerArray {
			row += c.State[i][j] + " "
		}
		tm.Println(row)
	}
}
