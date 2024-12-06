package main

import (
	"fmt"

	"github.com/davidgibsonmcr/adventofcode-2024/pkg/days"
)


func main() {
    rules := days.FixMiddlePages("./pkg/inputs/day5.rules", "./pkg/inputs/day5.txt")
    fmt.Printf("%d\n", rules)
}
