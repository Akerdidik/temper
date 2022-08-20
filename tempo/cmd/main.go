package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"quad"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	output := ""
	hasNextEl := false
	allElements := 0
	y := 0
	for _, value := range input {
		output += string(value)
		if value == '\n' {
			y++
		} else {
			allElements++
		}
	}
	if y == 0 || allElements == 0 {
		fmt.Print("Not a Raid function")
		return
	}
	x := allElements / y
	if quad.QuadA(x, y) == output {
		hasNextEl = true
		fmt.Printf("[quadA] [%d] [%d]", x, y)
	}
	if quad.QuadB(x, y) == output {
		if hasNextEl {
			fmt.Print(" || ")
		}
		hasNextEl = true
		fmt.Printf("[quadB] [%d] [%d]", x, y)
	}
	if quad.QuadC(x, y) == output {
		if hasNextEl {
			fmt.Print(" || ")
		}
		hasNextEl = true
		fmt.Printf("[quadC] [%d] [%d]", x, y)
	}
	if quad.QuadD(x, y) == output {
		if hasNextEl {
			fmt.Print(" || ")
		}
		hasNextEl = true
		fmt.Printf("[quadD] [%d] [%d]", x, y)
	}
	if quad.QuadE(x, y) == output {
		if hasNextEl {
			fmt.Print(" || ")
		}
		hasNextEl = true
		fmt.Printf("[quadE] [%d] [%d]", x, y)
	}
	if !hasNextEl {
		fmt.Print("Not a Raid function")
	}
}
