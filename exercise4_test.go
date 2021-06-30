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
