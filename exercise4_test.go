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

func TestStoreBan1(t *testing.T) {
	result := storeBan([]int{1, 2}, 1)
	if result[1][2] != 1 { //ox関係なく、おかれているかどうかを判定
		t.Error("TestStoreBan1 is failed")
	}
}
func TestStoreBan2(t *testing.T) {
	result := storeBan([]int{1, 1}, 2)
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
