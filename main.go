package main

import (
	"math/rand"
	"strconv"
	"time"

	tm "github.com/buger/goterm"
)

const cell = "x"
const space = " "

var population = 10

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
		tm.Println("Iteration: ", iteration)
		tm.Println("Population: ", population)
		c, moved = live(c)
		if moved {
			iteration++
		}
		tm.Flush()
		if !moved {
			tm.Println("End of simulation")
			break
		}
		time.Sleep(time.Second * 1)
	}
}

func born() Conway {
	c := Conway{Size: 10}

	rc := getRandomCells(population, c.Size)

	c.State = make([][]string, c.Size)
	for i := 0; i < c.Size; i++ {
		c.State[i] = make([]string, c.Size)
	}

	for i, innerArray := range c.State {
		for j := range innerArray {
			p := strconv.Itoa(i) + "-" + strconv.Itoa(j)
			if isIn(rc, p) {
				c.State[i][j] = cell
			} else {
				c.State[i][j] = space
			}
		}
	}

	return c
}

func getRandomCells(q int, s int) []string {
	rand.Seed(time.Now().UTC().UnixNano())
	rc := make([]string, q)
	for i := 0; i < q; i++ {
		for {
			r1 := rand.Intn(s)
			r2 := rand.Intn(s)
			p := strconv.Itoa(r1) + "-" + strconv.Itoa(r2)
			if !isIn(rc, p) {
				rc[i] = p
				break
			}
		}
	}
	return rc
}

func isIn(s []string, v string) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}
	return false
}

func getCells(state [][]string) []string {
	cells := []string{}
	for i, innerArray := range state {
		for j := range innerArray {
			if state[i][j] == cell {
				p := strconv.Itoa(i) + "-" + strconv.Itoa(j)
				cells = append(cells, p)
			}
		}
	}

	return cells
}

func live(c Conway) (Conway, bool) {
	cells := getCells(c.State)
	moved := false
	for i, innerArray := range c.State {
		for j := range innerArray {
			cl := c.State[i][j]
			lc := getLivingCells(cells, c.Size, i, j)
			if cl == cell {
				if !(lc == 2 || lc == 3) {
					moved = true
					c.State[i][j] = space
					population--
				}
			} else {
				if lc == 3 {
					moved = true
					c.State[i][j] = cell
					population++
				}
			}
		}
	}

	return c, moved
}

func getLivingCells(cells []string, s int, i int, j int) int {
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

			s := strconv.Itoa(ni) + "-" + strconv.Itoa(nj)
			if isIn(cells, s) {
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
