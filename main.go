package main

import (
	"fmt"
	"time"
)

// main prints out an iso datetime
func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
}
