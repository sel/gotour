package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	paint := func(x, y int) uint8 { return uint8((x * y) * 4) }
	img := make([][]uint8, dy)
	for y := range img {
		for x := 0; x < dx; x++ {
			img[y] = append(img[y], paint(x, y))
		}
	}
	return img
}

func main() {
	pic.Show(Pic)
}
