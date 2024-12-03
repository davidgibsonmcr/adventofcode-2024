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
		total += i1 * i2
	}
	return total

}

func FindWithConcat(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("couldn't open file ", err)
	}
	defer file.Close()
	var concat string
	result := make([][]string, 0)
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		concat += reader.Text()
	}
	halted := false
	for {
		foundDont, indexDont := checkDonts(concat)

		if foundDont {
			matched := re.FindAllStringSubmatch(concat[:indexDont], -1)
			result = append(result, matched...)
            concat = concat[indexDont:]
			halted = true
		}

		foundDo, indexDo := checkDos(concat)

		if foundDo {
            concat = concat[indexDo:]
            halted = false
		}

		if !foundDo && !foundDont {
			if halted {
				break
			} else {
				matched := re.FindAllStringSubmatch(concat, -1)
				result = append(result, matched...)
				break
			}
		}
	}

    total := 0

	for _, v := range result {
		i1, _ := strconv.Atoi(v[1])
		i2, _ := strconv.Atoi(v[2])
		total += i1 * i2
	}

	return total

}

func checkDonts(text string) (bool, int) {
	dontExp := regexp.MustCompile(`don't\(\)`)
	donts := dontExp.FindStringIndex(text)
	if donts == nil {
		return false, 0
	} else {
		return true, donts[1]
	}
}

func checkDos(text string) (bool, int) {
	doExp := regexp.MustCompile(`do\(\)`)
	dos := doExp.FindStringIndex(text)
	if dos == nil {
		return false, 0
	} else {
		return true, dos[1]
	}
}
