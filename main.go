package main

import (
	"fmt"

	caloriecounter "sambenning.dev/advent-of-code-2022/calorie_counter"
)

func main() {

	elves := caloriecounter.ScanInputFile("./calorie_counter/input.txt")

	highestElf := caloriecounter.GetElfWithMostCalories(elves)

	fmt.Println(highestElf)

}
