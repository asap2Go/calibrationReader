package main

import (
	"asap2Go/calibrationReader/a2l"
	"asap2Go/calibrationReader/ihex32"
	"fmt"
)

func main() {
	a, err := a2l.ParseFromFile("")
	h, err2 := ihex32.ParseFromFile("")

	if err == nil {
		fmt.Println(a.Project.Module[0].Characteristics["abc"])
	}
	if err2 == nil {
		fmt.Println(len(h.DataBytes))
	}
}
