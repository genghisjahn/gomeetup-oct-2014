package main

import (
	"fmt"
	"strconv"
)

type Thing struct {
	Id   int
	Name string
}
type ThingErr struct {
	Thing
}

func main() {
	thing := Thing{5, "Item"}
	if num, err := strconv.Atoi(thing); err != nil {
		PrintError(err)
	} else {
		fmt.Println("Num:", num)
	}
}

func PrintError(err error) {
	fmt.Printf("Error: %v\n", err)
}
