package main

import "fmt"

type Map map[int]interface{}

func ParseCard(card string) int {
	m := map[string]int{
		"ace":   11,
		"jack":  10,
		"queen": 10,
		"king":  10,
		"ten":   10,
		"nine":  9,
		"eight": 8,
		"seven": 7,
		"six":   6,
		"five":  5,
		"four":  4,
		"three": 3,
		"two":   2,
		"one":   1,
	}

	return m[card]

}

func IsBlackJack(num1 string, num2 string) bool {
	NumVal1 := ParseCard(num1)
	NumVal2 := ParseCard(num2)

	if NumVal1+NumVal2 == 21 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(ParseCard("ten"))
	fmt.Println(IsBlackJack("ten", "ace"))
}
