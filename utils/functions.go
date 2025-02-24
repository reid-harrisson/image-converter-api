package utils

import (
	"fmt"
	"image/color"
)

func ColorToHex(c color.Color) string {
	r, g, b, a := c.RGBA()
	// Normalize the values to 0-255 range
	return fmt.Sprintf("#%02x%02x%02x%02x", r>>8, g>>8, b>>8, a>>8)
}
