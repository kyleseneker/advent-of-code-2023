package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Part 1: ", getCalibrationValues("input.txt"))
	fmt.Println("Part 2: ", getCalibrationValues("input.txt"))
}

// getCalibrationValues retrieves the sum of all calibration values within a given input file
func getCalibrationValues(name string) int {
	input, _ := os.Open(name)
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	calibrationValues := 0

	for scanner.Scan() {
		firstDigit := getFirstDigit(scanner.Text())                    // get first digit (left to right)
		lastDigit := getLastDigit(scanner.Text())                      // get last digit (right to left)
		combined := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit) // combine first + last digit
		value, _ := strconv.Atoi(combined)                             // convert to integer
		calibrationValues += value                                     // add to running total
	}

	return calibrationValues
}

// getFirstDigit returns the first digit present in a given string, if one exists
func getFirstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		found, digit := containsSpelledDigit(s[:i])
		if found {
			return digit
		} else if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0')
		}
	}
	panic("No digit found in " + s)
}

// getLastDigit returns the last digit present in a given string, if one exists
func getLastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		found, digit := containsSpelledDigit(s[i:])
		if found {
			return digit
		} else if unicode.IsDigit(rune(s[i])) {
			fmt.Println(int(s[i]))
			fmt.Println(int(s[i] - '0'))
			return int(s[i] - '0')
		}
	}
	panic("No digit found in " + s)
}

var spelledDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// containsSpelledDigit validates if a given string contains a spelled digit
func containsSpelledDigit(s string) (bool, int) {
	for k, v := range spelledDigits {
		if strings.Contains(s, k) {
			return true, v
		}
	}
	return false, 0
}
