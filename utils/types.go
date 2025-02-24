package utils

import "fmt"

type RGB struct {
	R, G, B uint8
}

func NewRGB(hex string) RGB {
	// Parse the hex string
	var r, g, b uint8
	fmt.Sscanf(hex, "%02X%02X%02X", &r, &g, &b)
	return RGB{R: r, G: g, B: b}
}

type Analytic struct {
	Count int
	Color string
}
