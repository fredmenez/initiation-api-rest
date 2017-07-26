package main

import (
	"fmt"
	"time"
)

func main() {
	dir := 1
	iter := 1

	str := ""

	for {
		fmt.Println(str + "*")

		if iter%80 == 0 {
			dir = 1 - dir
		}

		iter++

		if dir == 1 {
			str += " "
		} else if len(str) > 0 {
			str = str[:len(str)-1]
		}

		time.Sleep(30 * time.Millisecond)
	}
}
