package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GenerateReports(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("couldn't open file ", err)
	}
	defer file.Close()

	result := make([][]int, 0)

	count := 0
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()

		s := strings.Split(line, " ")
		ints := make([]int, 0)
		for i := 0; i < len(s); i++ {
			val, err := strconv.Atoi(s[i])
			if err != nil {
				log.Fatal("error parsing report data")
			}
			ints = append(ints, val)
		}
		result = append(result, ints)
		count++
	}

	return result
}

func abs(val int) int {
	if val > 0 {
		return val
	} else {
		return -val
	}
}

func CheckReports(reports [][]int) int {
	safeCount := 0

	for i := 0; i < len(reports); i++ {
		fmt.Printf("Report %d: %v", i, reports[i])
		safe := isSafe(reports[i])
		if safe {
			safeCount++
		}
		fmt.Printf(" safecount: %d\n", safeCount)
	}
	return safeCount
}

func CheckDampen(reports [][]int) int {
	safeCount := 0
	for i := 0; i < len(reports); i++ {
		safe := isSafe(reports[i])
		if safe {
			safeCount++
		} else {
			errorLine := reports[i]
			fmt.Printf("Line: %v\n", errorLine)
			for j := 0; j < len(errorLine); j++ {
				var testReport []int
                for k := 0; k < len(errorLine); k++ {
                    if j != k {
                        testReport = append(testReport, errorLine[k])
                    }
                }
                dampSafe := safeWithDampen(testReport)
                if dampSafe {
                    safeCount++
                    break
                }
			}
		}
	}
	return safeCount
}

func isSafe(report []int) bool {
	down := false
	var diff int

	for i := 0; i < len(report)-1; i++ {
		v1 := report[i]
		v2 := report[i+1]

		if i == 0 {
			down = v1 > v2
		}
		if down {
			diff = v1 - v2
		} else {
			diff = v2 - v1
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func safeWithDampen(report []int) bool {
	if len(report) == 1 {
		return true
	}
	down := report[0] > report[1]
	for i := 0; i < len(report)-1; i++ {
		if report[i] == report[i+1] {
			return false
		}

		if down && report[i] < report[i+1] {
			return false
		}

		if !down && report[i] > report[i+1] {
			return false
		}
		diff := abs(report[i] - report[i+1])
		if diff > 3 || diff < 1 {
			return false
		}
	}
	return true
}
