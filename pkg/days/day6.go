package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FindLoop(input string) int {
	file, err := os.Open(input)

	if err != nil {
		log.Fatal("couldn't open file ", err)
	}

	defer file.Close()

	inputs := make([][]byte, 0)

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		var parsed []byte
		for i := 0; i < len(line); i++ {
			parsed = append(parsed, line[i])
		}
		inputs = append(inputs, parsed)
	}

	x, y := findStartPosition(inputs)
	startx := x
	starty := y
	fmt.Printf("start x %d - y %d\n", x, y)
	loops := 0
	grid := make([][]byte, len(inputs))
    currentDir := byte('^')
	for i, row := range inputs {
		grid[i] = make([]byte, len(row))
		copy(grid[i], row)
	}

	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs[0]); j++ {
			pathMap := make(map[int]byte, 0)
            currentDir = '^'
            for i, row := range grid {
                copy(inputs[i], row)
            }
            inputs[i][j] = '#'
            x, y = startx, starty

			fmt.Printf("SIMULATION %d %d\n", i, j)
			for {
				fmt.Printf("XY DEBUG %d %d\n", x, y)
				bounds, next := peek(x, y, currentDir, inputs)
				if bounds {
					fmt.Println("hit bounds 1")
					break
				}
				loc := x*len(inputs[0]) + y
				if v, ok := pathMap[loc]; ok {
					if v == currentDir {
						loops++
						fmt.Printf("reset: %d %d %d\n", x, y, loops)
						break
					}

				} else {
					pathMap[loc] = currentDir
				}

				if next == '#' {
					currentDir = nextDir(currentDir)
					b, _ := peek(x, y, currentDir, inputs)
					if b {
						log.Fatal("Something went wrong")
					}

				}
				offX, offY := dirOffset(currentDir)
				x += offX
				y += offY
				b, _ := peek(x, y, currentDir, inputs)
				if b {
					fmt.Println("hit bounds 2")
					break
				}
			}
		}
	}

	return loops
}

func FindDistinctPath(input string) int {
	file, err := os.Open(input)

	if err != nil {
		log.Fatal("couldn't open file ", err)
	}

	defer file.Close()

	inputs := make([][]byte, 0)

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		var parsed []byte
		for i := 0; i < len(line); i++ {
			parsed = append(parsed, line[i])
		}
		inputs = append(inputs, parsed)
	}

	x, y := findStartPosition(inputs)
	fmt.Printf("start x %d - y %d\n", x, y)
	currentDir := byte('^')

	for {
		bounds, next := peek(x, y, currentDir, inputs)
		if bounds {
			break
		}

		if next == '#' {
			currentDir = nextDir(currentDir)
			b, _ := peek(x, y, currentDir, inputs)
			if b {
				log.Fatal("Something went wrong")
			}

		}
		inputs[x][y] = byte('X')
		offX, offY := dirOffset(currentDir)
		x += offX
		y += offY
		bounds, next = peek(x, y, currentDir, inputs)
		if bounds {
			break
		}
	}
	unique := 1
	for i := range inputs {
		for j := range inputs[0] {
			if inputs[i][j] == 'X' {
				unique++
			}
		}
	}
	return unique
}

func peek(x, y int, dir byte, inputs [][]byte) (bounds bool, next byte) {
	offX, offY := dirOffset(dir)
	peekX := x + offX
	peekY := y + offY
	if peekX >= len(inputs) || peekY >= len(inputs[0]) || peekX < 0 || peekY < 0 {
		return true, next
	}
	next = inputs[peekX][peekY]
	return false, next

}

func findStartPosition(input [][]byte) (int, int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '^' {
				input[i][j] = 'X'
				return i, j
			}
		}
	}
	return -1, -1
}

func nextDir(dir byte) byte {
	dirMap := make(map[byte]byte)
	dirMap['^'] = '>'
	dirMap['>'] = 'v'
	dirMap['v'] = '<'
	dirMap['<'] = '^'
	return byte(dirMap[dir])
}

func dirOffset(dir byte) (int, int) {
	dirMap := make(map[byte][]int)
	dirMap['^'] = []int{-1, 0}
	dirMap['>'] = []int{0, 1}
	dirMap['v'] = []int{1, 0}
	dirMap['<'] = []int{0, -1}

	return dirMap[dir][0], dirMap[dir][1]
}
