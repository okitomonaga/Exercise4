package main

import (
	//"reflect"
	//"fmt"
	"testing"
)

func TestConvert1(t *testing.T) {
	result := conrvetIntSlice("1,2")
	if result[0] != 1 || result[1] != 2 { //resultと"hoge"の型を比較している
		t.Error("TestConvert1 is failed")
	}
}

func TestConvert2(t *testing.T) {
	result := conrvetIntSlice("2,3")
	if result[0] != 2 || result[1] != 3 { //resultと"hoge"の型を比較している
		t.Error("TestConvert2 is failed")
	}
}

func TestConvert3(t *testing.T) {
	result := conrvetIntSlice("2")
	if result[0] != -1 || result[1] != -1 { //resultと"hoge"の型を比較している
		t.Error("TestConvert3 is failed")
	}
}

func TestConvert4(t *testing.T) {
	result := conrvetIntSlice("e,3")
	if result[0] != -1 || result[1] != -1 { //resultと"hoge"の型を比較している
		t.Error("TestConvert4 is failed")
	}
}

func TestStoreBan1(t *testing.T) {
	result := storeBan([]int{1, 2}, 1, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}})
	if result[1][2] != 1 { //ox関係なく、おかれているかどうかを判定
		t.Error("TestStoreBan1 is failed")
	}
}
func TestStoreBan2(t *testing.T) {
	result := storeBan([]int{1, 1}, 2, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}})
	if result[1][1] != 2 { //ox関係なく、おかれているかどうかを判定
		t.Error("TestStoreBan2 is failed")
	}
}

func TestGenerateBanString1(t *testing.T) {
	ban := [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	result := generateBanString(ban)

	if result[0] != "o.." || result[1] != "..." || result[2] != "..." {
		t.Error("TestGenerateBanString1 is failed")
	}
}
func TestGenerateBanString2(t *testing.T) {
	ban := [][]int{
		{1, 1, 0},
		{0, 2, 0},
		{0, 0, 0},
	}
	result := generateBanString(ban)

	if result[0] != "oo." || result[1] != ".x." || result[2] != "..." {
		t.Error("TestGenerateBanString2 is failed")
	}
}

func TestCanPut1(t *testing.T) {
	ban := [][]int{
		{1, 1, 0},
		{0, 2, 0},
		{0, 0, 0},
	}
	input := []int{1, 1}
	result := canPut(input, ban)

	if result != false {
		t.Error("TestCanPut1 is failed")
	}
}

func TestCanPut2(t *testing.T) {
	ban := [][]int{
		{1, 1, 0},
		{0, 2, 0},
		{0, 0, 0},
	}
	input := []int{2, 2}
	result := canPut(input, ban)

	if result != true {
		t.Error("TestCanPut2 is failed")
	}
}

func TestCanPut3(t *testing.T) {
	ban := [][]int{
		{1, 1, 0},
		{0, 2, 0},
		{0, 0, 0},
	}
	input := []int{3, 3}
	result := canPut(input, ban)

	if result != false {
		t.Error("TestCanPut3 is failed")
	}
}

func TestCanPut4(t *testing.T) {
	ban := [][]int{
		{1, 1, 0},
		{0, 2, 0},
		{0, 0, 0},
	}
	input := []int{-1, -1}
	result := canPut(input, ban)

	if result != false {
		t.Error("TestCanPut4 is failed")
	}
}

func TestIsWin1(t *testing.T) {
	ban := [][]int{
		{1, 1, 1},
		{0, 0, 0},
		{0, 0, 0},
	}
	result := isWin(ban)

	if result != true {
		t.Error("TestIsWin1 is failed")
	}
}

func TestIsWin2(t *testing.T) {
	ban := [][]int{
		{0, 1, 1},
		{0, 0, 0},
		{0, 0, 0},
	}
	result := isWin(ban)

	if result != false {
		t.Error("TestIsWin2 is failed")
	}
}

func TestIsWin3(t *testing.T) {
	ban := [][]int{
		{1, 0, 0},
		{1, 0, 0},
		{1, 0, 0},
	}
	result := isWin(ban)

	if result != true {
		t.Error("TestIsWin3 is failed")
	}
}

func TestIsWin4(t *testing.T) {
	ban := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	result := isWin(ban)

	if result != true {
		t.Error("TestIsWin4 is failed")
	}
}

func TestIsWin5(t *testing.T) {
	ban := [][]int{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
	}
	result := isWin(ban)

	if result != true {
		t.Error("TestIsWin5 is failed")
	}
}
