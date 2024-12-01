package days

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func GenerateLists(filePath string) ([]int, []int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("couldn't open file ", err)
	}
	defer file.Close()

	var list1, list2 []int

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		string1 := line[0:5]
		string2 := line[8:13]

		i1, err := strconv.Atoi(string1)
		if err != nil {
			log.Fatal("error parsing int1", i1)
		}
		i2, err := strconv.Atoi(string2)
		if err != nil {
			log.Fatal("error parsing int2", i2)
		}
		list1 = append(list1, i1)
		list2 = append(list2, i2)

	}
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})
	return list1, list2
}
func SimilarityScore(filePath string) int {

	list1, list2 := GenerateLists(filePath)
	sm := make(map[int]int)
	for i := 0; i < len(list2); i++ {
		sm[list2[i]]++
	}
	count := 0
	for i := 0; i < len(list1); i++ {
		count += list1[i] * sm[list1[i]]
	}
	return count
}

func CalculateDistance(filePath string) int {
	list1, list2 := GenerateLists(filePath)
	distance := 0
	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			distance += list1[i] - list2[i]
		} else {
			distance += list2[i] - list1[i]
		}
	}
	return distance
}
