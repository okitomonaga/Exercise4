package main

import (
	//"reflect"
	//"fmt"
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

func main() {

}
