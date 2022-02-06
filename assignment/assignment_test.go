package assignment

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testAddUint32 struct {
	val1, val2, response uint32
	responseBool         bool
}

func TestAddUint32(t *testing.T) {
	/*
		Sum uint32 numbers, return uint32 sum value and boolean overflow flag
		cases need to pass:
			math.MaxUint32, 1 => 0, true
			1, 1 => 2, false
			42, 2701 => 2743, false
			42, math.MaxUint32 => 41, true
			4294967290, 5 => 4294967295, false
			4294967290, 6 => 0, true
			4294967290, 10 => 4, true
	*/

	var list = []testAddUint32{
		{val1: math.MaxUint32, val2: 1, response: 0, responseBool: true},
		{val1: 1, val2: 1, response: 2, responseBool: false},
		{val1: 42, val2: 2701, response: 2743, responseBool: false},
		{val1: 42, val2: math.MaxUint32, response: 41, responseBool: true},
		{val1: 4294967290, val2: 5, response: 4294967295, responseBool: false},
		{val1: 4294967290, val2: 6, response: 0, responseBool: true},
		{val1: 4294967290, val2: 10, response: 4, responseBool: true},
	}

	for i := 0; i < len(list); i++ {
		sum, overflow := AddUint32(list[i].val1, list[i].val2)

		assert.Equal(t, list[i].response, sum)
		assert.Equal(t, list[i].responseBool, overflow)

	}

}

func TestCeilNumber(t *testing.T) {
	/*
		Ceil the number within 0.25
		cases need to pass:
			42.42 => 42.50
			42 => 42
			42.01 => 42.25
			42.24 => 42.25
			42.25 => 42.25
			42.26 => 42.50
			42.55 => 42.75
			42.75 => 42.75
			42.76 => 43
			42.99 => 43
			43.13 => 43.25
	*/
	// map ile input ve output değerlerimizi oluşturduk.
	list := make(map[float64]float64)
	list[42.42] = 42.50
	list[42.00] = 42
	list[42.01] = 42.25
	list[42.24] = 42.25
	list[42.25] = 42.25
	list[42.26] = 42.50
	list[42.55] = 42.75
	list[42.75] = 42.75
	list[42.76] = 43
	list[42.99] = 43
	list[43.13] = 43.25

	//değerleri sırası ile test ettik
	for key, value := range list {
		point := CeilNumber(key)
		assert.Equal(t, value, point)
	}

}

func TestAlphabetSoup(t *testing.T) {
	/*
		String with the letters in alphabetical order.
		cases need to pass:
		 	"hello" => "ehllo"
			"" => ""
			"h" => "h"
			"ab" => "ab"
			"ba" => "ab"
			"bac" => "abc"
			"cba" => "abc"
	*/

	list := make(map[string]string)
	list["hello"] = "ehllo"
	list[""] = ""
	list["h"] = "h"
	list["ab"] = "ab"
	list["ba"] = "ab"
	list["bac"] = "abc"
	list["cba"] = "abc"

	//değerleri sırası ile test ettik
	for key, value := range list {
		result := AlphabetSoup(key)
		assert.Equal(t, value, result)
	}
}

type stringMaskTest struct {
	key, response string
	value         uint
}

func TestStringMask(t *testing.T) {
	/*
		Replace after n(uint) character of string with '*' character.
		cases need to pass:
			"!mysecret*", 2 => "!m********"
			"", n(any positive number) => "*"
			"a", 1 => "*"
			"string", 0 => "******"
			"string", 3 => "str***"
			"string", 5 => "strin*"
			"string", 6 => "******"
			"string", 7(bigger than len of "string") => "******"
			"s*r*n*", 3 => "s*r***"
	*/

	var list = []stringMaskTest{
		{key: "!mysecret*", value: 2, response: "!m********"},
		{key: "", value: 2, response: "*"},
		{key: "string", value: 0, response: "******"},
		{key: "string", value: 3, response: "str***"},
		{key: "string", value: 5, response: "strin*"},
		{key: "string", value: 6, response: "******"},
		{key: "string", value: 7, response: "******"},
		{key: "s*r*n*", value: 3, response: "s*r***"},
	}

	for i := 0; i < len(list); i++ {
		result := StringMask(list[i].key, list[i].value)

		assert.Equal(t, list[i].response, result)
	}

}

type wordSplit struct {
	value, response string
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	/*
		Your goal is to determine if the first element in the array can be split into two words,
		where both words exist in the dictionary(words variable) that is provided in the second element of array.

		cases need to pass:
			[2]string{"hellocat",words} => hello,cat
			[2]string{"catbat",words} => cat,bat
			[2]string{"yellowapple",words} => yellow,apple
			[2]string{"",words} => not possible
			[2]string{"notcat",words} => not possible
			[2]string{"bootcamprocks!",words} => not possible
	*/

	var list = []wordSplit{
		{value: "hellocat", response: "hello,cat"},
		{value: "catbat", response: "cat,bat"},
		{value: "yellowapple", response: "yellow,apple"},
		{value: "", response: "not possible"},
		{value: "notcat", response: "not possible"},
		{value: "bootcamprocks!", response: "not possible"},
	}

	for i := 0; i < len(list); i++ {
		result := WordSplit([2]string{list[i].value, words})
		assert.Equal(t, list[i].response, result)
	}

}

func TestVariadicSet(t *testing.T) {
	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks! => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/
	set1 := VariadicSet(4, 2, 5, 4, 2, 4)

	set2 := VariadicSet("bootcamp", "rocks!", "really", "rocks!")

	set3 := VariadicSet(1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first")

	assert.Equal(t, []interface{}{4, 2, 5}, set1)

	assert.Equal(t, []interface{}{"bootcamp", "rocks!", "really"}, set2)

	assert.Equal(t, []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}, set3)
}
