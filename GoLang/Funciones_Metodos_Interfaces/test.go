package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strconv"
)

func main() {
	a, vo, so, t := input()

	fmt.Println("")
	fmt.Println("Entered values:")
	fmt.Println("Acceleration:", a)
	fmt.Println("Initial velocity:", vo)
	fmt.Println("Initial displacement:", so)
	fmt.Println("Time (seconds):", t)

	fmt.Println("")

	fn := GenDisplaceFn(a, vo, so)
	displacement := fn(t)

	fmt.Printf("Displacement after %f seconds: %f", t, displacement)
	fmt.Println("")
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		result := (0.5 * a * math.Pow(t, 2)) + (vo * t) + so
		return result
	}
	return fn
}

func input() (float64, float64, float64, float64) {
	floatSize := 64
	scanner := bufio.NewReader(os.Stdin)

	fmt.Printf("Please enter acceleration: ")
	raw_a, _, _ := scanner.ReadLine()
	a, _ := strconv.ParseFloat(string(raw_a), floatSize)

	fmt.Printf("Please enter initial velocity: ")
	raw_vo, _, _ := scanner.ReadLine()
	vo, _ := strconv.ParseFloat(string(raw_vo), floatSize)

	fmt.Printf("Please enter initial displacement: ")
	raw_so, _, _ := scanner.ReadLine()
	so, _ := strconv.ParseFloat(string(raw_so), floatSize)

	fmt.Printf("Please enter time (seconds): ")
	raw_t, _, _ := scanner.ReadLine()
	t, _ := strconv.ParseFloat(string(raw_t), floatSize)

	return a, vo, so, t
}
