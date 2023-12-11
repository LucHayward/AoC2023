package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	Lines of text contain calibration values

Each line originally contained a **calibration value** that must be recovered.
Combine the first digit and the last digit to make a 2-digit number
Eg:
1abc2 => 12
pqr3stu8vwx -> 38
a1b2c3d4e5f -> 15
treb7uchet -> 77
Find the sum of these values => 142

Benchmark 1: go run main.go
Time (mean ± σ):     353.9 ms ±  22.2 ms    [User: 196.7 ms, System: 187.6 ms]
Range (min … max):   335.0 ms … 410.4 ms    10 runs

Benchmark 1: ./main
Time (mean ± σ):      11.4 ms ±  29.5 ms    [User: 1.4 ms, System: 1.6 ms]
Range (min … max):     0.0 ms … 120.7 ms    16 runs
*/
func main() {
	// Read in a file
	//testData := [4]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}

	file, err := os.Open("input.txt")
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
