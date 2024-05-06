package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertToStringAndConcat(intSlice []int) string {
	var reserveSlice []string
	for _, number := range intSlice {
		strInt := strconv.Itoa(number)
		reserveSlice = append(reserveSlice, strInt)
	}
	concatString := strings.Join(reserveSlice, "")
	return concatString
}

func decode(data string) (string, error) {
	dataSlice := strings.Split(data, "")
	var tempDecode []int
	breakpoint := 0

	tempDecode = append(tempDecode, 0)

	for index := range data {

		if dataSlice[index] == "L" {
			if tempDecode[index] == 0 {
				for j := len(tempDecode) - 1; j >= breakpoint; j-- {
					tempDecode[j] += 1
				}
			}
			tempDecode = append(tempDecode, 0)
		}

		if dataSlice[index] == "R" {
			breakpoint = index + 1
			next := tempDecode[index] + 1
			tempDecode = append(tempDecode, next)
		}

		if dataSlice[index] == "=" {
			current := tempDecode[index]
			tempDecode = append(tempDecode, current)
		}
	}

	result := convertToStringAndConcat(tempDecode)
	return result, nil
}

func main() {
	a := "LLRR="
	b := "==RLL"
	c := "=LLRR"
	d := "RRL=R"
	resultA, _ := decode(a)
	resultB, _ := decode(b)
	resultC, _ := decode(c)
	resultD, _ := decode(d)

	fmt.Printf("Decode %v to %v\n", a, resultA)
	fmt.Printf("Decode %v to %v\n", b, resultB)
	fmt.Printf("Decode %v to %v\n", c, resultC)
	fmt.Printf("Decode %v to %v\n", d, resultD)
}
