package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	puzzlePiece, err := readPuzzleFile("./day1/puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	mapOfConvertedInts, err := removeAllNonDigits(puzzlePiece)
	if err != nil {
		log.Print(err)
	}

	fmt.Println(mapOfConvertedInts)

	strippedMap := joinFirstAndLastInt(mapOfConvertedInts)

	fmt.Println(strippedMap)

	sumOfAllInts := addAllInts(strippedMap)

	fmt.Println(sumOfAllInts)

}

func readPuzzleFile(s string) (map[int]string, error) {
	output, err := os.Open(s)
	if err != nil {
		return nil, err
	}
	defer output.Close()

	lines := make(map[int]string)
	var index int

	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		index++
		lines[index] = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil

}

func removeAllNonDigits(puzzlePiece map[int]string) (map[int]string, error) {
	if len(puzzlePiece) == 0 {
		err := errors.New("Puzzle Piece Map is empty!")
		log.Println(err)
	}

	intMap := make(map[int]string)

	for i := range puzzlePiece {
		for k := range puzzlePiece[i] {
			if unicode.IsDigit(rune(puzzlePiece[i][k])) == true {
				intMap[i] += string(puzzlePiece[i][k])
			}
		}
	}

	if len(intMap) == 0 {
		err := errors.New("No characters were converted to integers...")
		return nil, err
	}

	return intMap, nil
}

func joinFirstAndLastInt(ints map[int]string) map[int]string {
	strippedInts := make(map[int]string)
	for i := range ints {
		if len(ints[i]) > 2 {
			strippedInts[i] += string(ints[i][0])
			strippedInts[i] += string(ints[i][len(ints[i])-1])
		} else {
			strippedInts[i] += ints[i]
		}
	}

	return strippedInts
}

func addAllInts(ints map[int]string) int {
	var summedInts int
	for i := range ints {
		num, err := strconv.Atoi(ints[i])
		if err != nil {
			log.Println(err)
		}
		summedInts += num
	}

	return summedInts
}
