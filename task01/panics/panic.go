package panics

import (
	"fmt"
)

func MakeAndHandlePanic() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovering after panic")
		}
	}()

	// Declaring map, but not initialized
	var panicMap map[string]string
	panicMap["hello"] = "world"
}
