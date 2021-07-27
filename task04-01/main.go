package main

import "fmt"

const MAX_NUM = 1000

func sumByStep(firstNumber, secondNumber int) int {
	values := make(chan int)
	result := make(chan int)

	for i := 0; i < secondNumber; i++ {
		go func(values chan int, result chan int, id int) {
			value := <-values
			value += 1
			if id == secondNumber-1 {
				result <- value
			} else {
				values <- value
			}
		}(values, result, i)
	}
	values <- firstNumber

	return <-result
}

func main() {
	num := sumByStep(0, 1000)
	fmt.Println(num)
}
