package main

import (
	"crypto/rand"
	"fmt"
	"log"
)

func main() {
	fmt.Printf("%s\t%s\t%s\t%s\n", "ArraySize", "LoopCount", "MaxMinReached", "Frequency")
	for i := 1; i < 2016; i += 10 {
		n := 0
		loop := 1000
		for j := 0; j < loop; j++ {
			arr, err := getRandBytes(i)
			if err != nil {
				log.Fatalf("get rand bytes error: %v", err)
			}

			max, min := getMaxMin(arr)
			if satisfy(max, min) {
				n++
			}
		}
		fmt.Printf("%4d\t%4d\t%4d\t%f\t\n", i, loop, n, float32(n)/float32(loop))
	}

}

// generate an array of n byte, each ranges from 0~255 randomly
func getRandBytes(n int) (arr []byte, err error) {
	arr = make([]byte, n)
	_, err = rand.Read(arr)
	return
}

func getMaxMin(arr []byte) (max byte, min byte) {
	max, min = arr[0], arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min
}

func satisfy(max byte, min byte) bool {
	if max == 255 && min == 0 {
		return true
	}
	return false
}
