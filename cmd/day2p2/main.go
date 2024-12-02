package main

import (
	"fmt"

	"github.com/davidgibsonmcr/adventofcode-2024/pkg/days"
)


func main() {
    reports := days.GenerateReports("./pkg/inputs/day2.txt")
    result := days.CheckDampen(reports)
    fmt.Printf("Result: %d", result)
}
