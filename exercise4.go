package main

import (
	//"reflect"
	"fmt"
	"strconv"
	"strings"
)

func conrvetIntSlice(str string) []int {

	var result []int
	slice := strings.Split(str, ",")
	r1, _ := strconv.Atoi(slice[0])
	r2, _ := strconv.Atoi(slice[1])
	result = append(result, r1)
	result = append(result, r2)
	return result
}

func storeBan(input []int, player int) [][]int {

	result := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	result[input[0]][input[1]] = player //

	return result
}

func generateBanString(ban [][]int) []string {
	result := make([]string, 3)
	for h := 0; h < 3; h++ {
		for w := 0; w < 3; w++ {

			switch ban[h][w] {
			case 0:
				result[h] += "."
			case 1:
				result[h] += "o"
			case 2:
				result[h] += "x"
			default:
			}

		}
	}

	return result
}

func main() {
	inputPut := make([]int, 2)
	ban := make([][]int, 9)
	outputString := make([]string, 3)

	fmt.Print("Player 1: Input (x,y) : ")

	var input string
	fmt.Scan(&input)

	copy(inputPut, conrvetIntSlice(input))
	copy(ban, storeBan(inputPut, 1))
	copy(outputString, generateBanString(ban))
	for i := 0; i < len(outputString); i++ {
		fmt.Println(outputString[i])
	}
}
