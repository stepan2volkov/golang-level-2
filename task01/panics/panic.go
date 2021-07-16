// Package panics implements a function for learning panic and recover
package panics

import (
	"fmt"
)

// MakeAndHandlePanic panics and recovers
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
