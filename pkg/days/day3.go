package days

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func FindInsructions(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("couldn't open file ", err)
	}
	defer file.Close()

	result := make([][]string, 0)
    total := 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
        match := re.FindAllStringSubmatch(line, -1)
        result = append(result, match...)
	}
    for _, v := range result {
        i1, _ := strconv.Atoi(v[1])
        i2, _ := strconv.Atoi(v[2])
        total += i1*i2
    }
    return total

}
