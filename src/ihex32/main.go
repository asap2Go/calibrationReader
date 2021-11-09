package ihex32

import (
	"fmt"
	"time"
)

func main() {
	var liste = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	var toCalc []string
	for _, a := range liste {
		for _, b := range liste {
			val := a + b
			toCalc = append(toCalc, val)
		}
	}

	start := time.Now()
	for _, elem := range toCalc {
		a, err := hexToByte(elem)
		if err != nil {
			fmt.Println(a)
		}
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(elapsed.Nanoseconds() / int64(len(toCalc)))

	start = time.Now()
	for _, elem := range toCalc {
		a, err := hexToByteSlice(elem)
		if err != nil {
			fmt.Println(a)
		}
	}
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Println(elapsed.Nanoseconds() / int64(len(toCalc)))

}
