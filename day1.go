package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`[^\d]+`)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers := re.ReplaceAllString(line, "")

		firstNumber := numbers[0] - '0'
		secondNumber := numbers[len(numbers)-1] - '0'

		num := firstNumber*10 + secondNumber
		sum += int(num)
	}

	fmt.Printf("%d\n", sum)
}
