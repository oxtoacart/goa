package main

import (
	"fmt"

	"github.com/oxtoacart/gob"
)

func main() {
	fmt.Printf("B says %s\n", gob.CallMe())
	fmt.Printf("C says %s\n", gob.CallMe())
}
