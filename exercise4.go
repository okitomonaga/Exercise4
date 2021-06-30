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

func storeBan(input []int, player int, ban [][]int) [][]int {

	result := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	//result := make([][]int, 9)

	copy(result, ban)

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
	//ban := make([][]int, 9)
	ban := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	outputString := make([]string, 3)

	var input string

	playernum := 1          ///先手から開始
	for i := 0; true; i++ { //無限ループ
		fmt.Printf("Player %d: Input (x,y) : ", playernum)
		fmt.Scan(&input)
		copy(inputPut, conrvetIntSlice(input))
		copy(ban, storeBan(inputPut, playernum, ban))
		if playernum == 1 {
			playernum = 2
		} else {
			playernum = 1
		}

		copy(outputString, generateBanString(ban))
		for i := 0; i < len(outputString); i++ {
			fmt.Println(outputString[i])
		}
		//勝利条件を満たしたときbreak
	}

	//fmt.Print("Player 1: Input (x,y) : ")

	//fmt.Scan(&input)

	//copy(inputPut, conrvetIntSlice(input))
	//copy(ban, storeBan(inputPut, 1))
	copy(outputString, generateBanString(ban))
	for i := 0; i < len(outputString); i++ {
		fmt.Println(outputString[i])
	}
}
