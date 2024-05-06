package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	result, err := solvePuzzle("./day1/puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func solvePuzzle(s string) (int, error) {
	bs, err := os.Open(s)
	if err != nil {
		return 0, err
	}
	defer bs.Close()

	lettersRemoved, err := extractDigits(bs)
	if err != nil {
		return 0, err
	}
	newPuzzleInput := keepFirstAndLast(lettersRemoved)
	solution := addAllNumbers(newPuzzleInput)

	return solution, nil
}

func extractDigits(bs *os.File) ([]string, error) {
	var numbers []string
	var newString strings.Builder
	scanner := bufio.NewScanner(bs)
	for scanner.Scan() {
		text := scanner.Text()
		for _, char := range text {
			if unicode.IsNumber(char) {
				newString.WriteRune(char)
			}
		}
		numbers = append(numbers, newString.String())
		newString.Reset()
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return numbers, nil
}

func keepFirstAndLast(s []string) []string {
	var firstAndLast []string
	for _, number := range s {
		if len(number) > 0 {
			firstAndLast = append(firstAndLast, string(number[0])+string(number[len(number)-1]))
		}
	}

	return firstAndLast
}

func addAllNumbers(s []string) int {
	var total int
	for _, j := range s {
		if n, err := strconv.Atoi(j); err == nil {
			total += n
		} else {
			log.Printf("Error during string conversion: %v\n", err)
		}
	}

	return total
}
