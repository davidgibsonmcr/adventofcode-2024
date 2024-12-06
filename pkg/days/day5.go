package days

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func FindMiddlePages(rules, input string) int {
	ruleMap := getRules(rules)
	file, err := os.Open(input)

	if err != nil {
		log.Fatal("couldn't open file ", err)
	}

	defer file.Close()

	inputs := make([][]int, 0)
	count := 0

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		splitLine := strings.Split(line, ",")
		var parsed []int
		for _, v := range splitLine {
			i, _ := strconv.Atoi(v)
			parsed = append(parsed, i)
		}
		inputs = append(inputs, parsed)
	}

	for _, v := range inputs {
		if ValidUpdate(v, ruleMap) {
			count += v[len(v)/2]
		}
	}
	return count
}

func FixMiddlePages(rules, input string) int {
	ruleMap := getRules(rules)
	file, err := os.Open(input)

	if err != nil {
		log.Fatal("couldn't open file ", err)
	}

	defer file.Close()

	inputs := make([][]int, 0)
	count := 0

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		splitLine := strings.Split(line, ",")
		var parsed []int
		for _, v := range splitLine {
			i, _ := strconv.Atoi(v)
			parsed = append(parsed, i)
		}
		inputs = append(inputs, parsed)
	}

	for _, v := range inputs {
		if !ValidUpdate(v, ruleMap) {
			fixed := FixUpdate(v, ruleMap)
			count += fixed[len(v)/2]
		}
	}
	return count
}

func FixUpdate(update []int, rules map[int][]int) []int {
	pos := make(map[int]int)
	for p, page := range update {
		pos[page] = p
	}

	for {
        swap := false
		for i := 0; i < len(update); i++ {
			page := update[i]
			before := rules[page]
			for j := i + 1; j < len(update); j++ {
				if slices.Contains(before, update[j]) {
                    swap = true
					update[i], update[j] = update[j], update[i]
				}
			}

		}
        if !swap {
            break
        }
	}
	return update
}

func ValidUpdate(update []int, rules map[int][]int) bool {
	pos := make(map[int]int)
	for p, page := range update {
		pos[page] = p
	}

	for key, before := range rules {
		if posKey, exists := pos[key]; exists {
			for _, bef := range before {
				if posBefore, exists := pos[bef]; exists {
					if posKey >= posBefore {
						return false

					}
				}
			}
		}
	}
	return true
}

func getRules(filePath string) map[int][]int {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("couldn't open file", err)
	}

	defer file.Close()

	rules := make(map[int][]int)

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := reader.Text()
		rule := strings.Split(line, "|")
		first, _ := strconv.Atoi(rule[0])
		second, _ := strconv.Atoi(rule[1])
		rules[first] = append(rules[first], second)
	}
	return rules
}
