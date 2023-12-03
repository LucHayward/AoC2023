package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Lines of text contain calibration values
// Each line originally contained a **calibration value** that must be recovered.
// Combine the first digit and the last digit to make a 2-digit number
// Eg:
// 1abc2 => 12
// pqr3stu8vwx -> 38
// a1b2c3d4e5f -> 15
// treb7uchet -> 77
// Find the sum of these values => 142
func main() {
	// Read in a file
	//testData := [4]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}

	file, err := os.Open("1/input.txt")
	check(err)
	defer file.Close()

	var nums []int
	var result int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			if 48 <= c && c <= 57 { // int
				nums = append(nums, int(c-48))
			}
		}
		i, err := strconv.Atoi(strconv.Itoa(nums[0]) + strconv.Itoa(nums[len(nums)-1]))
		check(err)
		result += i

		nums = nums[:0]
	}
	fmt.Println(result)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
