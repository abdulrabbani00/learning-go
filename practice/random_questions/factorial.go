package main

import "fmt"

func FirstFactorial(num int) int {

	if num == 1 {
		return 1
	} else {
		ret := num * FirstFactorial(num-1)
		return ret
	}

}

func main() {

	fmt.Println(FirstFactorial(8))

}
