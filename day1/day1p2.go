package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	result, err := solvePuzzle("./day1/inputtest.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func solvePuzzle(s string) (int, error) {
	bs, err := os.ReadFile(s)
	if err != nil {
		return 0, err
	}

	lettersToNumbers := translateWordsToNumbers(bs)
	fmt.Println(lettersToNumbers)

	return 0, nil

}

func translateWordsToNumbers(s []byte) []string {
	var translateStrings []string
	strToNumMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for i := range s {

	}

	fmt.Println(translateStrings)

	return nil
}
