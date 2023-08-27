package caloriecounter

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	inventorySize int
	calorieTotal  int
}

func ScanInputFile(filePath string) []Elf {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("could not open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var currentInventory int
	var calorieCount int

	var elves []Elf

	for scanner.Scan() {

		line := strings.Trim(scanner.Text(), " ")

		if line == "" {
			elves = append(elves, Elf{
				inventorySize: currentInventory,
				calorieTotal:  calorieCount,
			})
			currentInventory = 0
			calorieCount = 0
			continue
		}

		calorieVal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("could not convert line value to integer")
		}

		currentInventory++
		calorieCount += calorieVal

	}

	return elves

}

func GetElfWithMostCalories(elves []Elf) *Elf {

	var elfPtr *Elf

	for i, e := range elves {
		if i == 0 {
			elfPtr = &elves[0]
		} else {
			if elfPtr.calorieTotal < e.calorieTotal {
				elfPtr = &elves[i]
			}
		}
	}

	return elfPtr

}
