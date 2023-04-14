package main

import (
	"fmt"
)

func main() {
	var w, sh, p, expected float64
	fmt.Scanf("%f %f", &w, &sh)
	fmt.Scanf("%f", &p)
	expected = (w - (p * sh)) * 365
	fmt.Printf("%.2f", expected)
}
