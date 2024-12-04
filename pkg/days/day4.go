package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FindXmas(filePath string) int {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("couldn't open file ", err)
	}

	defer file.Close()

	grid := make([]string, 0)

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		grid = append(grid, line)
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'X' {
				count += checkSurround(grid, i, j, rows, cols)
			}
		}
	}

	return count
}

func checkSurround(grid []string, x, y, rows, cols int) int {
	index := 0
	xmas := "MAS"
	found := 0
	directions := getDirections()
	currentDir := 0
	tempX := x
	tempY := y
	for i := 0; i < 8; i++ {
		tempX = x + directions[currentDir][0]
		tempY = y + directions[currentDir][1]
		if tempX < 0 || tempX >= rows || tempY < 0 || tempY >= cols {
			currentDir++
			continue
		}

		if grid[tempX][tempY] == xmas[index] {
			tempX += directions[currentDir][0]
			tempY += directions[currentDir][1]
			if tempX < 0 || tempX >= rows || tempY < 0 || tempY >= cols {
				currentDir++
				continue
			}
			if grid[tempX][tempY] == xmas[index+1] {
				tempX += directions[currentDir][0]
				tempY += directions[currentDir][1]
				if tempX < 0 || tempX >= rows || tempY < 0 || tempY >= cols {
					currentDir++
					continue
				}
				if grid[tempX][tempY] == xmas[index+2] {
					fmt.Printf("Found! %d %d\n", x, y)
					found++
				}

			}

		}

		currentDir++
	}
	return found

}
func FindMas(filePath string) int {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("couldn't open file ", err)
	}

	defer file.Close()

	grid := make([]string, 0)

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		grid = append(grid, line)
	}
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'A' {
				count += checkMas(grid, i, j, rows, cols)
			}
		}
	}

	return count
}

func checkMas(grid []string, x, y, rows, cols int) int {

	if x > 0 && y > 0 && x < rows-1 && y < cols-1 {
		topLeft := grid[x-1][y-1]
		topRight := grid[x-1][y+1]
		bottomLeft := grid[x+1][y-1]
		bottomRight := grid[x+1][y+1]
		if ((topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M')) && ((topRight == 'M' && bottomLeft == 'S') || (topRight == 'S' && bottomLeft == 'M')) {
            fmt.Println("Found Mas!")
			return 1
		}
	}
    fmt.Println("Didn't Find mas")
	return 0
}

func getDirections() [][]int {
	return [][]int{
		//UP = 0
		{-1, 0},
		//UPRIGHT = 1
		{-1, 1},
		//RIGHT = 2
		{0, 1},
		//DOWNRIGHT = 3
		{1, 1},
		//DOWN = 4
		{1, 0},
		//DOWNLEFT = 5
		{1, -1},
		//LEFT = 6
		{0, -1},
		//UPLEFT = 7
		{-1, -1},
	}
}
