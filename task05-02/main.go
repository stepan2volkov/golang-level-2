package main

import (
	"fmt"
	"task05/set"
)

// 2. Реализуйте функцию для разблокировки мьютекса с помощью defer
func main() {
	var ms set.Set
	ms = set.NewMRWMutexSet()
	ms.Add(17)
	fmt.Println("Set has 7?", ms.Has(7))
	fmt.Println("Set has 17?", ms.Has(17))
}
